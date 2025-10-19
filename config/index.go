package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Database DatabaseConfig `mapstructure:"database"`
	Pusher   PusherConfig   `mapstructure:"pusher"`
	JwtKey   string         `mapstructure:"jwt-key"`
	Minio    MinioConfig    `mapstructure:"minio"`
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

type PusherConfig struct {
	AppID                       string `mapstructure:"app_id"`
	Key                         string `mapstructure:"key"`
	Secret                      string `mapstructure:"secret"`
	Host                        string `mapstructure:"host"`
	Secure                      bool   `mapstructure:"secure"`
	Cluster                     string `mapstructure:"cluster"`
	EncryptionMasterKey         string `mapstructure:"encryption_master_key"`
	EncryptionMasterKeyBase64   string `mapstructure:"encryption_master_key_base64"`
	OverrideMaxMessagePayloadKB int    `mapstructure:"override_max_message_payload_kb"`
}

type MinioConfig struct {
	Endpoint        string `mapstructure:"endpoint"`
	AccessKeyID     string `mapstructure:"access_key_id"`
	SecretAccessKey string `mapstructure:"secret_access_key"`
	UseSSL          bool   `mapstructure:"use_ssl"`
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
