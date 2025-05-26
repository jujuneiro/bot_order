package plenka_bot

import (
	"gopkg.in/yaml.v3"
	"gorm.io/datatypes"
	"gorm.io/gorm"
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
	gorm.Model
	Name  string         `gorm:"column:name;not null" json:"name"`
	Phone string         `gorm:"column:phone;not null" json:"phone"`
	Email string         `gorm:"column:email;not null" json:"email"`
	Cart  datatypes.JSON `gorm:"column:cart;type:jsonb;not null" json:"cart"`
}

type CartProduct struct {
	Product  string `json:"product"`
	Variants []struct {
		Color string `json:"color"`
		Count int    `json:"count"`
	} `json:"variants"`
}

type IncomingOrder struct {
	Cart  []CartProduct `json:"cart"`
	Name  string        `json:"name"`
	Phone string        `json:"phone"`
	Email string        `json:"email"`
}
