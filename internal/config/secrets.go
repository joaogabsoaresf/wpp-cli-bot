package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ZAPIConfigData struct {
	InstanceID    string `json:"instance_id"`
	InstanceToken string `json:"instance_token"`
	ClientToken   string `json:"client_token"`
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar o arquivo .env no init 2")
	}
}

// Z-API
func GetZAPIToken() string {
	return os.Getenv("Z_API_CLIENT_TOKEN")
}

func GetZAPIBaseURL() string {
	return os.Getenv("Z_API_BASE_URL")
}

func GetZAPIDefaultNumber() string {
	return os.Getenv("Z_API_DEFAULT_TEST_NUMBER")
}
