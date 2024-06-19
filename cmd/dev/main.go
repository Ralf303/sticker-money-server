package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/stickerMoneyAdmin/internal/api"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := api.Routes()

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
