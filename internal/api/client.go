package api

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/joaogabsoaresf/wpp-cli-bot/internal/models"
)

func SendReply(msgID, reply string) {
	fmt.Printf("Mock API: Replying to message ID %s with: %s\n", msgID, reply)
}

func ApiRecentChat(chatLimit int) ([]models.Chat, error) {
	if chatLimit == 0 {
		chatLimit = 5
	}

	file, err := os.Open("mock-responses/chats.json")
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir o arquivo JSON: %v", err)
	}

	defer file.Close()

	var chats []models.Chat
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&chats)
	if err != nil {
		return nil, fmt.Errorf("erro ao decodificar o JSON: %v", err)
	}

	time.Sleep(300 * time.Millisecond)

	if chatLimit > len(chats) {
		chatLimit = len(chats)
	}

	return chats[:chatLimit], nil
}

func ApiMessagesByID(ChatID int, messageLimit int) ([]models.Message, error) {
	var messageResponse []models.Message
	if messageLimit == 0 {
		messageLimit = 10
	}

	file, err := os.Open("mock-responses/messages.json")
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir o arquivo JSON: %v", err)
	}

	defer file.Close()

	var messages []models.Message
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&messages)
	if err != nil {
		return nil, fmt.Errorf("erro ao decodificar o JSON: %v", err)
	}

	time.Sleep(300 * time.Millisecond)
	for _, message := range messages {
		if message.ChatID == ChatID {
			messageResponse = append(messageResponse, message)
			if len(messageResponse) == messageLimit {
				break
			}
		}
	}

	return messageResponse, nil
}
