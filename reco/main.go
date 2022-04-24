package main

import (
	"log"
	"reco/models"
	"reco/routes"

	"github.com/Kagami/go-face"
)

func main() {
	db := models.SetupDB()
	db.AutoMigrate(&models.User{}, &models.UserDescriptor{})

	rec, err := face.NewRecognizer("./testdata/models")
	if err != nil {
		log.Fatalf("Can't init face recognizer: %v", err)

	}

	r := routes.SetupRoutes(db, rec)
	r.Run()
}
