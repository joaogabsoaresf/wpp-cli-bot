package messaging

import (
	"fmt"

	"github.com/joaogabsoaresf/wpp-cli-bot/internal/api"
	"github.com/joaogabsoaresf/wpp-cli-bot/internal/models"
)

type Message struct {
	ID      string
	ChatID  string
	Content string
}

func MockMessages() []Message {
	return []Message{
		{ID: "1", ChatID: "1", Content: "Hello!"},
		{ID: "2", ChatID: "1", Content: "How are you?"},
	}
}

func ListMessageByChatId(chat models.Chat) {
	fmt.Println("\nMensagens Recentes:")
	messages, err := api.ApiMessagesByID(chat.ID, 10)
	if err != nil {
		fmt.Printf("erro ao obter os recentes: %v", err)
		return
	}
	fmt.Printf("Mensagens Recentes do(a) %s:\n", chat.Name)
	for _, message := range messages {
		fmt.Printf("\033[32m%-20s %-40s\033[0m\n", chat.Name, message.Content)
	}
}
