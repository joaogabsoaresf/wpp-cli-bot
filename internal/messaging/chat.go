package messaging

import (
	"fmt"
	"strconv"

	"github.com/joaogabsoaresf/wpp-cli-bot/internal/api"
	"github.com/manifoldco/promptui"
)

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

func ListChatsWithSelection() {
	chats, err := api.ApiRecentChat(10)
	if err != nil {
		fmt.Printf("erro ao obter os recentes: %v\n", err)
		return
	}

	fmt.Println("\nConversas Recentes:")

	var items []string
	for _, chat := range chats {
		item := fmt.Sprintf("%d - %s - %s", chat.ID, chat.Name, chat.LastMsg)
		items = append(items, item)
	}

	prompt := promptui.Select{
		Label: "Selecione um Chat",
		Items: items,
		Templates: &promptui.SelectTemplates{
			Active:   "\033[32m{{ . | bold }}\033[0m", // Cor verde para a opção ativa
			Inactive: "\033[34m{{ . }}\033[0m",        // Cor azul para as opções inativas
		},
	}

	_, selectedChat, err := prompt.Run()
	if err != nil {
		fmt.Println("Erro ao selecionar o chat:", err)
		return
	}

	// Encontrar o chat selecionado e pegar o ID correto
	var selectedChatID string
	for _, chat := range chats {
		if fmt.Sprintf("%d - %s - %s", chat.ID, chat.Name, chat.LastMsg) == selectedChat {
			selectedChatID = strconv.Itoa(chat.ID)
			break
		}
	}

	// Se o chat foi encontrado, buscar as mensagens
	if selectedChatID != "" {
		for _, chat := range chats {
			if strconv.Itoa(chat.ID) == selectedChatID {
				ListMessageByChatId(chat)
				break
			}
		}
	}
}
