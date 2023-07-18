package telegramHandler

import telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (th *TelegramHandler) HelpCommand() {
	msg := telegram.NewMessage(th.update.Message.Chat.ID, "")
	msg.ReplyToMessageID = th.update.Message.MessageID

	msg.Text = "Чем могу вам помочь?"
	_, err := th.bot.Send(msg)
	if err != nil {
		return
	}
}
