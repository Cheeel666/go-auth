package config

import "github.com/spf13/viper"

// MongoDBConfig - config for mongo database.
type MongoDBConfig struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
}

// Config - application configuration.
type Config struct {
	MongoDB MongoDBConfig

	Env     string
	AppName string
}

func InitConfig() *Config {
	var cfg *Config = new(Config)

	// Init with default values.
	cfg.initDefault()

	// Init with file.
	cfg.initFileCfg()
	return cfg
}

func (c *Config) initDefault() {
	viper.SetDefault("mongodb.host", "localhost")
	viper.SetDefault("mongodb.port", 27017)
	viper.SetDefault("mongodb.database", "go-auth")
	viper.SetDefault("mongodb.username", "")
	viper.SetDefault("mongodb.password", "")

	viper.SetDefault("env", "dev")
	viper.SetDefault("appName", "go-auth")
}

func (c *Config) initFileCfg() {
	viper.AddConfigPath(".")
	viper.AddConfigPath("../")
	viper.AddConfigPath("config/")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		panic(err)
	}
}
