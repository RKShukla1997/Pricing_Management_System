package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	maxUploadSize = 10 << 20 // 10 MB
	uploadPath    = "./uploads"
)

// UploadResponse represents the response structure for successful uploads
type UploadResponse struct {
	Message  string `json:"message"`
	Filename string `json:"filename"`
	Size     int64  `json:"size"`
	Rows     int    `json:"rows"`
	Columns  int    `json:"columns"`
}

// HealthResponse represents the health check response
type HealthResponse struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}

// ErrorResponse represents error response structure
type ErrorResponse struct {
	Error string `json:"error"`
}

func main() {
	// Create uploads directory if it doesn't exist
	if err := os.MkdirAll(uploadPath, os.ModePerm); err != nil {
		log.Fatalf("Failed to create upload directory: %v", err)
	}

	// Set Gin to release mode (optional, comment out for debug mode)
	// gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	// Routes
	router.GET("/health", healthCheckHandler)
	router.POST("/upload", uploadCSVHandler)

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s...", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func healthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

func uploadCSVHandler(c *gin.Context) {
	// Set max memory for multipart forms (10MB)
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxUploadSize)

	// Get the file from form data
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "Error retrieving file from form data",
		})
		return
	}

	// Validate file size
	if fileHeader.Size > maxUploadSize {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "File too large. Maximum size is 10MB",
		})
		return
	}

	// Validate file extension
	if filepath.Ext(fileHeader.Filename) != ".csv" {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "Only CSV files are allowed",
		})
		return
	}

	// Open the uploaded file
	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "Error opening uploaded file",
		})
		return
	}
	defer file.Close()

	// Validate CSV content
	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "Invalid CSV file format",
		})
		return
	}

	if len(records) == 0 {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "CSV file is empty",
		})
		return
	}

	// Reset file pointer to beginning for saving
	file.Seek(0, 0)

	// Generate unique filename with timestamp
	timestamp := time.Now().Unix()
	filename := fmt.Sprintf("%d_%s", timestamp, fileHeader.Filename)
	destPath := filepath.Join(uploadPath, filename)

	// Create the destination file
	dst, err := os.Create(destPath)
	if err != nil {
		log.Printf("Error creating file: %v", err)
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: "Error saving file",
		})
		return
	}
	defer dst.Close()

	// Copy the uploaded file to destination
	bytesWritten, err := io.Copy(dst, file)
	if err != nil {
		log.Printf("Error copying file: %v", err)
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: "Error saving file",
		})
		return
	}

	log.Printf("File uploaded successfully: %s (size: %d bytes, rows: %d)", filename, bytesWritten, len(records))

	// Send success response
	c.JSON(http.StatusOK, UploadResponse{
		Message:  "File uploaded successfully",
		Filename: filename,
		Size:     bytesWritten,
		Rows:     len(records),
		Columns:  len(records[0]),
	})
}
