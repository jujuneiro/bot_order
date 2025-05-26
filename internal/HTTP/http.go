package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	dsn := "host=localhost user=postgres password=7-TPr-tooN dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&plenka_bot.Order{})

	router := gin.Default()

	router.POST("/user/create-order", func(c *gin.Context) {
		var input plenka_bot.IncomingOrder
		err := c.Bind(&input)
		if err != nil {
			c.JSON(400, gin.H{"message": "неверные данные"})
			return
		}

		message := fmt.Sprintf("Имя: %s\nТелефон: %s\nEmail: %s\n", input.Name, input.Phone, input.Email)
		for _, item := range input.Cart {
			message += fmt.Sprintf("Продукт: %s\n", item.Product)
			for _, variant := range item.Variants {
				message += fmt.Sprintf("Цвет: %s, Количество: %d\n", variant.Color, variant.Count)
			}
		}
		b.SendMessageToChannel(ctx, message)

		cartBytes, _ := json.Marshal(input.Cart)
		order := plenka_bot.Order{
			Name:  input.Name,
			Phone: input.Phone,
			Email: input.Email,
			Cart:  datatypes.JSON(cartBytes),
		}
		if err := db.Create(&order).Error; err != nil {
			c.JSON(500, gin.H{"message": "ошибка сохранения заказа"})
			return
		}

		c.JSON(201, gin.H{"message": "успешно"})
	})

	router.Run("localhost:8081")
	go b.Start(ctx)
}

func main() {
	http()
}
