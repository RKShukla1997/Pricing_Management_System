package models

import "time"

// PricingRecord represents a pricing record in the system
type PricingRecord struct {
	ID          int64     `json:"id" db:"id"`
	StoreID     string    `json:"store_id" db:"store_id"`
	SKU         string    `json:"sku" db:"sku"`
	ProductName string    `json:"product_name" db:"product_name"`
	Price       float64   `json:"price" db:"price"`
	Date        string    `json:"date" db:"date"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// UploadHistory tracks file upload history
type UploadHistory struct {
	ID           int64     `json:"id" db:"id"`
	Filename     string    `json:"filename" db:"filename"`
	UploadDate   time.Time `json:"upload_date" db:"upload_date"`
	Status       string    `json:"status" db:"status"` // success, failed, processing
	RecordsCount int       `json:"records_count" db:"records_count"`
	ErrorMessage string    `json:"error_message,omitempty" db:"error_message"`
}

// UploadResponse represents the response for file upload
type UploadResponse struct {
	Message  string `json:"message"`
	Filename string `json:"filename"`
	Size     int64  `json:"size"`
	Rows     int    `json:"rows"`
	Columns  int    `json:"columns"`
	UploadID int64  `json:"upload_id,omitempty"`
}

// ErrorResponse represents error response structure
type ErrorResponse struct {
	Error string `json:"error"`
}

// HealthResponse represents health check response
type HealthResponse struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
	Database  string `json:"database,omitempty"`
}

// CSVRecord represents a parsed CSV record
type CSVRecord struct {
	StoreID     string
	SKU         string
	ProductName string
	Price       float64
	Date        string
	LineNumber  int
}

// ValidationError represents a validation error
type ValidationError struct {
	Line    int    `json:"line"`
	Field   string `json:"field"`
	Message string `json:"message"`
}
