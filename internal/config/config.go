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
		v := viper.New()
		viper.AutomaticEnv()

		if err := v.ReadInConfig(); err != nil {
			panic(fmt.Errorf("failed to read the configuration file: %s", err))
		}

		err := v.Unmarshal(&Global)
		if err != nil {
			panic(err)
		}
	})

	return Global
}

func (c Config) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.DbUser, c.DbPass, c.HttpHost, c.DbPort, c.DbName)
}
