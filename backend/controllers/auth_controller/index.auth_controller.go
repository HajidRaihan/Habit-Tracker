package auth_controller

import (
	"gin-gonic-gorm/database"
	"gin-gonic-gorm/models"
	"gin-gonic-gorm/requests"
	"gin-gonic-gorm/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Login(ctx *gin.Context) {
	loginReq := new(requests.LoginRequest)

	errReq := ctx.ShouldBind(&loginReq)

	if errReq != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errReq.Error()})
		return
	}

	user := new(models.User)
	err := database.DB.Table("users").Where("email = ?", loginReq.Email).First(&user).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	passIsvalid := utils.VerifyPassword(loginReq.Password, *user.Password)

	if !passIsvalid {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
		return
	}

	claims := jwt.MapClaims{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	}

	token, errToken := utils.GenerateToken(&claims)

	if errToken != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errToken.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Login success",
		"token":   token,
	})

}

func Register(ctx *gin.Context) {
	registerReq := new(requests.RegisterRequest)

	errReq := ctx.ShouldBind(&registerReq)

	if errReq != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errReq.Error()})
		return
	}

	user := new(models.User)
	password := registerReq.Password

	hashedPassword, errHash := utils.HashPassword(password)

	if errHash != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"message": "Failed to hash password", "error": errHash.Error()})
		return
	}

	user.Name = &registerReq.Name
	user.Email = &registerReq.Email
	user.Password = &hashedPassword

	if err := database.DB.Table("users").Create(&user).Error; err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"message": "Failed to create user"})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "User created successfully",
	})
}
