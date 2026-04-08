package controllers

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shailendrapawar/book-store/internal/config"
	"github.com/shailendrapawar/book-store/internal/middlewares"
	"github.com/shailendrapawar/book-store/internal/services"
	"github.com/shailendrapawar/book-store/internal/utils"
)

type CartController interface {
	Create(ginContext *gin.Context)
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

	fmt.Print("reached")

	requestContext := ginContext.Request.Context()
	user := middlewares.CurrentUser(ginContext)

	res, err := c.cartService.Create(requestContext, user.UserID)
	if err != nil {
		utils.HandleErrorResponse(ginContext, 400, err.Error(), nil)
		return
	}

	utils.HandleSuccessResponse(ginContext, 200, "Cart created", res)
}
