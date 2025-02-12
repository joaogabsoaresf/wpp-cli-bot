package cmd

import (
	"fmt"

	"github.com/joaogabsoaresf/wpp-cli-bot/internal/messaging"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the WhatsApp bot",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting bot...")
		messaging.ProcessMessages()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
