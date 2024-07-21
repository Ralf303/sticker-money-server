package database

import (
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func GetAllUsers(startId int, filterBy string) []User {
	db, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	users := []User{}
	query := ""

	switch filterBy {
	case "no":
		query = `SELECT * FROM "Users" WHERE "id" >= $1 ORDER BY "id" ASC LIMIT 500`
	case "up":
		query = `SELECT * FROM "Users" ORDER BY "balance" DESC LIMIT 500 OFFSET $1`
	case "down":
		query = `SELECT * FROM "Users" ORDER BY "balance" ASC LIMIT 500 OFFSET $1`
	}

	err = db.Select(&users, query, startId)
	if err != nil {
		fmt.Println("ошибка GetAllUsers", err)
	}
	return users
}

func GetUser(id string) (User, error) {
	db, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var user User
	query := `SELECT * FROM "Users" WHERE "chatId" = $1`
	err = db.Get(&user, query, id)
	if err != nil {
		fmt.Println("ошибка GetUser", err)
		return User{}, err
	}
	return user, nil
}

func UpdateUserBalance(id string, value int) error {
	db, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	query := `UPDATE "Users" SET "balance"=$1 WHERE "chatId"=$2;`
	_, err = db.Exec(query, value, id)
	if err != nil {
		fmt.Println("Ошибка при обновлении балансa:", err)
		return err
	}
	return nil
}

func UpdateUserBan(id string, value bool) error {
	db, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	query := `UPDATE "Users" SET "isBan"=$1 WHERE "chatId"=$2;`
	_, err = db.Exec(query, value, id)
	if err != nil {
		fmt.Println("Ошибка при обновлении бана:", err)
		return err
	}
	return nil
}

func CountUsers() int {
	db, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var count int
	query := `SELECT COUNT(*) FROM "Users"`
	err = db.Get(&count, query)
	if err != nil {
		fmt.Println("ошибка CounUsers", err)
	}
	return count
}
