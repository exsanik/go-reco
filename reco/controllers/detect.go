package controllers

import (
	"log"
	"net/http"
	"reco/services"

	"github.com/gin-gonic/gin"
)

type DetectFaceInput struct {
	Photo string `json:"photo"`
}

func DetectFace(c *gin.Context) {
	log.Println("detect_face")
	var input DetectFaceInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	faces := services.FindFaces(input.Photo)

	c.JSON(http.StatusOK, gin.H{"data": faces})
}
