package bootstrap

import (
	"gin-gonic-gorm/configs/cors_config"
	"gin-gonic-gorm/database"
	"gin-gonic-gorm/routes"

	"github.com/gin-gonic/gin"
)

func BootstrapApp() {
	// Initialize the bootstrap process
	app := gin.Default()

	database.ConnectDatabase()

	app.Use(cors_config.CorsConfigContrib())
	routes.InitRoutes(app)
	app.Run(":8000")
}
