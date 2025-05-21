package plenka_bot

import (
	"gopkg.in/yaml.v3"
	"os"
)

func ParseConfig(roadConfig string) (config *Config, err error) {
	var body []byte
	if body, err = os.ReadFile(roadConfig); err != nil {
		return
	}
	if err = yaml.Unmarshal(body, &config); err != nil {
		return
	}

	return
}

type Config struct {
	BotToken string `yaml:"bot_token"`
}

type Order struct {
	Cart []struct {
		Product  string `gorm:"column:product;not null" json:"product"`
		Variants []struct {
			Color string `gorm:"column:color;not null" json:"color"`
			Count int    `gorm:"column:count;not null" json:"count"`
		} `json:"variants"`
	} `json:"cart"`
	Name  string `json:"name"`
	Phone string `gorm:"column:phone;not null" json:"phone"`
	Email string `json:"email"`
}
