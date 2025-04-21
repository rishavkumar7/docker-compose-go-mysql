package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	messageRouter := router.Group("/message")
	messageRouter.POST("/add", addMessage)
	messageRouter.GET("/fetch", fetchMessages)
	messageRouter.PUT("/update/:id", updateMessage)
	messageRouter.DELETE("/remove/:id", removeMessage)
	fmt.Println("Starting server on port 8501...")
	router.Run(":8501")
}

func addMessage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Message added successfully",
		"success": true,
	})
}

func fetchMessages(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Messages fetched successfully",
		"success": true,
	})
}

func updateMessage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Message updated successfully",
		"success": true,
	})
}

func removeMessage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Message removed successfully",
		"success": true,
	})
}
