package cmd

import (
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/s-vvardenfell/BooksStorage/books_storage"
	"github.com/s-vvardenfell/BooksStorage/config"
	"github.com/s-vvardenfell/BooksStorage/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var cnfg config.Config
var cfgFile string

const driver = "mysql"

var rootCmd = &cobra.Command{
	Use:   "BooksStorage",
	Short: "Simple books-authors storage",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := cmd.Flags().GetBool("create")
		cobra.CheckErr(err)

		if c { // создаем таблицы, если передан флаг -c
			server.CreateTables(driver, viper.GetString("DSN"))
		}

		p, err := cmd.Flags().GetBool("populate")
		cobra.CheckErr(err)

		if p { // заполняем таблицы данными, если передан флаг -p
			server.PopulateTables(driver, viper.GetString("DSN"))
		}

		grpcServ := grpc.NewServer()

		//получаем dsn из переменной окружения
		rcs := server.New(driver, viper.GetString("DSN"))
		books_storage.RegisterBooksStorageServer(grpcServ, rcs)

		lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cnfg.Host, cnfg.Port))
		if err != nil {
			logrus.Fatalf("failed to listen: %v", err)
		}

		// для использования, например, grpcui
		if cnfg.Reflect {
			reflection.Register(grpcServ)
		}

		logrus.Info("Starting gRPC listener on port " + cnfg.Port)
		if err := grpcServ.Serve(lis); err != nil {
			logrus.Fatalf("failed to serve: %v", err)
		}

		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT)
		<-sig
		grpcServ.GracefulStop()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile,
		"config", "", "config file (default is resources/config.yml)")

	rootCmd.Flags().BoolP("create", "c", false, "creates tables")
	rootCmd.Flags().BoolP("populate", "p", false, "populates tables")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		wd, err := os.Getwd()
		cobra.CheckErr(err)

		// конфиг по-умолчанию
		viper.AddConfigPath(filepath.Join(wd, "resources"))
		viper.SetConfigName("config")
		viper.SetConfigType("yml")
	}

	// получаем доступ к переменным окружения
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	} else {
		cobra.CheckErr(err)
	}

	if err := viper.Unmarshal(&cnfg); err != nil {
		cobra.CheckErr(err)
	}

	// если логи будут писаться в файл, создаем директорию
	if cnfg.Logrus.ToFile {
		if err := os.Mkdir(filepath.Dir(cnfg.Logrus.LogDir), 0644); err != nil && !errors.Is(err, os.ErrExist) {
			cobra.CheckErr(err)
		}

		file, err := os.OpenFile(cnfg.Logrus.LogDir, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err == nil {
			logrus.SetOutput(file)
		} else {
			cobra.CheckErr(err)
		}
	}

	// если логи нужны в формате json
	if cnfg.Logrus.ToJson {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}

	// уровень логирования
	logrus.SetLevel(logrus.Level(cnfg.Logrus.LogLvl))
}
