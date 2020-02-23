package models

import "github.com/jinzhu/gorm"

const (
	//USERS user database name
	USERS = "users"
	//POSTS post database name
	POSTS = "posts"
	//FEEDBACKS feedback database name
	FEEDBACKS = "feedbacks"
)

// GetAll users function
func GetAll(table string) interface{} {
	db := Connect()
	defer db.Close()
	switch table {
	case USERS:
		return db.Find(&[]User{}).Value
	case POSTS:
		return db.Find(&[]Post{}).Value
	case FEEDBACKS:
		return db.Find(&[]Feedback{}).Value
	}
	return nil
}

// GetByID looking for row with ID
func GetByID(table string, id uint64) interface{} {
	db := Connect()
	defer db.Close()
	switch table {
	case USERS:
		return db.Where("id = ?", id).Find(&User{}).Value
	case POSTS:
		// var post Post
		return db.Where("id = ?", id).Find(&Post{}).Value
		// db.Model(&post).Related(&post.User)
		// return post
	case FEEDBACKS:
		// var feedback Feedback
		return db.Where("id = ?", id).Find(&Feedback{}).Value
		// db.Model(&feedback).Related(&feedback.Users)
		// return feedback
	}
	return nil
}

// Delete is delete data from databases
func Delete(table string, id uint64) (int64, error) {
	db := Connect()
	defer db.Close()
	var rs *gorm.DB
	switch table {
	case USERS:
		rs = db.Where("id = ?", id).Delete(&User{})
	case POSTS:
		rs = db.Where("id = ?", id).Delete(&Post{})
	case FEEDBACKS:
		rs = db.Where("id = ?", id).Delete(&Feedback{})
	}
	return rs.RowsAffected, rs.Error
}
