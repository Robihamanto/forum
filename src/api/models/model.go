package models

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
	db.Connect()
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
