package routing

import (
	"net/http"
	"strconv"

	"github.com/Course/FamilyFinanceProject/common"
	"github.com/gin-gonic/gin"
)

var users []common.User

func CheckUserExists(user common.User) bool {
	for _, u := range users {
		if u.FirstName == user.FirstName && u.LastName == user.LastName {
			return true
		}
	}
	return false
}

func CreateUsersRouting(r *gin.Engine) {
	// Create User
	r.POST("/users", func(c *gin.Context) {
		var user common.User
		if err := c.BindJSON(&user); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if CheckUserExists(user) {
			c.AbortWithStatusJSON(http.StatusConflict, "user already exists")
			return
		}
		user.ID = len(users) + 1
		users = append(users, user)
		c.JSON(http.StatusCreated, user)
	})

	// Read all users
	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, users)
	})

	// Read single
	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, user := range users {
			if strconv.Itoa(user.ID) == id {
				c.JSON(http.StatusOK, user)
				return
			}
		}
		c.AbortWithStatus(http.StatusNotFound)
	})

	// Update user
	r.PUT("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updatedUser common.User
		if err := c.BindJSON(&updatedUser); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		for i, user := range users {
			if strconv.Itoa(user.ID) == id {
				updatedUser.ID = user.ID
				users[i] = updatedUser
				c.JSON(http.StatusOK, updatedUser)
				return
			}
		}
		c.AbortWithStatus(http.StatusNotFound)
	})

	// Delete user
	r.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		for i, user := range users {
			if strconv.Itoa(user.ID) == id {
				users = append(users[:i], users[i+1:]...)
				c.Status(http.StatusNoContent)
				return
			}
		}
		c.AbortWithStatus(http.StatusNotFound)
	})
}
