package utils

import "github.com/gin-gonic/gin"

// error response handler
func HandleErrorResponse(ctx *gin.Context, statusCode int, message string, data interface{}) interface{} {

	ctx.JSON(statusCode,
		gin.H{
			"success": false,
			"message": message,
			"data":    data,
		})

	return nil
}

//success repsonse ahndler

func HandleSuccessResponse(ctx *gin.Context, statusCode int, message string, data interface{}) interface{} {

	ctx.JSON(statusCode,
		gin.H{
			"success": true,
			"message": message,
			"data":    data,
		})

	return nil
}
