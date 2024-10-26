package handlers

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type HandlerManager struct {
	botID    string
	prefix   string
	handlers map[string]Handler
}

type Handler interface {
	Command() string
	ChatID() string
	Handle(s *discordgo.Session, m *discordgo.MessageCreate)
}

func NewHandlerManager(botID, prefix string) *HandlerManager {
	return &HandlerManager{
		botID:    botID,
		prefix:   prefix,
		handlers: make(map[string]Handler),
	}
}

func (hm *HandlerManager) AddHandler(handler Handler) {
	command := handler.Command()
	log.Printf("Adding handler: %s\n", command)
	hm.handlers[command] = handler
}

func (hm *HandlerManager) RegisterHandlers(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == hm.botID {
		return
	}

	content := m.Content
	channelID := m.ChannelID

	for _, handler := range hm.handlers {
		if chatID := handler.ChatID(); chatID != "" && channelID == chatID {
			handler.Handle(s, m)
			return
		}

		command := handler.Command()
		if strings.HasPrefix(content, hm.prefix+command) || content == "<@"+hm.botID+"> "+command {
			handler.Handle(s, m)
			return
		}
	}
}
