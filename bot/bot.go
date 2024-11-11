package bot

import (
	//tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/Maksim646/Bot/model"
)

type bot struct {
	apiKey   string
	userRepo model.UserRepository
}
