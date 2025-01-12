package routes

import (
	"gin-gonic-gorm/controllers/auth_controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(app *gin.Engine) {
	route := app

	route.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	authRoute := route.Group("auth")

	authRoute.POST("/register", auth_controller.Register)
	authRoute.POST("/login", auth_controller.Login)
}
