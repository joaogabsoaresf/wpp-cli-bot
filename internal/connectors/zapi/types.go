package zapi

import "github.com/joaogabsoaresf/wpp-cli-bot/internal/api"

type ZAPIClient struct {
	Client *api.Client
	Token  string
}

type ChatMetadata struct {
	MessagesUnread string                 `json:"messagesUnread"`
	Name           string                 `json:"name"`
	Phone          string                 `json:"phone"`
	ExtraFields    map[string]interface{} `json:",inline"`
}

type MessageResponse struct {
	MessageID string `json:"messageId"`
}

type ChatResponse struct {
	Pinned         string `json:"pinned"`
	MessagesUnread string `json:"messagesUnread"`
	Phone          string `json:"phone"`
	Name           string `json:"name"`
	IsGroup        bool   `json:"isGroup"`
}
