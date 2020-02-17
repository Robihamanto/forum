package model

//AutoMigrations for create db
func AutoMigrations() {
	db := Connect()
	defer db.Close()
	db.Debug().DropTableIfExists(&Users{}, &Posts{})
	db.Debug().AutoMigrate(&Users{}, &Posts{})
	db.Debug().Model(&Posts{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade")
}
