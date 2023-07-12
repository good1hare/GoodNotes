package telegramHandler

import (
	"MateMind/internal/usecase"
	"MateMind/pkg/logger"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type telegramHandler struct {
	user usecase.User
	l    logger.Interface
}

func Handle(bot *telegram.BotAPI, update telegram.Update, log logger.Interface, user usecase.User) {
	log.Info("[%s] %s", update.Message.From.UserName, update.Message.Text)

	msg := telegram.NewMessage(update.Message.Chat.ID, update.Message.Text)
	msg.ReplyToMessageID = update.Message.MessageID

	_, err := bot.Send(msg)
	if err != nil {
		return
	}
}
