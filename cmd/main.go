package main

import (
	"fmt"
	"log"
	"time"

	//_bot "github.com/Maksim646/Bot/bot"
	"github.com/Maksim646/Bot/database"
	_ "github.com/Maksim646/Bot/domain/user/repository/postgresql"

	//_userRepo "github.com/Maksim646/Bot/domain/user/repository/postgresql"
	//_userUsecase "github.com/Maksim646/Bot/domain/user/usecase"
	"github.com/heetch/sqalx"
	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
)

var config struct {
	TgBotSecretKey string `envconfig:"TGBOT_SECRET_KEY" default:"7655110388:AAGk_q4QlcIccS1MA4vHKM5FvFiHSnUbRVg"`
	PostgresURI    string `envconfig:"POSTGRES_URI" default:"postgres://postgres:postgres@localhost:5433/bot_db?sslmode=disable"`
	MigrationsDir  string `envconfig:"MIGRATIONS_DIR" default:"database/migrations"`
}

func main() {
	envconfig.MustProcess("", &config)
	time.Sleep(2 * time.Second)
	fmt.Println("db:", config.PostgresURI)

	sqlxConn, err := sqlx.Connect("postgres", config.PostgresURI)
	if err != nil {
		log.Fatal("cannot connect to postgres db: ", err)
	}

	migrator := database.NewMigrator(config.PostgresURI, config.MigrationsDir)
	if err := migrator.Apply(); err != nil {
		log.Fatal("cannot apply migrations: ", err, config.MigrationsDir)
	}

	defer sqlxConn.Close()

	log.Println("Successfully connected to the database")

	sqlxConn.SetMaxOpenConns(100)
	sqlxConn.SetMaxIdleConns(100)
	sqlxConn.SetConnMaxLifetime(5 * time.Minute)

	sqalxConn, err := sqalx.New(sqlxConn)
	if err != nil {
		log.Fatal("cannot create sqalx connection: ", err)
	}
	defer sqalxConn.Close()

	log.Println("db was ZBS")
	time.Sleep(2 * time.Minute)

}
