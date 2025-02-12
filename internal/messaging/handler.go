package messaging

import (
	"fmt"

	"github.com/joaogabsoaresf/wpp-cli-bot/internal/api"
)

func ProcessMessages() {
	messages := MockMessages()
	for _, msg := range messages {
		reply := fmt.Sprintf("Replying to: %s", msg.Content)
		api.SendReply(msg.ID, reply)
	}
}
