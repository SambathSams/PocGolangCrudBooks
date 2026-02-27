package routes

import (
	"go-crud-backend/controllers" // Adjust path to your project name

	"github.com/gin-gonic/gin"
)

func RegisterBookRoutes(router *gin.Engine) {
	bookRoutes := router.Group("/books")
	{
		bookRoutes.GET("", controllers.GetBooks)
		bookRoutes.GET("/:id", controllers.GetBookByID)
		bookRoutes.POST("", controllers.CreateBook)
		bookRoutes.PUT("/:id", controllers.UpdateBook)
		bookRoutes.DELETE("/:id", controllers.DeleteBook)
	}
}
