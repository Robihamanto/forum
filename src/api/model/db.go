package model

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	//USER user database name
	USER = "root"
	//PASS password database name
	PASS = "root"
	//HOST url
	HOST = "127.0.0.1"
	//PORT port of database
	PORT = 3306
	//DBNAME database name
	DBNAME = "gorm"
)

// Connect function for connecting do database
func Connect() *gorm.DB {
	URL := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	db, err := gorm.Open("mysql", URL)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return db
}
