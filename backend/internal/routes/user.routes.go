package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/shailendrapawar/book-store/internal/controllers"
)

func UserRoutes(r *gin.RouterGroup, db *sql.DB) {

	userRouter := r.Group("/users")

	//init controller with passing db
	userController := controllers.NewUserController(db)
	userRouter.GET("/:keyword", userController.Get)

}
