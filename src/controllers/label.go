package controllers

import (
	"budget4home/src/services"

	"github.com/gin-gonic/gin"
)

type LabelController struct {
	service services.ILabelService
}

func NewLabelController(r *gin.RouterGroup, service services.ILabelService) {
	controller := &LabelController{
		service: service,
	}
	r.GET("/labels", controller.Fetch)
}

func (this *LabelController) Fetch(c *gin.Context) {
	labels, _ := this.service.Fetch()
	c.JSON(200, labels)
	// c.JSON(200, gin.H{"data": labels})
}
