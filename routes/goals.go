package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"goals.com/models"
)

func getGoal(c *gin.Context) {
	var NewGoal models.Goal
	id := c.Param("id")
	coll := mgm.Coll(&NewGoal)
	_ = coll.FindByID(id, &NewGoal)
	if NewGoal.Name == "" {
		c.JSON(400, gin.H{"message": "goal not found"})
		return
	}
	c.JSON(200, gin.H{"message": "Goals successfully retrieved!", "goal": NewGoal})

}
func getGoals(c *gin.Context) {
	var goals []models.Goal
	err := mgm.Coll(&models.Goal{}).SimpleFind(&goals, bson.M{})
	if err != nil {
		c.JSON(400, gin.H{"message": "could not retrieve the goals", "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Goals successfully retrieved!", "goal": goals})
}

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

func deleteGoal(c *gin.Context) {
	var NewGoal models.Goal
	id := c.Param("id")
	fmt.Println(id)
	coll := mgm.Coll(&NewGoal)
	_ = coll.FindByID(id, &NewGoal)
	if NewGoal.Name == "" {
		c.JSON(400, gin.H{"message": "goal not found"})
		return
	}
	err := mgm.Coll(&NewGoal).Delete(&NewGoal)
	if err != nil {
		c.JSON(400, gin.H{"error": "could not delete data to database!"})
		return
	}
	c.JSON(200, gin.H{"message": NewGoal.Name + "Goal successfully Deleted!", "goal": NewGoal})
}
