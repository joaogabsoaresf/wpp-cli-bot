package messaging

import (
	"fmt"

	"github.com/joaogabsoaresf/wpp-cli-bot/internal/connectors/zapi"
	"github.com/manifoldco/promptui"
)

func ListRecentChats() {
	chats, err := zapi.GetLastChats()
	if err != nil {
		fmt.Printf("erro ao obter os recentes: %v", err)
		return
	}
	fmt.Println("Chats Recentes:")
	for i, chat := range chats {
		displayValue := chat.Name
		if displayValue == "" {
			displayValue = chat.Phone
		}
		if i%2 == 0 {
			fmt.Printf("\033[32m%-20s\n", displayValue)
		} else {
			fmt.Printf("\033[34m%-20s\n", displayValue)
		}
	}
}

func ListChatsWithSelection() {
	chats, err := zapi.GetLastChats()
	if err != nil {
		fmt.Printf("erro ao obter os recentes: %v", err)
		return
	}

	fmt.Println("\nChats:")

	var items []string
	for i, chat := range chats {
		displayValue := chat.Name
		if displayValue == "" {
			displayValue = chat.Phone
		}
		item := fmt.Sprintf("%d - %s", i, displayValue)
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
	for i, chat := range chats {
		displayValue := chat.Name
		if displayValue == "" {
			displayValue = chat.Phone
		}
		if fmt.Sprintf("%d - %s", i, displayValue) == selectedChat {
			selectedChatID = chat.Phone
			break
		}
	}

	// Se o chat foi encontrado, buscar as mensagens
	if selectedChatID != "" {
		for _, chat := range chats {
			if chat.Phone == selectedChatID {
				ListMessageByChatId(chat)
				break
			}
		}
	}
}
