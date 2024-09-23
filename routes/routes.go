package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {

	server.POST("/goals", createGoal)
	server.GET("/goal/:id", getGoal)
	server.GET("/goals", getGoals)
	server.DELETE("/goal/:id",deleteGoal)
	server.PUT("/goal/:id",updateGoal)
}
