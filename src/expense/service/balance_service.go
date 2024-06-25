package expenseService

import (
	expenseModels "github.com/NabinGrz/ExpenseTracker/src/expense/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func UpdateAmount(c *gin.Context, db *gorm.DB, isCash bool, amount float64) bool {
	id := "37025142-f400-4ed4-b8e8-a5b81f067ad4"

	var existing expenseModels.Balance
	if err := db.First(&existing, "id = ?", id).Error; err != nil {
		return false
	}

	if amount != 0 {
		if isCash {
			existing.CashAmount = existing.CashAmount - amount
		} else {
			existing.BankAmount = existing.BankAmount - amount
		}
	}

	if err := db.Save(&existing).Error; err != nil {
		return false
	}
	return true
}
func UpdateExpenseAmount(c *gin.Context, db *gorm.DB, isCash bool, oldAmount float64, newAmount float64) bool {
	id := "37025142-f400-4ed4-b8e8-a5b81f067ad4"

	var existing expenseModels.Balance
	if err := db.First(&existing, "id = ?", id).Error; err != nil {
		return false
	}

	if oldAmount != 0 && newAmount != 0 {
		if isCash {
			existing.CashAmount = (existing.CashAmount + oldAmount) - newAmount
		} else {
			existing.BankAmount = (existing.BankAmount + oldAmount) - newAmount
		}
	}

	if err := db.Save(&existing).Error; err != nil {
		return false
	}
	return true
}
func UpdateExpenseAmountWhenTypeChange(c *gin.Context, db *gorm.DB, isCash bool, oldAmount float64, newAmount float64) bool {
	id := "37025142-f400-4ed4-b8e8-a5b81f067ad4"

	var existing expenseModels.Balance
	if err := db.First(&existing, "id = ?", id).Error; err != nil {
		return false
	}

	if oldAmount != 0 && newAmount != 0 {
		if isCash {
			existing.BankAmount = existing.BankAmount + oldAmount
			existing.CashAmount = existing.CashAmount - newAmount
		} else {
			existing.CashAmount = existing.CashAmount + oldAmount
			existing.BankAmount = existing.BankAmount - newAmount
		}
	}

	if err := db.Save(&existing).Error; err != nil {
		return false
	}
	return true
}
