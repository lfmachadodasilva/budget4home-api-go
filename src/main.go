package main

import (
	"budget4home/src/controllers"
	"budget4home/src/repositories"
	"budget4home/src/services"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello world",
	})
}

func main() {
	fmt.Println("Hello world")

	r := gin.Default()
	r.GET("/", HomePage)

	api := r.Group("/v1")

	repo := repositories.NewLabelRepository()
	service := services.NewLabelService(repo)
	controllers.NewLabelController(api, service)

	if len(os.Getenv("PORT")) == 0 {
		// set default port
		os.Setenv("PORT", "5000")
	}
	r.Run(":" + os.Getenv("PORT"))
}
