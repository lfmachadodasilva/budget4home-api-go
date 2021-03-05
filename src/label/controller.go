package label

import (
	"budget4home/src/middleware"

	"github.com/gin-gonic/gin"
)

type LabelController struct {
	service ILabelService
}

func NewController(r *gin.RouterGroup, service ILabelService) {
	controller := &LabelController{
		service: service,
	}
	r.GET("/labels", middleware.AuthMiddleware, controller.Fetch)
	r.GET("/labels/full", middleware.AuthMiddleware, controller.FetchFull)
}

// GetLabels godoc
// @Security ApiKeyAuth
// @Summary Show a list of labels
// @Description get list of labels
// @Tags Labels
// @Accept  json
// @Produce  json
// @Success 200 {string} string "ok"
// @Router /labels [get]
func (this *LabelController) Fetch(c *gin.Context) {
	labels, _ := this.service.Fetch()
	c.JSON(200, labels)
	// c.JSON(200, gin.H{"data": labels})
}

// GetLabelsFull godoc
// @Summary Show a list of labels with full details
// @Tags Labels
// @Accept  json
// @Produce  json
// @Router /labels/full [get]
func (this *LabelController) FetchFull(c *gin.Context) {
	labels, _ := this.service.Fetch()
	c.JSON(200, labels)
	// c.JSON(200, gin.H{"data": labels})
}
