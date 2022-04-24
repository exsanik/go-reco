package routes

import (
	"reco/controllers"

	"github.com/Kagami/go-face"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT,DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func SetupRoutes(db *gorm.DB, rec *face.Recognizer) *gin.Engine {
	r := gin.Default()
	r.Use(CORSMiddleware())

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Set("rec", rec)
	})

	apiGroup := r.Group("/api")
	{
		apiGroup.POST("/detect", controllers.DetectFace)

		apiGroup.POST("/recognize", controllers.RecognizeFace)

		usersGroup := apiGroup.Group("/users")
		{
			usersGroup.GET("", controllers.GetAllUsers)
			usersGroup.POST("", controllers.CreateUser)
			usersGroup.PUT("/:id", controllers.UpdateUser)
			usersGroup.DELETE("/:id", controllers.DeleteUser)
		}
	}

	return r
}
