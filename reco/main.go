package main

import (
	"reco/models"
	"reco/routes"
)

func main() {
	db := models.SetupDB()
	db.AutoMigrate(&models.User{}, &models.UserDescriptor{})

	r := routes.SetupRoutes(db)
	r.Run()
}
