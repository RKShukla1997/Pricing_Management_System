from pydantic import BaseModel, Field, validator
from typing import Optional, List
from datetime import datetime
from datetime import date as DateType
from decimal import Decimal


# Pricing Record Schemas
class PricingRecordBase(BaseModel):
    store_id: str = Field(..., max_length=20, description="Store identifier")
    sku: str = Field(..., max_length=50, description="Product SKU")
    product_name: str = Field(..., max_length=200, description="Product name")
    price: Decimal = Field(..., gt=0, le=999999.99, description="Product price")
    date: DateType = Field(..., description="Effective date")


class PricingRecordCreate(PricingRecordBase):
    pass


class PricingRecordUpdate(BaseModel):
    store_id: Optional[str] = Field(None, max_length=20)
    sku: Optional[str] = Field(None, max_length=50)
    product_name: Optional[str] = Field(None, max_length=200)
    price: Optional[Decimal] = Field(None, gt=0, le=999999.99)
    date: Optional[DateType] = None


class PricingRecord(PricingRecordBase):
    id: int
    created_at: datetime
    updated_at: datetime

    class Config:
        from_attributes = True


# Presigned URL Schemas
class PresignedURLRequest(BaseModel):
    filename: str = Field(..., description="Name of the file to upload")
    content_type: str = Field(default="text/csv", description="MIME type of the file")
    
    @validator('filename')
    def validate_filename(cls, v):
        if not v.lower().endswith('.csv'):
            raise ValueError('Only CSV files are allowed')
        return v


class PresignedURLResponse(BaseModel):
    upload_url: str = Field(..., description="Presigned URL for upload")
    file_key: str = Field(..., description="S3 object key")
    expires_in: int = Field(..., description="URL expiration time in seconds")
    upload_id: str = Field(..., description="Unique upload identifier")


# Search Schemas
class PricingSearchParams(BaseModel):
    store_id: Optional[str] = Field(None, description="Filter by store ID")
    sku: Optional[str] = Field(None, description="Filter by SKU")
    product_name: Optional[str] = Field(None, description="Search in product name")
    min_price: Optional[Decimal] = Field(None, ge=0, description="Minimum price")
    max_price: Optional[Decimal] = Field(None, le=999999.99, description="Maximum price")
    date_from: Optional[DateType] = Field(None, description="Start date")
    date_to: Optional[DateType] = Field(None, description="End date")
    page: int = Field(default=1, ge=1, description="Page number")
    page_size: int = Field(default=10, ge=1, le=100, description="Items per page")
    sort_by: Optional[str] = Field(default="date", description="Sort field")
    sort_order: Optional[str] = Field(default="desc", description="Sort order (asc/desc)")


class PaginatedResponse(BaseModel):
    items: List[PricingRecord]
    total: int
    page: int
    page_size: int
    total_pages: int


# Upload History Schemas
class UploadHistory(BaseModel):
    id: int
    filename: str
    upload_date: datetime
    status: str
    records_count: int
    error_message: Optional[str] = None

    class Config:
        from_attributes = True


# Response Schemas
class MessageResponse(BaseModel):
    message: str
    detail: Optional[str] = None


class ErrorResponse(BaseModel):
    error: str
    detail: Optional[str] = None
    field: Optional[str] = None
