package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterSwaggerRoutes(router *gin.Engine) {
	// Redirect /api-docs to /api-docs/index.html
	router.GET("/api-docs", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/api-docs/index.html")
	})
	// router.GET("/api-docs/", func(c *gin.Context) {
	// 	c.Redirect(http.StatusMovedPermanently, "/api-docs/index.html")
	// })

	router.GET("/api-docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
