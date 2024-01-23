package main

import (
	"module/project1/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	// Define routes
	router.POST("/createUser", models.CreateUser)
	router.GET("/getUsers", models.GetUsers)
	router.GET("/user/:id", models.GetUserByID)
	router.GET("/userByEmail/:email", models.GetUserByEmail)
	router.PUT("/updateUser/:id", models.UpdateUserByID)
	router.DELETE("/deleteUser/:email", models.DeleteUserByEmail)

	router.Run("localhost:8080")

}
