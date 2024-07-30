package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetConfigName("dev")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error reading config file %w", err))
	}
}

func GetPort() string {
	return viper.GetString("port")
}

func GetEnv() string {
	return viper.GetString("environment")
}

func GetUrl() string {
	return viper.GetString("url")
}

func GetDBUrl() string {
	return viper.GetString("dbUrl")
}
