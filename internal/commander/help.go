package commander

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) Help(message *tgbotapi.Message) error {
	text := "/help - help\n/containers_list - list of running containers\n"
	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	_, err := c.bot.Send(msg)

	return err
}
