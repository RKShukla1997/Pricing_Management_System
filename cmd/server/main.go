package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"product-service/internal/config"
	"product-service/internal/handlers"
	"product-service/internal/middleware"
	"product-service/internal/services"
	"product-service/pkg/database"
)

func main() {
	// Load configuration
	cfg := config.Load()
	cfg.Validate()

	// Initialize database
	db, err := database.NewDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize schema
	if err := db.InitSchema(); err != nil {
		log.Fatalf("Failed to initialize schema: %v", err)
	}

	// Initialize services
	validator := services.NewCSVValidator()
	ingestionService := services.NewIngestionService(cfg.BatchSize)

	// Initialize handlers
	uploadHandler := handlers.NewUploadHandler(cfg, validator, ingestionService)
	healthHandler := handlers.NewHealthHandler(ingestionService)

	// Setup Gin router
	if cfg.LogLevel == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	// Apply middleware
	router.Use(middleware.Recovery())
	router.Use(middleware.Logger())
	router.Use(middleware.CORS())

	// Setup routes
	router.GET("/health", healthHandler.HandleHealth)
	router.POST("/upload", uploadHandler.HandleUpload)

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		log.Println("Shutting down server...")
		os.Exit(0)
	}()

	// Start server
	log.Printf("Server starting on port %s...", cfg.Port)
	log.Printf("Upload endpoint: http://localhost:%s/upload", cfg.Port)
	log.Printf("Health check: http://localhost:%s/health", cfg.Port)

	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
