from fastapi import APIRouter, HTTPException, status, Depends, Query
from sqlalchemy.orm import Session
from typing import Optional
from decimal import Decimal
from datetime import date
import math

from app.models.schemas import (
    PricingRecord,
    PricingRecordCreate,
    PricingRecordUpdate,
    PricingSearchParams,
    PaginatedResponse,
    MessageResponse
)
from app.services.database import get_db, db_service

router = APIRouter(prefix="/api/pricing", tags=["Pricing"])


@router.post(
    "/records",
    response_model=PricingRecord,
    status_code=status.HTTP_201_CREATED,
    summary="Create Pricing Record",
    description="Create a new pricing record"
)
async def create_pricing_record(
    record: PricingRecordCreate,
    db: Session = Depends(get_db)
):
    """
    Create a new pricing record in the database.
    
    **Parameters:**
    - **store_id**: Store identifier (max 20 chars)
    - **sku**: Product SKU (max 50 chars)
    - **product_name**: Product name (max 200 chars)
    - **price**: Product price (0.01 to 999999.99)
    - **date**: Effective date (YYYY-MM-DD)
    """
    try:
        db_record = db_service.create_record(db, record)
        return db_record
    except Exception as e:
        raise HTTPException(
            status_code=status.HTTP_500_INTERNAL_SERVER_ERROR,
            detail=f"Failed to create record: {str(e)}"
        )


@router.get(
    "/records/{record_id}",
    response_model=PricingRecord,
    summary="Get Pricing Record",
    description="Get a pricing record by ID"
)
async def get_pricing_record(
    record_id: int,
    db: Session = Depends(get_db)
):
    """Get a single pricing record by its ID"""
    db_record = db_service.get_record_by_id(db, record_id)
    
    if not db_record:
        raise HTTPException(
            status_code=status.HTTP_404_NOT_FOUND,
            detail=f"Record with ID {record_id} not found"
        )
    
    return db_record


@router.put(
    "/records/{record_id}",
    response_model=PricingRecord,
    summary="Update Pricing Record",
    description="Update an existing pricing record"
)
async def update_pricing_record(
    record_id: int,
    record_update: PricingRecordUpdate,
    db: Session = Depends(get_db)
):
    """
    Update an existing pricing record.
    
    **Parameters:**
    - **record_id**: ID of the record to update
    - **record_update**: Fields to update (only provided fields will be updated)
    
    **Returns:**
    - Updated pricing record
    """
    db_record = db_service.update_record(db, record_id, record_update)
    
    if not db_record:
        raise HTTPException(
            status_code=status.HTTP_404_NOT_FOUND,
            detail=f"Record with ID {record_id} not found"
        )
    
    return db_record


@router.delete(
    "/records/{record_id}",
    response_model=MessageResponse,
    summary="Delete Pricing Record",
    description="Delete a pricing record"
)
async def delete_pricing_record(
    record_id: int,
    db: Session = Depends(get_db)
):
    """Delete a pricing record by ID"""
    success = db_service.delete_record(db, record_id)
    
    if not success:
        raise HTTPException(
            status_code=status.HTTP_404_NOT_FOUND,
            detail=f"Record with ID {record_id} not found"
        )
    
    return MessageResponse(
        message="Record deleted successfully",
        detail=f"Record {record_id} has been removed"
    )


@router.get(
    "/search",
    response_model=PaginatedResponse,
    summary="Search Pricing Records",
    description="Search pricing records with filters and pagination"
)
async def search_pricing_records(
    store_id: Optional[str] = Query(None, description="Filter by store ID"),
    sku: Optional[str] = Query(None, description="Filter by SKU"),
    product_name: Optional[str] = Query(None, description="Search in product name"),
    min_price: Optional[Decimal] = Query(None, ge=0, description="Minimum price"),
    max_price: Optional[Decimal] = Query(None, le=999999.99, description="Maximum price"),
    date_from: Optional[date] = Query(None, description="Start date (YYYY-MM-DD)"),
    date_to: Optional[date] = Query(None, description="End date (YYYY-MM-DD)"),
    page: int = Query(1, ge=1, description="Page number"),
    page_size: int = Query(10, ge=1, le=100, description="Items per page"),
    sort_by: str = Query("date", description="Sort field (date, price, store_id, sku)"),
    sort_order: str = Query("desc", description="Sort order (asc, desc)"),
    db: Session = Depends(get_db)
):
    """
    Search pricing records with advanced filtering and pagination.
    
    **Query Parameters:**
    - **store_id**: Filter by exact store ID
    - **sku**: Filter by exact SKU
    - **product_name**: Search by product name (case-insensitive partial match)
    - **min_price**: Minimum price filter
    - **max_price**: Maximum price filter
    - **date_from**: Filter records from this date onwards
    - **date_to**: Filter records up to this date
    - **page**: Page number (default: 1)
    - **page_size**: Items per page (default: 10, max: 100)
    - **sort_by**: Sort field (date, price, store_id, sku)
    - **sort_order**: Sort order (asc or desc)
    
    **Example Queries:**
    - `/api/pricing/search?store_id=ST001` - All records for store ST001
    - `/api/pricing/search?sku=SKU-001` - All records for SKU-001
    - `/api/pricing/search?product_name=milk` - All products containing "milk"
    - `/api/pricing/search?min_price=5&max_price=50` - Products between $5 and $50
    - `/api/pricing/search?date_from=2024-01-01&date_to=2024-01-31` - January 2024 records
    - `/api/pricing/search?page=2&page_size=20&sort_by=price&sort_order=asc` - Second page, sorted by price
    
    **Returns:**
    - **items**: List of pricing records
    - **total**: Total number of matching records
    - **page**: Current page number
    - **page_size**: Items per page
    - **total_pages**: Total number of pages
    """
    search_params = PricingSearchParams(
        store_id=store_id,
        sku=sku,
        product_name=product_name,
        min_price=min_price,
        max_price=max_price,
        date_from=date_from,
        date_to=date_to,
        page=page,
        page_size=page_size,
        sort_by=sort_by,
        sort_order=sort_order
    )
    
    try:
        records, total = db_service.search_records(db, search_params)
        total_pages = math.ceil(total / page_size)
        
        return PaginatedResponse(
            items=records,
            total=total,
            page=page,
            page_size=page_size,
            total_pages=total_pages
        )
    except Exception as e:
        raise HTTPException(
            status_code=status.HTTP_500_INTERNAL_SERVER_ERROR,
            detail=f"Search failed: {str(e)}"
        )
