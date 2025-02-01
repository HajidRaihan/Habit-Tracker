package reminder_controller

import (
	"gin-gonic-gorm/database"
	"gin-gonic-gorm/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAll(ctx *gin.Context) {
	var reminder []models.Reminder

	err := database.DB.Table("reminders").Find(&reminder).Error

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success get reminder",
		"data":    reminder,
	})
}

func Create(ctx *gin.Context) {
	var reminder models.Reminder

	if err := ctx.ShouldBind(&reminder); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := time.Parse("15:04", reminder.ReminderTime)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Table("reminders").Create(&reminder).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Reminder created successfully",
		"data":    reminder,
	})

}

func GetById(ctx *gin.Context) {
	reminderId := ctx.Param("id")

	var reminder models.Reminder

	if err := database.DB.Table("reminders").Where("id = ?", reminderId).First(&reminder).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success get reminder",
		"data":    reminder,
	})
}

func GetByHabitId(ctx *gin.Context) {
	habitId := ctx.Param("id")
	var reminder []models.Reminder

	if err := database.DB.Table("reminders").Where("habit_id = ?", habitId).Find(&reminder).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success get reminder by habit id",
		"data":    reminder,
	})
}

func Update(ctx *gin.Context) {
	reminderId := ctx.Param("id")
	var reminder models.Reminder

	err := ctx.ShouldBind(&reminder)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := database.DB.Table("reminders").Where("id = ?", reminderId).Updates(&reminder)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update reminder",
		})
		return
	}

	if result.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "reminder not found",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "reminder updated successfully",
	})
}

func Delete(ctx *gin.Context) {
	reminderId := ctx.Param("id")

	// Hapus reminder dari database
	result := database.DB.Table("reminders").Where("id = ?", reminderId).Delete(&models.Reminder{})
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete reminder",
		})
		return
	}

	// Periksa apakah reminder benar-benar dihapus
	if result.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Reminder not found",
		})
		return
	}

	// Kirim response
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Reminder deleted successfully",
	})
}
