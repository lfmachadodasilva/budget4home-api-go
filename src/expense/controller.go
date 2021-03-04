package expense

import (
	"github.com/gin-gonic/gin"
)

type ExpenseController struct {
}

func NewController(r *gin.RouterGroup) {
	controller := &ExpenseController{}
	e := r.Group("/expenses")
	e.GET("/expenses", controller.Fetch)
	e.GET("/expenses/full", controller.FetchFull)
}

// GetExpenses godoc
// @Summary Show a list of expenses
// @Description get list of expenses
// @Tags expenses
// @Accept  json
// @Produce  json
// @Router /api/expenses [get]
func (this *ExpenseController) Fetch(c *gin.Context) {
	c.JSON(200, nil)
}

// GetExpensesFull godoc
// @Summary Show a list of expenses with full details
// @Tags expenses
// @Accept  json
// @Produce  json
// @Router /api/expenses/full [get]
func (this *ExpenseController) FetchFull(c *gin.Context) {
	c.JSON(200, nil)
}
