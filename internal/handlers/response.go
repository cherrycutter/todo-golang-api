package handlers

import (
	"github.com/cherrycutter/todo_app/pkg/logger"
	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

// newErrorResponse writes a message to logger and sends a JSON response containing the error message with the specified status code
func newErrorResponse(ctx *gin.Context, statusCode int, message string) {
	logger.Error.Println(message)
	ctx.JSON(statusCode, errorResponse{Message: message})
}
