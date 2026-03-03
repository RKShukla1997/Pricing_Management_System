package database

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "modernc.org/sqlite"
)

type DB struct {
	*sql.DB
	mu sync.Mutex
}

var globalDB *DB

// NewDB creates a new SQLite database connection
func NewDB(databaseURL string) (*DB, error) {
	sqlDB, err := sql.Open("sqlite", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Test the connection
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	db := &DB{DB: sqlDB}
	globalDB = db

	log.Printf("Database connection established: %s", databaseURL)
	return db, nil
}

// InitSchema creates the necessary tables (matching Python's SQLAlchemy schema)
func (db *DB) InitSchema() error {
	schema := `
	CREATE TABLE IF NOT EXISTS pricing_records (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		store_id TEXT NOT NULL,
		sku TEXT NOT NULL,
		product_name TEXT NOT NULL,
		price REAL NOT NULL,
		date TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE INDEX IF NOT EXISTS idx_store_id ON pricing_records(store_id);
	CREATE INDEX IF NOT EXISTS idx_sku ON pricing_records(sku);
	CREATE INDEX IF NOT EXISTS idx_date ON pricing_records(date);

	CREATE TABLE IF NOT EXISTS upload_history (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		filename TEXT NOT NULL,
		upload_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		status TEXT NOT NULL,
		records_count INTEGER DEFAULT 0,
		error_message TEXT
	);

	CREATE TABLE IF NOT EXISTS audit_logs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		action TEXT NOT NULL,
		record_id INTEGER,
		user_id TEXT,
		timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		details TEXT
	);
	`

	_, err := db.Exec(schema)
	if err != nil {
		return fmt.Errorf("failed to initialize schema: %w", err)
	}

	log.Println("Database schema initialized")
	return nil
}

// Close closes the database connection
func (db *DB) Close() error {
	if db.DB != nil {
		return db.DB.Close()
	}
	return nil
}

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

// InsertRecords inserts records into SQLite database
func InsertRecords(records []PricingRecord) error {
	if globalDB == nil || globalDB.DB == nil {
		return fmt.Errorf("database connection not initialized")
	}

	tx, err := globalDB.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`
		INSERT INTO pricing_records (store_id, sku, product_name, price, date, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	now := time.Now().Format("2006-01-02T15:04:05.000000")
	for _, record := range records {
		_, err := stmt.Exec(
			record.StoreID,
			record.SKU,
			record.ProductName,
			record.Price,
			record.Date,
			now,
			now,
		)
		if err != nil {
			return fmt.Errorf("failed to insert record: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	log.Printf("Inserted %d records into database", len(records))
	return nil
}

// CreateUploadHistory creates an upload history record
func CreateUploadHistory(filename string) (int64, error) {
	if globalDB == nil || globalDB.DB == nil {
		return 0, fmt.Errorf("database connection not initialized")
	}

	now := time.Now().Format("2006-01-02T15:04:05.000000")
	result, err := globalDB.Exec(`
		INSERT INTO upload_history (filename, upload_date, status, records_count)
		VALUES (?, ?, ?, ?)
	`, filename, now, "processing", 0)
	if err != nil {
		return 0, fmt.Errorf("failed to create upload history: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get insert id: %w", err)
	}

	return id, nil
}

// UpdateUploadHistory updates upload history
func UpdateUploadHistory(id int64, status string, recordsCount int, errorMessage string) error {
	if globalDB == nil || globalDB.DB == nil {
		return fmt.Errorf("database connection not initialized")
	}

	_, err := globalDB.Exec(`
		UPDATE upload_history 
		SET status = ?, records_count = ?, error_message = ?
		WHERE id = ?
	`, status, recordsCount, errorMessage, id)
	if err != nil {
		return fmt.Errorf("failed to update upload history: %w", err)
	}

	return nil
}

// GetRecordsCount returns total records count
func GetRecordsCount() int {
	if globalDB == nil || globalDB.DB == nil {
		return 0
	}

	var count int
	err := globalDB.QueryRow("SELECT COUNT(*) FROM pricing_records").Scan(&count)
	if err != nil {
		log.Printf("Failed to get records count: %v", err)
		return 0
	}

	return count
}
