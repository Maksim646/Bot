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
	"github.com/heetch/sqalx"
)

type AlertRepository struct {
	sqalxConn sqalx.Node
}

func New(sqalxConn sqalx.Node) model.IAlertRepository {
	return &AlertRepository{sqalxConn: sqalxConn}
}

func (r *AlertRepository) CreateAlert(ctx context.Context, Teacher string, SubjectOfStudy string, DataAlert time.Time) error {
	query, params, err := postgresql.Builder.Insert("alerts").
		Columns(
			"chat_id",
			"teacher",
			"subject_of_study",
			"data_alert",
		).
		Values(
			alert.ChatID,
			alert.Teacher,
			alert.SubjectOfStudy,
			alert.DataAlert,
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

func (r *AlertRepository) GetAlertsByChatID(ctx context.Context, chatID int64) ([]model.Alert, error) {
	var alerts []model.Alert
	query, params, err := postgresql.Builder.Select(
		"chat_id",
		"teacher",
		"subject_of_study",
		"data_alert",
	).
		From("alerts").
		Where(sq.Eq{"chat_id": chatID}).
		ToSql()
	if err != nil {
		fmt.Println("error:", err)
		return alerts, err
	}

	if err = r.sqalxConn.SelectContext(ctx, &alerts, query, params...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []model.Alert{}, nil
		}
		return nil, fmt.Errorf("ошибка при получении алерта: %w", err)
	}

	return alerts, nil
}