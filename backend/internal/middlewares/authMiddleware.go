package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/shailendrapawar/book-store/internal/config"
	"github.com/shailendrapawar/book-store/internal/utils"
)

// type AuthMiddleware struct {
// 	cfg *config.Config
// }

// func NewAuthMiddleware(cfg *config.Config) *AuthMiddleware {

// 	return &AuthMiddleware{
// 		cfg: cfg,
// 	}
// }

func AuthMiddleware(cfg *config.Config) gin.HandlerFunc {

	// ginContext.Next()

	return func(ctx *gin.Context) {
		// 1: get cookie
		tokenString, err := ctx.Cookie("token")
		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{
				"success": false,
				"message": "Unauthorised - no token",
			})
			return
		}

		// 2: if token found  validate it
		claims, err := utils.ValidateToken(tokenString, cfg.JWT.Secret)
		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{
				"success": false,
				"message": "Invalid token",
			})
			return
		}
		// 3: set user data in request
		ctx.Set("user", claims)
		ctx.Next()

	}
}
