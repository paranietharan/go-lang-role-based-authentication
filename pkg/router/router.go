package router

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StartServer(db *gorm.DB) {
	port := os.Getenv("PORT")

	if port == "" {
		log.Println("PORT variable not found in .env file")
	}

	router := gin.New()
	router.Use(gin.Logger())

	UserRoutes(router)
	AuthRoutes(router)

	router.GET("/api-1", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"success": "Access granted for api-1"})
	})

	router.GET("/api-2", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"success": "Access granted for api-2"})
	})

	router.Run(":" + port)
}
