package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type World struct{}

func (w *World) Command() string {
	return "world"
}

func (w *World) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	log.Println("World command received")

	_, _ = s.ChannelMessageSend(m.ChannelID, "Hello World!")

}

func (w *World) ChatID() string {
	return ""
}
