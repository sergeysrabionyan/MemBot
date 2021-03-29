package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

const YamlConfigPath = "configs/configs.yml"

type Config struct {
	TelegramToken  string `yaml:"telegram_token"`
	TelegramBotUrl string `yaml:"telegram_bot_url"`
}

func (c *Config) FromYaml(configPath string) (*Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(c); err != nil {
		return nil, err
	}
	return c, nil
}
