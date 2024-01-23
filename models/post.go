package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// import "time"

type User_Input struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Phone uint   `json:"phone_no" binding:"required"`
}

func CreateUser(c *gin.Context) {
	var input User_Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := User_Input{Name: input.Name, Email: input.Email, Phone: input.Phone}
	DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func GetUsers(c *gin.Context) {
	var users []User_Input

	// Retrieve all users from the User_Input table
	result := DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Respond with the user data
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func GetUserByID(c *gin.Context) {
	var user User_Input
	id := c.Param("id")

	// Retrieve user based on ID
	result := DB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Respond with the user data
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func GetUserByEmail(c *gin.Context) {
	var user User_Input
	email := c.Param("email")

	// Retrieve user based on Email
	result := DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Respond with the user data
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// UpdateUserByID updates a user based on the ID parameter
func UpdateUserByID(c *gin.Context) {
	var user User_Input
	id := c.Param("id")

	// Retrieve user based on ID
	result := DB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Bind the JSON request body to the User_Input struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the updated user to the database
	result = DB.Save(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user"})
		return
	}

	// Respond with the updated user data
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUserByEmail(c *gin.Context) {
	var user User_Input
	email := c.Param("email")

	// Retrieve user based on Email
	result := DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Delete the user from the database
	result = DB.Delete(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting user"})
		return
	}

	// Respond with a success message
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
