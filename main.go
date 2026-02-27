package main

import (
	"go-crud-backend/config"
	"go-crud-backend/docs"
	"go-crud-backend/logger"
	"go-crud-backend/routes"

	"github.com/gin-gonic/gin"
)

// @title           Book Management API
// @version         1.0
// @description     A sample CRUD APIs for managing books.
// @BasePath        /
func main() {
	cfg := config.LoadConfig()

	// This overwrites the "@host" comment
	docs.SwaggerInfo.Host = "localhost:" + cfg.Port

	router := gin.Default()

	routes.RegisterBookRoutes(router)
	routes.RegisterSwaggerRoutes(router)

	logger.Info("Running in %s mode on port %s", cfg.Env, cfg.Port)
	router.Run(":" + cfg.Port)
}
