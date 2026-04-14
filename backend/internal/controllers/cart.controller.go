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
	// Create(ginContext *gin.Context)
	Search(ginContext *gin.Context)
	Get(ginContext *gin.Context)
	GetByUserID(ginContext *gin.Context)

	AddItemToCart(ginContext *gin.Context)
	DeleteItemFromCart(ginContext *gin.Context)
	UpdateCartItem(ginContext *gin.Context)
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

// Search godoc
// @Summary      Search carts  (ADMIN)
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
func (c *cartControllerImpl) GetByUserID(ginContext *gin.Context) {
	requestContext := ginContext.Request.Context()
	user := middlewares.GetUserFromCTX(requestContext)

	res, err := c.cartService.GetByUserID(requestContext, user.UserID)
	if err != nil {
		utils.HandleErrorResponse(ginContext, 400, err.Error(), nil)
		return
	}

	utils.HandleSuccessResponse(ginContext, 200, "Cart Found", res)
}

// AddItemToCart godoc
// @Summary Add item to cart
// @Description Adds a book to user's cart. If item exists, quantity is updated.
// @Tags Cart-Items
// @Accept json
// @Produce json
// @Param request body adapters.AddItemToCartRequest true "Add Item Payload"
// @Success 200 {object} adapters.CartItem
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/carts/items [post]
func (c *cartControllerImpl) AddItemToCart(ginContext *gin.Context) {

	requestContext := ginContext.Request.Context()

	var req adapters.AddItemToCartRequest
	if err := ginContext.ShouldBindJSON(&req); err != nil {
		utils.HandleErrorResponse(ginContext, 400, err.Error(), nil)
		return
	}

	//call service
	user := middlewares.GetUserFromCTX(requestContext)
	res, err := c.cartService.AddItemToCart(requestContext, req, *user)

	if err != nil {
		utils.HandleErrorResponse(ginContext, 400, err.Error(), nil)
		return
	}

	utils.HandleSuccessResponse(ginContext, 200, "Item added to cart", res)

}

// DeleteItemFromCart godoc
// @Summary      Delete an item from the cart
// @Description  Delete an item from the cart by its ID
// @Tags         Cart-Items
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Item ID"
// @Success      200  {object} utils.SuccessResponse{data=nil} "Item deleted successfully"
// @Failure      400  {object} utils.ErrorResponse "Bad request or item not found"
// @Router       /api/v1/carts/items/{id} [delete]
func (c *cartControllerImpl) DeleteItemFromCart(ginContext *gin.Context) {

	reqContext := ginContext.Request.Context()
	id := ginContext.Param("id")
	user := middlewares.GetUserFromCTX(reqContext)

	res, err := c.cartService.DeleteItemFromCart(reqContext, id, *user)

	if err != nil {
		utils.HandleErrorResponse(ginContext, 400, err.Error(), nil)
		return
	}

	utils.HandleSuccessResponse(ginContext, 200, "Item deleted from cart", res)
}

// UpdateCartItem godoc
// @Summary Update cart item quantity
// @Description Update the quantity of a specific item using cart item ID
// @Tags Cart-Items
// @Accept json
// @Produce json
// @Param id path string true "Cart ID"
// @Param payload body adapters.UpdateCartItemRequest true "Update Cart Item Payload"
// @Success 200 {object} adapters.CartItem "Item updated successfully"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "Cart Item Not Found"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /api/v1/carts/items/{id} [put]
func (c *cartControllerImpl) UpdateCartItem(ginContext *gin.Context) {

	reqContext := ginContext.Request.Context()

	cartID := ginContext.Param("id")
	user := middlewares.GetUserFromCTX(reqContext)

	var req adapters.UpdateCartItemRequest
	if err := ginContext.ShouldBindJSON(&req); err != nil {
		utils.HandleErrorResponse(ginContext, 400, err.Error(), nil)
		return
	}

	payload := &adapters.UpdateCartItemPayload{
		CartID:   cartID,
		BookID:   req.BookID,
		Quantity: req.Quantity,
	}

	res, err := c.cartService.UpdateCartItem(reqContext, *payload, *user)

	if err != nil {
		utils.HandleErrorResponse(ginContext, 400, err.Error(), nil)
		return
	}

	utils.HandleSuccessResponse(ginContext, 200, "Item updated", res)

}
