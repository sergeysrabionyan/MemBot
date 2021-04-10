package main

import (
	"fmt"
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

	bot.StartListen()
}
