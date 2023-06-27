package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

// Configs names
type config struct {
	BotToken string

	PostgresUser         string
	PostgresPassword     string
	PostgresDataBaseName string
	PostgresHost         string
	PostgresPort         string
	PostgresSSLMode      string
}

var Config config

func init() {

	err := searchConfig()
	if err != nil {
		fmt.Println("Config Viper error: ", err)
		return
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		log.Printf("unable to decode into struct, %v\n", err)
	}
}

func searchConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	return viper.ReadInConfig()
}

//Configs getters

// GetToken get telegram bot token from config.*
func GetToken() string {
	return Config.BotToken
}

// GetPostgresUser get postgres user name from config.*
func GetPostgresUser() string {
	return Config.PostgresUser
}

// GetPostgresPassword get postgres user password from config.*
func GetPostgresPassword() string {
	return Config.PostgresPassword
}

// PostgresDataBaseName get postgres db name from config.*
func PostgresDataBaseName() string {
	return Config.PostgresDataBaseName
}

// GetPostgresHost get postgres db host from config.*
func GetPostgresHost() string {
	return Config.PostgresHost
}

// PostgresPort get postgres db port from config.*
func PostgresPort() string {
	return Config.PostgresPort
}

// PostgresSSLMode get postgres db sslmode from config.*
func PostgresSSLMode() string {
	return Config.PostgresSSLMode
}
