package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response represents a standardized API response
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// RespondWithSuccess sends a successful JSON response
func RespondWithSuccess(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, Response{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

// RespondWithError sends an error JSON response
func RespondWithError(c *gin.Context, code int, message string) {
	c.JSON(code, Response{
		Status: "error",
		Error:  message,
	})
}

// RespondWithValidationError sends a validation error response
func RespondWithValidationError(c *gin.Context, errors interface{}) {
	c.JSON(http.StatusBadRequest, Response{
		Status: "error",
		Error:  "Validation failed",
		Data:   errors,
	})
}

// RespondWithNotFound sends a 404 not found response
func RespondWithNotFound(c *gin.Context, message string) {
	if message == "" {
		message = "Resource not found"
	}
	RespondWithError(c, http.StatusNotFound, message)
}

// RespondWithInternalError sends a 500 internal server error response
func RespondWithInternalError(c *gin.Context, message string) {
	if message == "" {
		message = "Internal server error"
	}
	RespondWithError(c, http.StatusInternalServerError, message)
}
