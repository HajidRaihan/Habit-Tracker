package cors_config

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CORS Hanlder dengan gin
func CorsConfig(ctx *gin.Context) {

	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept, Authorization, X-Requested-With")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(http.StatusNoContent)
		return
	}

	ctx.Next()
}

var origins = []string{
	"*",
}

// CORS Hanlder dengan Contrib
func CorsConfigContrib() gin.HandlerFunc {
	config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	config.AllowOrigins = origins

	return cors.New(config)
}
