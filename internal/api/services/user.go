package services

import (
	"encoding/json"
	"net/http"
	"strconv"

	"example.com/stickerMoneyAdmin/internal/database"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		http.Error(w, "Отсутствует параметр id(id показывает с какого id надо возвращать юзера)", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Неверный параметр id(Нужно целочисленное)", http.StatusBadRequest)
		return
	}

	users := database.GetAllUsers(id)

	usersJSON, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(usersJSON)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		http.Error(w, "Отсутствует параметр id(id в базе, НЕ CHAT_ID)", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Неверный параметр id(Нужно целочисленное)", http.StatusBadRequest)
		return
	}

	user, err := database.GetUser(id)
	if err != nil {
		http.Error(w, "Юзер не найден", http.StatusBadRequest)
		return
	}
	userJSON, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(userJSON)
}

func UpdateBalance(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var req struct {
		ID     int `json:"id"`
		Amount int `json:"amount"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Ошибка при декодировании JSON", http.StatusBadRequest)
		return
	}

	_, err = database.GetUser(req.ID)
	if err != nil {
		http.Error(w, "Юзер не найден", http.StatusBadRequest)
		return
	}

	err = database.UpdateUserBalance(req.ID, req.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func UpdateBan(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var req struct {
		ID     int  `json:"id"`
		Status bool `json:"status"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Ошибка при декодировании JSON", http.StatusBadRequest)
		return
	}

	_, err = database.GetUser(req.ID)
	if err != nil {
		http.Error(w, "Юзер не найден", http.StatusBadRequest)
		return
	}

	err = database.UpdateUserBan(req.ID, req.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func CountUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	count := database.CountUsers()
	countSTR := strconv.Itoa(count)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(countSTR))
}
