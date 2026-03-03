from fastapi import FastAPI, status
from fastapi.middleware.cors import CORSMiddleware
from fastapi.responses import JSONResponse
from contextlib import asynccontextmanager

from app.config.settings import settings
from app.services.database import init_db
from app.api import upload, pricing


@asynccontextmanager
async def lifespan(app: FastAPI):
    """Lifespan events - startup and shutdown"""
    # Startup
    print("Initializing database...")
    init_db()
    print("Database initialized")
    yield
    # Shutdown
    print("Shutting down...")


# Create FastAPI app
app = FastAPI(
    title="Pricing Management API",
    description="""
    **FastAPI Microservice for Retail Pricing Management System**
    
    This service provides:
    - **Presigned URL generation** for direct S3 uploads
    - **Search API** with advanced filtering and pagination
    - **CRUD operations** for pricing records
    - **RESTful API** following best practices
    
    ## Architecture
    
    This is the **Python API Service** component that handles:
    1. User-facing REST API operations
    2. Presigned URL generation for CSV uploads
    3. Database queries and updates
    4. Authentication and authorization (future)
    
    The **Go Ingestion Service** handles:
    - CSV file processing
    - Validation
    - Batch inserts
    
    ## Usage
    
    1. **Upload Flow**: Generate presigned URL → Upload to S3 → Processing
    2. **Search Flow**: Query with filters → Get paginated results
    3. **Update Flow**: Update individual pricing records
    """,
    version="1.0.0",
    lifespan=lifespan,
    docs_url="/docs",
    redoc_url="/redoc"
)

# Configure CORS
app.add_middleware(
    CORSMiddleware,
    allow_origins=settings.ALLOWED_ORIGINS,
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# Include routers
app.include_router(upload.router)
app.include_router(pricing.router)


@app.get(
    "/",
    tags=["Root"],
    summary="Root Endpoint",
    description="API information and health check"
)
async def root():
    """Root endpoint - API information"""
    return {
        "service": "Pricing Management API",
        "version": "1.0.0",
        "status": "running",
        "docs": "/docs",
        "redoc": "/redoc",
        "endpoints": {
            "presigned_url": "/api/upload/presigned-url",
            "search": "/api/pricing/search",
            "records": "/api/pricing/records"
        }
    }


@app.get(
    "/health",
    tags=["Health"],
    summary="Health Check",
    description="Service health status"
)
async def health_check():
    """Health check endpoint"""
    return {
        "status": "healthy",
        "service": "pricing-api",
        "mock_mode": settings.MOCK_S3,
        "database": "connected"
    }


@app.exception_handler(Exception)
async def global_exception_handler(request, exc):
    """Global exception handler"""
    return JSONResponse(
        status_code=status.HTTP_500_INTERNAL_SERVER_ERROR,
        content={
            "error": "Internal server error",
            "detail": str(exc)
        }
    )


if __name__ == "__main__":
    import uvicorn
    
    uvicorn.run(
        "app.main:app",
        host=settings.HOST,
        port=settings.PORT,
        reload=True,
        log_level="info"
    )
