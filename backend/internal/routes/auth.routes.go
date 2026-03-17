package routes

import (
	"database/sql"

	"github.com/shailendrapawar/book-store/internal/config"
	"github.com/shailendrapawar/book-store/internal/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup, db *sql.DB, cfg *config.Config) {

	//group routes for auth
	authRouter := r.Group("/auth")

	//init controller with passing db
	authController := controllers.NewAuthController(db, cfg)

	authRouter.POST("/register", authController.Register)
	authRouter.POST("/login", authController.Login)

}
