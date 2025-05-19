package main

import (
	"context"
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
		data := &plenka_bot.Data{}
		//bind автоматом ставит 400.
		err = c.Bind(&data)
		if err != nil {
			c.JSON(400, gin.H{"message": "неверные данные"})
			return
		}
		b.SendMessageToChannel(ctx, data.Phone, data.Address, data.Comment)
		if err != nil {
			return
		}
	})

	router.Run("localhost:8081")
	go b.Start(ctx)

}
func main() {
	http()
}
