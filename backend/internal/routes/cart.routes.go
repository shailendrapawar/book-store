package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/shailendrapawar/book-store/internal/config"
	"github.com/shailendrapawar/book-store/internal/controllers"
	"github.com/shailendrapawar/book-store/internal/middlewares"
)

func CartRoutes(r *gin.RouterGroup, db *sql.DB, cfg *config.Config) {
	cartRouter := r.Group("/carts")

	//insert middlewares
	cartRouter.Use(
		middlewares.AuthMiddleware(cfg),
		// add further middlewares if needed
	)

	// Initialize controller
	CartController := controllers.NewCartController(db, cfg)

	// insert actual routes
	cartRouter.POST("/", CartController.Create)
	// cartRouter.GET("/", CartController.Search)
}
