package messaging

import (
	"fmt"

	"github.com/joaogabsoaresf/wpp-cli-bot/internal/api"
	"github.com/manifoldco/promptui"
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

func ListChatsWithSelection() {
	// Obtém os chats recentes (max. 10)
	chats, err := api.ApiRecentChat(10)
	if err != nil {
		fmt.Printf("erro ao obter os recentes: %v\n", err)
		return
	}

	// Exibe o título para a listagem de chats
	fmt.Println("\nConversas Recentes:")

	// Prepara as opções para o promptui
	var items []string
	for _, chat := range chats {
		item := fmt.Sprintf("%s - %s", chat.Name, chat.LastMsg)
		items = append(items, item)
	}

	// Função personalizada para alternar as cores das opções
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

	// Encontra o chat selecionado
	selectedID := -1
	for _, chat := range chats {
		if fmt.Sprintf("%d - %s", chat.ID, chat.Name) == selectedChat {
			selectedID = chat.ID
			break
		}
	}

	// Exibe a última mensagem do chat selecionado
	if selectedID != -1 {
		for _, chat := range chats {
			if chat.ID == selectedID {
				// Exibe a última mensagem com a mesma formatação
				fmt.Printf("\nÚltima Mensagem de %s: %s\n", chat.Name, chat.LastMsg)
				break
			}
		}
	}
}
