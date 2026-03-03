# Architecture Integration - Go + Python Services

## 👤 Frontend User Flow - File Upload

### Overview
The file upload process uses a **presigned URL pattern** for secure, scalable uploads directly to S3 storage, with asynchronous background processing.

### Step-by-Step Flow

```
┌─────────────────────────────────────────────────────────────────────┐
│                         FRONTEND APPLICATION                        │
│                      (React/Angular/Vue/CLI)                        │
└─────────────────────────────────────────────────────────────────────┘
```

#### **Phase 1: Request Upload Permission**

**1. User Action**: User selects CSV file in browser/app
   - File: `pricing-data-march-2024.csv`
   - File validation (optional): Check size, extension, format

**2. Frontend Request**: Get presigned URL from Python API
```javascript
// Example: JavaScript/TypeScript frontend code
const response = await fetch('http://localhost:8000/api/upload/presigned-url', {
  method: 'POST',
  headers: { 'Content-Type': 'application/json' },
  body: JSON.stringify({ 
    filename: 'pricing-data-march-2024.csv' 
  })
});

const uploadInfo = await response.json();
// Response: {
//   "upload_url": "http://localhost:8000/mock-upload/abc-123-def",
//   "file_key": "uploads/20240303_120000_abc-123-def_pricing-data-march-2024.csv",
//   "expires_in": 3600,
//   "upload_id": "abc-123-def-456"
// }
```

**3. Backend Processing** (Python API Service - Port 8000):
   - Validates filename
   - Generates unique file key: `uploads/{timestamp}_{uuid}_{filename}`
   - Creates presigned URL (expires in 1 hour)
   - Returns upload credentials to frontend

---

#### **Phase 2: Direct Upload to Storage**

**4. Frontend Upload**: Upload file directly to S3 using presigned URL
```javascript
// Frontend uploads directly to storage (bypasses backend)
const uploadResponse = await fetch(uploadInfo.upload_url, {
  method: 'PUT',
  body: selectedFile,  // The actual CSV file
  headers: {
    'Content-Type': 'text/csv'
  }
});

if (uploadResponse.ok) {
  console.log('File uploaded successfully!');
  // Store upload_id for status tracking
  localStorage.setItem('uploadId', uploadInfo.upload_id);
}
```

**Why Direct Upload?**
- ✅ Reduces load on backend servers
- ✅ Faster upload (direct to storage)
- ✅ Secure (presigned URL expires)
- ✅ Scalable (no backend bottleneck)
- ✅ Cost-effective (no data transfer through backend)

---

#### **Phase 3: Background Processing (Asynchronous)**

**5. Event Trigger**: S3 sends notification to Go Ingestion Service
   - S3 Event: `ObjectCreated:Put`
   - Notification method: Webhook / SNS / SQS / Lambda
   - Event payload includes: bucket, file_key, timestamp

**6. Go Service Processing** (Port 8080):
```go
// Go service receives event and processes file
// - Downloads file from S3
// - Validates CSV structure and data types
// - Streams large files (memory efficient)
// - Parses records line by line
// - Validates business rules (price > 0, required fields)
// - Batch inserts to database (1000 records/batch)
// - Updates upload_history table with status
```

**Processing Steps**:
   - ✅ Download CSV from S3
   - ✅ Validate CSV structure (headers: store_id, sku, product_name, price, date)
   - ✅ Stream parse (handles large files)
   - ✅ Validate each record
   - ✅ Batch insert to PostgreSQL (1000 records/batch)
   - ✅ Update status: `pending` → `processing` → `completed` / `failed`

---

#### **Phase 4: Status Tracking & Results**

**7. Frontend Status Check**: Poll for processing status
```javascript
// Option 1: Polling (check every 5 seconds)
const checkStatus = setInterval(async () => {
  const status = await fetch(
    `http://localhost:8000/api/upload/status/${uploadInfo.upload_id}`
  );
  const data = await status.json();
  
  if (data.status === 'completed') {
    console.log(`✅ Processed ${data.records_processed} records`);
    clearInterval(checkStatus);
    // Refresh the pricing data grid
    loadPricingData();
  } else if (data.status === 'failed') {
    console.error(`❌ Upload failed: ${data.error_message}`);
    clearInterval(checkStatus);
  } else {
    console.log(`⏳ Processing... ${data.status}`);
  }
}, 5000);

// Option 2: WebSocket (real-time updates)
const ws = new WebSocket('ws://localhost:8000/ws/upload-status');
ws.onmessage = (event) => {
  const status = JSON.parse(event.data);
  updateProgressBar(status.progress);
};
```

**8. View Results**: Search and display imported data
```javascript
// Search for newly imported records
const results = await fetch(
  'http://localhost:8000/api/pricing/search?date_from=2024-03-03&sort_by=created_at&sort_order=desc'
);
const pricingData = await results.json();

// Display in data grid/table
renderDataGrid(pricingData.items);
```

---

### Complete Frontend Flow Diagram

```
┌──────────────────────────────────────────────────────────────────┐
│  USER ACTIONS                                                    │
└──────────────────────────────────────────────────────────────────┘
    │
    │ 1️⃣ Select CSV file
    ▼
┌──────────────────────────────────────────────────────────────────┐
│  FRONTEND (React/Angular/Vue)                                    │
│  - File validation (size, type)                                  │
│  - User feedback (loading spinner)                               │
└──────────────────────────────────────────────────────────────────┘
    │
    │ 2️⃣ POST /api/upload/presigned-url
    │    { filename: "pricing.csv" }
    ▼
┌──────────────────────────────────────────────────────────────────┐
│  PYTHON API (Port 8000)                                          │
│  ✅ Generate presigned URL                                       │
│  ✅ Create upload_id                                             │
│  ✅ Set expiration (1 hour)                                      │
└──────────────────────────────────────────────────────────────────┘
    │
    │ 3️⃣ Returns: { upload_url, upload_id, expires_in }
    ▼
┌──────────────────────────────────────────────────────────────────┐
│  FRONTEND                                                        │
│  - Display upload progress bar                                   │
│  - PUT file to presigned URL                                     │
└──────────────────────────────────────────────────────────────────┘
    │
    │ 4️⃣ Direct upload (no backend involved)
    ▼
┌──────────────────────────────────────────────────────────────────┐
│  S3 / CLOUD STORAGE                                              │
│  ✅ File stored: uploads/20240303_120000_uuid_pricing.csv        │
│  ✅ Triggers event notification                                  │
└──────────────────────────────────────────────────────────────────┘
    │
    │ 5️⃣ Event: ObjectCreated
    ▼
┌──────────────────────────────────────────────────────────────────┐
│  GO INGESTION SERVICE (Port 8080)                                │
│  ✅ Download file from S3                                        │
│  ✅ Validate CSV structure                                       │
│  ✅ Stream parse (memory efficient)                              │
│  ✅ Validate each record                                         │
│  ✅ Batch insert (1000/batch)                                    │
│  ✅ Update upload_history status                                 │
└──────────────────────────────────────────────────────────────────┘
    │
    │ 6️⃣ Processing complete
    ▼
┌──────────────────────────────────────────────────────────────────┐
│  POSTGRESQL DATABASE                                             │
│  ✅ pricing_records table (10,000 new rows)                      │
│  ✅ upload_history table (status: completed)                     │
│  ✅ audit_logs table (processing events)                         │
└──────────────────────────────────────────────────────────────────┘
    │
    │ 7️⃣ Poll status or receive WebSocket update
    ▼
┌──────────────────────────────────────────────────────────────────┐
│  FRONTEND                                                        │
│  GET /api/upload/status/{upload_id}                              │
│  - Show success message                                          │
│  - Display record count                                          │
│  - Refresh data grid                                             │
└──────────────────────────────────────────────────────────────────┘
    │
    │ 8️⃣ View imported data
    ▼
┌──────────────────────────────────────────────────────────────────┐
│  PYTHON API (Port 8000)                                          │
│  GET /api/pricing/search?date_from=2024-03-03                    │
│  - Returns paginated results                                     │
│  - User can search, filter, sort                                 │
└──────────────────────────────────────────────────────────────────┘
```

---

### Frontend Integration Examples

#### **React Example**

```typescript
// components/FileUpload.tsx
import React, { useState } from 'react';

export const FileUpload: React.FC = () => {
  const [file, setFile] = useState<File | null>(null);
  const [uploading, setUploading] = useState(false);
  const [uploadId, setUploadId] = useState<string>('');
  const [status, setStatus] = useState<string>('');

  const handleFileSelect = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files && e.target.files[0]) {
      setFile(e.target.files[0]);
    }
  };

  const uploadFile = async () => {
    if (!file) return;
    
    setUploading(true);
    
    try {
      // Step 1: Get presigned URL
      const urlResponse = await fetch('http://localhost:8000/api/upload/presigned-url', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ filename: file.name })
      });
      
      const { upload_url, upload_id } = await urlResponse.json();
      setUploadId(upload_id);
      
      // Step 2: Upload file directly to S3
      await fetch(upload_url, {
        method: 'PUT',
        body: file,
        headers: { 'Content-Type': 'text/csv' }
      });
      
      // Step 3: Poll for status
      pollStatus(upload_id);
      
    } catch (error) {
      console.error('Upload failed:', error);
      setUploading(false);
    }
  };

  const pollStatus = (id: string) => {
    const interval = setInterval(async () => {
      const response = await fetch(`http://localhost:8000/api/upload/status/${id}`);
      const data = await response.json();
      
      setStatus(data.status);
      
      if (data.status === 'completed' || data.status === 'failed') {
        clearInterval(interval);
        setUploading(false);
      }
    }, 3000);
  };

  return (
    <div className="upload-container">
      <input type="file" accept=".csv" onChange={handleFileSelect} />
      <button onClick={uploadFile} disabled={!file || uploading}>
        {uploading ? 'Uploading...' : 'Upload CSV'}
      </button>
      {status && <div className="status">Status: {status}</div>}
    </div>
  );
};
```

#### **Angular Example**

```typescript
// services/upload.service.ts
import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, interval } from 'rxjs';
import { switchMap, takeWhile } from 'rxjs/operators';

@Injectable({ providedIn: 'root' })
export class UploadService {
  private apiUrl = 'http://localhost:8000';

  constructor(private http: HttpClient) {}

  async uploadFile(file: File): Promise<string> {
    // Step 1: Get presigned URL
    const urlResponse = await this.http.post<any>(
      `${this.apiUrl}/api/upload/presigned-url`,
      { filename: file.name }
    ).toPromise();

    // Step 2: Upload to S3
    await this.http.put(urlResponse.upload_url, file, {
      headers: { 'Content-Type': 'text/csv' }
    }).toPromise();

    return urlResponse.upload_id;
  }

  watchUploadStatus(uploadId: string): Observable<any> {
    return interval(3000).pipe(
      switchMap(() => 
        this.http.get(`${this.apiUrl}/api/upload/status/${uploadId}`)
      ),
      takeWhile((status: any) => 
        status.status !== 'completed' && status.status !== 'failed',
        true // include final value
      )
    );
  }
}
```

#### **Vue 3 Example**

```typescript
// composables/useFileUpload.ts
import { ref } from 'vue';

export function useFileUpload() {
  const uploading = ref(false);
  const uploadProgress = ref(0);
  const uploadStatus = ref('');

  const uploadFile = async (file: File) => {
    uploading.value = true;
    uploadProgress.value = 0;

    try {
      // Get presigned URL
      const response = await fetch('http://localhost:8000/api/upload/presigned-url', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ filename: file.name })
      });

      const { upload_url, upload_id } = await response.json();

      // Upload file
      const uploadResponse = await fetch(upload_url, {
        method: 'PUT',
        body: file
      });

      if (!uploadResponse.ok) {
        throw new Error('Upload failed');
      }

      uploadProgress.value = 100;

      // Monitor status
      monitorStatus(upload_id);

    } catch (error) {
      console.error('Upload error:', error);
      uploading.value = false;
    }
  };

  const monitorStatus = (uploadId: string) => {
    const interval = setInterval(async () => {
      const response = await fetch(
        `http://localhost:8000/api/upload/status/${uploadId}`
      );
      const data = await response.json();
      
      uploadStatus.value = data.status;

      if (data.status === 'completed' || data.status === 'failed') {
        clearInterval(interval);
        uploading.value = false;
      }
    }, 3000);
  };

  return {
    uploading,
    uploadProgress,
    uploadStatus,
    uploadFile
  };
}
```

---

### Key Considerations for Frontend Developers

#### **1. Error Handling**
```javascript
try {
  const response = await fetch(presignedUrl, {
    method: 'PUT',
    body: file
  });
  
  if (!response.ok) {
    throw new Error(`Upload failed: ${response.status}`);
  }
} catch (error) {
  // Show user-friendly error message
  showNotification('Upload failed. Please try again.', 'error');
  
  // Log for debugging
  console.error('Upload error:', error);
  
  // Optionally retry
  retryUpload(file);
}
```

#### **2. File Validation**
```javascript
const validateFile = (file: File): boolean => {
  // Check file type
  if (!file.name.endsWith('.csv')) {
    alert('Please select a CSV file');
    return false;
  }
  
  // Check file size (e.g., max 50MB)
  if (file.size > 50 * 1024 * 1024) {
    alert('File too large. Maximum size is 50MB');
    return false;
  }
  
  return true;
};
```

#### **3. Progress Tracking**
```javascript
const uploadWithProgress = async (file: File, url: string) => {
  const xhr = new XMLHttpRequest();
  
  return new Promise((resolve, reject) => {
    xhr.upload.addEventListener('progress', (e) => {
      if (e.lengthComputable) {
        const percentComplete = (e.loaded / e.total) * 100;
        updateProgressBar(percentComplete);
      }
    });
    
    xhr.addEventListener('load', () => resolve(xhr.response));
    xhr.addEventListener('error', () => reject(xhr.statusText));
    
    xhr.open('PUT', url);
    xhr.send(file);
  });
};
```

#### **4. Status Polling Best Practices**
```javascript
const pollWithExponentialBackoff = async (uploadId: string) => {
  let delay = 2000; // Start with 2 seconds
  const maxDelay = 30000; // Max 30 seconds
  const maxAttempts = 50;
  let attempts = 0;

  while (attempts < maxAttempts) {
    const status = await checkStatus(uploadId);
    
    if (status === 'completed' || status === 'failed') {
      return status;
    }
    
    await sleep(delay);
    delay = Math.min(delay * 1.5, maxDelay);
    attempts++;
  }
  
  throw new Error('Status check timeout');
};
```

---

## 🏗️ System Overview

This pricing management system consists of TWO microservices:

```
┌─────────────────────────────────────────────────────────┐
│                     USER/FRONTEND                        │
│                  (Angular/React/CLI)                     │
└────────────────┬───────────────────────┬─────────────────┘
                 │                       │
                 │ 1. Get URL            │ 3. Search/CRUD
                 ▼                       ▼
        ┌─────────────────┐    ┌──────────────────┐
        │  Python API     │    │   Python API     │
        │  (Port 8000)    │    │   (Port 8000)    │
        │  Presigned URLs │    │   Search/CRUD    │
        └────────┬────────┘    └────────┬─────────┘
                 │                      │
                 │ 2. Direct Upload     │ Query/Update
                 ▼                      ▼
        ┌──────────────────┐  ┌────────────────────┐
        │   S3 Storage     │  │   PostgreSQL DB    │
        │   (CSV Files)    │  │  (Pricing Records) │
        └────────┬─────────┘  └─────────▲──────────┘
                 │                      │
                 │ Event Notification   │ Batch Insert
                 ▼                      │
        ┌─────────────────────────────────┐
        │    Go Ingestion Service         │
        │         (Port 8080)              │
        │   - Event Listener               │
        │   - CSV Validator                │
        │   - Stream Parser                │
        │   - Batch Insert                 │
        └──────────────────────────────────┘
```

---

## 🔄 Complete Data Flow

### Upload Flow (Asynchronous)

```
Step 1: User initiates upload
   └─> Frontend calls: POST http://localhost:8000/api/upload/presigned-url
       Request: { "filename": "pricing.csv" }
       
Step 2: Python API generates presigned URL
   └─> Returns: { "upload_url": "s3://...", "upload_id": "uuid" }
   
Step 3: Frontend uploads directly to S3
   └─> PUT to presigned URL (bypasses backend servers)
   
Step 4: S3 triggers event notification
   └─> Event sent to Go service (webhook/SQS/SNS)
   
Step 5: Go service processes file
   └─> Downloads from S3
   └─> Validates CSV structure
   └─> Streams and parses records
   └─> Inserts in batches (1000 records/batch)
   └─> Updates upload_history status
   
Step 6: User checks status
   └─> Frontend polls: GET http://localhost:8000/api/upload/status/{upload_id}
   └─> Or searches for new records immediately
```

### Search Flow (Synchronous)

```
Step 1: User searches for pricing data
   └─> Frontend calls: GET http://localhost:8000/api/pricing/search?store_id=ST001
   
Step 2: Python API queries database
   └─> Filters applied (store, SKU, price range, dates)
   └─> Pagination applied (page, page_size)
   └─> Sorting applied (by price, date, etc.)
   
Step 3: Results returned to user
   └─> Returns: { items: [...], total: 100, page: 1, total_pages: 10 }
```

### Update Flow (Synchronous)

```
Step 1: User edits a pricing record
   └─> Frontend calls: PUT http://localhost:8000/api/pricing/records/123
       Request: { "price": 24.99, "date": "2024-03-03" }
       
Step 2: Python API validates and updates
   └─> Validates business rules
   └─> Updates record in database
   └─> Logs audit trail
   
Step 3: Updated record returned
   └─> Returns: { id: 123, ...updated_fields, updated_at: "..." }
```

---

## 🚀 Running Both Services

### Terminal 1: Start Go Ingestion Service

```powershell
# Navigate to project root
cd c:\Workspace\Personal_Project\golang-project-product-service

# Run Go service
go run cmd/server/main.go

# Or build and run
go build -o product-service.exe cmd/server/main.go
.\product-service.exe
```

**Go Service Running**: http://localhost:8080
- Health: http://localhost:8080/health
- Upload: http://localhost:8080/upload

### Terminal 2: Start Python API Service

```powershell
# Navigate to Python service
cd python-api-service

# Activate virtual environment
.\venv\Scripts\activate

# Run Python service
uvicorn app.main:app --reload --port 8000
```

**Python Service Running**: http://localhost:8000
- Docs: http://localhost:8000/docs
- Health: http://localhost:8000/health
- Presigned URL: http://localhost:8000/api/upload/presigned-url
- Search: http://localhost:8000/api/pricing/search

---

## 📊 Service Responsibilities

### Python API Service (Port 8000)

**Purpose**: User-facing REST API for frontend applications

**Responsibilities**:
- ✅ Generate presigned URLs for S3 uploads
- ✅ Search pricing records (with filters, pagination, sorting)
- ✅ CRUD operations (Create, Read, Update, Delete)
- ✅ Authentication & Authorization (future)
- ✅ Rate limiting (future)
- ✅ API documentation (Swagger/ReDoc)

**Technology Stack**:
- FastAPI (Python web framework)
- SQLAlchemy (ORM)
- Pydantic (data validation)
- Boto3 (AWS SDK)

**Database Access**: Read + Write (for CRUD operations)

**Why Python**:
- Rapid API development
- Rich ecosystem for web APIs
- Easy integration with ML/AI models (future)
- Excellent documentation generation

### Go Ingestion Service (Port 8080)

**Purpose**: Background CSV processing and data ingestion

**Responsibilities**:
- ✅ Listen for S3 event notifications
- ✅ Download CSV files from S3
- ✅ Validate CSV structure and data
- ✅ Stream-parse large CSV files
- ✅ Batch insert records (1000/batch)
- ✅ Handle concurrent processing
- ✅ Update upload history status
- ✅ Error logging and retry logic

**Technology Stack**:
- Gin (Go web framework)
- encoding/csv (standard library)
- Goroutines (concurrency)
- sync.RWMutex (thread-safe storage)

**Database Access**: Write-heavy (batch inserts)

**Why Go**:
- High performance for I/O operations
- Excellent concurrency with goroutines
- Low memory footprint
- Fast CSV processing
- Efficient for long-running processes

---

## 🗄️ Shared Database

Both services access the **same database**:

### Development (SQLite)
```
golang-project-product-service/
├── pricing.db                    # Shared SQLite database
├── python-api-service/
│   └── app/
│       └── main.py              # Uses: sqlite:///./pricing.db
└── cmd/server/main.go           # Uses: pricing.db
```

### Production (PostgreSQL)

**Python Service** (`.env`):
```env
DATABASE_URL=postgresql+asyncpg://user:pass@db-host:5432/pricing_db
```

**Go Service** (`.env`):
```env
DATABASE_URL=postgresql://user:pass@db-host:5432/pricing_db
```

### Database Schema

```sql
-- Pricing Records (main table)
CREATE TABLE pricing_records (
    id INTEGER PRIMARY KEY,
    store_id VARCHAR(20) NOT NULL,
    sku VARCHAR(50) NOT NULL,
    product_name VARCHAR(200) NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_store_id (store_id),
    INDEX idx_sku (sku),
    INDEX idx_date (date)
);

-- Upload History (tracking)
CREATE TABLE upload_history (
    id INTEGER PRIMARY KEY,
    filename VARCHAR(255) NOT NULL,
    upload_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status VARCHAR(50) NOT NULL,  -- 'processing', 'success', 'failed'
    records_count INTEGER DEFAULT 0,
    error_message TEXT
);

-- Audit Logs (future)
CREATE TABLE audit_logs (
    id INTEGER PRIMARY KEY,
    table_name VARCHAR(100),
    record_id INTEGER,
    action VARCHAR(50),  -- 'INSERT', 'UPDATE', 'DELETE'
    old_value TEXT,
    new_value TEXT,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

---

## 🧪 End-to-End Testing

### Test Complete Upload Flow

```powershell
# Step 1: Start both services
# Terminal 1: go run cmd/server/main.go
# Terminal 2: cd python-api-service; uvicorn app.main:app --reload --port 8000

# Step 2: Generate presigned URL
$presignedRequest = @{
    filename = "test.csv"
    content_type = "text/csv"
} | ConvertTo-Json

$presignedResponse = Invoke-RestMethod `
    -Uri "http://localhost:8000/api/upload/presigned-url" `
    -Method Post -Body $presignedRequest -ContentType "application/json"

Write-Host "Upload URL: $($presignedResponse.upload_url)"
Write-Host "Upload ID: $($presignedResponse.upload_id)"

# Step 3: Upload file to Go service (mock for now)
$file = Get-Item "test.csv"
$uploadForm = @{ file = $file }

$uploadResponse = Invoke-RestMethod `
    -Uri "http://localhost:8080/upload" `
    -Method Post -Form $uploadForm

Write-Host "Upload successful: $($uploadResponse.message)"
Write-Host "Records imported: $($uploadResponse.rows - 1)"

# Step 4: Search for imported data via Python API
$searchResponse = Invoke-RestMethod `
    -Uri "http://localhost:8000/api/pricing/search?page=1&page_size=10"

Write-Host "Total records in DB: $($searchResponse.total)"
Write-Host "Records on page: $($searchResponse.items.Count)"

# Step 5: Update a record
if ($searchResponse.items.Count -gt 0) {
    $recordId = $searchResponse.items[0].id
    $updateBody = @{
        price = 99.99
    } | ConvertTo-Json
    
    $updatedRecord = Invoke-RestMethod `
        -Uri "http://localhost:8000/api/pricing/records/$recordId" `
        -Method Put -Body $updateBody -ContentType "application/json"
    
    Write-Host "Updated record $recordId - New price: $($updatedRecord.price)"
}
```

---

## 🔧 Configuration

### Python Service Configuration

**File**: `python-api-service/.env`

```env
# Server
HOST=0.0.0.0
PORT=8000

# S3 Configuration
MOCK_S3=true                              # Set to false for production
AWS_ACCESS_KEY_ID=your-key
AWS_SECRET_ACCESS_KEY=your-secret
S3_BUCKET_NAME=pricing-csv-uploads
PRESIGNED_URL_EXPIRATION=3600

# Database
DATABASE_URL=sqlite:///./pricing.db       # or PostgreSQL URL

# CORS (for frontend)
ALLOWED_ORIGINS=http://localhost:4200,http://localhost:3000
```

### Go Service Configuration

**File**: `.env` (in project root)

```env
# Server
PORT=8080

# Database
DATABASE_URL=pricing.db                   # or PostgreSQL URL

# Upload
UPLOAD_PATH=./uploads
MAX_UPLOAD_SIZE=10485760
BATCH_SIZE=1000

# S3 (for future)
USE_S3=false
S3_BUCKET=pricing-csv-uploads
```

---

## 🚦 Service Health Checks

### Check Both Services

```powershell
# Check Go service
Invoke-RestMethod "http://localhost:8080/health"

# Check Python service
Invoke-RestMethod "http://localhost:8000/health"
```

### Expected Responses

**Go Service**:
```json
{
  "status": "healthy",
  "timestamp": "2024-03-03T10:00:00Z",
  "database": "connected (in-memory)",
  "record_count": 42
}
```

**Python Service**:
```json
{
  "status": "healthy",
  "service": "pricing-api",
  "mock_mode": true,
  "database": "connected"
}
```

---

## 🎯 Production Deployment

### Deployment Checklist

**Python API Service**:
- [ ] Set `MOCK_S3=false`
- [ ] Configure real AWS credentials
- [ ] Use PostgreSQL database
- [ ] Enable HTTPS
- [ ] Add authentication (JWT)
- [ ] Configure rate limiting
- [ ] Set up monitoring (Prometheus)
- [ ] Enable CORS for production domain

**Go Ingestion Service**:
- [ ] Configure S3 event notifications
- [ ] Set up webhook or SQS listener
- [ ] Use PostgreSQL database
- [ ] Configure connection pooling
- [ ] Set up error alerting
- [ ] Enable structured logging
- [ ] Configure auto-scaling
- [ ] Set up health check endpoints

**Infrastructure**:
- [ ] Deploy both services separately
- [ ] Use managed database (RDS/Cloud SQL)
- [ ] Configure S3 bucket with versioning
- [ ] Set up CloudWatch/monitoring
- [ ] Configure load balancers
- [ ] Enable auto-scaling
- [ ] Set up CI/CD pipelines

---

## 📚 API Documentation

### Python API (Interactive)
- Swagger UI: http://localhost:8000/docs
- ReDoc: http://localhost:8000/redoc

### Go API (Manual)
- Health: `GET /health`
- Upload: `POST /upload` (multipart/form-data)

---

## ✅ Benefits of This Architecture

1. **Separation of Concerns**: CRUD APIs separate from heavy processing
2. **Language Optimization**: Right tool for the right job
3. **Independent Scaling**: Scale services based on load
4. **Async Processing**: Non-blocking uploads
5. **Fault Tolerance**: Services can fail independently
6. **Maintainability**: Clear boundaries and responsibilities
7. **Performance**: Go for I/O, Python for API flexibility

---

**Both services are ready to run!** 🎉

Start them in separate terminals and test the complete flow.
