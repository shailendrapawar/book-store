package controllers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/shailendrapawar/book-store/internal/adapters"
	"github.com/shailendrapawar/book-store/internal/config"
	"github.com/shailendrapawar/book-store/internal/services"
	"github.com/shailendrapawar/book-store/internal/utils"
)

type BookController interface {
}

type bookControllerImpl struct {
	bookService services.BookService
	cfg         *config.Config
}

func NewBookController(db *sql.DB, cfg *config.Config) BookController {
	return &bookControllerImpl{
		bookService: services.NewBookService(db),
		cfg:         cfg,
	}
}

func (c *bookControllerImpl) Create(ginContext *gin.Context) {

	var req adapters.CreateBookRequest
	requestContext := ginContext.Request.Context()

	//validate request
	if err := ginContext.ShouldBindJSON(&req); err != nil {
		utils.HandleErrorResponse(ginContext, 400, err.Error(), nil)
		return
	}

	book, err := c.bookService.Create(requestContext, &req)
	if err != nil {
		utils.HandleErrorResponse(ginContext, 400, err.Error(), nil)
		return
	}

	utils.HandleSuccessResponse(ginContext, 200, "Book created successfully", book)
}
