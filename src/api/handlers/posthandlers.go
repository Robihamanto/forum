package handlers

import (
	"api/models"
	"api/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

// GetPosts func to retrieve post data
func GetPosts(w http.ResponseWriter, r *http.Request) {
	posts := models.GetPosts()
	utils.ToJSON(w, posts, http.StatusOK)
}

// GetPost retrieve post details
func GetPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	//post := models.GetByID(models.POSTS, id)
	post := models.GetPost(id)
	if post == nil {
		utils.ToJSON(w, utils.JSONMessageWriter("No data with that ID"), http.StatusNotFound)
		return
	}
	utils.ToJSON(w, post, http.StatusOK)
}

// PutPost func to update post data
func PutPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	body := utils.BodyParser(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	var post models.Post
	err := json.Unmarshal(body, &post)
	if err != nil {
		utils.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	post.ID = uint64(id)
	_, err = models.UpdatePost(post)
	if err != nil {
		utils.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	utils.ToJSON(w, utils.JSONMessageWriter("Post successfully updated!"), http.StatusOK)
}

// DeletePost delete post from db
func DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	_, err := models.Delete(models.POSTS, id)
	if err != nil {
		utils.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	utils.ToJSON(w, utils.JSONMessageWriter("Post deleted from database"), http.StatusOK)
}
