package bot

import (
	"context"
	"fmt"
	"strings"

	"github.com/Maksim646/Bot/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

type Tgbot struct {
	bot         *tgbotapi.BotAPI
	apiKey      string
	userRepo    model.IUserRepository
	userUsecase model.IUserUsecase
}

func New(apiKey string, userRepo model.IUserRepository, userUsecase model.IUserUsecase) (*Tgbot, error) {
	bot, err := tgbotapi.NewBotAPI(apiKey)
	fmt.Println("инит")
	return &Tgbot{
		bot:         bot,
		apiKey:      apiKey,
		userRepo:    userRepo,
		userUsecase: userUsecase,
	}, err
}

func (b *Tgbot) ListenUpdates() {
	u := tgbotapi.NewUpdate(0)
	u.AllowedUpdates = []string{"messages"}
	u.Timeout = 60
	updates := b.bot.GetUpdatesChan(u)
	zap.L().Info("Start listening for tg messages")

	for update := range updates {
		if err := b.ProcessMessage(update); err != nil {
			zap.L().Error("error process tg messages", zap.Error(err))
		}
	}
}

func (b *Tgbot) ProcessMessage(tgUpdate tgbotapi.Update) error {
	ctx := context.Background()

	if tgUpdate.Message != nil {
		zap.L().Info(fmt.Sprintf("[%s %s %s][%v] --> %s",
			tgUpdate.Message.From.UserName,
			tgUpdate.Message.From.ID))

		if tgUpdate.Message.Text == model.START_CMD {
			return b.StartCommandHandler(ctx, tgUpdate, nil)
		}

		if strings.Contains(tgUpdate.Message.Text, ":") {
			credentials := strings.Split(tgUpdate.Message.Text, ":")
			if len(credentials) == 2 {
				login := credentials[0]
				password := credentials[1]
				return b.SaveCredentials(ctx, tgUpdate, login, password)
			}
		}

		return model.ErrUnknownBotCommand
	}

	return model.ErrUnknownBotCommand
}

func (b *Tgbot) StartCommandHandler(ctx context.Context, update tgbotapi.Update, data []string) error {
	_, err := b.userUsecase.GetUserByTgID(ctx, update.Message.From.ID)
	if err != nil {
		login := "defaultLogin"
		password := "defaultPass"
		err := b.userRepo.CreateUserByTg(ctx, update.Message.From.UserName, update.Message.Chat.ID, login, password)
		if err != nil {
			return model.ErrUserNotFound
		}
		if err := b.HelloMessage(ctx, update.Message.Chat.ID); err != nil {
			return err
		}
		return b.AskForCredentials(ctx, update.Message.Chat.ID)
	}

	return b.AskForCredentials(ctx, update.Message.Chat.ID)
}

func (b *Tgbot) HelloMessage(ctx context.Context, chatID int64) error {
	newMsg := tgbotapi.NewMessage(chatID, model.HelloTgBotMessage)
	newMsg.ParseMode = model.ParseModeHTML
	_, err := b.bot.Send(newMsg)

	return err
}

func (b *Tgbot) AskForCredentials(ctx context.Context, chatID int64) error {
	askMsg := "Введите логин и пароль в виде 'логин:пароль'"
	newMsg := tgbotapi.NewMessage(chatID, askMsg)
	newMsg.ParseMode = model.ParseModeHTML
	_, err := b.bot.Send(newMsg)

	return err
}

func (b *Tgbot) SaveCredentials(ctx context.Context, update tgbotapi.Update, login string, password string) error {
	err := b.userRepo.CreateUserByTg(ctx, update.Message.From.UserName, update.Message.Chat.ID, login, password)
	if err != nil {
		return err
	}

	successMsg := "Ваши логин и пароль успешно сохранены!"
	newMsg := tgbotapi.NewMessage(update.Message.Chat.ID, successMsg)
	newMsg.ParseMode = model.ParseModeHTML
	_, err = b.bot.Send(newMsg)

	return err
}