package cmd

import (
	"github.com/joaogabsoaresf/wpp-cli-bot/internal/messaging"

	"github.com/spf13/cobra"
)

// listChatsCmd represents the listChats command
var listChatsCmd = &cobra.Command{
	Use:   "list-chats",
	Short: "Lists recent chats",
	Long:  `This command will list all the recent chats with the last message preview.`,
	Run: func(cmd *cobra.Command, args []string) {
		messaging.ListRecentChats()
	},
}

func init() {
	rootCmd.AddCommand(listChatsCmd)
}
