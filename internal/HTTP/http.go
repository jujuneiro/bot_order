package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	plenka_bot "untitled"
	"untitled/internal/bot"
)

type HttpServer struct {
	Router *gin.Engine
	B      *bot.Bot
	DB     *gorm.DB
	Ctx    context.Context
}

func NewHttpServer(ctx context.Context, b *bot.Bot, db *gorm.DB) (srv *HttpServer) {
	srv = &HttpServer{
		Router: gin.Default(),
		B:      b,
		DB:     db,
		Ctx:    ctx,
	}
	return
}

func (h *HttpServer) Start() {
	h.Router.POST("/user/create-order", func(c *gin.Context) {
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
		h.B.SendMessageToChannel(h.Ctx, message)

		cartBytes, _ := json.Marshal(input.Cart)
		order := plenka_bot.Order{
			Name:  input.Name,
			Phone: input.Phone,
			Email: input.Email,
			Cart:  datatypes.JSON(cartBytes),
		}
		if err := h.DB.Create(&order).Error; err != nil {
			c.JSON(500, gin.H{"message": "ошибка сохранения заказа"})
			return
		}

		c.JSON(201, gin.H{"message": "успешно"})
	})

	h.Router.Run("localhost:8081")
}
