package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Mode            string `mapstructure:"MODE"`
	Domain          string `mapstructure:"DOMAIN"`
	Port            string `mapstructure:"PORT"`
	BaseURL         string `mapstructure:"BASE_URL"`
	DbHost          string `mapstructure:"DB_HOST"`
	DbDriver        string `mapstructure:"DB_DRIVER"`
	DbUser          string `mapstructure:"DB_USER"`
	DbPassword      string `mapstructure:"DB_PASSWORD"`
	DbName          string `mapstructure:"DB_NAME"`
	DbPort          string `mapstructure:"DB_PORT"`
	ImageBucketName string `mapstructure:"IMAGE_BUCKET_NAME"`
	JwtPrivatePath  string `mapstructure:"JWT_PRIVATE_PATH"`
	JwtPublicPath   string `mapstructure:"JWT_PPUBLIC_PATH"`
}

func New() *Config {
	return &Config{
		Mode:            os.Getenv("MODE"),
		Domain:          os.Getenv("DOMAIN"),
		Port:            os.Getenv("PORT"),
		BaseURL:         os.Getenv("BASE_URL"),
		DbHost:          os.Getenv("DB_HOST"),
		DbDriver:        os.Getenv("DB_DRIVER"),
		DbUser:          os.Getenv("DB_USER"),
		DbPassword:      os.Getenv("DB_PASSWORD"),
		DbName:          os.Getenv("DB_NAME"),
		DbPort:          os.Getenv("DB_PORT"),
		ImageBucketName: os.Getenv("IMAGE_BUCKET_NAME"),
		JwtPrivatePath:  os.Getenv("JWT_PRIVATE_PATH"),
		JwtPublicPath:   os.Getenv("JWT_PUBLIC_PATH"),
	}
}

// RepositoryConfig interface
type RepositoryConfig interface {
	RepositoryConfig() string
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
