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
	donation    string
	gitHub      string
}

func NewTelegramHandler(bot *telegram.BotAPI, update telegram.Update, log logger.Interface, userUseCase usecase.User, donation string, gitHub string) *TelegramHandler {
	return &TelegramHandler{
		bot:         bot,
		update:      update,
		log:         log,
		userUseCase: userUseCase,
		donation:    donation,
		gitHub:      gitHub,
	}
}

func (th *TelegramHandler) Handle() {

	th.log.Info("[%s] %s", th.update.Message.From.UserName, th.update.Message.Text)

	th.register()

	command := th.update.Message.Text

	//get secret command from redis
	secretCommand := "secret"

	switch command {
	case "/start":
		th.StartCommand()

	case "🗒️Быстрая заметка":
		th.QuickNoteCommand()

	case "📒Заметка":
		th.NoteCommand()

	case "🎗️Напоминание":
		th.ReminderCommand()

	case "❓Помощь":
		th.HelpCommand()

	case "⚙️Настройки":
		th.SettingsNoteCommand()

	case "☕Разработчику на кофе":
		th.DonationCommand()

	case secretCommand:
		th.SecretNoteCommand()

	default:
		th.DefaultAnswer()
	}
}

func (th *TelegramHandler) DefaultAnswer() {
	msg := telegram.NewMessage(th.update.Message.Chat.ID, "")
	msg.ReplyToMessageID = th.update.Message.MessageID

	var defaultKeyboard = telegram.NewReplyKeyboard(
		telegram.NewKeyboardButtonRow(
			telegram.NewKeyboardButton("🗒️Быстрая заметка"),
			telegram.NewKeyboardButton("📒Заметка"),
			telegram.NewKeyboardButton("🎗️Напоминание"),
		),
		telegram.NewKeyboardButtonRow(
			telegram.NewKeyboardButton("❓Помощь"),
			telegram.NewKeyboardButton("⚙️Настройки"),
			telegram.NewKeyboardButton("☕Разработчику на кофе"),
		),
		telegram.NewKeyboardButtonRow(
			telegram.NewKeyboardButton("🔐Секретная заметка"),
		),
		//Отправить описание серкетной заметки и предложить назначить секретное слово для создания секретной заметки
	)

	msg.ReplyMarkup = defaultKeyboard

	msg.Text = "Что нужно сделать?"
	_, err := th.bot.Send(msg)
	if err != nil {
		th.log.Error(err)
	}
}

func (th *TelegramHandler) register() {
	e := entity.User{
		UserName:     th.update.Message.Chat.UserName,
		ChatId:       th.update.Message.Chat.ID,
		FirstName:    th.update.Message.From.FirstName,
		LastName:     th.update.Message.From.LastName,
		LanguageCode: th.update.Message.From.LanguageCode,
	}

	_, err := th.userUseCase.CreateUser(e)
	if err != nil {
		th.log.Error(err)
	}
}
