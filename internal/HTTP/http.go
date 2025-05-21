package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	bot2 "untitled/internal/bot"
	plenka_bot "untitled/template"
)

func http() {
	var config *plenka_bot.Config
	var err error
	if config, err = plenka_bot.ParseConfig("./template/config.yaml"); err != nil {
		panic(err)
	}
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	b, err := bot2.NewBot(config.BotToken)
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	router.POST("/user/create-order", func(c *gin.Context) {
		data := &plenka_bot.Order{}
		//bind автоматом ставит 400.
		err = c.Bind(data)
		if err != nil {
			c.JSON(400, gin.H{"message": "неверные данные"})
			return
		}

		message := fmt.Sprintf("Имя: %s\nТелефон: %s\nEmail: %s\n", data.Name, data.Phone, data.Email)
		for _, item := range data.Cart {
			message += fmt.Sprintf("Продукт: %s\n", item.Product)
			for _, variant := range item.Variants {
				message += fmt.Sprintf("Цвет: %s, Количество: %d\n", variant.Color, variant.Count)
			}
		}
		b.SendMessageToChannel(ctx, message)

	})
	router.Run("localhost:8081")
	go b.Start(ctx)

}
func main() {
	http()
}
