package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Token  string
	Prefix string
	ChatID string
}

func NewConfig() *Config {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("BOT_TOKEN")
	prefix := os.Getenv("BOT_PREFIX")
	chatID := os.Getenv("BOT_CHAT_ID")

	if token == "" || prefix == "" {
		log.Fatal("BOT_TOKEN or BOT_PREFIX is not set")
	}

	return &Config{
		Token:  token,
		Prefix: prefix,
		ChatID: chatID,
	}
}
