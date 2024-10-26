package commands

import (
	"context"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/gabehamasaki/go-rem/internal/ollama"
)

type Talk struct {
	history          []ollama.Message
	maxHistoryLength int
	chatID           string
}

func NewTalk(maxHistoryLength int, chatID string) *Talk {
	return &Talk{
		history:          make([]ollama.Message, 0, maxHistoryLength),
		maxHistoryLength: maxHistoryLength,
		chatID:           chatID,
	}
}

func (t *Talk) Command() string {
	return "talk"
}

func (t *Talk) ChatID() string {
	return t.chatID
}

func (t *Talk) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	prompt := t.extractPrompt(m)
	if prompt == "" {
		s.ChannelMessageSendReply(m.ChannelID, "Please provide a message to talk about.", &discordgo.MessageReference{
			MessageID: m.Message.ID,
		})
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := s.ChannelTyping(m.ChannelID, discordgo.WithContext(ctx)); err != nil {
		log.Printf("Error setting typing status: %v", err)
	}

	log.Printf("Received message from %s\n", m.Author.Username)

	response, err := t.getResponse(ctx, prompt)
	if err != nil {
		log.Printf("Error getting response: %v", err)
		s.ChannelMessageSendReply(m.ChannelID, "Sorry, I encountered an error while processing your request.", &discordgo.MessageReference{
			MessageID: m.Message.ID,
		})
		return
	}

	t.updateHistory(prompt, response)

	if err := t.sendResponse(s, m, response); err != nil {
		log.Printf("Error sending message: %v", err)
		_, err = s.ChannelMessageSendReply(m.ChannelID, "Sorry, I encountered an error while sending my response.", &discordgo.MessageReference{
			MessageID: m.Message.ID,
		})
		if err != nil {
			log.Printf("Error sending error message: %v", err)
		}
		// Add reaction separately
		err = s.MessageReactionAdd(m.ChannelID, m.ID, "❌")
		if err != nil {
			log.Printf("Error adding reaction: %v", err)
		}
		return
	}

	log.Printf("Sent message to %s\n", m.Author.Username)
}

func (t *Talk) extractPrompt(m *discordgo.MessageCreate) string {
	if m.ChannelID == t.ChatID() {
		return m.Content
	}

	messages := strings.Split(m.Content, " ")
	if len(messages) < 2 {
		return ""
	}
	return strings.Join(messages[1:], " ")
}

func (t *Talk) getResponse(ctx context.Context, prompt string) (string, error) {
	responseChan := make(chan string)
	errChan := make(chan error)

	go func() {
		response, err := ollama.SendPrompt(prompt, t.history)
		if err != nil {
			errChan <- err
			return
		}
		responseChan <- response
	}()

	select {
	case response := <-responseChan:
		return response, nil
	case err := <-errChan:
		return "", err
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func (t *Talk) updateHistory(prompt, response string) {
	t.addToHistory(ollama.Message{Role: "user", Content: prompt})
	t.addToHistory(ollama.Message{Role: "assistant", Content: response})
}

func (t *Talk) addToHistory(message ollama.Message) {
	if len(t.history) >= t.maxHistoryLength {
		t.history = t.history[1:]
	}
	t.history = append(t.history, message)
}

func (t *Talk) sendResponse(s *discordgo.Session, m *discordgo.MessageCreate, response string) error {
	// Send the reply message
	_, err := s.ChannelMessageSendReply(m.ChannelID, response, &discordgo.MessageReference{
		MessageID: m.Message.ID,
	})

	return err
}
