package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shailendrapawar/book-store/internal/config"
	"github.com/shailendrapawar/book-store/internal/db"
	"github.com/shailendrapawar/book-store/internal/routes"
)

// @title            Bookstore API
// @version         1.0
// @description     REST API for  Bookstore
// @host            localhost:8080
// @BasePath        /
func main() {

	// 1: init DB connection
	cfg := config.Load()
	dbConn := db.ConnectDB(cfg)

	//1.2:close connection when app is shutdown
	defer dbConn.Close()

	// 2: create gin server
	router := gin.Default()

	//3: init routes
	routes.InitRoutes(router, dbConn, cfg)

	//add health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "server working",
			"success": true,
			"status":  "ok",
		})
	})

	log.Println("Server starting on port " + cfg.App.Port)
	if err := router.Run(":" + cfg.App.Port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

}
