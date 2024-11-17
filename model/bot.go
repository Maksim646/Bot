package model

import "errors"

var (
	ErrUnknownBotCommand = errors.New("unknown bot command")
	ErrUserNotFound      = errors.New("user not found")
)

const (
	START_CMD = "/start"

	ParseModeHTML     = "html"
	HelloTgBotMessage = "Скорее начинай учиться!"
)
