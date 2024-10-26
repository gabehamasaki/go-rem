package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/gabehamasaki/go-rem/internal/config"
	"github.com/gabehamasaki/go-rem/internal/handlers"
)

type Bot struct {
	session  *discordgo.Session
	bot      *discordgo.User
	config   *config.Config
	handlers *handlers.HandlerManager
}

func NewBot(config *config.Config) (*Bot, error) {
	session, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		return nil, err
	}
	user, err := session.User("@me")
	if err != nil {
		return nil, err
	}

	bot := &Bot{
		session: session,
		config:  config,
		bot:     user,
	}

	handlers := handlers.NewHandlerManager(user.ID, config.Prefix)
	bot.handlers = handlers

	return bot, nil
}

func (b *Bot) Start() error {

	b.session.AddHandler(b.handlers.RegisterHandlers)

	err := b.session.Open()
	if err != nil {
		return err
	}

	log.Println(b.bot.Username, "is online!")
	return nil
}

func (b *Bot) AddHandler(handler handlers.Handler) {
	b.handlers.AddHandler(handler)
}

func (b *Bot) Stop() error {
	return b.session.Close()
}

func (b *Bot) GetBotID() string {
	return b.bot.ID
}

func (b *Bot) GetPrefix() string {
	return b.config.Prefix
}
