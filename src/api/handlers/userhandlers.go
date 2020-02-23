package handlers

import (
	"api/models"
	"api/utils"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
		resp := map[string]string{
			"message": err.Error(),
		}
		utils.ToJSON(w, resp, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJSON(w, "User successfully added", http.StatusCreated)
}

// GetUsers func to retrieve users data
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := models.GetAll(models.USERS)
	utils.ToJSON(w, users, http.StatusOK)
}

// GetUser func to retrieve users data
func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	user := models.GetByID(models.USERS, id)
	utils.ToJSON(w, user, http.StatusOK)
}

// PutUser is editing user data
func PutUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	body := utils.BodyParser(r)
	var user models.User
	err := json.Unmarshal(body, &user)
	if err != nil {
		utils.ToJSON(w, err.Error, http.StatusUnprocessableEntity)
		return
	}
	user.ID = uint32(id)
	_, err = models.UpdateUser(user)
	if err != nil {
		utils.ToJSON(w, err.Error, http.StatusUnprocessableEntity)
		return
	}

	resp := map[string]interface{}{
		"message": "Success edit users",
		"user":    user,
	}

	utils.ToJSON(w, resp, http.StatusOK)
}

// DeleteUser is delete the user from database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	_, err := models.Delete(models.USERS, id)
	if err != nil {
		log.Println(err.Error())
		utils.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	resp := map[string]interface{}{
		"message": "Success delete user data",
	}
	utils.ToJSON(w, resp, http.StatusOK)
}
