package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName string
	AppEnv  string
	Database
}

type Database struct {
	Mysql MysqlConfig
	Mongo MongoConfig
}

type MysqlConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
}

type MongoConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
}

var (
	cfg *Config
)

func GetConfig() *Config {
	return &Config{
		AppName: os.Getenv("APP_NAME"),
		AppEnv:  GetEnv(),
		Database: Database{
			Mysql: MysqlConfig{
				Host:     os.Getenv("DB_MYSQL_HOST"),
				User:     os.Getenv("DB_MYSQL_USER"),
				Password: os.Getenv("DB_MYSQL_PASSWORD"),
				Name:     os.Getenv("DB_MYSQL_NAME"),
				Port:     os.Getenv("DB_MYSQL_PORT"),
			},
			Mongo: MongoConfig{
				Host:     os.Getenv("DB_MONGO_HOST"),
				User:     os.Getenv("DB_MONGO_USER"),
				Password: os.Getenv("DB_MONGO_PASSWORD"),
				Name:     os.Getenv("DB_MONGO_NAME"),
				Port:     os.Getenv("DB_MONGO_PORT"),
			},
		},
	}
}

func InitConfig() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetEnv() string {
	switch os.Getenv("APP_ENV") {
	case "production":
		return "production"
	case "staging":
		return "staging"
	default:
		return "development"
	}
}

func GetAppName() string {
    if cfg == nil {
        cfg = GetConfig()
    }
    return cfg.AppName
}
