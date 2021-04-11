package app

import (
	"fmt"
	"telegramBot/internal/telegram"
	"telegramBot/pkg/config"
)

var Config config.Config

func Run(configPath string) {
	_, err := Config.FromYaml(configPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	client := telegram.NewClient(Config.TelegramBotUrl + Config.TelegramToken)
	bot := telegram.InitBot(client)

	bot.StartListen()
}
