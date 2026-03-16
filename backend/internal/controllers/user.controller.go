package controllers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/shailendrapawar/book-store/internal/adapters"
	"github.com/shailendrapawar/book-store/internal/services"
	"github.com/shailendrapawar/book-store/internal/utils"
)

type UserController interface {
	Register(ginContext *gin.Context)
}

type UserControllerImpl struct {
	userService services.UserService
}

func NewUserController(db *sql.DB) UserController {
	return &UserControllerImpl{
		userService: services.NewUserService(db),
	}
}

// @Summary      Register a new user
// @Description  Create a new user account
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        body  body  adapters.RegisterRequest  true  "Register payload"
// @Success      201   {object}  adapters.User
// @Failure      400   {object}  object
// @Failure      500   {object}  object
// @Router       /api/v1/users/register [post]
func (c *UserControllerImpl) Register(ginContext *gin.Context) {

	//1 validation
	var req adapters.RegisterRequest
	if err := ginContext.ShouldBindJSON(&req); err != nil {
		//validation error
		utils.HandleErrorResponse(ginContext, 400, err.Error(), nil)
	}

	user, err := c.userService.Create(ginContext.Request.Context(), req)

	if err != nil {
		utils.HandleErrorResponse(ginContext, 400, err.Error(), nil)
		return
	}
	utils.HandleSuccessResponse(ginContext, 201, "User Registered", user)
}
