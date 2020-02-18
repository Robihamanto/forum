package models

import (
	"time"
)

//Feedback struct
type Feedback struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	UserID    uint32    `gorm:"not null" json:"user_id"`
	Users     User      `json:"user"`
	PostID    uint64    `gorm:"not null" json:"post_id"`
	Posts     Post      `json:"post"`
	Comment   string    `gorm:"type:varchar(255)" json:"comment"`
	CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
	DeletedAt time.Time `gorm:"default:current_timestamp()" json:"deleted_at"`
}
