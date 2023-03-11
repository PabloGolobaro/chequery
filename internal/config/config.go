package config

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

var Global Config
var once sync.Once

type Config struct {
	// the data source name (PostgresDSN) for connecting to the database. required.
	DbUser   string `mapstructure:"DB_USER"`
	DbPass   string `mapstructure:"DB_PASS"`
	DbHost   string `mapstructure:"DB_HOST"`
	DbPort   string `mapstructure:"DB_PORT"`
	DbName   string `mapstructure:"DB_NAME"`
	HttpHost string `mapstructure:"SERVER_HOST"`
	HttpPort string `mapstructure:"SERVER_PORT"`
}

func Load() Config {

	once.Do(func() {
		_ = viper.New()
		viper.AutomaticEnv()

		Global.DbUser = viper.GetString("DB_USER")
		Global.DbPass = viper.GetString("DB_PASS")
		Global.DbHost = viper.GetString("DB_HOST")
		Global.DbPort = viper.GetString("DB_PORT")
		Global.DbName = viper.GetString("DB_NAME")
		Global.HttpHost = viper.GetString("SERVER_HOST")
		Global.HttpPort = viper.GetString("SERVER_PORT")
	})

	return Global
}

func (c Config) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", c.DbUser, c.DbPass, c.HttpHost, c.DbPort, c.DbName)
}
