package messaging

import (
	"fmt"

	"github.com/joaogabsoaresf/wpp-cli-bot/internal/api"
)

type Chat struct {
	ID      int
	Name    string
	LastMsg string
}

func ListRecentChats() {
	chats, err := api.ApiRecentChat(10)
	if err != nil {
		fmt.Printf("erro ao obter os recentes: %v", err)
		return
	}
	fmt.Println("Chats Recentes:")
	for i, chat := range chats {
		if i%2 == 0 {
			// Linha ímpar (cor verde)
			fmt.Printf("\033[32m%-20s %-40s\033[0m\n", chat.Name, chat.LastMsg)
		} else {
			// Linha par (cor azul)
			fmt.Printf("\033[34m%-20s %-40s\033[0m\n", chat.Name, chat.LastMsg)
		}
	}
}
