package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rishavkumar7/docker-compose-go-mysql/database"
)

var db *sql.DB

func main() {
	db = database.ConnectDb()

	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS messages (id INT AUTO_INCREMENT PRIMARY KEY, content TEXT)`)
	if err != nil {
		log.Fatalf("Error while creating table: %v\n", err)
		return
	}

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
	var message struct {
		Content string `json:"content" binding:"required"`
	}
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	_, err := db.Exec(`INSERT INTO messages (content) VALUES (?)`, message.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while adding message",
			"success": false,
			"error":   err.Error(),
		})
		return
	}

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
