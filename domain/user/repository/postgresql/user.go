package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	postgresql "github.com/Maksim646/Bot/database/migrations/sql"
	"github.com/Maksim646/Bot/model"
	sq "github.com/Masterminds/squirrel"
	"go.uber.org/zap"

	//sq "github.com/Masterminds/squirrel"
	"github.com/heetch/sqalx"
)

type UserRepository struct {
	sqalxConn sqalx.Node
}

func New(sqalxConn sqalx.Node) model.IUserRepository {
	return &UserRepository{sqalxConn: sqalxConn}
}

func (r *UserRepository) CreateUserByTg(ctx context.Context, userName string, chatID int64) error {
	query, params, err := postgresql.Builder.Insert("users").
		Columns(
			"user_name",
			"chat_id",
		).
		Values(
			userName,
			chatID,
		).
		ToSql()
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = r.sqalxConn.ExecContext(ctx, query, params...)
	fmt.Println(postgresql.BuildQuery(query, params))
	return err
}

func (r *UserRepository) GetUserByTgID(ctx context.Context, chatID int64) (model.User, error) {
	var user model.User
	query, params, err := postgresql.Builder.Select(
		"users.id",
		"users.user_name",
		"users.chat_id",
	).
		From("users").
		Where(sq.Eq{"users.chat_id": chatID}).
		ToSql()
	if err != nil {
		fmt.Println("error:", err)
		return user, err
	}
	//zap.L().Debug(postgresql.BuildQuery(query, params))

	if err = r.sqalxConn.GetContext(ctx, &user, query, params...); err != nil {
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return model.User{}, fmt.Errorf("пользователь с ID %d не найден", chatID)
			}
			return model.User{}, fmt.Errorf("ошибка при получении пользователя: %w", err)
		}
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, userLogin string, userPassword string, chatID int64) error {
	builder := postgresql.Builder.Update("users")

	if userLogin != "" {
		builder = builder.Set("user_login", userLogin)
	}

	if userPassword != "" {
		builder = builder.Set("user_password", userPassword)
	}

	query, params, err := builder.
		Where(sq.Eq{"chat_id": chatID}).
		ToSql()
	if err != nil {
		return err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))
	_, err = r.sqalxConn.ExecContext(ctx, query, params...)
	return err
}
