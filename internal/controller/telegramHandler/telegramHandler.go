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

	if !th.update.Message.IsCommand() { // ignore any non-command Messages
		return
	}

	command := th.update.Message.Command()
	switch command {
	case "start":
		th.HelloAnswer()
	case "help":
		th.HelpAnswer()
	default:
		th.DefaultAnswer()
	}

	//msg := telegram.NewMessage(th.update.Message.Chat.ID, th.update.Message.Text)
	//msg.ReplyToMessageID = th.update.Message.MessageID
	//
	//_, err := th.bot.Send(msg)
	//if err != nil {
	//	return
	//}
}

func (th *TelegramHandler) register() {
	e := entity.User{UserName: th.update.Message.Chat.UserName, ChatId: th.update.Message.Chat.ID}

	_, err := th.userUseCase.CreateUser(e)
	if err != nil {
		return
	}
}

func (th *TelegramHandler) HelloAnswer() {
	msg := telegram.NewMessage(th.update.Message.Chat.ID, "")
	msg.ReplyToMessageID = th.update.Message.MessageID

	msg.Text = "Здравствуйте, буду рад вам помочь!"
	_, err := th.bot.Send(msg)
	if err != nil {
		return
	}
}

func (th *TelegramHandler) DefaultAnswer() {
	msg := telegram.NewMessage(th.update.Message.Chat.ID, "")
	msg.ReplyToMessageID = th.update.Message.MessageID

	var defaultKeyboard = telegram.NewInlineKeyboardMarkup(
		telegram.NewInlineKeyboardRow(
			//telegram.NewInlineKeyboardButtonURL("1.com", "http://1.com"),
			//telegram.NewInlineKeyboardButtonData("2", "2"),
			//telegram.NewInlineKeyboardButtonData("3", "3"),
			telegram.InlineKeyboardButton{Text: "Создать быструю заметку"},
		),
		//telegram.NewInlineKeyboardRow(
		//	telegram.NewInlineKeyboardButtonData("4", "4"),
		//	telegram.NewInlineKeyboardButtonData("5", "5"),
		//	telegram.NewInlineKeyboardButtonData("6", "6"),
		//),
	)

	msg.ReplyMarkup = defaultKeyboard

	msg.Text = "Что нужно сделать?"
	_, err := th.bot.Send(msg)
	if err != nil {
		return
	}
}

func (th *TelegramHandler) HelpAnswer() {
	msg := telegram.NewMessage(th.update.Message.Chat.ID, "")
	msg.ReplyToMessageID = th.update.Message.MessageID

	msg.Text = "Чем могу вам помочь?"
	_, err := th.bot.Send(msg)
	if err != nil {
		return
	}
}
