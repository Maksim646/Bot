package bot

import (
	"context"
	"fmt"
	"strings"
	_ "time"

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

		commandData := strings.Split(tgUpdate.Message.Text, " ")
		if commandData[0] != model.START_CMD {
			return model.ErrUnknownBotCommand
		}

		return b.StartCommandHandler(ctx, tgUpdate, commandData)

	}

	return model.ErrUnknownBotCommand
}

func (b *Tgbot) StartCommandHandler(ctx context.Context, update tgbotapi.Update, data []string) error {
	_, err := b.userUsecase.GetUserByTgID(ctx, update.Message.From.ID)
	if err != nil {
		fmt.Println("1")

		err := b.userRepo.CreateUserByTg(ctx, update.Message.From.UserName, update.Message.Chat.ID)
		fmt.Println("2")
		if err != nil {
			return model.ErrUserNotFound
		}
		//zap.L().Info(fmt.Sprintf("create user from telegram with ID=[%s]", userID))
		fmt.Println("3")
		return b.HelloMessage(ctx, update.Message.Chat.ID)
	}

	// if user.Status.String != model.UserActive {
	// 	return model.ErrUserPermissionDenied
	// }

	return b.HelloMessage(ctx, update.Message.Chat.ID)
}

func (b *Tgbot) HelloMessage(ctx context.Context, chatID int64) error {
	newMsg := tgbotapi.NewMessage(chatID, model.HelloTgBotMessage)
	newMsg.ParseMode = model.ParseModeHTML
	_, err := b.bot.Send(newMsg)

	return err
}
