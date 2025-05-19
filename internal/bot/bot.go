package bot

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
)

type Bot struct {
	*bot.Bot
}

func NewBot(botToken string) (newBot Bot, err error) {
	b, err := bot.New(botToken)
	if err != nil {
		return
	}
	newBot.Bot = b
	return
}

func (b *Bot) SendMessageToChannel(ctx context.Context, phone int, address string, comment string) (err error) {
	message := fmt.Sprintf("Телефон: %d\nАдрес: %s\nКомментарий: %s", phone, address, comment)
	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: -1002642818193,
		Text:   message,
	})
	if err != nil {
		return
	}
	return
}
