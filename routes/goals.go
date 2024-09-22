package routes

import (
	"fmt"

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
	var goal models.Goal
	_ = c.ShouldBindBodyWith(&goalReq, binding.JSON)
	fmt.Println(goalReq)
	goalData := models.NewGoal(goalReq.Description, goalReq.Name)
	mgm.Coll(&goal).Create(goalData)
	c.JSON(200, gin.H{"message": "goal saved", "goal": goalData})
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
