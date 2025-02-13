package connectors

import (
	"encoding/json"
	"fmt"

	"github.com/joaogabsoaresf/wpp-cli-bot/internal/api"
	"github.com/joaogabsoaresf/wpp-cli-bot/internal/config"
)

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

func NewZAPIClient(baseURL, token string) *ZAPIClient {
	return &ZAPIClient{
		Client: api.NewClient(baseURL),
		Token:  token,
	}
}

func (z *ZAPIClient) GetTextMetaData(phoneNumber string) (*ChatMetadata, error) {
	endpoint := "/chats/" + phoneNumber

	headers := map[string]string{
		"Client-Token": z.Token,
		"Content-Type": "application/json",
	}

	response, err := z.Client.Get(endpoint, headers)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar metadados da mensagem %w", err)
	}

	var result ChatMetadata
	if err := json.Unmarshal(response, &result); err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("erro ao decodificar resposta JSON: %w", err)
	}

	return &result, nil
}

func GetMetaDataFromDefault() {
	zapiClient := NewZAPIClient(config.GetZAPIBaseURL(), config.GetZAPIToken())

	response, err := zapiClient.GetTextMetaData(config.GetZAPIDefaultNumber())
	if err != nil {
		fmt.Printf("erro ao obter metadados: %v", err)
		return
	}

	fmt.Printf("Metadados do Chat:\n")
	fmt.Printf("Name: %s\n", response.Name)
	fmt.Printf("Phone: %s\n", response.Phone)
}
