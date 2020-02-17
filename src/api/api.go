package api

import (
	"api/model"
)

// Run database model
func Run() {
	model.AutoMigrations()
}
