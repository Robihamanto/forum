package models

import (
	"time"
)

//Feedback struct
type Feedback struct {
	ID        uint64     `gorm:"primary_key;auto_increment" json:"id"`
	UserID    uint32     `gorm:"not null" json:"user_id"`
	User      User       `json:"user"`
	PostID    uint64     `gorm:"not null" json:"post_id"`
	Post      Post       `json:"post"`
	Comment   string     `gorm:"type:varchar(255)" json:"comment"`
	CreatedAt time.Time  `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:current_timestamp()" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// NewFeedback create new feedback or comment
func NewFeedback(feedback Feedback) error {
	db := Connect()
	defer db.Close()
	return db.Create(&feedback).Error
}

// GetFeedBacks is retrieving all feedback
func GetFeedBacks() []Feedback {
	db := Connect()
	defer db.Close()
	var feedbacks []Feedback
	db.Order("id ASC").Find(&feedbacks)
	return feedbacks
}

// GetFeedbacksByPost is retrieving all feedback
func GetFeedbacksByPost(post Post) []Feedback {
	db := Connect()
	defer db.Close()
	var feedbacks []Feedback
	db.Where("post_id = ?", post.ID).Find(&feedbacks)
	for i := range feedbacks {
		db.Model(&feedbacks[i]).Related(&feedbacks[i].User)
		db.Model(&feedbacks[i]).Related(&post)
	}
	return feedbacks
}

// GetFeedBack is retrieving feedback by id
func GetFeedBack(id uint64) interface{} {
	db := Connect()
	defer db.Close()
	feedback := db.Where("id = ?", id).Find(&Feedback{}).Value
	return feedback
}

// UpdateFeedback create new feedback or comment
func UpdateFeedback(feedback Feedback) (int64, error) {
	db := Connect()
	defer db.Close()
	rs := db.Model(&feedback).Where("id = ?", feedback.ID).UpdateColumns(
		map[string]interface{}{
			"comment": feedback.Comment,
		},
	)
	return rs.RowsAffected, rs.Error
}

// DeleteFeedback is deleting feedback with id
func DeleteFeedback(feedback Feedback) error {
	db := Connect()
	defer db.Close()
	if rs := db.Where("id = ?", feedback.ID).Delete(&Feedback{}); rs.Error != nil {
		return rs.Error
	}
	return nil
}
