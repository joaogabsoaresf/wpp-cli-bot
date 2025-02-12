package messaging

import (
	"fmt"
	"time"

	"github.com/joaogabsoaresf/wpp-cli-bot/internal/api"
	"github.com/joaogabsoaresf/wpp-cli-bot/internal/models"
	"github.com/manifoldco/promptui"
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

func showLoading(duration time.Duration) {
	frames := []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}
	done := make(chan bool)

	go func() {
		for i := 0; ; i++ {
			select {
			case <-done:
				return
			default:
				// Exibe o frame atual do loading
				fmt.Printf("\r\033[KCarregando mensagens %s", frames[i%len(frames)])
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	// Simula o tempo de carregamento do WebSocket
	time.Sleep(duration)
	done <- true
	fmt.Printf("\r\033[KMensagens carregadas!\n")
}

func ListMessageByChatId(chat models.Chat) {
	showLoading(3 * time.Second)

	// Limpa o terminal antes de exibir as mensagens
	fmt.Print("\033[H\033[2J")

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

	for {
		prompt := promptui.Prompt{
			Label: "Digite sua mensagem (ou :exit para sair)",
			Validate: func(input string) error {
				if input == ":exit" {
					return nil
				}
				return nil
			},
		}

		message, err := prompt.Run()
		if err != nil {
			fmt.Printf("erro ao capturar a mensagem: %v\n", err)
			return
		}

		if message == ":exit" {
			fmt.Println("Saindo do modo de digitação...")
			return
		}

		fmt.Print("\033[1A\033[K")

		fmt.Printf("\033[34m%-20s %-40s\033[0m\n", "Eu", message)
	}
}
