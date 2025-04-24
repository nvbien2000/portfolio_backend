package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck is a handler for the health check endpoint
func HealthCheck(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

// PingHandler is a handler for the ping endpoint
func PingHandler(c *gin.Context) {
	c.String(http.StatusOK, "Pong")
}
