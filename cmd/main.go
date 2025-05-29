package main

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"os/signal"
	plenka_bot "untitled"
	http "untitled/internal/HTTP"
	bot2 "untitled/internal/bot"
	"untitled/internal/store"
)

func main() {
	//BOT
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
	//DB (через Store)
	dsn := "host=localhost user=postgres password=root dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Moscow"
	stor, err := store.NewStore(dsn)
	if err != nil {
		panic(err)
	}

	stor.DB.AutoMigrate(&plenka_bot.Order{})

	//http
	srv := http.NewHttpServer(ctx, b, stor.DB)
	srv.Start()
	go b.Start(ctx)
}
