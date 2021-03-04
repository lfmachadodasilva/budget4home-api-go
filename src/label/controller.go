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
	l := r.Group("/labels")
	l.GET("/labels", controller.Fetch)
	l.GET("/labels/full", controller.FetchFull)
}

// GetLabels godoc
// @Summary Show a list of labels
// @Description get list of labels
// @Accept  json
// @Produce  json
// @Success 200 {string} string "ok"
// @Router /api/labels [get]
func (this *LabelController) Fetch(c *gin.Context) {
	labels, _ := this.service.Fetch()
	c.JSON(200, labels)
	// c.JSON(200, gin.H{"data": labels})
}

// GetLabelsFull godoc
// @Summary Show a list of labels with full details
// @Accept  json
// @Produce  json
// @Router /api/labels/full [get]
func (this *LabelController) FetchFull(c *gin.Context) {
	labels, _ := this.service.Fetch()
	c.JSON(200, labels)
	// c.JSON(200, gin.H{"data": labels})
}
