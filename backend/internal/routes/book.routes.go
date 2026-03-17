package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/shailendrapawar/book-store/internal/config"
	"github.com/shailendrapawar/book-store/internal/middlewares"
)

func BookRoutes(r *gin.RouterGroup, db *sql.DB, cfg *config.Config) {

	bookRouter := r.Group("/books")

	bookRouter.Use(
		middlewares.AuthMiddleware(cfg),
		// add further middlewares if needed
	)

}
