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
type Data struct {
	Phone   int    `gorm:"column:phone;not null" json:"phone"`
	Address string `gorm:"column:address;not null" json:"address"`
	Comment string `gorm:"column:comment;not null" json:"comment"`
}
