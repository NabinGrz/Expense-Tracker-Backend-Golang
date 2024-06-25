package main

import (
	router "github.com/NabinGrz/ExpenseTracker/src"
	initializers "github.com/NabinGrz/ExpenseTracker/src/config"
	databaseService "github.com/NabinGrz/ExpenseTracker/src/database"
)

func init() {
	initializers.LoadEnvVariables()
	databaseService.DBConnection()
}

func main() {
	r := router.Router(databaseService.DB)

	r.Run()
}
