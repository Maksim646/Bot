package model

import (
	"database/sql"
)

type User struct {
	ID       sql.NullInt32  `db:"id"`
	userName sql.NullString `db:"user_name"`
	chatID   sql.NullString `db:"chat_id"`
	tgID     sql.NullString `db:"tg_id"`
}

type IUserRepository interface {
}

type IUserUsecase interface {
}
