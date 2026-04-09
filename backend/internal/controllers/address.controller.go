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
	Search(ginContext *gin.Context)
	Get(ginContext *gin.Context)
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

// Search godoc
// @Summary      Search addresses
// @Description  Search addresses with filters like user, state, city, district, and deletion status
// @Tags         Address
// @Accept       json
// @Produce      json
// @Param        userID     query     string  false  "User ID"
// @Param        state      query     string  false  "State"     example(Uttarakhand)
// @Param        city       query     string  false  "City"      example(Haldwani)
// @Param        district   query     string  false  "District"  example(nainital)
// @Param        isDeleted  query     bool    false  "Is Deleted" example(false)
// @Param        page       query     int     false  "Page number"     example(1)
// @Param        limit      query     int     false  "Items per page"  example(10)
// @Success      200        {object}  utils.SuccessResponse{data=[]adapters.Address}
// @Failure      400        {object}  utils.ErrorResponse
// @Failure      401        {object}  utils.ErrorResponse
// @Router       /api/v1/addresses [get]
func (c *addressControllerImpl) Search(ginContext *gin.Context) {

	requestContext := ginContext.Request.Context()

	pagination := utils.HandlePagination(ginContext)

	var filters adapters.SearchAddressFilters

	if userID := ginContext.Query("userID"); userID != "" {
		filters.UserID = &userID
	}
	if state := ginContext.Query("state"); state != "" {
		filters.State = &state
	}
	if city := ginContext.Query("city"); city != "" {
		filters.City = &city
	}
	if district := ginContext.Query("district"); district != "" {
		filters.District = &district
	}
	if isDeleted := ginContext.Query("isDeleted"); isDeleted != "" {
		val := false
		if isDeleted == "true" {
			val = true
			filters.IsDeleted = &val
		}
	}

	addresses, err := c.addressService.Search(requestContext, filters, pagination)
	if err != nil {
		utils.HandleErrorResponse(ginContext, 400, err.Error(), nil)
		return
	}

	utils.HandleSuccessResponse(ginContext, 200, "Addresses Fetched", addresses)

}

// Get godoc
// @Summary      Get address by ID
// @Description  Retrieve a single address by its UUID
// @Tags         Address
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Address UUID"  format(uuid)
// @Success      200  {object}  utils.SuccessResponse{data=adapters.Address}
// @Failure      400  {object}  utils.ErrorResponse
// @Failure      404  {object}  utils.ErrorResponse
// @Failure      401  {object}  utils.ErrorResponse
// @Router       /api/v1/addresses/{id} [get]
func (c *addressControllerImpl) Get(ginContext *gin.Context) {

	reqContext := ginContext.Request.Context()
	id := ginContext.Param("id")
	if id == "" || !utils.IsUUID(id) {
		utils.HandleErrorResponse(ginContext, 400, "Invalid id", nil)
		return
	}

	address, err := c.addressService.Get(reqContext, id)
	if err != nil {
		utils.HandleErrorResponse(ginContext, 400, err.Error(), nil)
		return
	}

	utils.HandleSuccessResponse(ginContext, 200, "Address Found", address)
}
