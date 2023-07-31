package telegramHandler

import telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (th *TelegramHandler) SettingsNoteCommand() {
	msg := telegram.NewMessage(th.update.Message.Chat.ID, "")
	msg.ReplyToMessageID = th.update.Message.MessageID

	msg.Text = "Функция еще в разработке👨‍💻"
	_, err := th.bot.Send(msg)
	if err != nil {
		th.log.Error(err)
	}
}
