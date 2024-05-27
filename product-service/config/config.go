package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	DBHost          string
	DBPort          string
	DBUser          string
	DBPassword      string
	DBName          string
	MinioEndpoint   string
	MinioAccessKey  string
	MinioSecretKey  string
	MinioBucketName string
}

var AppConfig Config

func LoadConfig() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	AppConfig = Config{
		DBHost:          viper.GetString("DB_HOST"),
		DBPort:          viper.GetString("DB_PORT"),
		DBUser:          viper.GetString("DB_USER"),
		DBPassword:      viper.GetString("DB_PASSWORD"),
		DBName:          viper.GetString("DB_NAME"),
		MinioEndpoint:   viper.GetString("MINIO_ENDPOINT"),
		MinioAccessKey:  viper.GetString("MINIO_ACCESS_KEY"),
		MinioSecretKey:  viper.GetString("MINIO_SECRET_KEY"),
		MinioBucketName: viper.GetString("MINIO_BUCKET_NAME"),
	}
}
