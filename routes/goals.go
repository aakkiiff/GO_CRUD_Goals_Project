package routes

import (
	"github.com/gin-gonic/gin"
	"goals.com/models"
)

func hello(context *gin.Context) {
	context.JSON(200, gin.H{"message": "hello from golang"})
}

func createGoal(c *gin.Context) {
	var goal models.Goal
	err := c.ShouldBindJSON(&goal)
	if err != nil {
		c.JSON(400, gin.H{"message": "could not parse the data", "error": err.Error()})
		return
	}
	err = goal.Save()
	if err != nil {
		c.JSON(400, gin.H{"message": "could not save the data", "error": err.Error()})
		return
	}
	c.JSON(201, gin.H{"message": "goal saved", "goal": goal})
}
