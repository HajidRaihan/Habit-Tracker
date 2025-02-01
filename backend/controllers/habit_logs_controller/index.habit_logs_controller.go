package habit_logs_controller

import (
	"errors"
	"gin-gonic-gorm/database"
	"gin-gonic-gorm/models"
	"gin-gonic-gorm/requests"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetAll(ctx *gin.Context) {
	log := new([]models.HabitLog)

	if err := database.DB.Table("habit_logs").Find(log).Error; err != nil {
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

	if err := database.DB.Table("habit_logs").Create(log).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Habit log created successfully",
		"data":    log,
	})

}

func GetById(ctx *gin.Context) {
	habitLogId := ctx.Param("id")

	// log := new(models.HabitLog)
	var log models.HabitLog

	if err := database.DB.Table("habit_logs").Where("id = ?", habitLogId).First(&log).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Habit log not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    log,
	})

}

func GetLogByHabitsId(ctx *gin.Context) {
	habitId := ctx.Param("id")

	var log []models.HabitLog

	err := database.DB.Table("habit_logs").Where("habit_id = ?", habitId).Find(&log).Error

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Habit log not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    log,
	})

}

func Update(ctx *gin.Context) {
	habitLogId := ctx.Param("id")

	var log models.HabitLog

	if err := ctx.ShouldBind(&log); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := database.DB.Table("habit_logs").Where("id = ?", habitLogId).Updates(&log)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update log",
		})
		return
	}

	if result.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "log not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Habit log updated successfully",
	})
}

func Delete(ctx *gin.Context) {
	habitLogId := ctx.Param("id")

	err := database.DB.Table("habit_logs").Where("id = ?", habitLogId).Delete(&models.HabitLog{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "Habit log not found",
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to delete habit log",
			})
		}
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Habit log deleted successfully",
	})

}
