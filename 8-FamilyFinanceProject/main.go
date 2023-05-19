package main

import (
	"github.com/Course/Golang-Course/8-FamilyFinanceProject/routing"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new default instance of the Gin framework router
	r := gin.Default()

	// Register routes for user-related functionality
	routing.CreateUsersRouting(r)

	// Register routes for income-related functionality
	routing.CreateIncomeRouting(r)

	// Run the server on port 8080
	r.Run(":8080")
}
