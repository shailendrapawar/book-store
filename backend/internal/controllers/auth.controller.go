package controllers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/shailendrapawar/book-store/internal/adapters"
	"github.com/shailendrapawar/book-store/internal/services"
	"github.com/shailendrapawar/book-store/internal/utils"
)

type AuthController interface {
	Register(ginContext *gin.Context)
	Login(ginContext *gin.Context)
}

type authControllerImpl struct {
	authService services.AuthService
}

func NewAuthController(db *sql.DB) AuthController {
	return &authControllerImpl{
		authService: services.NewAuthService(db),
		// userService:services.NewUserService(db)
	}
}

// @Summary      Register a new user
// @Description  Create a new user account
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        body  body  adapters.RegisterRequest  true  "Register payload"
// @Success      201   {object}  adapters.User
// @Failure      400   {object}  object
// @Failure      500   {object}  object
// @Router       /api/v1/auth/register [post]
func (c *authControllerImpl) Register(ginContext *gin.Context) {

	//1 validation
	var req adapters.RegisterRequest
	if err := ginContext.ShouldBindJSON(&req); err != nil {
		//validation error
		utils.HandleErrorResponse(ginContext, 400, err.Error(), nil)
	}

	user, err := c.authService.Register(ginContext.Request.Context(), req)

	if err != nil {
		utils.HandleErrorResponse(ginContext, 400, err.Error(), nil)
		return
	}
	utils.HandleSuccessResponse(ginContext, 201, "User Registered", user)
}

func (c *authControllerImpl) Login(ginContext *gin.Context) {

	var req *adapters.LoginRequest
	var requestContext = ginContext.Request.Context()

	if err := ginContext.ShouldBindJSON(&req); err != nil {
		// validation error
		utils.HandleErrorResponse(ginContext, 400, "Invalid payload", nil)
		return
	}

	// call service
	user, err := c.authService.Login(requestContext, req)

	if err != nil {
		//user dosent exists
		utils.HandleErrorResponse(ginContext, 404, err.Error(), nil)
		return
	}

	//return response
	utils.HandleSuccessResponse(ginContext, 201, "User Logged in", user)

}
