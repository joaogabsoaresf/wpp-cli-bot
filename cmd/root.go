package cmd

import (
	"fmt"
	"os"

	"github.com/joaogabsoaresf/wpp-cli-bot/internal/config"
	"github.com/joaogabsoaresf/wpp-cli-bot/internal/connectors/zapi"
	"github.com/joaogabsoaresf/wpp-cli-bot/internal/messaging"
	"github.com/joaogabsoaresf/wpp-cli-bot/internal/utils"
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
				zapi.GetMetaDataFromDefault()
			case "Sair":
				fmt.Println("Saindo... Até logo!")
				return // Sai do loop e encerra o programa
			default:
				fmt.Println("Opção inválida")
			}
		}
	},
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the WhatsApp bot configuration",
	Long:  `This command initializes the necessary configurations for the WhatsApp bot.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Inicializando configuração...")
		instanceIdPrompt := promptui.Prompt{
			Label: "Digite o ID da Instância",
			Validate: func(s string) error {
				if s == "" {
					return fmt.Errorf("ID da instância não pode ser vazio")
				}
				return nil
			},
		}
		instanceID, err := instanceIdPrompt.Run()
		if err != nil {
			fmt.Printf("Erro ao capturar Instance ID: %v\n", err)
			return
		}
		instanceTokenPrompt := promptui.Prompt{
			Label: "Digite o Token da Instância",
			Validate: func(input string) error {
				if input == "" {
					return fmt.Errorf("instance Token não pode estar vazio")
				}
				return nil
			},
		}
		instanceToken, err := instanceTokenPrompt.Run()
		if err != nil {
			fmt.Printf("Erro ao capturar Instance Token: %v\n", err)
			return
		}

		clientTokenPrompt := promptui.Prompt{
			Label: "Digite o Client Token",
			Validate: func(input string) error {
				if input == "" {
					return fmt.Errorf("client Token não pode estar vazio")
				}
				return nil
			},
		}
		clientToken, err := clientTokenPrompt.Run()
		if err != nil {
			fmt.Printf("Erro ao capturar Client Token: %v\n", err)
			return
		}

		fmt.Println("\nConfiguração capturada com sucesso:")
		configData := config.ZAPIConfigData{
			InstanceID:    instanceID,
			InstanceToken: instanceToken,
			ClientToken:   clientToken,
		}

		filePath := "z_api_config_file.json"

		err = utils.SaveToJson(configData, filePath)
		if err != nil {
			fmt.Println("Erro ao salvar arquivo JSON:", err)
			return
		}

		fmt.Println("Configuração salva com sucesso!")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
