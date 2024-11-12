package main

import (
	_ "fmt"
	"log"
	"time"

	_bot "github.com/Maksim646/Bot/bot"
	"github.com/Maksim646/Bot/database"
	_ "github.com/Maksim646/Bot/domain/user/repository/postgresql"
	_userRepo "github.com/Maksim646/Bot/domain/user/repository/postgresql"
	_userUsecase "github.com/Maksim646/Bot/domain/user/usecase"
	"github.com/heetch/sqalx"
	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
)

var config struct {
	TgBotSecretKey string `envconfig:"TGBOT_SECRET_KEY" default:"7655110388:AAGk_q4QlcIccS1MA4vHKM5FvFiHSnUbRVg"`
	PostgresURI    string `envconfig:"POSTGRES_URI" default:"postgres://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable"`
	MigrationsDir  string `envconfig:"MIGRATIONS_DIR" default:"C:/Users/Максим/Downloads/Bot/Bot/database/migrations"`
}

func main() {

	envconfig.MustProcess("", &config)
	//fmt.Println("config:", config.PostgresURI)
	time.Sleep(2 * time.Second)
	migrator := database.NewMigrator(config.PostgresURI, config.MigrationsDir)
	if err := migrator.Apply(); err != nil {
		log.Fatal("cannot apply migrations: ", err)
	}

	sqlxConn, err := sqlx.Connect("postgres", config.PostgresURI)
	if err != nil {
		log.Fatal("cannot connect to postgres db: ", err)
	}

	sqlxConn.SetMaxOpenConns(100)
	sqlxConn.SetMaxIdleConns(100)
	sqlxConn.SetConnMaxLifetime(5 * time.Minute)

	defer sqlxConn.Close()

	sqalxConn, err := sqalx.New(sqlxConn)
	if err != nil {
		log.Fatal("cannot connect to postgres db: ", err)
	}
	defer sqalxConn.Close()

	log.Default().Println("db was ZBS")

	userRepo := _userRepo.New(sqalxConn)
	userUsecase := _userUsecase.New(userRepo)
	_, err = _bot.New(config.TgBotSecretKey, userRepo, userUsecase)

}
