package config

import (
	"os"
	"time"
)

type Config struct {
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`
	DBUrl          string `mapstructure:"DATABASE_URL"`

	Port string `mapstructure:"PORT"`

	JwtSecret    string        `mapstructure:"JWT_SECRET"`
	JwtExpiresIn time.Duration `mapstructure:"JWT_EXPIRED_IN"`
	JwtMaxAge    int           `mapstructure:"JWT_MAXAGE"`

	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`

	OpenAIURL    string `mapstructure:"OPENAI_URL"`
	OpenAIAPIKey string `mapstructure:"OPENAI_API_KEY"`

	HumanPalURL    string `mapstructure:"HUMANPAL_URL"`
	HumanPalAPIKey string `mapstructure:"HUMANPAL_API_KEY"`
}

// func LoadConfig(path string) (config Config, err error) {
// 	viper.AddConfigPath(path)
// 	viper.SetConfigType("env")
// 	viper.SetConfigName("app")

// 	viper.AutomaticEnv()

// 	err = viper.ReadInConfig()
// 	if err != nil {
// 		return
// 	}

// 	err = viper.Unmarshal(&config)
// 	return
// }

func LoadConfig(path string) (config Config, err error) {
	return Config{
		DBUserName:     os.Getenv("POSTGRES_USER"),
		DBHost:         os.Getenv("POSTGRES_HOST"),
		DBUserPassword: os.Getenv("POSTGRES_PASSWORD"),
		DBName:         os.Getenv("POSTGRES_DB"),
		DBPort:         os.Getenv("POSTGRES_PORT"),
		DBUrl:          os.Getenv("DATABASE_URL"),
		JwtSecret:      os.Getenv("JWT_SECRET"),
		Port:           os.Getenv("PORT"),
	}, nil
}
