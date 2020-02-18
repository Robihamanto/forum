package handlers

import (
	"api/models"
	"api/utils"
	"encoding/json"
	"net/http"
)

// PostUser is creating new users with new data
func PostUser(w http.ResponseWriter, r *http.Request) {
	body := utils.BodyParser(r)
	var user models.User
	err := json.Unmarshal(body, &user)
	if err != nil {
		utils.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	err = models.NewUser(user)
	if err != nil {
		utils.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	utils.ToJSON(w, "User successfully added", http.StatusCreated)
}
