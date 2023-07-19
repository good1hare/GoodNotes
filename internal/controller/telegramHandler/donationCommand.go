package telegramHandler

import telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (th *TelegramHandler) DonationCommand() {
	msg := telegram.NewMessage(th.update.Message.Chat.ID, "Спасибо вам за вклад 💖")

	var numericKeyboard = telegram.NewInlineKeyboardMarkup(
		telegram.NewInlineKeyboardRow(
			telegram.NewInlineKeyboardButtonURL("Поддержать проект", th.donation),
		),
	)
	msg.ReplyMarkup = numericKeyboard

	_, err := th.bot.Send(msg)
	if err != nil {
		th.log.Error(err)
	}
}
