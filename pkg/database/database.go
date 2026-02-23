package database

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"
)

type DB struct {
	*sql.DB
	mu sync.Mutex
}

// NewDB creates a new in-memory database connection
func NewDB(databaseURL string) (*DB, error) {
	// Use in-memory map-based storage for simplicity
	db := &DB{}

	log.Println("Database connection established (in-memory)")
	return db, nil
}

// InitSchema creates the necessary tables
func (db *DB) InitSchema() error {
	log.Println("Database schema initialized (in-memory)")
	return nil
}

// Close closes the database connection
func (db *DB) Close() error {
	return nil
}

// In-memory storage
var (
	pricingRecords []PricingRecord
	uploadHistory  []UploadHistory
	recordsLock    sync.RWMutex
	historyLock    sync.RWMutex
	nextRecordID   int64 = 1
	nextHistoryID  int64 = 1
)

type PricingRecord struct {
	ID          int64
	StoreID     string
	SKU         string
	ProductName string
	Price       float64
	Date        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UploadHistory struct {
	ID           int64
	Filename     string
	UploadDate   time.Time
	Status       string
	RecordsCount int
	ErrorMessage string
}

// InsertRecords inserts records into in-memory storage
func InsertRecords(records []PricingRecord) error {
	recordsLock.Lock()
	defer recordsLock.Unlock()

	for _, record := range records {
		record.ID = nextRecordID
		nextRecordID++
		pricingRecords = append(pricingRecords, record)
	}

	return nil
}

// CreateUploadHistory creates an upload history record
func CreateUploadHistory(filename string) (int64, error) {
	historyLock.Lock()
	defer historyLock.Unlock()

	history := UploadHistory{
		ID:         nextHistoryID,
		Filename:   filename,
		UploadDate: time.Now(),
		Status:     "processing",
	}
	nextHistoryID++

	uploadHistory = append(uploadHistory, history)
	return history.ID, nil
}

// UpdateUploadHistory updates upload history
func UpdateUploadHistory(id int64, status string, recordsCount int, errorMessage string) error {
	historyLock.Lock()
	defer historyLock.Unlock()

	for i := range uploadHistory {
		if uploadHistory[i].ID == id {
			uploadHistory[i].Status = status
			uploadHistory[i].RecordsCount = recordsCount
			uploadHistory[i].ErrorMessage = errorMessage
			return nil
		}
	}

	return fmt.Errorf("upload history not found")
}

// GetRecordsCount returns total records count
func GetRecordsCount() int {
	recordsLock.RLock()
	defer recordsLock.RUnlock()
	return len(pricingRecords)
}
