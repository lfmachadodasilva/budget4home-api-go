package expense

import (
	"budget4home/src/middleware"

	"github.com/gin-gonic/gin"
)

type ExpenseController struct {
}

func NewController(r *gin.RouterGroup) {
	controller := &ExpenseController{}
	r.GET("/expenses", middleware.AuthMiddleware, controller.Fetch)
	r.GET("/expenses/full", middleware.AuthMiddleware, controller.FetchFull)
}

// GetExpenses godoc
// @Security ApiKeyAuth
// @Summary Show a list of expenses
// @Description get list of expenses
// @Tags expenses
// @Accept  json
// @Produce  json
// @Router /expenses [get]
func (this *ExpenseController) Fetch(c *gin.Context) {
	c.JSON(200, nil)
}

// GetExpensesFull godoc
// @Security ApiKeyAuth
// @Summary Show a list of expenses with full details
// @Tags expenses
// @Accept  json
// @Produce  json
// @Router /expenses/full [get]
func (this *ExpenseController) FetchFull(c *gin.Context) {
	c.JSON(200, nil)
}
