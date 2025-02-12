package api

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Chat struct {
	ID      int    `json:"ID"`
	Name    string `json:"Name"`
	LastMsg string `json:"LastMsg"`
}

func SendReply(msgID, reply string) {
	fmt.Printf("Mock API: Replying to message ID %s with: %s\n", msgID, reply)
}

func ApiRecentChat(chatLimit int) ([]Chat, error) {
	if chatLimit == 0 {
		chatLimit = 5
	}

	file, err := os.Open("mock-responses/chats.json")
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir o arquivo JSON: %v", err)
	}

	defer file.Close()

	var chats []Chat
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
