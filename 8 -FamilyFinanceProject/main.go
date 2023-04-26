package main

import (
	"github.com/Course/FamilyFinanceProject/routing"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routing.CreateUsersRouting(r)

	// Run the server
	r.Run(":8080")
}
