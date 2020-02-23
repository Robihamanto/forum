package models

import (
	"api/security"
	"time"
)

//User struct
type User struct {
	ID        uint32     `gorm:"primary_key;auto_increment" json:"id"`
	Username  string     `gorm:"type:varchar(20); not null; unique_index" json:"username"`
	Email     string     `gorm:"type:varchar(40); not null; unique_index" json:"email"`
	Password  string     `gorm:"type:varchar(60); not null;" json:"password"`
	Posts     []Post     `gorm:"ForeignKey:UserID" json:"posts"`
	Feedbacks []Feedback `gorm:"ForeignKey:UserID" json:"feedbacks"`
	CreatedAt time.Time  `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:current_timestamp()" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

//NewUser create new user
func NewUser(user User) error {
	db := Connect()
	defer db.Close()
	var err error
	user.Password, err = security.Hash(user.Password)
	if err != nil {
		return err
	}
	err = db.Create(&user).Error
	return err
}

// UpdateUser update user data
func UpdateUser(user User) (int64, error) {
	db := Connect()
	defer db.Close()
	rs := db.Model(&user).Where("id = ?", user.ID).UpdateColumns(
		map[string]interface{}{
			"username": user.Username,
			"email":    user.Email,
		},
	)
	return rs.RowsAffected, rs.Error
}
