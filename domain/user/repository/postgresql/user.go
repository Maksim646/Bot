package postgresql

import (
	"github.com/Maksim646/Bot/model"
	//sq "github.com/Masterminds/squirrel"
	"github.com/heetch/sqalx"
)

type UserRepository struct {
	sqalxConn sqalx.Node
}

func New(sqalxConn sqalx.Node) model.IUserRepository {
	return &UserRepository{sqalxConn: sqalxConn}
}
