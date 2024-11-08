package main

import (
	"assignment/model"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World!")

	server := gin.Default()

	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World! This project is created to demonstrate API creation in Go with Gin. API : - /Create-Student -> POST request -> Use Postman or Curl to send student creation request. Make sure to send json with ID,Name,Age and Email",
		})
	})

	server.POST("/Create-Student", func(c *gin.Context) {

		var student *model.Student

		if err := c.ShouldBindJSON(&student); err != nil {
			c.JSON(400, gin.H{
				"message": "Invalid request",
			})
			return
		}

		fmt.Print("Student : ", student)

		model.Students = append(model.Students, *student)

		c.JSON(200, gin.H{
			"message": "Hello, World!",
			"data":    model.Students,
		})
	})

	server.Run(":10000")
}
