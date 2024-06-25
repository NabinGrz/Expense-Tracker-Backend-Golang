package databaseService

import (
	"fmt"
	"os"

	expenseModels "github.com/NabinGrz/ExpenseTracker/src/expense/models"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func DBConnection() {

	var err error

	var host = os.Getenv("HOST")
	var port = os.Getenv("DB_PORT")
	var user = os.Getenv("DB_USER")
	var dbName = os.Getenv("DB_NAME")
	var sslmode = os.Getenv("DB_SSLMODE")
	var password = os.Getenv("PASSWORD")

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		host, port, user, dbName, sslmode, password)
	DB, err = gorm.Open("postgres", dsn)
	DB.LogMode(true)
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect database")

	}
	if DB != nil {
		fmt.Println("Successfully connected to database.....")
	}
	DB.AutoMigrate(&expenseModels.Category{}, &expenseModels.Expense{}, &expenseModels.Balance{})

}
