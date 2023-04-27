package main

import (
	"github.com/Course/Golang-Course/8-FamilyFinanceProject/routing"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routing.CreateUsersRouting(r)
	routing.CreateIncomeRouting(r)

	// Run the server
	r.Run(":8080")
}
