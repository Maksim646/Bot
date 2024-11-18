package model

import (
	"context"
	"database/sql"
)

type User struct {
	ID          sql.NullInt32  `db:"id"`
	UserName    sql.NullString  `db:"user_name"`
	ChatID      sql.NullString  `db:"chatid"`
	IDKafedrs   sql.NullInt32   `db:"id_kaf"`
	UserLogin   sql.NullString   `db:"user_Login"`
	UserPassword sql.NullString  `db:"user_Password"`
}

type IUserRepository interface {
	CreateUserByTg(ctx context.Context, userName string, chatID int64, userLogin string, userPassword string) error
	GetUserByTgID(ctx context.Context, userID int64) (User , error)
}

type IUserUsecase interface {
	CreateUserByTg(ctx context.Context, userName string, chatID int64, userLogin string, userPassword string) error
	GetUserByTgID(ctx context.Context, userID int64) (User , error)
}