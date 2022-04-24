package controllers

import (
	"net/http"
	"reco/models"
	"reco/services"
	"strconv"

	"github.com/Kagami/go-face"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllUsers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	users, err := models.FindAllUsers(db)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

type CreateUserInput struct {
	Photo string `json:"photo"`
	Name  string `json:"name"`
}

func CreateUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rec := c.MustGet("rec").(*face.Recognizer)
	descriptor := services.GetFaceDescriptor(rec, input.Photo)
	descriptorBytes := services.DescriptorToBytes(*descriptor)

	user := models.User{Name: input.Name, UserDescriptors: []models.UserDescriptor{{Descriptor: descriptorBytes[:], Photo: input.Photo}}}
	db.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

type UpdateUserInput struct {
	Photo string `json:"photo"`
}

func UpdateUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rec := c.MustGet("rec").(*face.Recognizer)
	descriptor := services.GetFaceDescriptor(rec, input.Photo)
	descriptorBytes := services.DescriptorToBytes(*descriptor)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := models.FindUserById(db, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Model(&user).Association("UserDescriptors").Append(&models.UserDescriptor{Descriptor: descriptorBytes[:], Photo: input.Photo})

	c.JSON(http.StatusCreated, gin.H{"data": "ok"})
}

func DeleteUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	userId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := models.FindUserById(db, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Unscoped().Where("user_id = ?", userId).Delete(&models.UserDescriptor{})
	db.Unscoped().Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}
