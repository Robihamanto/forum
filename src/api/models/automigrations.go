package models

//AutoMigrations for create db
func AutoMigrations() {
	db := Connect()
	defer db.Close()
	db.Debug().DropTableIfExists(&User{}, &Post{}, &Feedback{})
	db.Debug().AutoMigrate(&User{}, &Post{}, &Feedback{})
	db.Debug().Model(&Post{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade")
	db.Debug().Model(&Feedback{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade")
	db.Debug().Model(&Feedback{}).AddForeignKey("post_id", "posts(id)", "cascade", "cascade")
}
