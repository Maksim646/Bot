package usecase

import (
	"context"

	"github.com/Maksim646/Bot/model"
)

type Usecase struct {
	userRepository model.IAlertRepository
}

func New(alertRepository model.IAlertRepository) model.IAlertUsecase {
	return &AlertUsecase{
		alertRepository: alertRepository,
	}
}

func (u *AlertUsecase) CreateAlert(ctx context.Context, chatID int64,teacher string, subjectOfStudy string, dataAlert time.Time) error {
	return u.alertRepository.CreateAlert(ctx, teacher, subjectOfStudy, dataAlert)
}
