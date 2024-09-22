package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {

	server.POST("/goals", createGoal)
	server.GET("/goals/:id", getGoal)
	server.GET("/goals", getGoals)
}
