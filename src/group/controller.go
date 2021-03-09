package group

import (
	"budget4home/src/middleware"

	"github.com/gin-gonic/gin"
)

// Controller group controller
type Controller struct {
	service IGroupService
}

// NewController constructor
func NewController(r *gin.RouterGroup, service IGroupService) {
	controller := &Controller{
		service: service,
	}
	r.GET("/groups", middleware.AuthMiddleware, controller.Get)
	r.GET("/groups/full", middleware.AuthMiddleware, controller.GetFull)
	r.POST("/groups", middleware.AuthMiddleware, controller.Post)
	r.PUT("/groups", middleware.AuthMiddleware, controller.Put)
	r.DELETE("/groups", middleware.AuthMiddleware, controller.Delete)
}

// Get get all groups
// GetGroups godoc
// @Security ApiKeyAuth
// @Tags Groups
// @Accept json
// @Produce json
// @Success 200 {object} []group.Group
// @Router /groups [get]
func (con *Controller) Get(c *gin.Context) {
	groups, _ := con.service.Fetch()
	c.JSON(200, groups)
}

// GetFull godoc
// @Security ApiKeyAuth
// @Tags Groups
// @Accept json
// @Produce json
// @Success 200 {object} []group.FullResponse
// @Router /groups/full [get]
func (con *Controller) GetFull(c *gin.Context) {
	groups, _ := con.service.FetchFull(c)
	c.JSON(200, groups)
}

// Post godoc
// @Security ApiKeyAuth
// @Tags Groups
// @Accept json
// @Produce json
// @Param group body group.AddRequest true "group info"
// @Success 200 {object} []group.Response
// @Router /groups [post]
func (con *Controller) Post(c *gin.Context) {
	con.service.Add(c)
	c.JSON(200, nil)
}

// Put godoc
// @Security ApiKeyAuth
// @Tags Groups
// @Accept json
// @Produce json
// @Router /groups [put]
func (con *Controller) Put(c *gin.Context) {
	c.JSON(200, nil)
}

// Delete godoc
// @Security ApiKeyAuth
// @Tags Groups
// @Accept json
// @Produce json
// @Router /groups [delete]
func (con *Controller) Delete(c *gin.Context) {
	c.JSON(200, nil)
}
