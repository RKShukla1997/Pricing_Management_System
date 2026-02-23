# Go CSV Ingestion Service - Test Instructions

## Running the Service

1. Start the server:
```bash
go run cmd/server/main.go
```

Or build and run:
```bash
go build -o product-service.exe cmd/server/main.go
.\product-service.exe
```

The server will start on port 8080 with the following endpoints:
- **Health Check**: `GET http://localhost:8080/health`
- **Upload CSV**: `POST http://localhost:8080/upload`

## Testing with curl

### 1. Health Check
```bash
curl http://localhost:8080/health
```

Expected Response:
```json
{
  "status": "healthy",
  "timestamp": "2026-02-20T13:00:00Z",
  "database": "connected (in-memory)",
  "record_count": 0
}
```

### 2. Upload CSV File
```bash
curl -X POST http://localhost:8080/upload \
  -F "file=@test.csv" \
  -H "Content-Type: multipart/form-data"
```

### PowerShell Alternative
```powershell
$form = @{
    file = Get-Item -Path "test.csv"
}
Invoke-RestMethod -Uri "http://localhost:8080/upload" -Method Post -Form $form
```

Expected Response:
```json
{
  "message": "File uploaded successfully",
  "filename": "1234567890_test.csv",
  "size": 1024,
  "rows": 101,
  "columns": 5,
  "upload_id": 1
}
```

### 3. Check Health After Upload
```bash
curl http://localhost:8080/health
```

Should show updated `record_count` with the number of records inserted.

## Testing with Postman/Insomnia

1. Create a new POST request to `http://localhost:8080/upload`
2. Go to Body → form-data
3. Add a key named `file` with type `File`
4. Select your CSV file
5. Send the request

## CSV Format

The CSV file must have the following structure:

```csv
Store ID,SKU,Product Name,Price,Date
ST001,SKU123,Product A,19.99,2024-01-15
ST002,SKU456,Product B,29.99,2024-01-16
```

### Required Columns:
1. **Store ID** - Max 20 characters
2. **SKU** - Max 50 characters
3. **Product Name** - Max 200 characters
4. **Price** - Decimal value between 0.01 and 999999.99
5. **Date** - Format: YYYY-MM-DD

## Features Demonstrated

✅ **CSV Upload** - Multipart form-data file upload
✅ **Streaming Validation** - Line-by-line CSV validation
✅ **Batch Processing** - Inserts 1000 records per batch
✅ **Error Handling** - Detailed validation error messages
✅ **In-Memory Storage** - Thread-safe concurrent data storage
✅ **Health Monitoring** - Live health checks with record counts
✅ **Middleware** - Logging, CORS, and recovery middleware
✅ **Clean Architecture** - Separation of concerns (handlers, services, models)

## Configuration

Edit `.env` file or set environment variables:

```
PORT=8080
DATABASE_URL=pricing.db
UPLOAD_PATH=./uploads
MAX_UPLOAD_SIZE=10485760
BATCH_SIZE=1000
MAX_GOROUTINES=10
LOG_LEVEL=info
```

## Architecture Notes

This implementation follows the architecture described in the README:
- **Event-driven design** - Ready to integrate with S3 event triggers
- **Streaming CSV parser** - Handles large files without loading entire file into memory
- **Batch inserts** - Efficient database operations with configurable batch size
- **Validation** - Comprehensive data validation with detailed error messages
- **Middleware** - Request logging, CORS support, panic recovery
