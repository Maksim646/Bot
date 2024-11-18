package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	postgresql "github.com/Maksim646/Bot/database/migrations/sql"
	"github.com/Maksim646/Bot/model"
	sq "github.com/Masterminds/squirrel"

	//sq "github.com/Masterminds/squirrel"
	"github.com/heetch/sqalx"
)

type UserRepository struct {
	sqalxConn sqalx.Node
}

func New(sqalxConn sqalx.Node) model.IUserRepository {
	return &UserRepository{sqalxConn: sqalxConn}
}

func (r *UserRepository) CreateUserByTg(ctx context.Context, userName string, chatID int64, userLogin string, userPassword string) error {
	query, params, err := postgresql.Builder.Insert("users").
		Columns(
			"user_name",
			"chatid",
			"user_login",
			"user_password",
		).
		Values(
			userName,
			chatID,
			userLogin,
			userPassword,
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

func (r *UserRepository) GetUserByTgID(ctx context.Context, userID int64) (model.User, error) {
	var user model.User
	query, params, err := postgresql.Builder.Select(
		"users.id",
		"users.user_name",
		"users.chatid",
		"users.user_Login",
		"users.user_Password",
	).
		From("users").
		Where(sq.Eq{"users.id": userID}).
		ToSql()
	if err != nil {
		return user, err
	}
	//zap.L().Debug(postgresql.BuildQuery(query, params))

	if err = r.sqalxConn.GetContext(ctx, &user, query, params...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, err
		}
		return user, err
	}

	return user, nil
}
