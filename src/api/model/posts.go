package model

import (
	"time"
)

//Post struct
type Posts struct {
	ID          uint32    `gorm:"primary_key;auto_increment" json:"id"`
	UserID      uint32    `json:"user_id"`
	Users       Users     `json:"user"`
	Desciptions string    `gorm:"type:varchar(255)" json:"description"`
	ImageURL    string    `gorm:"type:varchar(255)" json:"image_url"`
	Subtitle    string    `gorm:"type:varchar(100)" json:"subtitle"`
	CreatedAt   time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
	DeletedAt   time.Time `gorm:"default:current_timestamp()" json:"deleted_at"`
}
