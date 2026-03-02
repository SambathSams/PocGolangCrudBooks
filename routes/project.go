package routes

import (
	"go-crud-backend/controllers" // Adjust path to your project name

	"github.com/gin-gonic/gin"
)

func RegisterProjectRoutes(router *gin.Engine) {
	projectRoutes := router.Group("/projects")
	{
		projectRoutes.GET("", controllers.GetProjects)
		projectRoutes.GET("/:id", controllers.GetProjectByID)
		projectRoutes.POST("", controllers.CreateProject)
		projectRoutes.PUT("/:id", controllers.UpdateProject)
		projectRoutes.DELETE("/:id", controllers.DeleteProject)
	}
}
