package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/shailendrapawar/book-store/internal/config"
	"github.com/shailendrapawar/book-store/internal/controllers"
	"github.com/shailendrapawar/book-store/internal/middlewares"
)

func UserRoutes(r *gin.RouterGroup, db *sql.DB, cfg *config.Config) {

	userRouter := r.Group("/users")
	//set middlewares
	userRouter.Use(middlewares.AuthMiddleware(cfg))

	//init controller with passing db
	userController := controllers.NewUserController(db, cfg)
	userRouter.GET("/:keyword", userController.Get)

}
