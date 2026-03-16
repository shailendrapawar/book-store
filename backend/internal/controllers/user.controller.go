package controllers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/shailendrapawar/book-store/internal/services"
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
// @Param        body  body  object  true  "User registration details"
// @Success      201   {object}  object
// @Failure      400   {object}  object
// @Router       /api/v1/users/register [post]
func (c *UserControllerImpl) Register(ginContext *gin.Context) {
	ginContext.JSON(200,
		gin.H{
			"hello": "keind",
		})
}
