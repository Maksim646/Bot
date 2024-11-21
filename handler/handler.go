package handler

import (
	"time"

	"github.com/Maksim646/Bot/bot"
	"github.com/Maksim646/Bot/model"
	"github.com/patrickmn/go-cache"
)

type Handler struct {
	cache                   *cache.Cache
	telegramBotSecretKey    string
	telegramInitDataExpired int
	userUsecase             model.IUserUsecase
	AlertUsecase			model.IAlertUsecase
	bot                     bot.Tgbot
}

func New(
	version string,
	userUsecase model.IUserUsecase,
	AlertUsecase model.IAlertUsecase,
	bot bot.Tgbot,
) *Handler {

	h := &Handler{
		cache:       cache.New(cache.NoExpiration, time.Minute),
		userUsecase: userUsecase,
		bot:         bot,
	}
	return h
}
