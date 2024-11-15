package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	rbBot "github.com/kirimi/rb_bot/internal/commander"
)

func main() {
	_ = godotenv.Load()
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}

	telegramToken := os.Getenv("RB_BOT_TOKEN")

	bot, err := tgbotapi.NewBotAPI(telegramToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	cmdr := rbBot.NewCommander(bot)

	u := tgbotapi.UpdateConfig{Timeout: 60}

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			var cmd rbBot.Command

			switch update.Message.Command() {
			case rbBot.HelpCmd:
				cmd = cmdr.Help
			case rbBot.ContainersListCmd:
				cmd = cmdr.ContainersList
			default:
				cmd = cmdr.Default
			}

			err := cmd(update.Message)
			if err != nil {
				log.Println("Sending reply error")
			}
		}
	}
}
