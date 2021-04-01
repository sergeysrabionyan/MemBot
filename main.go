package main

import (
	"fmt"
	"log"
	"strings"
	"telegramBot/internal/telegram"
	"telegramBot/pkg/config"
	"telegramBot/pkg/parser"
)

var Config config.Config

func main() {
	_, err := Config.FromYaml(config.YamlConfigPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	botUrl := Config.TelegramBotUrl + Config.TelegramToken
	offset := 0
	for {
		updates, err := telegram.GetUpdates(botUrl, offset)
		if err != nil {
			log.Println(err)
		}
		for _, update := range updates {
			if strings.Contains(update.Message.Text, "/mem") {
				memName := strings.Trim(strings.Trim(update.Message.Text, "/mem"), " ")
				err = findAndSendMem(botUrl, memName, update.Message.Chat.ChatId)
			}
			if err != nil {
				log.Println(err)
			}
			offset = update.Id + 1
		}
		fmt.Println(updates)
	}
}

func findAndSendMem(botUrl string, name string, chatId int) error {
	imageUrl, err := parser.GetRandomImageUrl(name)
	if err != nil {
		fmt.Println(err)
	}
	if imageUrl == "" {
		return nil
	}
	fmt.Println("Идет отправка изображения")
	err = telegram.SendImage(chatId, imageUrl, botUrl)
	if err != nil {
		return err
	}
	return nil
}
