package bootstrap

import (
	"gin-gonic-gorm/database"
	"gin-gonic-gorm/routes"

	"github.com/gin-gonic/gin"
)

func BootstrapApp() {
	// Initialize the bootstrap process
	app := gin.Default()

	database.ConnectDatabase()
	routes.InitRoutes(app)
	app.Run(":8000")
}
