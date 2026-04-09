package controllers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/shailendrapawar/book-store/internal/adapters"
	"github.com/shailendrapawar/book-store/internal/config"
	"github.com/shailendrapawar/book-store/internal/middlewares"
	"github.com/shailendrapawar/book-store/internal/services"
	"github.com/shailendrapawar/book-store/internal/utils"
)

type CartController interface {
	Create(ginContext *gin.Context)
	Search(ginContext *gin.Context)
	Get(ginContext *gin.Context)
	GetByID(ginContext *gin.Context)
}

type cartControllerImpl struct {
	cartService services.CartService
	cfg         *config.Config
}

func NewCartController(db *sql.DB, cfg *config.Config) CartController {
	return &cartControllerImpl{
		cartService: services.NewCartService(db),
		cfg:         cfg,
	}
}

// Create godoc
// @Summary      Create cart
// @Description  Creates a cart for the authenticated user
// @Tags         Cart
// @Accept       json
// @Produce      json
// @Success      200 {object} utils.SuccessResponse{data=adapters.Cart}
// @Failure      400 {object} utils.ErrorResponse{data=nil}
// @Router       /api/v1/carts [post]
func (c *cartControllerImpl) Create(ginContext *gin.Context) {

	requestContext := ginContext.Request.Context()

	res, err := c.cartService.Create(requestContext)
	if err != nil {
		utils.HandleErrorResponse(ginContext, 400, err.Error(), nil)
		return
	}

	utils.HandleSuccessResponse(ginContext, 200, "Cart created", res)
}

// Search godoc
// @Summary      Search carts
// @Description  Search carts with optional filters and pagination
// @Tags         Cart
// @Accept       json
// @Produce      json
// @Param        status   query     string  false  "Filter by cart status"
// @Param        page     query     int     false  "Page number for pagination"
// @Param        limit    query     int     false  "Number of items per page"
// @Success      200 {object} utils.SuccessResponse{data=[]adapters.Cart} "Successful search result"
// @Failure      400 {object} utils.ErrorResponse{data=nil} "Bad request"
// @Router       /api/v1/carts [get]
func (c *cartControllerImpl) Search(ginContext *gin.Context) {

	requestContext := ginContext.Request.Context()

	//1:extract pagination data
	pagination := utils.HandlePagination(ginContext)

	//2: extract filters
	var filters adapters.CartSearchFilters

	// user id
	if userId := ginContext.Query("userId"); userId != "" {
		filters.UserID = userId
	}

	// status
	if status := ginContext.Query("status"); status != "" {
		filters.Status = status
	}

	res, err := c.cartService.Search(requestContext, filters, pagination)

	if err != nil {
		utils.HandleErrorResponse(ginContext, 400, err.Error(), nil)
		return
	}

	utils.HandleSuccessResponse(ginContext, 200, "Cart Found", res)
}

// Get godoc
// @Summary      Get a cart by ID (ADMIN ONLY)
// @Description  Retrieve a single cart using its UUID
// @Tags         Cart
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Cart UUID"
// @Success      200  {object} utils.SuccessResponse{data=adapters.Cart} "Cart retrieved successfully"
// @Failure      400  {object} utils.ErrorResponse "Bad request or cart not found"
// @Router       /api/v1/carts/{id} [get]
func (c *cartControllerImpl) Get(ginContext *gin.Context) {

	requestContext := ginContext.Request.Context()
	id := ginContext.Param("id")

	res, err := c.cartService.Get(requestContext, id)
	if err != nil {
		utils.HandleErrorResponse(ginContext, 400, err.Error(), nil)
		return
	}

	utils.HandleSuccessResponse(ginContext, 200, "Cart Found", res)
}

// GetCartByUser godoc
// @Summary      Get cart for logged-in user
// @Description  Retrieve the cart associated with the authenticated user
// @Tags         Cart
// @Accept       json
// @Produce      json
// @Success      200  {object} utils.SuccessResponse{data=adapters.Cart} "Cart retrieved successfully"
// @Failure      400  {object} utils.ErrorResponse{data=nil} "Bad request or cart not found"
// @Router       /api/v1/carts/me [get]
func (c *cartControllerImpl) GetByID(ginContext *gin.Context) {
	requestContext := ginContext.Request.Context()
	user := middlewares.GetUserFromCTX(requestContext)

	res, err := c.cartService.GetByUserID(requestContext, user.UserID)
	if err != nil {
		utils.HandleErrorResponse(ginContext, 400, err.Error(), nil)
		return
	}

	utils.HandleSuccessResponse(ginContext, 200, "Cart Found", res)
}
