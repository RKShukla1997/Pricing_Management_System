package services

import (
	"fmt"
	"log"
	"time"

	"product-service/internal/models"
	"product-service/pkg/database"
)

// IngestionService handles CSV ingestion and database operations
type IngestionService struct {
	batchSize int
}

// NewIngestionService creates a new ingestion service
func NewIngestionService(batchSize int) *IngestionService {
	return &IngestionService{
		batchSize: batchSize,
	}
}

// CreateUploadHistory creates an upload history record
func (s *IngestionService) CreateUploadHistory(filename string) (int64, error) {
	return database.CreateUploadHistory(filename)
}

// UpdateUploadHistory updates the upload history record
func (s *IngestionService) UpdateUploadHistory(id int64, status string, recordsCount int, errorMessage string) error {
	return database.UpdateUploadHistory(id, status, recordsCount, errorMessage)
}

// BatchInsertRecords inserts records in batches
func (s *IngestionService) BatchInsertRecords(records []models.CSVRecord) error {
	now := time.Now()
	inserted := 0
	batch := []database.PricingRecord{}

	for _, record := range records {
		dbRecord := database.PricingRecord{
			StoreID:     record.StoreID,
			SKU:         record.SKU,
			ProductName: record.ProductName,
			Price:       record.Price,
			Date:        record.Date,
			CreatedAt:   now,
			UpdatedAt:   now,
		}
		batch = append(batch, dbRecord)
		inserted++

		// Insert batch
		if len(batch) >= s.batchSize {
			if err := database.InsertRecords(batch); err != nil {
				return fmt.Errorf("failed to insert batch: %w", err)
			}
			log.Printf("Inserted batch of %d records", len(batch))
			batch = []database.PricingRecord{}
		}
	}

	// Insert remaining records
	if len(batch) > 0 {
		if err := database.InsertRecords(batch); err != nil {
			return fmt.Errorf("failed to insert final batch: %w", err)
		}
		log.Printf("Inserted final batch of %d records", len(batch))
	}

	log.Printf("Successfully inserted %d records", inserted)
	return nil
}

// ProcessCSVFile processes the entire CSV file
func (s *IngestionService) ProcessCSVFile(filename string, records []models.CSVRecord) (int64, error) {
	// Create upload history
	uploadID, err := s.CreateUploadHistory(filename)
	if err != nil {
		return 0, err
	}

	// Insert records
	if err := s.BatchInsertRecords(records); err != nil {
		s.UpdateUploadHistory(uploadID, "failed", 0, err.Error())
		return uploadID, err
	}

	// Update upload history
	if err := s.UpdateUploadHistory(uploadID, "success", len(records), ""); err != nil {
		return uploadID, err
	}

	return uploadID, nil
}

// GetRecordsCount returns the total number of records
func (s *IngestionService) GetRecordsCount() int {
	return database.GetRecordsCount()
}
