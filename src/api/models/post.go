package models

import (
	"time"
)

//Post struct
type Post struct {
	ID         uint64     `gorm:"primary_key;auto_increment" json:"id"`
	UserID     uint32     `gorm:"not null" json:"user_id"`
	User       User       `json:"user"`
	Desciption string     `gorm:"type:varchar(255)" json:"description"`
	ImageURL   string     `gorm:"type:varchar(255)" json:"image_url"`
	Subtitle   string     `gorm:"type:varchar(100)" json:"subtitle"`
	Feedbacks  []Feedback `gorm:"ForeignKey:PostID" json:"feedbacks"`
	CreatedAt  time.Time  `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt  time.Time  `gorm:"default:current_timestamp()" json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
}

//NewPost create new post
func NewPost(post Post) error {
	db := Connect()
	defer db.Close()
	return db.Create(&post).Error
}

//UpdatePost updating post
func UpdatePost(post Post) (int64, error) {
	db := Connect()
	defer db.Close()
	rs := db.Model(&post).Where("id = ? ", post.ID).UpdateColumns(
		map[string]interface{}{
			"desciption": post.Desciption,
			"image_url":  post.ImageURL,
			"subtitle":   post.Subtitle,
		},
	)
	return rs.RowsAffected, rs.Error
}

// GetPosts is retrieving all post bt users
func GetPosts() []Post {
	db := Connect()
	defer db.Close()
	var posts []Post
	db.Order("id ASC").Find(&posts)
	for i := range posts {
		db.Model(&posts[i]).Related(&posts[i].User)
		posts[i].Feedbacks = GetFeedbacksByPost(posts[i])
	}
	return posts
}

// GetPost get single post
func GetPost(id uint64) interface{} {
	db := Connect()
	defer db.Close()
	var post Post
	post.ID = id
	db.Where("id = ?", id).Find(&post)
	db.Model(&post).Related(&post.User)
	post.Feedbacks = GetFeedbacksByPost(post)
	return post
}
