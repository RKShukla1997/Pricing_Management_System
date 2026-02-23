package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"product-service/internal/models"
	"product-service/internal/services"
)

type HealthHandler struct {
	ingestionService *services.IngestionService
}

func NewHealthHandler(ingestionService *services.IngestionService) *HealthHandler {
	return &HealthHandler{
		ingestionService: ingestionService,
	}
}

// HandleHealth handles health check requests
func (h *HealthHandler) HandleHealth(c *gin.Context) {
	recordCount := h.ingestionService.GetRecordsCount()

	response := models.HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now().Format(time.RFC3339),
		Database:  "connected (in-memory)",
	}

	c.JSON(http.StatusOK, gin.H{
		"status":       response.Status,
		"timestamp":    response.Timestamp,
		"database":     response.Database,
		"record_count": recordCount,
	})
}
