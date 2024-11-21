package bot

import (
	"context"
	"fmt"
	"io"
	"strings"

	"bytes"
	"encoding/json"
	"net/http"
	"net/url"

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

type alert struct {
	alertRepo model.IAlertRepository 
	AlertUsecase model.IAlertUsecase
}

const (
	tokenURL    = "https://sso.guap.ru:8443/realms/master/protocol/openid-connect/token"
	clientID    = "prosuai"
	redirectURI = "https://pro.guap.ru/oauth/callback"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
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
			fmt.Println("ща отправлю hello")
			b.StartCommandHandler(ctx, tgUpdate, nil)
			user, err := b.userUsecase.GetUserByTgID(ctx, tgUpdate.Message.From.ID)
			if err != nil {
				return model.ErrUserNotFound
			}
			return b.AskForCredentials(ctx, user.ChatID.Int64)
		}

		if strings.Contains(tgUpdate.Message.Text, ":") {
			credentials := strings.Split(tgUpdate.Message.Text, ":")
			if len(credentials) == 2 {
				login := credentials[0]
				password := credentials[1]
				b.SaveCredentials(ctx, tgUpdate, login, password)
				user, err := b.userUsecase.GetUserByTgID(ctx, tgUpdate.Message.From.ID)
				if err != nil {
					return model.ErrUserNotFound
				}
				token, err := b.AuthHandler(ctx, user.UserLogin.String, user.UserPassword.String)
				fmt.Println("token, но скорее всего ошибка", token, err)
			}
		}

		return model.ErrUnknownBotCommand
	}

	return model.ErrUnknownBotCommand
}

func (b *Tgbot) StartCommandHandler(ctx context.Context, update tgbotapi.Update, data []string) error {
	_, err := b.userUsecase.GetUserByTgID(ctx, update.Message.From.ID)
	if err != nil {
		fmt.Println("error:", err)
		err := b.userRepo.CreateUserByTg(ctx, update.Message.From.UserName, update.Message.Chat.ID)
		if err != nil {
			fmt.Println("ошибка create user")
			return model.ErrUserNotFound
		}
		fmt.Println("попал втрой раз")
		if err != nil {
			fmt.Println("ошибка hellomessage?")
			return err
		}
	}

	return b.HelloMessage(ctx, update.Message.Chat.ID)
}

func (b *Tgbot) HelloMessage(ctx context.Context, chatID int64) error {
	newMsg := tgbotapi.NewMessage(chatID, model.HelloTgBotMessage)
	newMsg.ParseMode = model.ParseModeHTML
	_, err := b.bot.Send(newMsg)

	return err
}

func (b *Tgbot) AskForCredentials(ctx context.Context, chatID int64) error {
	askMsg := model.AksLogin
	newMsg := tgbotapi.NewMessage(chatID, askMsg)
	newMsg.ParseMode = model.ParseModeHTML
	_, err := b.bot.Send(newMsg)

	return err
}

func (b *Tgbot) SaveCredentials(ctx context.Context, update tgbotapi.Update, login string, password string) error {
	fmt.Println("request save credentials")
	err := b.userUsecase.UpdateUser(ctx, login, password, update.Message.From.ID)
	if err != nil {
		fmt.Println("error update login", err)
		return err
	}

	successMsg := model.YourDataWasSaved
	newMsg := tgbotapi.NewMessage(update.Message.Chat.ID, successMsg)
	newMsg.ParseMode = model.ParseModeHTML
	_, err = b.bot.Send(newMsg)

	return err
}

func (b *Tgbot) AuthHandler(ctx context.Context, username string, password string) (*TokenResponse, error) {
	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("username", username)
	data.Set("password", password)
	data.Set("client_id", "prosuai")
	data.Set("redirect_uri", redirectURI)

	req, err := http.NewRequest("POST", tokenURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("ошибка при создании запроса: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении запроса: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("ошибка при получении токена: %s, ответ сервера: %s", resp.Status, string(body))
	}

	var tokenResponse TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return nil, fmt.Errorf("ошибка при декодировании ответа: %w", err)
	}

	return &tokenResponse, nil
}
