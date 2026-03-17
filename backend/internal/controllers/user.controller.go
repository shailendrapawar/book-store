package controllers

import (
	"database/sql"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/shailendrapawar/book-store/internal/services"
	"github.com/shailendrapawar/book-store/internal/utils"
)

type UserController interface {
	Get(ginContext *gin.Context)
}

type UserControllerImpl struct {
	userService services.UserService
}

func NewUserController(db *sql.DB) UserController {
	return &UserControllerImpl{
		userService: services.NewUserService(db),
	}
}

// @Summary      Get user by ID or email
// @Description  Get user by ID or email as identifier
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        keyword  path  string  true  "User ID or Email"
// @Success      200      {object}  adapters.User
// @Failure      400      {object}  object
// @Failure      404      {object}  object
// @Router       /api/v1/users/{keyword} [get]
func (c *UserControllerImpl) Get(ginContext *gin.Context) {

	//get and validate keyword
	keyword := ginContext.Param("keyword")

	if strings.Trim(keyword, "") == "" {
		utils.HandleErrorResponse(ginContext, 400, "Invalid keyword", nil)
		return
	}

	user, err := c.userService.Get(ginContext.Request.Context(), keyword)

	if err != nil {
		utils.HandleErrorResponse(ginContext, 400, err.Error(), nil)
		return
	}

	utils.HandleSuccessResponse(ginContext, 200, "User found", user)

}
