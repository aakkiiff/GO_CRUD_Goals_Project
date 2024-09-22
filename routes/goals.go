package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"goals.com/models"
)

func createGoal(c *gin.Context) {
	var NewGoal models.Goal
	err := c.ShouldBindJSON(&NewGoal)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid request payload"})
		return
	}
	// fmt.Println(NewGoal)

	err = mgm.Coll(&NewGoal).Create(&NewGoal)
	if err != nil {
		c.JSON(500, gin.H{"error": "could not save data to database!"})
		return
	}
	c.JSON(200, gin.H{"message": "Goal saved!", "goal": NewGoal})
}
