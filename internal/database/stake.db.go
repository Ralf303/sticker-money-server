package database

import (
	"fmt"
	"log"
)

func GetStakes() (Stake, error) {
	db, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var stake Stake
	query := `SELECT * FROM "Stakes" WHERE "id" = 1`
	err = db.Get(&stake, query)
	if err != nil {
		fmt.Println("ошибка GetStakes", err)
		return Stake{}, err
	}
	return stake, nil
}

func ChangeStakes(raw string, value int) error {
	db, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	query := fmt.Sprintf(`UPDATE "Stakes" SET "%s"=$1 WHERE "id"=1;`, raw)
	_, err = db.Exec(query, value)
	if err != nil {
		fmt.Println("Ошибка при обновлении ставки:", err)
		return err
	}
	return nil
}
