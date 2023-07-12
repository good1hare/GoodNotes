package telegramHandler

import (
	"MateMind/internal/entity"
	"MateMind/internal/usecase"
	"MateMind/pkg/logger"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type telegramHandler struct {
	bot         *telegram.BotAPI
	update      telegram.Update
	log         logger.Interface
	userUseCase usecase.User
}

func Handle(bot *telegram.BotAPI, update telegram.Update, log logger.Interface, user usecase.User) {
	th := &telegramHandler{bot, update, log, user}

	log.Info("[%s] %s", update.Message.From.UserName, update.Message.Text)

	th.register()

	msg := telegram.NewMessage(update.Message.Chat.ID, update.Message.Text)
	msg.ReplyToMessageID = update.Message.MessageID

	_, err := bot.Send(msg)
	if err != nil {
		return
	}
}

func (th *telegramHandler) register() {
	e := entity.User{UserName: th.update.ChatMember.Chat.UserName, ChatId: th.update.ChatMember.Chat.ID}

	e, err := th.userUseCase.CreateUser(e)
	if err != nil {
		return
	}
}
