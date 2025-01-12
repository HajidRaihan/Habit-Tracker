package routes

import (
	"gin-gonic-gorm/controllers/auth_controller"
	"gin-gonic-gorm/controllers/habit_controller"
	"gin-gonic-gorm/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoutes(app *gin.Engine) {
	route := app.Group("/api")

	route.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	authRoute := route.Group("auth")
	habitRoute := route.Group("habit", middleware.AuthMiddleware)

	habitRoute.GET("/all", habit_controller.GetAllHabits)
	habitRoute.POST("/create", habit_controller.Create)

	authRoute.POST("/register", auth_controller.Register)
	authRoute.POST("/login", auth_controller.Login)
}
