package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/shailendrapawar/book-store/docs"
	"github.com/shailendrapawar/book-store/internal/config"
)

func InitRoutes(r *gin.Engine, db *sql.DB, cfg *config.Config) {

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	rootRouter := r.Group("/api/v1") //basic grouping

	//init all routes here=======>
	UserRoutes(rootRouter, db, cfg)
	AuthRoutes(rootRouter, db, cfg)
	BookRoutes(rootRouter, db, cfg)
	CartRoutes(rootRouter, db, cfg)
}
