package services

import (
	"encoding/json"
	"net/http"
	"time"

	"example.com/stickerMoneyAdmin/internal/database"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth/v5"
)

var tokenAuth *jwtauth.JWTAuth

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	type body struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	var creds body

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Unable to parse JSON(sever error)", http.StatusBadRequest)
		return
	}

	err = database.Register(creds.Login, creds.Password)
	if err != nil {
		http.Error(w, "юзер не админ", http.StatusInternalServerError)
		return
	}

	claims := jwt.MapClaims{"exp": time.Now().Add(time.Hour * 24 * 7).Unix()}
	_, tokenString, _ := tokenAuth.Encode(claims)
	w.Write([]byte(tokenString))
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	login := r.URL.Query().Get("login")
	password := r.URL.Query().Get("password")

	userPassword, err := database.GetPassword(login)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if password != userPassword {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	claims := jwt.MapClaims{"exp": time.Now().Add(time.Hour * 24 * 7).Unix()}
	_, tokenString, _ := tokenAuth.Encode(claims)
	w.Write([]byte(tokenString))
}
