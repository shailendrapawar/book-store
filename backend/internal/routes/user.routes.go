package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/shailendrapawar/book-store/internal/config"
	"github.com/shailendrapawar/book-store/internal/controllers"
)

func UserRoutes(r *gin.RouterGroup, db *sql.DB, cfg *config.Config) {

	userRouter := r.Group("/users")

	//init controller with passing db
	userController := controllers.NewUserController(db, cfg)
	userRouter.GET("/:keyword", userController.Get)

}
