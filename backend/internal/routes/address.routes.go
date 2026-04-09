package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/shailendrapawar/book-store/internal/config"
	"github.com/shailendrapawar/book-store/internal/controllers"
	"github.com/shailendrapawar/book-store/internal/middlewares"
)

func AddressRoutes(r *gin.RouterGroup, db *sql.DB, cfg *config.Config) {

	//group routes
	addressRouter := r.Group("/addresses")

	//insert common middleware
	addressRouter.Use(middlewares.AuthMiddleware(cfg))

	AddressController := controllers.NewAddressController(db, cfg)

	//actual routes
	addressRouter.POST("/", AddressController.Create)
	addressRouter.GET("/", AddressController.Search)
	addressRouter.GET("/:id", AddressController.Get)

}
