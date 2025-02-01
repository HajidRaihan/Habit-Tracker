package routes

import (
	"gin-gonic-gorm/controllers/auth_controller"
	"gin-gonic-gorm/controllers/habit_controller"
	"gin-gonic-gorm/controllers/habit_logs_controller"
	"gin-gonic-gorm/controllers/reminder_controller"
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
	reminderRoute := route.Group("reminder", middleware.AuthMiddleware)

	habitRoute.GET("/all", habit_controller.GetAllHabits)
	habitRoute.POST("/create", habit_controller.Create)
	habitRoute.POST("/update/:id", habit_controller.Update)

	logRoute.GET("/", habit_logs_controller.GetAll)
	logRoute.GET("/:id", habit_logs_controller.GetById)
	logRoute.GET("/habit/:id", habit_logs_controller.GetLogByHabitsId)
	logRoute.POST("/create/:id", habit_logs_controller.Create)
	logRoute.PUT("/update/:id", habit_logs_controller.Update)
	logRoute.DELETE("/delete/:id", habit_logs_controller.Delete)

	reminderRoute.GET("/", reminder_controller.GetAll)
	reminderRoute.GET("/:id", reminder_controller.GetById)
	reminderRoute.GET("/habit/:id", reminder_controller.GetByHabitId)
	reminderRoute.POST("/create", reminder_controller.Create)
	reminderRoute.PUT("/update/:id", reminder_controller.Update)
	reminderRoute.DELETE("/delete/:id", reminder_controller.Delete)

	authRoute.POST("/register", auth_controller.Register)
	authRoute.POST("/login", auth_controller.Login)
}
