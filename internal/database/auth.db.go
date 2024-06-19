package database

import (
	"errors"
	"fmt"
	"log"
)

func Register(id string, password string) error {
	db, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var admin Admin
	query := `SELECT * FROM "Admins" WHERE "id" = $1`
	err = db.Get(&admin, query, id)
	if err != nil {
		fmt.Println("ошибка регистрации", err)
		return err
	}

	if admin.Password != "none" {
		fmt.Println("ошибка регистрации: пользователь уже зарегистрирован")
		return errors.New("пользователь уже зарегистрирован")
	}

	query = `UPDATE "Admins" SET "password" = $1 WHERE "id" = $2`
	_, err = db.Exec(query, password, id)
	if err != nil {
		fmt.Println("ошибка регистрации", err)
		return err
	}

	return nil
}

func GetPassword(login string) (string, error) {
	db, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var admin Admin
	query := `SELECT * FROM "Admins" WHERE "id" = $1`
	err = db.Get(&admin, query, login)
	if err != nil {
		fmt.Println("ошибка получения пароля", err)
		return "", err
	}

	return admin.Password, nil
}
