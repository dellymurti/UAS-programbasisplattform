package userhandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var userList []User

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   userList,
	})
}

func Show(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var user User
	for _, u := range userList {
		if u.ID == id {
			user = u
			break
		}
	}

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   user,
	})
}

func Create(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser.ID = len(userList) + 1

		userList = append(userList, newUser)

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   newUser,
	})
}

func Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var existingUser User
	for i, u := range userList {
		if u.ID == id {
			existingUser = u

						var updatedUser User
			if err := c.ShouldBindJSON(&updatedUser); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			updatedUser.ID = existingUser.ID
			userList[i] = updatedUser

			c.JSON(http.StatusOK, gin.H{
				"status": "success",
				"data":   updatedUser,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

		for i, u := range userList {
		if u.ID == id {
			// Remove the user from the list (replace this with your actual data removal logic).
			userList = append(userList[:i], userList[i+1:]...)

			c.JSON(http.StatusOK, gin.H{"status": "success"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}
