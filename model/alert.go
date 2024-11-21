package model

import (
	"context"
	"database/sql"
)

type Alert struct {
	ChatID         sql.NullInt64  `db:"chat_id"`
	Teacher        sql.NullString  `db:"teacher"`
	SubjectOfStudy sql.NullString  `db:"subject_of_study"`
	DataAlert      sql.NullTime    `db:"data_alert"`
}

type IAlertRepository interface {
	CreateAlert(ctx context.context, alert Alert) error
}

type IAlertUsecase interface {
	CreateAlert(ctx context.context, alert Alert) error
}
