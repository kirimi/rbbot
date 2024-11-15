package commander

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Commander struct {
	bot *tgbotapi.BotAPI
}

func NewCommander(bot *tgbotapi.BotAPI) *Commander {
	return &Commander{bot: bot}
}

type Command func(message *tgbotapi.Message) error
