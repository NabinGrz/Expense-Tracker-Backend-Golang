package expensesController

import (
	"net/http"

	expenseModels "github.com/NabinGrz/ExpenseTracker/src/expense/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func UpdateAmount(c *gin.Context, db *gorm.DB, isCash bool) {
	id := "37025142-f400-4ed4-b8e8-a5b81f067ad4"
	var update expenseModels.Balance
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	var existing expenseModels.Balance
	if err := db.First(&existing, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Balance Not Found"})
		return
	}

	if isCash {
		if update.CashAmount != 0 {
			existing.CashAmount = update.CashAmount
		}
	} else {
		if update.BankAmount != 0 {
			existing.BankAmount = update.BankAmount
		}
	}

	if err := db.Save(&existing).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update amount"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully updated amount"})
}
