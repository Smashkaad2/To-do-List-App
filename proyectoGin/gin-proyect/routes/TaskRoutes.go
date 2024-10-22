package routes

import (
	"github.com/gin-gonic/gin"
	"gin-proyect/controllers"
)

func TaskRouter(router *gin.Engine) {

	routes := router.Group("api/v1/tasks")
	routes.POST("", controllers.TaskCreate)
}