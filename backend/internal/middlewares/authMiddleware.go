package middlewares

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/shailendrapawar/book-store/internal/adapters"
	"github.com/shailendrapawar/book-store/internal/config"
	"github.com/shailendrapawar/book-store/internal/utils"
)

const ContextUserKey = "user"

func AuthMiddleware(cfg *config.Config) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		// 1: get cookie
		tokenString, err := ctx.Cookie("token")
		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{
				"success": false,
				"message": "Unauthorized - no token",
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
		ctx.Set(ContextUserKey, claims)

		// 🔹 Inject into Go context.Request.Context
		reqCtx := context.WithValue(ctx.Request.Context(), ContextUserKey, claims)
		ctx.Request = ctx.Request.WithContext(reqCtx)

		ctx.Next()

	}
}

func GetUserFromCTX(ctx context.Context) *adapters.Claims {
	if ctx == nil {
		return nil
	}
	user, ok := ctx.Value(ContextUserKey).(*adapters.Claims)
	if !ok {
		return nil
	}
	return user
}
