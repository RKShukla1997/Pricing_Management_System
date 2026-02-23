package services

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"product-service/internal/models"
)

// CSVValidator validates CSV files
type CSVValidator struct {
	requiredHeaders []string
}

// NewCSVValidator creates a new CSV validator
func NewCSVValidator() *CSVValidator {
	return &CSVValidator{
		requiredHeaders: []string{"Store ID", "SKU", "Product Name", "Price", "Date"},
	}
}

// ValidateHeaders checks if CSV has required headers
func (v *CSVValidator) ValidateHeaders(headers []string) error {
	if len(headers) != len(v.requiredHeaders) {
		return fmt.Errorf("expected %d columns, got %d", len(v.requiredHeaders), len(headers))
	}

	for i, required := range v.requiredHeaders {
		if strings.TrimSpace(headers[i]) != required {
			return fmt.Errorf("column %d: expected '%s', got '%s'", i+1, required, headers[i])
		}
	}

	return nil
}

// ValidateRecord validates a single CSV record
func (v *CSVValidator) ValidateRecord(record []string, lineNumber int) (*models.CSVRecord, error) {
	if len(record) != 5 {
		return nil, fmt.Errorf("line %d: expected 5 columns, got %d", lineNumber, len(record))
	}

	// Validate Store ID
	storeID := strings.TrimSpace(record[0])
	if storeID == "" {
		return nil, fmt.Errorf("line %d: Store ID cannot be empty", lineNumber)
	}
	if len(storeID) > 20 {
		return nil, fmt.Errorf("line %d: Store ID exceeds 20 characters", lineNumber)
	}

	// Validate SKU
	sku := strings.TrimSpace(record[1])
	if sku == "" {
		return nil, fmt.Errorf("line %d: SKU cannot be empty", lineNumber)
	}
	if len(sku) > 50 {
		return nil, fmt.Errorf("line %d: SKU exceeds 50 characters", lineNumber)
	}

	// Validate Product Name
	productName := strings.TrimSpace(record[2])
	if productName == "" {
		return nil, fmt.Errorf("line %d: Product Name cannot be empty", lineNumber)
	}
	if len(productName) > 200 {
		return nil, fmt.Errorf("line %d: Product Name exceeds 200 characters", lineNumber)
	}

	// Validate Price
	priceStr := strings.TrimSpace(record[3])
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return nil, fmt.Errorf("line %d: invalid price format: %s", lineNumber, priceStr)
	}
	if price < 0.01 || price > 999999.99 {
		return nil, fmt.Errorf("line %d: price must be between 0.01 and 999999.99", lineNumber)
	}

	// Validate Date (YYYY-MM-DD format)
	dateStr := strings.TrimSpace(record[4])
	if _, err := time.Parse("2006-01-02", dateStr); err != nil {
		return nil, fmt.Errorf("line %d: invalid date format (expected YYYY-MM-DD): %s", lineNumber, dateStr)
	}

	return &models.CSVRecord{
		StoreID:     storeID,
		SKU:         sku,
		ProductName: productName,
		Price:       price,
		Date:        dateStr,
		LineNumber:  lineNumber,
	}, nil
}

// ParseCSV parses and validates CSV file
func (v *CSVValidator) ParseCSV(reader io.Reader) ([]models.CSVRecord, []models.ValidationError, error) {
	csvReader := csv.NewReader(reader)
	csvReader.TrimLeadingSpace = true

	var records []models.CSVRecord
	var validationErrors []models.ValidationError
	lineNumber := 0

	for {
		record, err := csvReader.Read()
		lineNumber++

		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, nil, fmt.Errorf("error reading CSV at line %d: %w", lineNumber, err)
		}

		// First line should be headers
		if lineNumber == 1 {
			if err := v.ValidateHeaders(record); err != nil {
				return nil, nil, fmt.Errorf("invalid CSV headers: %w", err)
			}
			continue
		}

		// Validate record
		csvRecord, err := v.ValidateRecord(record, lineNumber)
		if err != nil {
			validationErrors = append(validationErrors, models.ValidationError{
				Line:    lineNumber,
				Field:   "record",
				Message: err.Error(),
			})
			continue
		}

		records = append(records, *csvRecord)
	}

	if len(records) == 0 {
		return nil, nil, fmt.Errorf("CSV file is empty or contains only headers")
	}

	return records, validationErrors, nil
}
