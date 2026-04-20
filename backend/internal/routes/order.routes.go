package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/shailendrapawar/book-store/internal/config"
	"github.com/shailendrapawar/book-store/internal/controllers"
	"github.com/shailendrapawar/book-store/internal/middlewares"
)

func OrderRoutes(r *gin.RouterGroup, db *sql.DB, cfg *config.Config) {

	orderRouter := r.Group("/orders")
	//set middlewares
	orderRouter.Use(middlewares.AuthMiddleware(cfg))

	//init controller with passing db
	orderController := controllers.NewOrderController(db, cfg)
	orderRouter.POST("/", orderController.Create)
	orderRouter.GET("/:id", orderController.Get)

	orderRouter.GET("/me", orderController.GetUserOrder)

}
