package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"telegramBot/internal/telegram"
	"telegramBot/pkg/config"
)

var Config config.Config

func main() {
	_, err := Config.FromYaml(config.YamlConfigPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	client := telegram.NewClient(Config.TelegramBotUrl + Config.TelegramToken)
	bot := telegram.InitBot(client)

	ctx, cancel := context.WithCancel(context.Background())

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func(c <-chan os.Signal) {
		log.Printf("system call:%+v", <-c)
		cancel()
	}(c)

	bot.Start(ctx)
}
