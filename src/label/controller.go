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
	r.GET("/labels", controller.Get)
	r.GET("/labels/full", middleware.AuthMiddleware, controller.GetFull)
	r.POST("/labels", middleware.AuthMiddleware, controller.Post)
	r.PUT("/labels", middleware.AuthMiddleware, controller.Put)
	r.DELETE("/labels", middleware.AuthMiddleware, controller.Delete)
}

// GetLabels godoc
// @Tags Labels
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /labels [get]
func (this *LabelController) Get(c *gin.Context) {
	labels, _ := this.service.Fetch()
	c.JSON(200, labels)
}

// GetLabelsFull godoc
// @Tags Labels
// @Accept json
// @Produce json
// @Router /labels/full [get]
func (this *LabelController) GetFull(c *gin.Context) {
	labels, _ := this.service.Fetch()
	c.JSON(200, labels)
}

// PostLabels godoc
// @Tags Labels
// @Accept json
// @Produce json
// @Router /labels [post]
func (this *LabelController) Post(c *gin.Context) {
	c.JSON(200, nil)
}

// PutLabels godoc
// @Tags Labels
// @Accept json
// @Produce json
// @Router /labels [put]
func (this *LabelController) Put(c *gin.Context) {
	c.JSON(200, nil)
}

// DeleteLabels godoc
// @Tags Labels
// @Accept json
// @Produce json
// @Router /labels [delete]
func (this *LabelController) Delete(c *gin.Context) {
	c.JSON(200, nil)
}
