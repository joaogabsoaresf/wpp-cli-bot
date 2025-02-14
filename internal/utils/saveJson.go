package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func SaveToJson(data interface{}, filePath string) error {
	jsonData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return fmt.Errorf("erro ao converter dados para JSON: %v", err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("erro ao criar arquivo: %v", err)
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		return fmt.Errorf("erro ao escreer no arquivo: %v", err)
	}

	fmt.Printf("Arquivo JSON salvo com sucesso em %s\n", filePath)
	return nil
}
