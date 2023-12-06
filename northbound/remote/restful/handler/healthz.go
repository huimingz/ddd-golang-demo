package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthyHandler health checks for the server.
// guide: https://tools.ietf.org/id/draft-inadarei-api-health-check-01.html#rfc.section.3
func HealthyHandler(ctx *gin.Context) {
	ctx.Header("Cache-Control", "max-age=3600")
	ctx.Header("Content-Type", "application/health+json")

	ctx.JSON(http.StatusOK, gin.H{
		"status":      "pass",
		"version":     "1",
		"releaseID":   "1.0.0",
		"description": "health of scan-management service",
	})
}
