package cmd

import (
	"fmt"
	"os"

	"github.com/joaogabsoaresf/wpp-cli-bot/internal/connectors"
	"github.com/joaogabsoaresf/wpp-cli-bot/internal/messaging"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wpp-cli-bot",
	Short: "A simple WhatsApp bot using Cobra CLI",
	Long:  `This is a CLI tool to interact with a WhatsApp bot, reply to messages, and more.`,
	Run: func(cmd *cobra.Command, args []string) {
		for {
			prompt := promptui.Select{
				Label: "Escolha uma opção",
				Items: []string{
					"Listar Chats Recentes",
					"Selecionar Chat Recente",
					"Metadados do Número Default",
					"Sair", // Adicionando a opção de sair
				},
			}

			_, result, err := prompt.Run()

			if err != nil {
				fmt.Println("Erro ao fazer a seleção:", err)
				os.Exit(1)
			}

			switch result {
			case "Listar Chats Recentes":
				messaging.ListRecentChats()
			case "Selecionar Chat Recente":
				messaging.ListChatsWithSelection()
			case "Metadados do Número Default":
				connectors.GetMetaDataFromDefault()
			case "Sair":
				fmt.Println("Saindo... Até logo!")
				return // Sai do loop e encerra o programa
			default:
				fmt.Println("Opção inválida")
			}
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
