package usecase

import (
	"context"

	"github.com/Maksim646/Bot/model"
)

type Usecase struct {
	userRepository model.IUserRepository
}

func New(userRepository model.IUserRepository) model.IUserUsecase {
	return &Usecase{
		userRepository: userRepository,
	}
}

func (u *Usecase) CreateUserByTg(ctx context.Context, userName string, chatID int64, userLogin string, userPassword string) error {
	return u.userRepository.CreateUserByTg(ctx, userName, chatID, userLogin, userPassword)
}

func (u *Usecase) GetUserByTgID(ctx context.Context, userID int64) (model.User, error) {
	return u.userRepository.GetUserByTgID(ctx, userID)
}
