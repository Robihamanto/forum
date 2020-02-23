package handlers

import (
	"api/models"
	"api/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// PostFeedback is creating new post with new data
func PostFeedback(w http.ResponseWriter, r *http.Request) {
	body := utils.BodyParser(r)
	var feedback models.Feedback
	err := json.Unmarshal(body, &feedback)
	if err != nil {
		utils.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	err = models.NewFeedback(feedback)
	if err != nil {
		utils.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	utils.ToJSON(w, "Feedback successfully added", http.StatusCreated)
}

// GetFeedbacks func to retrieve feedback data
func GetFeedbacks(w http.ResponseWriter, r *http.Request) {
	feedbacks := models.GetFeedBacks()
	utils.ToJSON(w, feedbacks, http.StatusOK)
}

// GetFeedback func to retrieve feedback data
func GetFeedback(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	feedback := models.GetFeedBack(id)
	utils.ToJSON(w, feedback, http.StatusOK)
}

// GetFeedBacksByPost func to retrieve feedback data
func GetFeedBacksByPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	feedback := models.GetFeedBack(id)
	utils.ToJSON(w, feedback, http.StatusOK)
}

// DeleteFeedback func to felete feedback data
func DeleteFeedback(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	_, err := models.Delete(models.FEEDBACKS, id)
	if err != nil {
		utils.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// PutFeedback func to felete feedback data
func PutFeedback(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	body := utils.BodyParser(r)
	var feedback models.Feedback
	err := json.Unmarshal(body, &feedback)
	if err != nil {
		utils.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	feedback.ID = id
	rows, err := models.UpdateFeedback(feedback)
	if err != nil {
		utils.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	utils.ToJSON(w, rows, http.StatusOK)
}
