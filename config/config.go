package config

type Config struct {
	Host    string `mapstructure:"serv_host"`
	Port    string `mapstructure:"serv_port"`
	Reflect bool   `mapstructure:"reflection"`
	Logrus  Logrus `mapstructure:"logrus"`
}

type Logrus struct {
	LogLvl int    `mapstructure:"log_level"`
	ToFile bool   `mapstructure:"to_file"`
	ToJson bool   `mapstructure:"to_json"`
	LogDir string `mapstructure:"log_dir"`
}
