package config

import (
	"errors"
	"github.com/cherrycutter/todo_app/pkg/logger"
	"github.com/spf13/viper"
	"sync"
)

type Config struct {
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
}

var (
	config Config
	once   sync.Once
)

// LoadConfig loads a new config .json file from directory configs
func LoadConfig() Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("json")
		viper.AddConfigPath(".")
		viper.AddConfigPath("./configs")
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			logger.Error.Fatalf("error reading config file, %s", err)
		}

		config = Config{
			DBHost: viper.GetString("DBHost"),
			DBPort: viper.GetString("DBPort"),
			DBUser: viper.GetString("DBUser"),
			DBPass: viper.GetString("DBPass"),
			DBName: viper.GetString("DBName"),
		}

		if err := validateConfig(config); err != nil {
			logger.Error.Fatal(err)
		}
		logger.Info.Println("Configuration loaded successfully")
	})
	return config
}

// validateConfig validates a config file for an empty fields
func validateConfig(config Config) error {
	if config.DBHost == "" || config.DBPort == "" || config.DBUser == "" || config.DBPass == "" || config.DBName == "" {
		return errors.New("missing field in config file")
	}
	return nil
}
