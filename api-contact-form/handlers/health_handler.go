// Package handlers contains the HTTP handler implementations for various endpoints.
//
// The HealthHandler provides a health check endpoint to verify
// that the API is running correctly.
//
// Author: Tri Wicaksono
// Website: https://triwicaksono.com
package handlers

import (
	"net/http"

	"api-contact-form/responses"

	"github.com/gin-gonic/gin"
)

// HealthHandler handles HTTP requests related to health checks.
type HealthHandler struct{}

// NewHealthHandler creates and returns a new instance of HealthHandler.
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// HealthCheck responds with a JSON message confirming that the API is operational.
//
// Example Response:
// {
//     "code": "SUCCESS",
//     "message": "API is running."
// }
func (h *HealthHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, responses.APIResponse{
		Code:    "SUCCESS",
		Message: "API is running.",
	})
}
