package routes

import "github.com/gin-gonic/gin"

func hello(context *gin.Context) {
	context.JSON(200, gin.H{"message": "hello from golang"})
}
