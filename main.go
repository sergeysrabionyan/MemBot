package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"telegramBot/domain"
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
	client := telegram.NewClient(Config.TelegramBotUrl + Config.TelegramToken)
	offset := 0

	for {
		updates, err := client.GetUpdates(offset)
		if err != nil {
			log.Println(err)
			continue
		}
		for _, update := range updates {
			offset = update.Id + 1
			// Todo подумать над механизмом распознавания команд, текущая реализация - мусор
			if strings.Contains(update.Message.Text, "/mem") {
				memName := strings.TrimSpace(strings.Trim(update.Message.Text, "/mem"))
				go func(memName string, update domain.Update) {
					url, err := findMemUrl(memName)
					if err != nil {
						fmt.Println(err)
						return
					}
					fmt.Println("Идет отправка изображения")
					err = client.SendImage(update.Message.Chat.ChatId, url)
					if err != nil {
						fmt.Println(err)
						return
					}
				}(memName, update)
			}
		}
		fmt.Println(updates)
	}
}

func findMemUrl(name string) (string, error) {
	imageUrl, err := parser.GetRandomImageUrl(name)
	count := 0
	for imageUrl == "" && count < 10 {
		imageUrl, err = parser.GetRandomImageUrl(name)
		count++
		if err != nil {
			fmt.Println(err)
		}
	}
	if err != nil {
		return "", err
	}
	if imageUrl == "" {
		return "", errors.New("не найдено изображение")
	}
	return imageUrl, nil
}
