package user

import (
	"budget4home/src/middleware"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service IUserService
}

func NewController(r *gin.RouterGroup, service IUserService) {
	controller := &UserController{
		service: service,
	}
	r.GET("/users", middleware.AuthMiddleware, controller.Fetch)
}

// GetLabels godoc
// @Security ApiKeyAuth
// @Summary Show a list of users
// @Description get list of users
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200 {object} []user.User
// @Router /users [get]
func (this *UserController) Fetch(c *gin.Context) {
	users, _ := this.service.Fetch(c)
	c.JSON(200, users)
}
