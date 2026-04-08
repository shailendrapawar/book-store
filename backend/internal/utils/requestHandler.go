package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shailendrapawar/book-store/internal/adapters"
)

type SuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// error response handler
func HandleErrorResponse(ctx *gin.Context, statusCode int, message string, data interface{}) interface{} {

	ctx.JSON(statusCode, ErrorResponse{
		Success: false,
		Message: message,
		Data:    data,
	})

	return nil
}

//success repsonse ahndler

func HandleSuccessResponse(ctx *gin.Context, statusCode int, message string, data interface{}) interface{} {

	ctx.JSON(statusCode, SuccessResponse{
		Success: true,
		Message: message,
		Data:    data,
	})

	return nil
}

func HandlePagination(ctx *gin.Context) adapters.PaginationRequest {

	pagination := adapters.PaginationRequest{}

	page := ctx.DefaultQuery("page", "1")
	if page == "0" {
		page = "1"
	}
	limit := ctx.DefaultQuery("limit", "10")

	//covert to numeric
	pagination.Page, _ = strconv.Atoi(page)
	pagination.Limit, _ = strconv.Atoi(limit)
	pagination.Offset = (pagination.Page - 1) * pagination.Limit

	return pagination
}
