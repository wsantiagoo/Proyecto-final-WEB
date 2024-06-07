package models

type Car struct {
	ID           int    `json:"id"`
	Type         string `json:"type"`
	Transmission string `json:"transmission"`
	Fuel         string `json:"fuel"`
	Model        string `json:"model"`
	Brand        string `json:"brand"`
}
