package telegramHandler

import telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (th *TelegramHandler) SecretNoteCommand() {
	msg := telegram.NewMessage(th.update.Message.Chat.ID, "")
	msg.ReplyToMessageID = th.update.Message.MessageID

	msg.Text = "Секретная функция еще в разработке👨‍💻 но этого никто не узнает🤫"
	_, err := th.bot.Send(msg)
	if err != nil {
		th.log.Error(err)
	}
}
