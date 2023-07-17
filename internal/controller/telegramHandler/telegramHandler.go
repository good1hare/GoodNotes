package telegramHandler

import (
	"MateMind/internal/entity"
	"MateMind/internal/usecase"
	"MateMind/pkg/logger"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramHandler struct {
	bot         *telegram.BotAPI
	update      telegram.Update
	log         logger.Interface
	userUseCase usecase.User
}

func NewTelegramHandler(bot *telegram.BotAPI, update telegram.Update, log logger.Interface, userUseCase usecase.User) *TelegramHandler {
	return &TelegramHandler{
		bot:         bot,
		update:      update,
		log:         log,
		userUseCase: userUseCase,
	}
}

func (th *TelegramHandler) Handle() {

	th.log.Info("[%s] %s", th.update.Message.From.UserName, th.update.Message.Text)

	th.register()

	msg := telegram.NewMessage(th.update.Message.Chat.ID, th.update.Message.Text)
	msg.ReplyToMessageID = th.update.Message.MessageID

	_, err := th.bot.Send(msg)
	if err != nil {
		return
	}
}

func (th *TelegramHandler) register() {
	e := entity.User{UserName: th.update.Message.Chat.UserName, ChatId: th.update.Message.Chat.ID}

	_, err := th.userUseCase.CreateUser(e)
	if err != nil {
		return
	}
}
