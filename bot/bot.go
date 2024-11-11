package bot

import (
	"github.com/Maksim646/Bot/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Tgbot struct {
	bot         *tgbotapi.BotAPI
	apiKey      string
	userRepo    model.IUserRepository
	userUsecase model.IUserUsecase
}

func New(apiKey string, userRepo model.IUserRepository, userUsecase model.IUserUsecase) (*Tgbot, error) {
	bot, err := tgbotapi.NewBotAPI(apiKey)
	return &Tgbot{
		bot:         bot,
		apiKey:      apiKey,
		userRepo:    userRepo,
		userUsecase: userUsecase,
	}, err
}
