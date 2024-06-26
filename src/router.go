package router

import (
	cloudinaryController "github.com/NabinGrz/ExpenseTracker/src/cloudinary/controller"
	expensesController "github.com/NabinGrz/ExpenseTracker/src/expense/controller"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Router(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	expenseApi := router.Group("/api/v1")

	expense := expenseApi.Group("/expense")
	expense.POST("/createCategory", func(ctx *gin.Context) { expensesController.CreateCategory(ctx, db) })
	expense.POST("/createExpense", func(ctx *gin.Context) { expensesController.CreateExpense(ctx, db) })
	expense.GET("/getExpenseDetail/:id", func(ctx *gin.Context) { expensesController.GetExpenseDetail(ctx, db) })
	expense.PUT("/updateExpense/:id", func(ctx *gin.Context) { expensesController.UpdateExpense(ctx, db) })
	expense.GET("/getExpenses", func(ctx *gin.Context) { expensesController.GetExpenses(ctx, db) })
	expense.GET("/getSelectedDateExpenses", func(ctx *gin.Context) { expensesController.GetSelectedDateExpenses(ctx, db) })
	expense.GET("/getSelectedRangeExpenses", func(ctx *gin.Context) { expensesController.GetSelectedRangeExpenses(ctx, db) })
	expense.GET("/searchExpense", func(ctx *gin.Context) { expensesController.SearchExpense(ctx, db) })
	expense.GET("/searchExpenseWithCategory", func(ctx *gin.Context) { expensesController.SpecificCategoryExpenses(ctx, db) })

	balance := expenseApi.Group("/balance")
	balance.PUT("/updateCash", func(ctx *gin.Context) { expensesController.UpdateAmount(ctx, db, true) })
	balance.PUT("/updateBankAmount", func(ctx *gin.Context) { expensesController.UpdateAmount(ctx, db, false) })

	//!! FILE UPLOAD
	expenseApi.POST("/upload-file", func(ctx *gin.Context) { cloudinaryController.UpdatePOSTImage(ctx) })
	return router
}
