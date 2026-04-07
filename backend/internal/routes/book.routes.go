package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/shailendrapawar/book-store/internal/config"
	"github.com/shailendrapawar/book-store/internal/controllers"
	"github.com/shailendrapawar/book-store/internal/middlewares"
)

func BookRoutes(r *gin.RouterGroup, db *sql.DB, cfg *config.Config) {

	// group routes
	bookRouter := r.Group("/books")

	//insert middlewares
	bookRouter.Use(
		middlewares.AuthMiddleware(cfg),
		// add further middlewares if needed
	)

	// Initialize controller
	BookController := controllers.NewBookController(db, cfg)

	// insert actual routes
	bookRouter.POST("/", BookController.Create)
	bookRouter.GET("/:id", BookController.Get)
	bookRouter.GET("/", BookController.Search)
	bookRouter.PUT("/:id", BookController.Update)
}
