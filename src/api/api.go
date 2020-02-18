package api

import "api/models"

// Run database model
func Run() {
	models.AutoMigrations()
	listen(9000)
}
