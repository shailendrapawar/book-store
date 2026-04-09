package controllers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/shailendrapawar/book-store/internal/adapters"
	"github.com/shailendrapawar/book-store/internal/config"
	"github.com/shailendrapawar/book-store/internal/services"
	"github.com/shailendrapawar/book-store/internal/utils"
)

type AddressController interface {
	Create(ginContext *gin.Context)
}

type addressControllerImpl struct {
	addressService services.AddressService
	cfg            *config.Config
}

func NewAddressController(db *sql.DB, cfg *config.Config) AddressController {
	return &addressControllerImpl{
		addressService: services.NewAddressService(db),
		cfg:            cfg,
	}
}

// Create godoc
// @Summary      Create a new address
// @Description  Create a new address for the logged-in user
// @Tags         Address
// @Accept       json
// @Produce      json
// @Param        request  body      adapters.CreateAddressRequest  true  "Address Payload"
// @Success      201      {object}  utils.SuccessResponse{data=adapters.Address}
// @Failure      400      {object}  utils.ErrorResponse
// @Router       /api/v1/addresses [post]
func (c *addressControllerImpl) Create(ginContext *gin.Context) {

	reqContext := ginContext.Request.Context()

	var req adapters.CreateAddressRequest
	if err := ginContext.ShouldBindJSON(&req); err != nil {
		utils.HandleErrorResponse(ginContext, 400, err.Error(), nil)
		return
	}

	address, err := c.addressService.Create(reqContext, &req)
	if err != nil {
		utils.HandleErrorResponse(ginContext, 400, err.Error(), nil)
		return
	}

	utils.HandleSuccessResponse(ginContext, 201, "Address Created", address)

}
