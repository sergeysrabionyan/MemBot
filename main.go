package main

import (
	"telegramBot/internal/app"
	"telegramBot/pkg/config"
)

func main() {
	app.Run(config.YamlConfigPath)
}
