package services

import (
	"encoding/json"
	"net/http"

	"example.com/stickerMoneyAdmin/internal/database"
)

func GetStakes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	stakes, err := database.GetStakes()
	if err != nil {
		http.Error(w, "ставки не найдены ", http.StatusBadRequest)
		return
	}

	stakesJSON, err := json.Marshal(stakes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(stakesJSON)

}

func UpdateStake(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var req struct {
		Raw   string `json:"raw"`
		Value int    `json:"value"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Ошибка при декодировании JSON", http.StatusBadRequest)
		return
	}

	err = database.ChangeStakes(req.Raw, req.Value)
	if err != nil {
		http.Error(w, "Ставка не изменилась/не найдена", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
