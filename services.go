package main

import (
	// "time"
	"github.com/gin-gonic/gin"
	"net/http"
)
func getTasks(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
	return
}

func getTask(c *gin.Context){
	id := c.Param("id")
		for _, task := range tasks {
			if task.ID == id {
				c.JSON(http.StatusOK, gin.H{"task": task})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not Found!"})
		return
}

func createTask(c *gin.Context){
	var newTask Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tasks = append(tasks, newTask)
	c.JSON(http.StatusOK, gin.H{"message": "Task Created"})
	return
}

func updateTask(c *gin.Context){
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
}

func deleteTask(c *gin.Context){
	id := c.Param("id")

	for i, val := range tasks {
		if val.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Task Deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task Not Found"})
	return
}