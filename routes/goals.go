package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"goals.com/models"
)

func hello(context *gin.Context) {
	context.JSON(200, gin.H{"message": "hello from golang"})
}

func createGoal(c *gin.Context) {
	var goalReq models.Goal

	// Bind request body to goalReq and handle any errors
	if err := c.ShouldBindJSON(&goalReq); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	// Create a new goal from the request data
	goalData := models.NewGoal(goalReq.Description, goalReq.Name)

	// Save the goal to the database and handle potential errors
	if err := mgm.Coll(goalData).Create(goalData); err != nil {
		c.JSON(500, gin.H{"error": "Failed to save goal"})
		return
	}

	// Return a success response
	c.JSON(200, gin.H{"message": "Goal saved", "goal": goalData})
}

func getGoals(c *gin.Context) {
	var goal models.Goal
	var goals []models.Goal
	err := mgm.Coll(&goal).SimpleFind(&goals, bson.M{})
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"goal": goals})
}

func editGoals(c *gin.Context) {
	var goalReq models.Goal
	var goal models.Goal
	_ = c.ShouldBindBodyWith(&goalReq, binding.JSON)
	objId := primitive.ObjectID(goalReq.ID)

	err := mgm.Coll(&goal).FindByID(objId, &goal)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
	}

	goal.Name = goalReq.Name
	goal.Description = goalReq.Description

	err = mgm.Coll(&goalReq).Update(&goal)

	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "goal updated", "goal": goal})
}

func deleteGoals(c *gin.Context) {
	var goalReq models.Goal
	var goal models.Goal
	_ = c.ShouldBindBodyWith(&goalReq, binding.JSON)

	mgm.Coll(&goal).DeleteOne(mgm.Ctx(), bson.M{"_id": goalReq.ID})
	c.JSON(200, gin.H{"message": "goal deleted"})
}
