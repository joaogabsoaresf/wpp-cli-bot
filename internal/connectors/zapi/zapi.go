package zapi

import (
	"encoding/json"
	"fmt"

	"github.com/joaogabsoaresf/wpp-cli-bot/internal/api"
	"github.com/joaogabsoaresf/wpp-cli-bot/internal/config"
)

func NewZAPIClient(baseURL, token string) *ZAPIClient {
	return &ZAPIClient{
		Client: api.NewClient(baseURL),
		Token:  token,
	}
}

func (z *ZAPIClient) GetHeaders() map[string]string {
	return map[string]string{
		"Client-Token": z.Token,
	}
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

func (z *ZAPIClient) GetTextMetaData(phoneNumber string) (*ChatMetadata, error) {
	endpoint := "/chats/" + phoneNumber

	headers := z.GetHeaders()

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

func (z *ZAPIClient) SendTextMessage(phone string, message string) error {
	// Validação de entrada
	if phone == "" {
		return fmt.Errorf("o número de telefone não pode estar vazio")
	}
	if message == "" {
		return fmt.Errorf("o texto da mensagem não pode estar vazio")
	}

	endpoint := "/send-text/"
	headers := z.GetHeaders()

	// Cria o corpo da requisição
	body := map[string]string{
		"phone":   phone,
		"message": message,
	}

	response, err := z.Client.Post(endpoint, headers, body)
	if err != nil {
		return fmt.Errorf("erro ao enviar mensagem: %w", err)
	}

	var apiResponse MessageResponse
	if err := json.Unmarshal(response, &apiResponse); err != nil {
		return fmt.Errorf("erro ao decodificar resposta da API: %w", err)
	}

	if apiResponse.MessageID == "" {
		fmt.Println("Erro ao enviar mensagem, tente novamente mais tarde.'")
	}

	return nil
}

func SendMsg(phone string, message string) {
	zapiClient := NewZAPIClient(config.GetZAPIBaseURL(), config.GetZAPIToken())

	err := zapiClient.SendTextMessage(phone, message)
	if err != nil {
		fmt.Printf("erro ao enviar mensagem: %v\n", err)
		return
	}
}

func (z *ZAPIClient) GetChats() ([]ChatResponse, error) {
	endpoint := "/chats?page=1&pageSize=20"

	headers := z.GetHeaders()

	response, err := z.Client.Get(endpoint, headers)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar metadados da mensagem %w", err)
	}

	var result []ChatResponse
	if err := json.Unmarshal(response, &result); err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("erro ao decodificar resposta JSON: %w", err)
	}

	return result, nil
}

func GetLastChats() ([]ChatResponse, error) {
	zapiClient := NewZAPIClient(config.GetZAPIBaseURL(), config.GetZAPIToken())

	response, err := zapiClient.GetChats()
	if err != nil {
		fmt.Printf("erro ao obter metadados: %v", err)
		return nil, err
	}

	return response, nil
}
