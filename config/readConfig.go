package config

import (
	"log"

	"github.com/spf13/viper"

)

var Config *Configuration

type Configuration struct {
	Server          ServerConfiguration
	Database        DatabaseConfiguration
	LogConfig       LogConfig
	AWSConfig       AWSConfiguration
	DBLogConfig     DbLogConfig
	DBSalesTracking DbSalesTracking
}

type ServerConfiguration struct {
	AppName string
	Port    string
	Secret  string
	Mode    string
	Env     string
}

type DatabaseConfiguration struct {
	Driver       string
	Dbname       string
	Username     string
	Password     string
	Host         string
	Port         string
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
}

type LogConfig struct {
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	MaxSize    int    `json:"maxsize"`
	MaxAge     int    `json:"max_age"`
	MaxBackups int    `json:"max_backups"`
}

type AWSConfiguration struct {
	AWSRegion          string
	AWSAccessKeyId     string
	AWSSecretAccessKey string
	BucketName1        string
	BucketName2        string
	BucketName3        string
}

type DbLogConfig struct {
	Driver       string
	Dbname       string
	Username     string
	Password     string
	Host         string
	Port         string
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
}

type DbSalesTracking struct {
	Driver       string
	Dbname       string
	Username     string
	Password     string
	Host         string
	Port         string
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
}

// Setup SetupDB initialize configuration
func Setup(configPath string) {
	var configuration *Configuration

	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("Unable to decode into structs, %v", err)
	}

	Config = configuration
}

func GetConfig() *Configuration {
	return Config
}
