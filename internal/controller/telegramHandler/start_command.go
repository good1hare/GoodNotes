package telegramHandler

import telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (th *TelegramHandler) StartCommand() {
	sticker := telegram.NewSticker(th.update.Message.Chat.ID, telegram.FileID("CAACAgIAAxkBAAKoG2TLsNnQLUCTwzMAAZn4KcKt7j3ftwACAQEAAladvQoivp8OuMLmNC8E"))
	_, err := th.bot.Send(sticker)
	if err != nil {
		th.log.Error(err)
	}

	msg := telegram.NewMessage(th.update.Message.Chat.ID, "")
	msg.Text = "Здравствуйте, буду рад вам помочь!"
	_, err = th.bot.Send(msg)
	if err != nil {
		th.log.Error(err)
	}
}
