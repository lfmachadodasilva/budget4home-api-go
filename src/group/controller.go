package group

import (
	"budget4home/src/middleware"

	"github.com/gin-gonic/gin"
)

type GroupController struct {
	service IGroupService
}

func NewController(r *gin.RouterGroup, service IGroupService) {
	controller := &GroupController{
		service: service,
	}
	r.GET("/groups", middleware.AuthMiddleware, controller.Get)
	r.GET("/groups/full", middleware.AuthMiddleware, controller.GetFull)
	r.POST("/groups", middleware.AuthMiddleware, controller.Post)
	r.PUT("/groups", middleware.AuthMiddleware, controller.Put)
	r.DELETE("/groups", middleware.AuthMiddleware, controller.Delete)
}

// GetGroups godoc
// @Security ApiKeyAuth
// @Tags Groups
// @Accept json
// @Produce json
// @Success 200 {object} []group.Group
// @Router /groups [get]
func (this *GroupController) Get(c *gin.Context) {
	groups, _ := this.service.Fetch()
	c.JSON(200, groups)
}

// GetFull get all groups full
// GetGroupsFull godoc
// @Security ApiKeyAuth
// @Tags Groups
// @Accept json
// @Produce json
// @Success 200 {object} []group.FullResponse
// @Router /groups/full [get]
func (con *GroupController) GetFull(c *gin.Context) {
	groups, _ := con.service.FetchFull(c)
	c.JSON(200, groups)
}

// PostGroups godoc
// @Security ApiKeyAuth
// @Tags Groups
// @Accept json
// @Produce json
// @Router /groups [post]
func (this *GroupController) Post(c *gin.Context) {
	c.JSON(200, nil)
}

// PutGroups godoc
// @Security ApiKeyAuth
// @Tags Groups
// @Accept json
// @Produce json
// @Router /groups [put]
func (this *GroupController) Put(c *gin.Context) {
	c.JSON(200, nil)
}

// DeleteGroups godoc
// @Security ApiKeyAuth
// @Tags Groups
// @Accept json
// @Produce json
// @Router /groups [delete]
func (this *GroupController) Delete(c *gin.Context) {
	c.JSON(200, nil)
}
