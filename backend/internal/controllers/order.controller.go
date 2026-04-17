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

type OrderController interface {
	Create(ginContext *gin.Context)
}

type orderControllerImpl struct {
	orderService services.OrderService
	cartService  services.CartService
}

func NewOrderController(db *sql.DB, cfg *config.Config) OrderController {
	return &orderControllerImpl{
		orderService: services.NewOrderService(db),
		cartService:  services.NewCartService(db),
	}
}

// CreateOrder godoc
// @Summary      Create a new order
// @Description  Creates an order from the user's cart with optional coupon
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request  body      adapters.CreateOrderRequest  true  "Create Order Request"
// @Success      200      {object}  object  "Order created successfully"
// @Failure      400      {object}  object  "Invalid request / cart empty / business error"
// @Failure      401      {object}  object  "Unauthorized"
// @Router       /api/v1/orders [post]
func (o *orderControllerImpl) Create(ginContext *gin.Context) {

	//1: get request-context and user
	reqContext := ginContext.Request.Context()
	user := middlewares.GetUserFromCTX(reqContext)

	//2: extract payload
	var req adapters.CreateOrderRequest
	if err := ginContext.ShouldBindJSON(&req); err != nil {
		utils.HandleErrorResponse(ginContext, 400, err.Error(), nil)
		return
	}

	//3: find user cart
	userCart, err := o.cartService.GetByUserID(reqContext, user.UserID)
	if err != nil || len(userCart.Items) == 0 {
		//return back if cart not found or empty cart
		utils.HandleErrorResponse(ginContext, 400, "Invalid cart", nil)
		return
	}

	//4: call service
	res, err := o.orderService.Create(reqContext, &req, userCart, user)

	if err != nil {
		utils.HandleErrorResponse(ginContext, 400, err.Error(), nil)
		return
	}

	utils.HandleSuccessResponse(ginContext, 200, "Order created successfully", res)

}
