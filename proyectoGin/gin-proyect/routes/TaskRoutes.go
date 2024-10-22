package routes

import (
	"github.com/gin-gonic/gin"
	"gin-proyect/controllers"
)

func TaskRouter(router *gin.Engine) {

	routes := router.Group("api/v1/tasks")
	routes.POST("", controllers.TaskCreate)
	routes.GET("", controllers.TaskGet)
	routes.GET("/:id", controllers.TaskGetById)
	routes.PUT("/:id", controllers.TaskUpdate)
	routes.DELETE("/:id", controllers.TaskDelete)
}