package main

import (
	// "fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Task struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	DueDate time.Time `json:"due_date"`
	Status string `json:"status"`
}

var tasks = []Task{
	Task{ID: "1", Title: "Task 1", Description: "Description 1", DueDate: time.Now(), Status: "pending"},
	Task{ID: "2", Title: "Task 2", Description: "Description 2", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
	Task{ID: "3", Title: "Task 3", Description: "Description 3", DueDate: time.Now().AddDate(0, 0, 2), Status: "completed"},
}
func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/tasks", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{"tasks": tasks})
	})


	router.GET("/tasks:id", func(c *gin.Context){
		id := c.Param("id")
		for _, task := range tasks {
			if task.ID == id {
				c.JSON(http.StatusOK, gin.H{"task": task})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not Found!"})

	})

	router.PUT("/tasks/:id", func(c *gin.Context){
		id := c.Param("id")

		var updatedTask Task

		if err := c.ShouldBindJSON(&updatedTask); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		for i , task := range tasks {
			if task.ID == id {
				if updatedTask.Title != ""{
					tasks[i].Title = updatedTask.Title
				}
				if updatedTask.Description != ""{
					tasks[i].Description = updatedTask.Description
				}

				c.JSON(http.StatusOK, gin.H{"message": "Task Updated"})
				return 
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Task Not Found"})
		return
	})

	router.DELETE("/tasks/:id", func(c *gin.Context){
		id := c.Param("id")

		for i, val := range tasks {
			if val.ID == id {
				tasks = append(tasks[:i], tasks[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "Task Deleted"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Task Not Found"})
	})

	router.POST("/tasks", func(c *gin.Context){
		var newTask Task
		if err := c.ShouldBindJSON(&newTask); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		tasks = append(tasks, newTask)
		c.JSON(http.StatusOK, gin.H{"message": "Task Created"})
	})
	router.Run(":8080")
}