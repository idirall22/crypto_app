package config

import (
	"os"
)

// RepositoryConfig interface
type RepositoryConfig interface {
	RepositoryConfig() string
}

type Config struct {
	Port             string `mapstructure:"PORT"`
	GMailPassword    string `mapstructure:"GMAIL_PASSWORD"`
	GMailEmail       string `mapstructure:"GMAIL_EMAIL"`
	GMailSMTPPort    string `mapstructure:"GMAIL_SMTP_PORT"`
	GMailSMTP        string `mapstructure:"GMAIL_SMTP"`
	JwtPrivatePath   string `mapstructure:"JWT_PRIVATE_PATH"`
	JwtPublicPath    string `mapstructure:"JWT_PPUBLIC_PATH"`
	RabbitMQUser     string `mapstructure:"RABBITMQ_USER"`
	RabbitMQHost     string `mapstructure:"RABBITMQ_HOST"`
	RabbitMQPassword string `mapstructure:"RABBITMQ_PASSWORD"`
	RabbitMQPort     string `mapstructure:"RABBITMQ_PORT"`
}

func New() *Config {
	return &Config{
		Port:             os.Getenv("PORT"),
		GMailPassword:    os.Getenv("GMAIL_PASSWORD"),
		GMailEmail:       os.Getenv("GMAIL_EMAIL"),
		GMailSMTPPort:    os.Getenv("GMAIL_SMTP_PORT"),
		GMailSMTP:        os.Getenv("GMAIL_SMTP"),
		JwtPrivatePath:   os.Getenv("JWT_PRIVATE_PATH"),
		JwtPublicPath:    os.Getenv("JWT_PPUBLIC_PATH"),
		RabbitMQUser:     os.Getenv("RABBITMQ_USER"),
		RabbitMQHost:     os.Getenv("RABBITMQ_HOST"),
		RabbitMQPassword: os.Getenv("RABBITMQ_PASSWORD"),
		RabbitMQPort:     os.Getenv("RABBITMQ_PORT"),
	}
}
