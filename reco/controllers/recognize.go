package controllers

import (
	"log"
	"net/http"
	"reco/models"
	"reco/services"

	"github.com/Kagami/go-face"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RecognizeFaceInput struct {
	Photo string `json:"photo"`
}

func RecognizeFace(c *gin.Context) {
	log.Println("RecognizeFace")
	var input RecognizeFaceInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	users, err := models.FindAllUsers(db)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var descriptors []models.UserDescriptor

	for _, user := range *users {
		descriptors = append(descriptors, user.UserDescriptors...)
	}

	rec := c.MustGet("rec").(*face.Recognizer)

	descriptor := services.GetFaceDescriptor(rec, input.Photo)
	descriptorId := services.FindThresholdUserID(rec, descriptors, *descriptor, 0.4)

	if descriptorId == -1 {
		c.JSON(http.StatusOK, gin.H{"data": nil})
		return
	}

	recognizedDescriptor, err := models.FindDescriptorById(db, descriptorId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recognizedUser, err := models.FindUserById(db, int(recognizedDescriptor.UserID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": gin.H{"user": recognizedUser, "descriptor": recognizedDescriptor}})
}
