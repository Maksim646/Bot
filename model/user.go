package model

import (
	"database/sql"
)

type userRepository struct {
	userName sql.NullString `db:"user_name"`
	chatID sql.NullString `db:"chat_id"`
	tgID sql.NullString `db:"tg_id"`
}
