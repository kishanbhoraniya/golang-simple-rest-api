package main

import (
	"example/rest-api/controllers"
	"fmt"

	"github.com/gin-gonic/gin"
)

const portNumber = ":9091"

func main() {
	// creates a Gin router with default middleware: logger and recovery middleware
	r := gin.Default()

	// Routes
	r.GET("/post", controllers.GetAllPosts)
	r.GET("/post/:id", controllers.GetPost)
	r.POST("/post", controllers.CreatePost)
	r.PUT("/post/:id", controllers.UpdatePost)
	r.DELETE("/post/:id", controllers.DeletePost)

	// Run the server on port
	fmt.Printf("Starting application on port %v\n", portNumber)
	r.Run(portNumber)
}
