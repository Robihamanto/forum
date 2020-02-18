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
	DeletedAt  time.Time  `gorm:"default:current_timestamp()" json:"deleted_at"`
}

//NewPost create new post
func NewPost(post Post) error {
	db := Connect()
	defer db.Close()
	return db.Create(&post).Error
}
