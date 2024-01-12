// main.go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

var DB *bun.DB

func main() {
	fmt.Println("Task Manager API Project")

	// Initialize the database
	db, err := InitializeDB()
	if err != nil {
		fmt.Println("Failed to initialize the database:", err)
		return
	}
	defer db.Close()

	DB = db
	// Set up Gin router
	router := gin.Default()

	// Define routes
	router.GET("/", HomePage)
	router.GET("/tasks", GetTasks)
	router.GET("/tasks/:id", GetTask)
	router.DELETE("/tasks/:id", RemoveTask)
	router.POST("/tasks", AddTask)
	router.PUT("/tasks/:id", UpdateTask)

	// Start the server
	router.Run(":8080")
}

// package main

// import (
// 	// "fmt"
// 	// "net/http"
// 	"time"

// 	"github.com/gin-gonic/gin"
// )

// type Task struct{
// 	ID string `json:"id"`
// 	Title string `json:"title"`
// 	Description string `json:"description"`
// 	DueDate time.Time `json:"due_date"`
// 	Status string `json:"status"`
// }

// var tasks = []Task{
// 	{ID: "1", Title: "Task 1", Description: "Description 1", DueDate: time.Now(), Status: "pending"},
// 	{ID: "2", Title: "Task 2", Description: "Description 2", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
// 	{ID: "3", Title: "Task 3", Description: "Description 3", DueDate: time.Now().AddDate(0, 0, 2), Status: "completed"},
// }
// func main() {
// 	router := gin.Default()

// 	router.GET("/ping", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"message": "pong",
// 		})
// 	})

// 	router.GET("/tasks", GetTasks)
// 	router.GET("/tasks/:id", GetTask)

// 	router.PUT("/tasks/:id", UpdateTask)

// 	router.DELETE("/tasks/:id", DeleteTask)

// 	router.POST("/tasks", CreateTask)
// 	router.Run(":8080")
// }