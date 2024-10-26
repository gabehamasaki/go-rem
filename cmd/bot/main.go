package main

import (
	"log"

	"github.com/gabehamasaki/go-rem/internal/bot"
	"github.com/gabehamasaki/go-rem/internal/commands"
	"github.com/gabehamasaki/go-rem/internal/config"
)

func main() {
	config := config.NewConfig()

	bot, err := bot.NewBot(config)
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	bot.AddHandler(&commands.World{})
	bot.AddHandler(commands.NewTalk(64, config.ChatID))

	bot.Start()
	defer bot.Stop()
	<-make(chan struct{})
}
