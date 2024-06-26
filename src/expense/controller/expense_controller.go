package expensesController

import (
	"log"
	"net/http"
	"strings"
	"time"

	expenseModels "github.com/NabinGrz/ExpenseTracker/src/expense/models"
	expenseService "github.com/NabinGrz/ExpenseTracker/src/expense/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

func SpecificCategoryExpenses(c *gin.Context, db *gorm.DB) {
	categoryId := c.Query("categoryId")

	if err := uuid.Validate(categoryId); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Please enter a validate category ID"})
		return
	}
	var expenses []expenseModels.Expense
	if err := db.Preload("Category").Where("category_id = ?", categoryId).Find(&expenses).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	var totalAmount float64
	for _, ex := range expenses {
		totalAmount += ex.Amount
	}
	c.JSON(http.StatusOK, gin.H{"expenses": expenses, "total_count": len(expenses), "total_amount": totalAmount})
}

func SearchExpense(c *gin.Context, db *gorm.DB) {
	keyword := c.Query("keyword")

	if keyword == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Please enter a keyword to search expense"})
		return
	}
	searchPattern := "%" + keyword + "%"
	var expenses []expenseModels.Expense

	if err := db.Joins("JOIN categories ON categories.id = expenses.category_id").
		Where("expenses.name ILIKE ? OR categories.name ILIKE ?", searchPattern, searchPattern).
		Preload("Category").
		Find(&expenses).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"expenses": expenses, "total_count": len(expenses)})

}
func GetSelectedRangeExpenses(c *gin.Context, db *gorm.DB) {
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	var startDate, endDate time.Time
	var err error

	if startDateStr != "" {
		startDate, err = time.Parse("2006-01-02", startDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date format. Use yyyy-mm-dd."})
			return
		}
	}

	if endDateStr != "" {
		endDate, err = time.Parse("2006-01-02", endDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date format. Use yyyy-mm-dd."})
			return
		}
		// Adjust the end date to include the whole day
		endDate = endDate.Add(24 * time.Hour)
	} else {
		// Default to today's date if end date is not provided
		endDate = time.Now().Add(24 * time.Hour)
	}

	// Validate date range
	if !startDate.IsZero() && !endDate.IsZero() && startDate.After(endDate) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Start date cannot be after end date."})
		return
	}
	var expenses []expenseModels.Expense
	// dateStr := expense.CreatedAt.Format("2006-01-02")

	if err := db.Preload("Category").Where("created_at BETWEEN ? AND ?", startDate, endDate).Find(&expenses).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	var totalAmount float64
	for _, ex := range expenses {
		totalAmount += ex.Amount
	}
	c.JSON(http.StatusOK, gin.H{"expenses": expenses, "total_count": len(expenses), "total_amount": totalAmount})
}
func GetSelectedDateExpenses(c *gin.Context, db *gorm.DB) {
	// Example date input provided by user
	selectedDate := c.Query("date")
	parsedDate, err := time.Parse("2006-01-02", selectedDate)
	if err != nil {
		log.Fatal(err)
	}
	var expenses []expenseModels.Expense
	// dateStr := expense.CreatedAt.Format("2006-01-02")

	if err := db.Preload("Category").Where("DATE(created_at) = ?", parsedDate).Find(&expenses).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	var totalAmount float64
	for _, ex := range expenses {
		totalAmount += ex.Amount
	}
	c.JSON(http.StatusOK, gin.H{"expenses": expenses, "total_count": len(expenses), "total_amount": totalAmount})
}

func GetExpenseDetail(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Expense Id must be valid UUID"})
		return
	}

	var expense expenseModels.Expense

	if err := db.Preload("Category").Find(&expense, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, expense)

}
func GetExpenses(c *gin.Context, db *gorm.DB) {

	var expenses []expenseModels.Expense
	if err := db.Preload("Category").Find(&expenses).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	var totalAmount float64
	for _, ex := range expenses {
		totalAmount += ex.Amount
	}
	c.JSON(http.StatusOK, gin.H{"expenses": expenses, "total_count": len(expenses), "total_amount": totalAmount})

}

func CreateExpense(c *gin.Context, db *gorm.DB) {
	var input expenseModels.Expense

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	expense := expenseModels.Expense{
		Name:       strings.TrimSpace(input.Name),
		CreatedAt:  time.Now(),
		Amount:     input.Amount,
		CategoryID: input.CategoryID,
		IsCash:     input.IsCash,
	}
	// Create the expense
	if err := db.Create(&expense).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	expenseService.UpdateAmount(c, db, expense.IsCash, input.Amount)

	c.JSON(http.StatusOK, gin.H{"message": "Successfully created expense"})
}

func UpdateExpense(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Expense Id must be valid UUID"})
		return
	}

	var update expenseModels.Expense
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	var existing expenseModels.Expense
	if err := db.First(&existing, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Expense Not Found"})
		return
	}
	oldAmount := existing.Amount

	if update.Name != "" {
		existing.Name = strings.TrimSpace(update.Name)
	}
	if update.Amount != 0 {
		existing.Amount = update.Amount
	}
	if update.CategoryID.String() != "00000000-0000-0000-0000-000000000000" {
		existing.CategoryID = update.CategoryID
	}

	if err := db.Save(&existing).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update expense"})
		return
	}
	newAmount := update.Amount

	// isBankNow := !update.IsCash
	if existing.IsCash == update.IsCash {
		expenseService.UpdateExpenseAmount(c, db, update.IsCash, oldAmount, newAmount)
		c.JSON(http.StatusOK, gin.H{"message": "Successfully updated expense"})
	} else {

		expenseService.UpdateExpenseAmountWhenTypeChange(c, db, update.IsCash, oldAmount, newAmount)
		c.JSON(http.StatusOK, gin.H{"message": "Successfully updated expense"})

	}

}
