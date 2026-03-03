# Angular Frontend - Quick Start Guide

## 🚀 Complete System Setup

### Terminal 1 - Go Service (Port 8080)
```powershell
cd c:\Workspace\Personal_Project\golang-project-product-service
.\product-service.exe
```

### Terminal 2 - Python API (Port 8000)
```powershell
cd c:\Workspace\Personal_Project\golang-project-product-service\python-api-service
& .\venv\Scripts\python.exe -m uvicorn app.main:app --reload --host 0.0.0.0 --port 8000
```

### Terminal 3 - Angular Frontend (Port 4200)
```powershell
cd c:\Workspace\Personal_Project\golang-project-product-service\pricing-frontend
ng serve --open
```

## 📱 Application URLs

- **Angular Frontend**: http://localhost:4200
- **Go Service**: http://localhost:8080
- **Python API**: http://localhost:8000
- **API Docs**: http://localhost:8000/docs

## 🎯 User Flow Demo

### 1. Upload CSV File
1. Navigate to http://localhost:4200/upload
2. Select upload method:
   - **Direct Upload**: Sends file directly to Go service
   - **Presigned URL**: Uses Python API mock presigned URL
3. Choose CSV file (see format below)
4. Click "Upload File"
5. Wait for success message
6. Click "View All Records →"

### 2. View & Edit Records
1. Navigate to http://localhost:4200/records
2. Use filters to search:
   - Store ID
   - SKU  
   - Product Name
   - Price range
   - Date range
3. Click "Apply Filters"
4. To edit a record:
   - Click ✏️ (edit icon)
   - Modify fields
   - Click ✅ (save)
5. To delete a record:
   - Click 🗑️ (delete icon)
   - Confirm deletion

## 📄 Sample CSV File

Create `test-products.csv`:

```csv
Store ID,SKU,Product Name,Price,Date
STORE001,LAPTOP001,Dell Laptop,899.99,2024-01-15
STORE001,MOUSE001,Wireless Mouse,29.99,2024-01-15
STORE002,KEYBOARD001,Mechanical Keyboard,129.99,2024-01-16
STORE002,MONITOR001,27" Monitor,349.99,2024-01-16
STORE003,HEADSET001,Gaming Headset,79.99,2024-01-17
```

**Important**: Save with UTF-8 encoding (no BOM)

## 🔄 Data Flow

```
┌─────────────────────────────────────────────────────────┐
│  Angular Frontend (localhost:4200)                      │
│  - Upload Component                                     │
│  - Records Component                                    │
└────────────┬──────────────────────────┬─────────────────┘
             │                          │
             │ CSV Upload               │ REST API Calls
             │ (multipart/form-data)    │ (JSON)
             ↓                          ↓
┌────────────────────────┐   ┌─────────────────────────────┐
│  Go Service (8080)     │   │  Python API (8000)          │
│  - CSV Validation      │   │  - Search & Filter          │
│  - Batch Processing    │   │  - CRUD Operations          │
│  - Write to SQLite     │   │  - Presigned URLs           │
└────────────┬───────────┘   └────────────┬────────────────┘
             │                            │
             └──────────┬─────────────────┘
                        ↓
            ┌──────────────────────┐
            │  SQLite Database     │
            │  pricing.db          │
            └──────────────────────┘
```

## 🎨 Features Overview

### Upload Page Features
✅ File selection with visual feedback
✅ Upload method toggle (Direct/Presigned)
✅ Real-time upload status
✅ Success/Error notifications
✅ CSV format guidelines
✅ Direct navigation to records

### Records Page Features
✅ Responsive data table
✅ Multi-field search filters
✅ Inline editing
✅ Delete with confirmation
✅ Pagination
✅ Empty state handling
✅ Loading indicators
✅ Error handling

## 🐛 Troubleshooting

### Issue: Angular app not loading
**Solution**: Check if port 4200 is available
```powershell
Get-NetTCPConnection -LocalPort 4200 -ErrorAction SilentlyContinue
# If in use, stop the process or use different port
ng serve --port 4300
```

### Issue: CORS errors in browser console
**Solution**: Backend CORS is already configured for `localhost:4200`. If using different port, update backend CORS settings.

### Issue: "Failed to load records"
**Check**:
```powershell
# Python API health
curl http://localhost:8000/health

# Go service health  
curl http://localhost:8080/health
```

### Issue: Upload fails
**Check**:
1. CSV format matches required headers
2. File is UTF-8 encoded (no BOM)
3. Go service is running
4. Check browser console for errors

## 📊 Testing the System

### Quick Test Script
```powershell
# 1. Check all services
curl http://localhost:8080/health
curl http://localhost:8000/health
curl http://localhost:4200

# 2. Upload test file
curl.exe -X POST http://localhost:8080/upload -F "file=@test-products.csv"

# 3. Query via Python API
curl "http://localhost:8000/api/pricing/search"

# 4. Filter by store
curl "http://localhost:8000/api/pricing/search?store_id=STORE001"
```

## 🎬 Demo Script for Interview

**1. Introduction (30 seconds)**
- "This is a microservices-based pricing management system"
- "Go for CSV ingestion, Python for REST API, Angular for UI"
- "All services share a SQLite database for strong consistency"

**2. Upload Demo (1 minute)**
- Open upload page
- Show dual upload options
- Upload sample CSV
- Explain: "Go validates headers, processes in batches of 1000"
- Show success message

**3. Records Demo (2 minutes)**
- Navigate to records page
- "Python API provides rich query interface"
- Demo filters: "Let's filter by STORE001"
- Edit a record: "Inline editing with instant validation"
- Show pagination
- Delete a record with confirmation

**4. Architecture Explanation (1 minute)**
- "CQRS pattern: Go writes, Python reads"
- "Single bounded context - pricing management"
- "Shared database for strong consistency"
- "Can scale to PostgreSQL or event-driven when needed"

**5. Code Walkthrough (1 minute)**
- Show Angular service with typed interfaces
- Show Go CSV streaming parser
- Show Python FastAPI endpoints
- Highlight: "Type-safe across all layers"

## 📈 Next Steps / Extensions

Future enhancements to discuss:
- [ ] WebSocket for real-time updates
- [ ] Export filtered results to CSV
- [ ] Bulk edit multiple records
- [ ] File upload history dashboard
- [ ] Data visualization charts
- [ ] User authentication & authorization
- [ ] Audit log viewer
- [ ] CSV template generator

## 🎯 Success Criteria

System is working correctly when:
✅ Can upload CSV via Angular → Go service
✅ Records appear in Angular table immediately
✅ Filters work correctly
✅ Can edit records and see updates
✅ Can delete records
✅ All three services running without errors
✅ No CORS errors in browser console

## 📝 Notes

- Angular uses standalone components (modern approach)
- Services use HttpClient with RxJS observables
- Type-safe models for all API responses
- Responsive design works on mobile/tablet
- Error handling at every level
- Loading states for better UX

---

**Your system is ready for demo! 🎉**

Open http://localhost:4200 and start uploading CSV files!
