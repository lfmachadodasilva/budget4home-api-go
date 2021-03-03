package label

import (
	"github.com/gin-gonic/gin"
)

type LabelController struct {
	service ILabelService
}

func NewController(r *gin.RouterGroup, service ILabelService) {
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
