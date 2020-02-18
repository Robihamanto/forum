package handlers

import (
	"api/models"
	"api/utils"
	"encoding/json"
	"net/http"
)

// PostPost is creating new post with new data
func PostPost(w http.ResponseWriter, r *http.Request) {
	body := utils.BodyParser(r)
	var post models.Post
	err := json.Unmarshal(body, &post)
	if err != nil {
		utils.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	err = models.NewPost(post)
	if err != nil {
		utils.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	utils.ToJSON(w, "Post successfully added", http.StatusCreated)
}

// GetFeedbacks func to retrieve feedback data
func GetFeedbacks(w http.ResponseWriter, r *http.Request) {
	feedbacks := models.GetAll(models.FEEDBACK)
	utils.ToJSON(w, feedbacks, http.StatusCreated)
}
