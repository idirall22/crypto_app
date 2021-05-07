package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// RepositoryConfig interface
type RepositoryConfig interface {
	RepositoryConfig() string
}

type Config struct {
	Domain         string `mapstructure:"DOMAIN"`
	Port           string `mapstructure:"PORT"`
	BaseURL        string `mapstructure:"BASE_URL"`
	DbHost         string `mapstructure:"DB_HOST"`
	DbDriver       string `mapstructure:"DB_DRIVER"`
	DbUser         string `mapstructure:"DB_USER"`
	DbPassword     string `mapstructure:"DB_PASSWORD"`
	DbName         string `mapstructure:"DB_NAME"`
	DbPort         string `mapstructure:"DB_PORT"`
	JwtPrivatePath string `mapstructure:"JWT_PRIVATE_PATH"`
	JwtPublicPath  string `mapstructure:"JWT_PPUBLIC_PATH"`
}

func New() *Config {
	return &Config{
		Domain:         os.Getenv("DOMAIN"),
		Port:           os.Getenv("PORT"),
		BaseURL:        os.Getenv("BASE_URL"),
		DbHost:         os.Getenv("DB_HOST"),
		DbDriver:       os.Getenv("DB_DRIVER"),
		DbUser:         os.Getenv("DB_USER"),
		DbPassword:     os.Getenv("DB_PASSWORD"),
		DbName:         os.Getenv("DB_NAME"),
		DbPort:         os.Getenv("DB_PORT"),
		JwtPrivatePath: os.Getenv("JWT_PRIVATE_PATH"),
		JwtPublicPath:  os.Getenv("JWT_PUBLIC_PATH"),
	}
}

// RepositoryConfig get postgres string.
func (c Config) RepositoryConfig() (string, string) {
	return c.DbDriver, fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=%s",
		c.DbHost, c.DbPort, c.DbUser, c.DbPassword, c.DbName, "disable")
}

// LoadConfig load env file for test
func LoadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("example")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
