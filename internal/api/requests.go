package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewClient(baseUrl string) *Client {
	return &Client{
		BaseURL: baseUrl,
		HTTPClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func (c *Client) Get(endpoint string, headers map[string]string) ([]byte, error) {
	url := c.BaseURL + endpoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar requisicao GET: %w", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer requisicao GET: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler corpo da resposta GET: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erro na requisição: status %d, corpo: %s", resp.StatusCode, string(body))
	}

	return body, nil
}

func (c *Client) Post(endpoint string, headers map[string]string, body interface{}) ([]byte, error) {
	url := c.BaseURL + endpoint

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("erro ao serializar corpo da requisição: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("erro ao criar requisição POST: %w", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer requisição POST: %w", err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler corpo da resposta: %w", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("erro na requisição: status %d, corpo: %s", resp.StatusCode, string(responseBody))
	}

	return responseBody, nil
}

func (c *Client) Put(endpoint string, headers map[string]string, body interface{}) ([]byte, error) {
	url := c.BaseURL + endpoint

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("erro ao serializar corpo da requisição: %w", err)
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("erro ao criar requisição PUT: %w", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer requisição PUT: %w", err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler corpo da resposta: %w", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return nil, fmt.Errorf("erro na requisição: status %d, corpo: %s", resp.StatusCode, string(responseBody))
	}

	return responseBody, nil
}

func (c *Client) Delete(endpoint string, headers map[string]string) error {
	url := c.BaseURL + endpoint

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("erro ao criar requisição DELETE: %w", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("erro ao fazer requisição DELETE: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("erro ao ler corpo da resposta: %w", err)
		}
		return fmt.Errorf("erro na requisição: status %d, corpo: %s", resp.StatusCode, string(body))
	}

	return nil
}
