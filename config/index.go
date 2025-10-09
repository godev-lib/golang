package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Database DatabaseConfig `mapstructure:"database"`
	JwtKey   string         `mapstructure:"jwt-key"`
}

type AppConfig struct {
	Name     string `mapstructure:"name"`
	Host     string `mapstructure:"host"`
	GrpcPort string `mapstructure:"grpc_port"`
	HttpPort string `mapstructure:"http_port"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
}

func NewConfig() *Config {
	s := &Config{}
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("error load config: ", err.Error())
		return nil
	}

	if err := viper.Unmarshal(&s); err != nil {
		fmt.Println("error map config to struct: ", err.Error())
		return nil
	}

	return s
}
