package model

import (
	"context"
	"database/sql"
)

type User struct {
	ID        sql.NullInt32  `db:"id"`
	userName  sql.NullString `db:"user_name"`
	chatID    sql.NullString `db:"chatid"`
	IDKafedrs sql.NullInt32  `db:"id_kaf"`
}

type IUserRepository interface {
	CreateUserByTg(ctx context.Context, userName string, chatID int64) error
	GetUserByTgID(ctx context.Context, userID int64) (User, error)
}

type IUserUsecase interface {
	CreateUserByTg(ctx context.Context, userName string, chatID int64) error
	GetUserByTgID(ctx context.Context, userID int64) (User, error)
}
