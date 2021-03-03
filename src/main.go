package main

import (
	"fmt"

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
	r.Run()
}
