package usecase

import (
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
