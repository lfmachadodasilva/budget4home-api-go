package expense

import (
	"budget4home/src/middleware"

	"github.com/gin-gonic/gin"
)

type ExpenseController struct {
}

func NewController(r *gin.RouterGroup) {
	controller := &ExpenseController{}
	r.GET("/expenses", middleware.AuthMiddleware, controller.Get)
	r.GET("/expenses/full", middleware.AuthMiddleware, controller.GetFull)
	r.POST("/expenses", middleware.AuthMiddleware, controller.Post)
	r.PUT("/expenses", middleware.AuthMiddleware, controller.Put)
	r.DELETE("/expenses", middleware.AuthMiddleware, controller.Delete)
}

// GetExpenses godoc
// @Security ApiKeyAuth
// @Tags Expenses
// @Accept json
// @Produce json
// @Router /expenses [get]
func (this *ExpenseController) Get(c *gin.Context) {
	c.JSON(200, nil)
}

// GetExpensesFull godoc
// @Security ApiKeyAuth
// @Tags Expenses
// @Accept json
// @Produce json
// @Router /expenses/full [get]
func (this *ExpenseController) GetFull(c *gin.Context) {
	c.JSON(200, nil)
}

// PostExpenses godoc
// @Security ApiKeyAuth
// @Tags Expenses
// @Accept json
// @Produce json
// @Router /expenses [post]
func (this *ExpenseController) Post(c *gin.Context) {
	c.JSON(200, nil)
}

// PutExpenses godoc
// @Security ApiKeyAuth
// @Tags Expenses
// @Accept json
// @Produce json
// @Router /expenses [put]
func (this *ExpenseController) Put(c *gin.Context) {
	c.JSON(200, nil)
}

// DeleteExpenses godoc
// @Security ApiKeyAuth
// @Tags Expenses
// @Accept json
// @Produce json
// @Router /expenses [delete]
func (this *ExpenseController) Delete(c *gin.Context) {
	c.JSON(200, nil)
}
