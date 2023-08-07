package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string `yaml:"env" env-default:"local"`
	HTTPServer `yaml:"http"`
	Storage    `yaml:"storage"`
	Telegram   `yaml:"telegram"`
	Notion     `yaml:"notion"`
}

type HTTPServer struct {
	Host string `yaml:"host" env-default:"localhost"`
	Port int    `yaml:"port" env-default:"8080"`
}

type Storage struct {
	MongoDB struct {
		Path string `yaml:"path" env-default:"mongodb://test:test@127.0.0.1:27017"`
	} `yaml:"mongodb"`
}

type Telegram struct {
	WebhookURL string `yaml:"webhook_url"`
	BotToken   string `yaml:"bot_token"`
}

type Notion struct {
	OAuth struct {
		ClientID     string `yaml:"client_id"`
		ClientSecret string `yaml:"client_secret"`
		RedirectURI  string `yaml:"redirect_uri"`
	} `yaml:"oauth"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if len(configPath) == 0 {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
