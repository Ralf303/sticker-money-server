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

	certPath := "./fullchain.pem"
	keyPath := "./privkey.pem"

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServeTLS(":8080", certPath, keyPath, router))
}
