package expensesController

import (
	"net/http"
	"strings"

	expenseModels "github.com/NabinGrz/ExpenseTracker/src/expense/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func CreateCategory(c *gin.Context, db *gorm.DB) {
	var input expenseModels.Category

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	category := expenseModels.Category{
		Name:  strings.TrimSpace(input.Name),
		Image: strings.TrimSpace(input.Image),
	}
	// Create the category
	if err := db.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully created category", "data": category})

}
