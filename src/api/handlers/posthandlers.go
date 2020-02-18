package handlers

import (
	"api/models"
	"api/utils"
	"encoding/json"
	"net/http"
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

// GetPosts func to retrieve post data
func GetPosts(w http.ResponseWriter, r *http.Request) {
	posts := models.GetAll(models.POSTS)
	utils.ToJSON(w, posts, http.StatusCreated)
}
