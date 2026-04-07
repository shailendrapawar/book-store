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
	Create(ginContext *gin.Context)
	Get(ginContext *gin.Context)
	Update(ginContext *gin.Context)
	Search(ginContext *gin.Context)
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

// Create godoc
// @Summary      Create/Add a new book
// @Description  Create/Add a new book with the provided details to system
// @Tags         Books
// @Accept       json
// @Produce      json
// @Param        request  body      adapters.CreateBookRequest  true  "Create Book Request"
// @Success      200      {object}  adapters.Book
// @Failure      400      {object}  object
// @Router       /api/v1/books [post]
func (c *bookControllerImpl) Create(ginContext *gin.Context) {

	var req adapters.CreateBookRequest
	requestContext := ginContext.Request.Context()

	//validate request
	if err := ginContext.ShouldBindJSON(&req); err != nil {
		utils.HandleErrorResponse(ginContext, 400, err.Error(), nil)
		return
	}

	if utils.IsISBN(req.Isbn) == false {
		utils.HandleErrorResponse(ginContext, 400, "invalid ISBN number", nil)
		return
	}

	book, err := c.bookService.Create(requestContext, &req)
	if err != nil {
		utils.HandleErrorResponse(ginContext, 400, err.Error(), nil)
		return
	}

	utils.HandleSuccessResponse(ginContext, 200, "Book created successfully", book)
}

// Get godoc
// @Summary      Get a book by UUID or ISBN
// @Description  Retrieve a book using either a UUID (e.g. 550e8400-e29b-41d4-a716-446655440000)
// @Description  or an ISBN-10 (e.g. 0134190440) or ISBN-13 (e.g. 9780134190440)
// @Tags         Books
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Book UUID or ISBN-10/ISBN-13"
// @Success      200  {object}  adapters.Book
// @Failure      400  {object}  object
// @Router       /api/v1/books/{id} [get]
func (c *bookControllerImpl) Get(ginContext *gin.Context) {

	reqContext := ginContext.Request.Context()
	id := ginContext.Param("id")
	if id == "" || !utils.IsUUID(id) {
		utils.HandleErrorResponse(ginContext, 400, "Invalid id", nil)
		return
	}

	book, err := c.bookService.Get(reqContext, id)
	if err != nil {
		utils.HandleErrorResponse(ginContext, 400, err.Error(), nil)
		return
	}

	utils.HandleSuccessResponse(ginContext, 200, "Book Found", book)

}

// Update godoc
// @Summary      Update a book
// @Description  Update a book by id/UUID
// @Tags         Books
// @Accept       json
// @Produce      json
// @Param        id       path      string                    true  "Book UUID "
// @Param        request  body      adapters.UpdateBookRequest   true  "Update Book Request"
// @Success      200  {object}  adapters.Book
// @Failure      400  {object}  object
// @Router       /api/v1/books/{id} [put]
func (c *bookControllerImpl) Update(ginContext *gin.Context) {
	//basic validation
	reqContext := ginContext.Request.Context()
	id := ginContext.Param("id")
	if id == "" || !utils.IsUUID(id) {
		utils.HandleErrorResponse(ginContext, 400, "Invalid id", nil)
		return
	}

	var req adapters.UpdateBookRequest

	//validate request
	if err := ginContext.ShouldBindJSON(&req); err != nil {
		utils.HandleErrorResponse(ginContext, 400, err.Error(), nil)
		return
	}

	book, err := c.bookService.Update(reqContext, id, req)
	if err != nil {
		utils.HandleErrorResponse(ginContext, 400, err.Error(), nil)
		return
	}

	utils.HandleSuccessResponse(ginContext, 200, "Book Updated", book)

}

// Search godoc
// @Summary      Search Books
// @Description  Search books with optional pagination and keyword filtering
// @Tags         Books
// @Accept       json
// @Produce      json
// @Param        page     query     int     false  "Page number, default 1"
// @Param        limit    query     int     false  "Items per page, default 10"
// @Param        search   query     string  false  "Search keyword"
// @Success      200  {object}  adapters.SearchBooksResponse
// @Failure      400  {object}  object
// @Router       /api/v1/books [get]
func (c *bookControllerImpl) Search(ginContext *gin.Context) {
	reqContext := ginContext.Request.Context()

	//extract pagination data
	pagination := utils.HandlePagination(ginContext)

	res, err := c.bookService.Search(reqContext, pagination)

	if err != nil {
		utils.HandleErrorResponse(ginContext, 400, err.Error(), nil)
		return
	}

	utils.HandleSuccessResponse(ginContext, 200, "Books Found", res)
}
