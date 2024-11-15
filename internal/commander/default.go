package commander

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) Default(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "I don't know what is: "+message.Text)
	msg.ReplyToMessageID = message.MessageID
	_, err := c.bot.Send(msg)

	return err
}
