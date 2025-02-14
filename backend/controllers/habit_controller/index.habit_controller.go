package habit_controller

import (
	"encoding/json"
	"errors"
	"gin-gonic-gorm/database"
	"gin-gonic-gorm/models"
	"gin-gonic-gorm/requests"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetAllHabits(ctx *gin.Context) {
	habits := new([]models.Habit)
	err := database.DB.Table("habits").Find(&habits).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Habits retrieved successfully",
		"data":    habits,
	})
}

func GetDetailHabit(ctx *gin.Context) {
	habitId := ctx.Param("id")
	userIdStr := ctx.MustGet("user_id").(string)
	userId, err := uuid.Parse(userIdStr)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//verify that habits is users

	habit := new(models.Habit)
	err = database.DB.Table("habits").Where("id = ?", habitId).Where("user_id = ?", userId).First(&habit).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Habit retrieved successfully",
		"data":    habit,
	})

}

func Create(ctx *gin.Context) {
	userIdStr := ctx.MustGet("user_id").(string)

	userId, err := uuid.Parse(userIdStr)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized invalid user ID",
		})
		return
	}

	user := new(models.User)

	if err := database.DB.Table("users").Where("id = ?", userId).First(&user).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized user not found",
			"error":   err.Error(),
		})
		return
	}

	habitReq := new(requests.HabitRequest)

	errReq := ctx.ShouldBind(&habitReq)

	if errReq != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errReq.Error()})
		return
	}
	habit := new(models.Habit)

	habit.UserID = userId
	habit.Name = &habitReq.Name
	habit.Goal = &habitReq.Goal

	timeJson := json.RawMessage(habitReq.Time)
	habit.Time = &timeJson

	if err := database.DB.Table("habits").Create(&habit).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Habit created successfully",
	})
}

func Update(ctx *gin.Context) {
	userIdStr := ctx.MustGet("user_id").(string)
	habitId := ctx.Param("id")

	userId, _ := uuid.Parse(userIdStr)

	user := new(models.User)

	if err := database.DB.Table("users").Where("id = ?", userId).First(&user).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized user not found",
		})
		return
	}

	habit := new(models.Habit)

	if err := database.DB.Table("habits").Where("id = ?", habitId).First(&habit).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized habit not found",
		})
		return
	}

	habitReq := new(requests.HabitRequest)

	errReq := ctx.ShouldBind(&habitReq)

	if errReq != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errReq.Error()})
		return
	}

	habit.Name = &habitReq.Name
	habit.Goal = &habitReq.Goal

	timeJson := json.RawMessage(habitReq.Time)
	habit.Time = &timeJson

	if err := database.DB.Table("habits").Where("id = ?", habitId).Updates(&habit).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Habit updated successfully",
	})
}

func Delete(ctx *gin.Context) {
	// Ambil user_id dari context (setelah autentikasi)
	userIdStr := ctx.MustGet("user_id").(string)
	habitId := ctx.Param("id")

	// Parse user_id menjadi UUID
	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user ID",
		})
		return
	}

	// Periksa apakah habit ada dan milik user tersebut
	habit := new(models.Habit)
	if err := database.DB.Table("habits").Where("id = ?", habitId).Where("user_id = ?", userId).First(habit).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "Habit not found",
			})
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to fetch habit",
			})
		}
		return
	}

	// Hapus habit dari database
	result := database.DB.Table("habits").Where("id = ?", habitId).Where("user_id = ?", userId).Delete(&models.Habit{})
	if result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to delete habit",
		})
		return
	}

	// Periksa apakah habit benar-benar dihapus
	if result.RowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Habit not found",
		})
		return
	}

	// Kirim response
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Habit deleted successfully",
	})
}
