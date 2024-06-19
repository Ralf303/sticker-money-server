package database

import (
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib" // Импорт драйвера для PostgreSQL
	"github.com/jmoiron/sqlx"
)

type User struct {
	Id        int    `db:"id"`
	ChatId    string `db:"chatId"`
	Username  string `db:"username"`
	FirstName string `db:"firstName"`
	Balance   int    `db:"balance"`
	Demo      int    `db:"demo"`
	IsBan     bool   `db:"isBan"`
}

type Stake struct {
	Id      int `db:"id"`
	Jackpot int `db:"jackpot"`
	Bar     int `db:"bar"`
	Berries int `db:"berries"`
	Lemons  int `db:"lemons"`
	Odd     int `db:"odd"`
	Correct int `db:"correct"`
}

type Admin struct {
	Id       int    `db:"id"`
	Status   string `db:"status"`
	Password string `db:"password"`
}

func connect() (*sqlx.DB, error) {
	dbUrl := os.Getenv("DATABASE_URL")

	db, err := sqlx.Connect("pgx", dbUrl)
	if err != nil {
		return nil, err
	}
	fmt.Println("Database connected")
	return db, nil
}
