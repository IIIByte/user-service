package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type AppConfig struct {
	ServerConfig   *ServerConfig
	PostgresConfig *PostgresConfig
}

// ServerConfig - конфигурация для запуска сервера
type ServerConfig struct {
	Addr string
	Port string
}

// PostgresConfig - конфигурация для PostgreSQL
type PostgresConfig struct {
	Host           string
	Port           int
	User           string
	Password       string
	DBName         string
	MaxConnections int
}

func Load() (*AppConfig, error) {
	appConfig := &AppConfig{}

	v := viper.New()

	// Указываем, что работаем с .env файлом
	v.SetConfigName(".env")
	v.SetConfigType("env")
	v.AddConfigPath("../")

	// Читаем .env файл
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("Config (Load): failed load config to struct: %v", err)
	}

	// Автоматически учитываем переменные окружения (без префикса)
	v.AutomaticEnv()

	appConfig.ServerConfig = &ServerConfig{
		Addr: v.GetString("SERVER_ADDR"),
		Port: v.GetString("SERVER_PORT"),
	}

	appConfig.PostgresConfig = &PostgresConfig{
		Host:           v.GetString("POSTGRES_HOST"),
		Port:           v.GetInt("POSTGRES_PORT"),
		User:           v.GetString("POSTGRES_USER"),
		Password:       v.GetString("POSTGRES_PASSWORD"),
		DBName:         v.GetString("POSTGRES_DBNAME"),
		MaxConnections: v.GetInt("POSTGRES_MAX_CONNECTIONS"),
	}

	return appConfig, nil
}
