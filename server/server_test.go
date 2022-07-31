package server

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	storage "github.com/s-vvardenfell/BooksStorage/books_storage"
)

func RunServer(driver, dsn, host, port string) {
	grpcServ := grpc.NewServer()
	serv := New(driver, dsn)
	storage.RegisterBooksStorageServer(grpcServ, serv)

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}

	logrus.Info("Starting gRPC listener on port " + port)
	if err := grpcServ.Serve(lis); err != nil {
		logrus.Fatalf("failed to serve: %v", err)
	}
}

func Test_Service(t *testing.T) {
	viper.AutomaticEnv()

	// получаем dsn из окружения
	go RunServer("mysql", viper.GetString("DSN"), "localhost", "50052")
	time.Sleep(time.Second)

	cnxn, err := sql.Open("mysql", viper.GetString("DSN"))
	require.NoError(t, err)

	if err := CreateTables(cnxn); err == nil {
		if err := PopulateTables(cnxn); err != nil {
			require.NoError(t, err)
		}
	}

	defer func() {
		if err := CleanUp(cnxn); err != nil {
			require.NoError(t, err)
		}

		if err := cnxn.Close(); err != nil {
			require.NoError(t, err)
		}
	}()

	client := NewClientStub("localhost", "50052")
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	t.Log("\tGetting books by author name")
	{
		author := "Достаевский Ф.М."
		res, err := client.GetBooksByAuthor(ctx, &storage.Author{AuthorName: author})
		require.NoError(t, err)
		require.NotNil(t, res)
		require.Equal(t, "Идиот", res.BookNames[0])
	}

	t.Log("\tGetting author by book title")
	{
		book := "Идиот"
		res, err := client.GetAuthorsByBook(ctx, &storage.Book{BookName: book})
		require.NoError(t, err)
		require.NotNil(t, res)
		require.Equal(t, "Достаевский Ф.М.", res.AuthorNames[0])
	}

	t.Log("\tGetting book authors (more that 1 author)")
	{
		book := "Понедельник начинается в субботу"
		res, err := client.GetAuthorsByBook(ctx, &storage.Book{BookName: book})
		require.NoError(t, err)
		require.NotNil(t, res)
		require.Equal(t, 2, len(res.AuthorNames))
	}

	t.Log("\tGetting books by wrong author name")
	{
		author := "Толкин Д.Р.Р."
		res, err := client.GetBooksByAuthor(ctx, &storage.Author{AuthorName: author})
		require.NoError(t, err)
		require.Empty(t, res.BookNames)
	}

	t.Log("\tGetting author by wrong book title")
	{
		book := "Властелин колец"
		res, err := client.GetAuthorsByBook(ctx, &storage.Book{BookName: book})
		require.NoError(t, err)
		require.Empty(t, res.AuthorNames)
	}
}
