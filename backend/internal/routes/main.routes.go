package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/shailendrapawar/book-store/docs"
)

func InitRoutes(r *gin.Engine, db *sql.DB) {

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	rootRouter := r.Group("/api/v1") //basic grouping

	//init all routes here=======>
	UserRoutes(rootRouter, db)
	AuthRoutes(rootRouter, db)
}
