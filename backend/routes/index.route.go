package routes

import (
	"gin-gonic-gorm/controllers/auth_controller"
	"gin-gonic-gorm/controllers/habit_controller"
	"gin-gonic-gorm/controllers/habit_logs_controller"
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
	logRoute := route.Group("log", middleware.AuthMiddleware)

	habitRoute.GET("/all", habit_controller.GetAllHabits)
	habitRoute.POST("/create", habit_controller.Create)
	habitRoute.POST("/update/:id", habit_controller.Update)

	logRoute.GET("/", habit_logs_controller.GetAll)
	logRoute.POST("/create/:id", habit_logs_controller.Create)

	authRoute.POST("/register", auth_controller.Register)
	authRoute.POST("/login", auth_controller.Login)
}
