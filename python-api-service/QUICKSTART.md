# Quick Start Guide - Python FastAPI Service

## 🚀 Setup and Run (5 minutes)

### Step 1: Install Python Dependencies

```powershell
# Navigate to Python service directory
cd python-api-service

# Create virtual environment
python -m venv venv

# Activate virtual environment
.\venv\Scripts\activate  # Windows PowerShell

# Install dependencies
pip install -r requirements.txt
```

### Step 2: Create .env file

```powershell
# Copy example env file
Copy-Item .env.example .env

# Edit .env if needed (optional - defaults work for testing)
```

### Step 3: Run the Server

```powershell
# Start FastAPI server
uvicorn app.main:app --reload --port 8000
```

Server will be running at: **http://localhost:8000**

### Step 4: Test the APIs

Open your browser:
- **API Docs**: http://localhost:8000/docs (Interactive Swagger UI)
- **Health Check**: http://localhost:8000/health

---

## 🧪 API Testing Examples

### Test 1: Generate Presigned URL

```powershell
$body = @{
    filename = "test.csv"
    content_type = "text/csv"
} | ConvertTo-Json

$response = Invoke-RestMethod -Uri "http://localhost:8000/api/upload/presigned-url" `
    -Method Post -Body $body -ContentType "application/json"

Write-Host "Presigned URL Response:" -ForegroundColor Green
$response | ConvertTo-Json
```

### Test 2: Create a Pricing Record

```powershell
$body = @{
    store_id = "ST001"
    sku = "SKU-TEST-001"
    product_name = "Test Product"
    price = 19.99
    date = "2024-03-03"
} | ConvertTo-Json

$response = Invoke-RestMethod -Uri "http://localhost:8000/api/pricing/records" `
    -Method Post -Body $body -ContentType "application/json"

Write-Host "Created Record:" -ForegroundColor Green
$response | ConvertTo-Json
```

### Test 3: Search Records

```powershell
# Search all records
$response = Invoke-RestMethod -Uri "http://localhost:8000/api/pricing/search"
Write-Host "Search Results:" -ForegroundColor Green
$response | ConvertTo-Json -Depth 5

# Search by store
$response = Invoke-RestMethod -Uri "http://localhost:8000/api/pricing/search?store_id=ST001"

# Search with price range
$response = Invoke-RestMethod -Uri "http://localhost:8000/api/pricing/search?min_price=10&max_price=50"

# Search with pagination
$response = Invoke-RestMethod -Uri "http://localhost:8000/api/pricing/search?page=1&page_size=20"
```

### Test 4: Update a Record

```powershell
# Update record ID 1
$body = @{
    price = 24.99
    product_name = "Updated Product Name"
} | ConvertTo-Json

$response = Invoke-RestMethod -Uri "http://localhost:8000/api/pricing/records/1" `
    -Method Put -Body $body -ContentType "application/json"

Write-Host "Updated Record:" -ForegroundColor Green
$response | ConvertTo-Json
```

### Test 5: Get Specific Record

```powershell
# Get record by ID
$response = Invoke-RestMethod -Uri "http://localhost:8000/api/pricing/records/1"
Write-Host "Record Details:" -ForegroundColor Green
$response | ConvertTo-Json
```

### Test 6: Delete a Record

```powershell
$response = Invoke-RestMethod -Uri "http://localhost:8000/api/pricing/records/1" -Method Delete
Write-Host "Delete Response:" -ForegroundColor Green
$response | ConvertTo-Json
```

---

## 📝 Complete Test Script

Save this as `test-api.ps1`:

```powershell
Write-Host "`n=== Testing Python FastAPI Service ===" -ForegroundColor Cyan

# Test health check
Write-Host "`n1. Testing Health Check..." -ForegroundColor Yellow
try {
    $health = Invoke-RestMethod -Uri "http://localhost:8000/health"
    Write-Host "✓ Server is healthy" -ForegroundColor Green
    $health | ConvertTo-Json
} catch {
    Write-Host "✗ Server not running!" -ForegroundColor Red
    exit 1
}

# Test presigned URL generation
Write-Host "`n2. Testing Presigned URL Generation..." -ForegroundColor Yellow
$urlBody = @{
    filename = "test_upload.csv"
    content_type = "text/csv"
} | ConvertTo-Json

try {
    $presigned = Invoke-RestMethod -Uri "http://localhost:8000/api/upload/presigned-url" `
        -Method Post -Body $urlBody -ContentType "application/json"
    Write-Host "✓ Presigned URL generated" -ForegroundColor Green
    $presigned | ConvertTo-Json
} catch {
    Write-Host "✗ Failed to generate presigned URL" -ForegroundColor Red
}

# Create test records
Write-Host "`n3. Creating Test Records..." -ForegroundColor Yellow
$testRecords = @(
    @{
        store_id = "ST001"
        sku = "TEST-001"
        product_name = "Test Product 1"
        price = 19.99
        date = "2024-03-01"
    },
    @{
        store_id = "ST002"
        sku = "TEST-002"
        product_name = "Test Product 2"
        price = 29.99
        date = "2024-03-02"
    },
    @{
        store_id = "ST001"
        sku = "TEST-003"
        product_name = "Test Product 3"
        price = 39.99
        date = "2024-03-03"
    }
)

$createdIds = @()
foreach ($record in $testRecords) {
    try {
        $body = $record | ConvertTo-Json
        $response = Invoke-RestMethod -Uri "http://localhost:8000/api/pricing/records" `
            -Method Post -Body $body -ContentType "application/json"
        $createdIds += $response.id
        Write-Host "✓ Created record ID: $($response.id)" -ForegroundColor Green
    } catch {
        Write-Host "✗ Failed to create record" -ForegroundColor Red
    }
}

# Test search - all records
Write-Host "`n4. Testing Search - All Records..." -ForegroundColor Yellow
try {
    $search = Invoke-RestMethod -Uri "http://localhost:8000/api/pricing/search"
    Write-Host "✓ Found $($search.total) records" -ForegroundColor Green
    Write-Host "Items on page: $($search.items.Count)"
} catch {
    Write-Host "✗ Search failed" -ForegroundColor Red
}

# Test search - by store
Write-Host "`n5. Testing Search - By Store..." -ForegroundColor Yellow
try {
    $search = Invoke-RestMethod -Uri "http://localhost:8000/api/pricing/search?store_id=ST001"
    Write-Host "✓ Found $($search.total) records for ST001" -ForegroundColor Green
} catch {
    Write-Host "✗ Store search failed" -ForegroundColor Red
}

# Test search - price range
Write-Host "`n6. Testing Search - Price Range..." -ForegroundColor Yellow
try {
    $search = Invoke-RestMethod -Uri "http://localhost:8000/api/pricing/search?min_price=20&max_price=40"
    Write-Host "✓ Found $($search.total) records in price range" -ForegroundColor Green
} catch {
    Write-Host "✗ Price range search failed" -ForegroundColor Red
}

# Test update
if ($createdIds.Count -gt 0) {
    Write-Host "`n7. Testing Update..." -ForegroundColor Yellow
    $updateBody = @{
        price = 99.99
        product_name = "Updated Test Product"
    } | ConvertTo-Json
    
    try {
        $updated = Invoke-RestMethod -Uri "http://localhost:8000/api/pricing/records/$($createdIds[0])" `
            -Method Put -Body $updateBody -ContentType "application/json"
        Write-Host "✓ Updated record ID: $($createdIds[0])" -ForegroundColor Green
        Write-Host "New price: $($updated.price)"
    } catch {
        Write-Host "✗ Update failed" -ForegroundColor Red
    }
}

# Test get by ID
if ($createdIds.Count -gt 0) {
    Write-Host "`n8. Testing Get by ID..." -ForegroundColor Yellow
    try {
        $record = Invoke-RestMethod -Uri "http://localhost:8000/api/pricing/records/$($createdIds[0])"
        Write-Host "✓ Retrieved record ID: $($record.id)" -ForegroundColor Green
    } catch {
        Write-Host "✗ Get by ID failed" -ForegroundColor Red
    }
}

# Test delete
if ($createdIds.Count -gt 0) {
    Write-Host "`n9. Testing Delete..." -ForegroundColor Yellow
    try {
        $response = Invoke-RestMethod -Uri "http://localhost:8000/api/pricing/records/$($createdIds[0])" `
            -Method Delete
        Write-Host "✓ Deleted record ID: $($createdIds[0])" -ForegroundColor Green
    } catch {
        Write-Host "✗ Delete failed" -ForegroundColor Red
    }
}

Write-Host "`n=== All Tests Complete ===" -ForegroundColor Cyan
Write-Host "Visit http://localhost:8000/docs for interactive API documentation" -ForegroundColor Yellow
```

Run with:
```powershell
.\test-api.ps1
```

---

## 🔗 Integration with Go Service

Both services can run simultaneously:

- **Python API**: http://localhost:8000 (User-facing CRUD & Search)
- **Go Ingestion**: http://localhost:8080 (CSV Processing)

They share the same `pricing.db` SQLite database.

### Full Upload Flow:

1. Call Python API: `POST /api/upload/presigned-url`
2. Get presigned URL
3. Upload CSV to S3 (or mock endpoint)
4. S3 triggers Go ingestion service
5. Query Python API: `GET /api/pricing/search` to see imported data

---

## 📊 API Endpoints Summary

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Health check |
| POST | `/api/upload/presigned-url` | Generate upload URL |
| GET | `/api/upload/status/{id}` | Check upload status |
| POST | `/api/pricing/records` | Create record |
| GET | `/api/pricing/records/{id}` | Get record by ID |
| PUT | `/api/pricing/records/{id}` | Update record |
| DELETE | `/api/pricing/records/{id}` | Delete record |
| GET | `/api/pricing/search` | Search with filters |

---

## ✅ Success Indicators

You should see:
- ✅ Server starts without errors
- ✅ Database file `pricing.db` created
- ✅ Swagger UI accessible at `/docs`
- ✅ All test endpoints return 200 OK
- ✅ Mock presigned URLs generated
- ✅ CRUD operations work
- ✅ Search with filters returns results

---

## 🐛 Troubleshooting

### Error: "Module not found"
```powershell
pip install -r requirements.txt
```

### Error: "Port already in use"
```powershell
# Use different port
uvicorn app.main:app --reload --port 8001
```

### Error: "Database locked"
```powershell
# Close Go service if running on same database
# Or use different database files
```

---

**Happy Testing! 🎉**
