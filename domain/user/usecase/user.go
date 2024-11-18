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

func (u *Usecase) CreateUserByTg(ctx context.Context, userName string, chatID int64) error {
	return u.userRepository.CreateUserByTg(ctx, userName, chatID)
}

func (u *Usecase) GetUserByTgID(ctx context.Context, chatID int64) (model.User, error) {
	return u.userRepository.GetUserByTgID(ctx, chatID)
}

func (u *Usecase) UpdateUser(ctx context.Context, userLogin string, userPassword string, userID int64) error {
	return u.userRepository.UpdateUser(ctx, userLogin, userPassword, userID)
}
