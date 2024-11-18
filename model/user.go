package model

import (
	"context"
	"database/sql"
)

type User struct {
	ID           sql.NullInt32  `db:"id"`
	UserName     sql.NullString `db:"user_name"`
	ChatID       sql.NullInt64  `db:"chat_id"`
	UserLogin    sql.NullString `db:"user_login"`
	UserPassword sql.NullString `db:"user_password"`
}

type IUserRepository interface {
	CreateUserByTg(ctx context.Context, userName string, chatID int64) error
	GetUserByTgID(ctx context.Context, chatID int64) (User, error)
	UpdateUser(ctx context.Context, userLogin string, userPassword string, userID int64) error
}

type IUserUsecase interface {
	CreateUserByTg(ctx context.Context, userName string, chatID int64) error
	GetUserByTgID(ctx context.Context, chatID int64) (User, error)
	UpdateUser(ctx context.Context, userLogin string, userPassword string, userID int64) error
}
