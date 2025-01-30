package habit_logs_controller

import (
	"gin-gonic-gorm/database"
	"gin-gonic-gorm/models"
	"gin-gonic-gorm/requests"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAll(ctx *gin.Context) {
	log := new([]models.HabitLog)

	if err := database.DB.Table("habit_logs").Find(&log).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   log,
	})
}

func Create(ctx *gin.Context) {
	habitId := ctx.Param("id")

	log := new(models.HabitLog)

	logReq := new(requests.HabitLogRequest)
	log.ID = uuid.New()
	log.HabitID, _ = uuid.Parse(habitId)

	err := ctx.ShouldBind(&logReq)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.LogDate = time.Time(logReq.LogDate)
	log.Progress = logReq.Progress
	log.Status = logReq.Status

	if err := database.DB.Table("habit_logs").Create(&log).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Habit log created successfully",
		"data":    log,
	})

}
