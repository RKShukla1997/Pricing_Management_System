package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"

	"product-service/internal/config"
	"product-service/internal/models"
	"product-service/internal/services"
)

type UploadHandler struct {
	cfg              *config.Config
	validator        *services.CSVValidator
	ingestionService *services.IngestionService
}

func NewUploadHandler(cfg *config.Config, validator *services.CSVValidator, ingestionService *services.IngestionService) *UploadHandler {
	return &UploadHandler{
		cfg:              cfg,
		validator:        validator,
		ingestionService: ingestionService,
	}
}

// HandleUpload handles CSV file upload
func (h *UploadHandler) HandleUpload(c *gin.Context) {
	// Set max memory
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, h.cfg.MaxUploadSize)

	// Get file from form
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Error retrieving file from form data",
		})
		return
	}

	// Validate file size
	if fileHeader.Size > h.cfg.MaxUploadSize {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: fmt.Sprintf("File too large. Maximum size is %d bytes", h.cfg.MaxUploadSize),
		})
		return
	}

	// Validate file extension
	if filepath.Ext(fileHeader.Filename) != ".csv" {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Only CSV files are allowed",
		})
		return
	}

	// Open uploaded file
	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Error opening uploaded file",
		})
		return
	}
	defer file.Close()

	// Parse and validate CSV
	records, validationErrors, err := h.validator.ParseCSV(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: fmt.Sprintf("Invalid CSV file: %v", err),
		})
		return
	}

	// If there are validation errors, return them
	if len(validationErrors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":             "CSV validation failed",
			"validation_errors": validationErrors,
		})
		return
	}

	// Generate unique filename
	timestamp := time.Now().Unix()
	filename := fmt.Sprintf("%d_%s", timestamp, fileHeader.Filename)
	destPath := filepath.Join(h.cfg.UploadPath, filename)

	// Save file to disk (for audit/backup)
	file.Seek(0, 0)
	dst, err := os.Create(destPath)
	if err != nil {
		log.Printf("Error creating file: %v", err)
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Error saving file",
		})
		return
	}
	defer dst.Close()

	bytesWritten, err := io.Copy(dst, file)
	if err != nil {
		log.Printf("Error copying file: %v", err)
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Error saving file",
		})
		return
	}

	// Process CSV and insert into database
	uploadID, err := h.ingestionService.ProcessCSVFile(filename, records)
	if err != nil {
		log.Printf("Error processing CSV: %v", err)
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: fmt.Sprintf("Error processing CSV: %v", err),
		})
		return
	}

	log.Printf("File uploaded successfully: %s (size: %d bytes, rows: %d)", filename, bytesWritten, len(records))

	// Send success response
	c.JSON(http.StatusOK, models.UploadResponse{
		Message:  "File uploaded successfully",
		Filename: filename,
		Size:     bytesWritten,
		Rows:     len(records) + 1, // +1 for header
		Columns:  5,
		UploadID: uploadID,
	})
}
