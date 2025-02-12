package models

type Chat struct {
	ID      int    `json:"ID"`
	Name    string `json:"Name"`
	LastMsg string `json:"LastMsg"`
}
