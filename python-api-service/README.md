# Python API Service - Pricing Management System

FastAPI microservice for handling pricing data operations.

## 🚀 Features

### 1. **Presigned URL Generation**
- Generate presigned URLs for direct S3 uploads
- Eliminates large file handling from API servers
- Event-driven architecture integration
- Mock mode for testing without AWS

### 2. **Advanced Search API**
- Filter by store, SKU, product name, price range, date range
- Pagination support (configurable page size)
- Sorting by any field (asc/desc)
- Case-insensitive partial text search

### 3. **CRUD Operations**
- Create pricing records
- Read/Get records by ID
- Update existing records
- Delete records
- Full validation on all operations

## 📁 Project Structure

```
python-api-service/
├── app/
│   ├── api/
│   │   ├── upload.py          # Presigned URL endpoints
│   │   └── pricing.py         # CRUD & Search endpoints
│   ├── models/
│   │   ├── schemas.py         # Pydantic models
│   │   └── database.py        # SQLAlchemy models
│   ├── services/
│   │   ├── storage.py         # S3 service
│   │   └── database.py        # Database service
│   ├── config/
│   │   └── settings.py        # Configuration
│   └── main.py                # FastAPI app
├── requirements.txt
├── .env.example
└── README.md
```

## 🛠️ Setup

### 1. Install Dependencies

```bash
cd python-api-service
python -m venv venv
.\venv\Scripts\activate  # Windows
# or
source venv/bin/activate  # Linux/Mac

pip install -r requirements.txt
```

### 2. Configure Environment

Copy `.env.example` to `.env` and configure:

```env
# Mock mode for testing without AWS
MOCK_S3=true

# Database (SQLite by default)
DATABASE_URL=sqlite:///./pricing.db

# CORS for Angular/React frontend
ALLOWED_ORIGINS=http://localhost:4200,http://localhost:3000
```

### 3. Run the Server

```bash
# Development mode with auto-reload
uvicorn app.main:app --reload --port 8000

# Or using Python directly
python -m app.main
```

Server will start at: **http://localhost:8000**

## 📖 API Documentation

### Interactive Documentation

- **Swagger UI**: http://localhost:8000/docs
- **ReDoc**: http://localhost:8000/redoc

### Endpoints Overview

#### 1. Presigned URL Generation

**POST** `/api/upload/presigned-url`

```bash
curl -X POST "http://localhost:8000/api/upload/presigned-url" \
  -H "Content-Type: application/json" \
  -d '{
    "filename": "pricing_data.csv",
    "content_type": "text/csv"
  }'
```

Response:
```json
{
  "upload_url": "http://localhost:8000/mock-upload/uuid",
  "file_key": "uploads/20240303_120000_uuid_pricing_data.csv",
  "expires_in": 3600,
  "upload_id": "uuid"
}
```

#### 2. Search Records

**GET** `/api/pricing/search`

```bash
# Search by store
curl "http://localhost:8000/api/pricing/search?store_id=ST001"

# Search by product name
curl "http://localhost:8000/api/pricing/search?product_name=milk"

# Search with price range
curl "http://localhost:8000/api/pricing/search?min_price=5&max_price=50"

# Search with date range and pagination
curl "http://localhost:8000/api/pricing/search?date_from=2024-01-01&date_to=2024-01-31&page=1&page_size=20"

# Complex search with sorting
curl "http://localhost:8000/api/pricing/search?store_id=ST001&min_price=10&sort_by=price&sort_order=asc"
```

Response:
```json
{
  "items": [
    {
      "id": 1,
      "store_id": "ST001",
      "sku": "SKU-001",
      "product_name": "Milk 1L",
      "price": 3.99,
      "date": "2024-01-15",
      "created_at": "2024-03-03T10:00:00",
      "updated_at": "2024-03-03T10:00:00"
    }
  ],
  "total": 1,
  "page": 1,
  "page_size": 10,
  "total_pages": 1
}
```

#### 3. Get Record by ID

**GET** `/api/pricing/records/{record_id}`

```bash
curl "http://localhost:8000/api/pricing/records/1"
```

#### 4. Update Record

**PUT** `/api/pricing/records/{record_id}`

```bash
curl -X PUT "http://localhost:8000/api/pricing/records/1" \
  -H "Content-Type: application/json" \
  -d '{
    "price": 4.99,
    "date": "2024-03-03"
  }'
```

#### 5. Delete Record

**DELETE** `/api/pricing/records/{record_id}`

```bash
curl -X DELETE "http://localhost:8000/api/pricing/records/1"
```

#### 6. Create Record

**POST** `/api/pricing/records`

```bash
curl -X POST "http://localhost:8000/api/pricing/records" \
  -H "Content-Type: application/json" \
  -d '{
    "store_id": "ST001",
    "sku": "SKU-001",
    "product_name": "New Product",
    "price": 9.99,
    "date": "2024-03-03"
  }'
```

## 🧪 Testing

### PowerShell Testing Script

```powershell
# Test presigned URL generation
$body = @{
    filename = "test.csv"
    content_type = "text/csv"
} | ConvertTo-Json

Invoke-RestMethod -Uri "http://localhost:8000/api/upload/presigned-url" `
    -Method Post -Body $body -ContentType "application/json"

# Search records
Invoke-RestMethod -Uri "http://localhost:8000/api/pricing/search?store_id=ST001"

# Update record
$updateBody = @{
    price = 5.99
} | ConvertTo-Json

Invoke-RestMethod -Uri "http://localhost:8000/api/pricing/records/1" `
    -Method Put -Body $updateBody -ContentType "application/json"
```

## 🏗️ Architecture Integration

### With Go Ingestion Service

1. **Frontend** calls Python API for presigned URL
2. **Frontend** uploads CSV directly to S3 using presigned URL
3. **S3** triggers event notification
4. **Go Service** (port 8080) processes CSV file
5. **Go Service** inserts records into shared database
6. **Frontend** queries Python API for search/CRUD operations

### Database Sharing

Both services can share the same database:

**SQLite (Development):**
```env
# Python API
DATABASE_URL=sqlite:///./pricing.db

# Go Service
DATABASE_URL=pricing.db
```

**PostgreSQL (Production):**
```env
# Python API
DATABASE_URL=postgresql+asyncpg://user:pass@localhost/pricing_db

# Go Service  
DATABASE_URL=postgresql://user:pass@localhost:5432/pricing_db
```

## 🔐 Security Features

### Planned (Not Yet Implemented)

- [ ] JWT authentication
- [ ] API key validation
- [ ] Rate limiting
- [ ] Request validation
- [ ] SQL injection prevention (SQLAlchemy handles this)
- [ ] CORS configuration (already configured)

## 📊 Mock Mode vs Production

### Mock Mode (Default)

```env
MOCK_S3=true
```

- No AWS credentials required
- Returns mock presigned URLs
- Perfect for local development
- No S3 charges

### Production Mode

```env
MOCK_S3=false
AWS_ACCESS_KEY_ID=your-key
AWS_SECRET_ACCESS_KEY=your-secret
S3_BUCKET_NAME=your-bucket
```

- Real S3 integration
- Actual presigned URLs
- Event notifications configured
- Production-ready

## 🚀 Deployment

### Docker (Recommended)

```dockerfile
FROM python:3.11-slim

WORKDIR /app
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

COPY app ./app
EXPOSE 8000

CMD ["uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "8000"]
```

### Cloud Deployment

- AWS ECS/Fargate
- Azure Container Apps
- Google Cloud Run
- Heroku
- Railway

## 📝 Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `HOST` | Server host | 0.0.0.0 |
| `PORT` | Server port | 8000 |
| `MOCK_S3` | Use mock S3 | true |
| `DATABASE_URL` | Database connection | sqlite:///./pricing.db |
| `S3_BUCKET_NAME` | S3 bucket name | pricing-csv-uploads |
| `PRESIGNED_URL_EXPIRATION` | URL expiry (seconds) | 3600 |
| `ALLOWED_ORIGINS` | CORS origins | localhost:4200,... |

## 🎯 Next Steps

1. **Authentication**: Add JWT authentication
2. **Caching**: Implement Redis caching for search
3. **Async Database**: Switch to async SQLAlchemy
4. **WebSockets**: Real-time upload status
5. **Monitoring**: Add Prometheus metrics
6. **Logging**: Structured logging with correlation IDs

---

**Status**: ✅ Ready for testing and integration  
**API Docs**: http://localhost:8000/docs  
**Python Version**: 3.11+  
**Framework**: FastAPI 0.109.0
