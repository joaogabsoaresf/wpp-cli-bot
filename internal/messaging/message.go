package messaging

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
