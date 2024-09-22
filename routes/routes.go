package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/hello", hello)
	server.POST("/create-goal", createGoal)
	server.POST("/edit-goal", editGoals)
	server.DELETE("/delete-goal", deleteGoals)
	server.GET("/goals", getGoals)
}
