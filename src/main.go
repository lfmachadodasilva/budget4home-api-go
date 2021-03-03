package main

import (
	"budget4home/src/config"
	"budget4home/src/label"
	"os"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello world",
	})
}

func main() {
	db := config.PrepareDb()

	r := gin.Default()
	r.GET("/", HomePage)

	api := r.Group("/v1")

	repo := label.NewRepository(db)
	service := label.NewService(repo)
	label.NewController(api, service)

	if len(os.Getenv("PORT")) == 0 {
		// set default port
		os.Setenv("PORT", "5000")
	}
	r.Run(":" + os.Getenv("PORT"))
}
