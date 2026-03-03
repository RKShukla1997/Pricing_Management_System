from fastapi import APIRouter, HTTPException, status
from app.models.schemas import (
    PresignedURLRequest, 
    PresignedURLResponse, 
    MessageResponse
)
from app.services.storage import storage_service

router = APIRouter(prefix="/api/upload", tags=["Upload"])


@router.post(
    "/presigned-url",
    response_model=PresignedURLResponse,
    status_code=status.HTTP_200_OK,
    summary="Generate Presigned URL",
    description="Generate a presigned URL for direct CSV file upload to S3"
)
async def generate_presigned_url(request: PresignedURLRequest):
    """
    Generate a presigned URL for uploading CSV files directly to S3.
    
    **Flow:**
    1. Frontend calls this endpoint with filename
    2. Backend generates presigned URL with expiration
    3. Frontend uploads file directly to S3 using the URL
    4. S3 triggers event notification
    5. Go ingestion service processes the file
    
    **Parameters:**
    - **filename**: Name of the CSV file to upload
    - **content_type**: MIME type (default: text/csv)
    
    **Returns:**
    - **upload_url**: Presigned URL for PUT request
    - **file_key**: S3 object key for tracking
    - **expires_in**: URL expiration time in seconds
    - **upload_id**: Unique identifier for this upload
    """
    try:
        result = storage_service.generate_presigned_url(
            filename=request.filename,
            content_type=request.content_type
        )
        
        return PresignedURLResponse(**result)
        
    except Exception as e:
        raise HTTPException(
            status_code=status.HTTP_500_INTERNAL_SERVER_ERROR,
            detail=f"Failed to generate presigned URL: {str(e)}"
        )


@router.get(
    "/status/{upload_id}",
    response_model=MessageResponse,
    summary="Get Upload Status",
    description="Check the status of a file upload by upload ID"
)
async def get_upload_status(upload_id: str):
    """
    Check the processing status of an uploaded file.
    
    This would typically query the upload_history table to get status.
    For now, returns a mock response.
    """
    # TODO: Query database for actual status
    return MessageResponse(
        message="Upload status retrieved",
        detail=f"Upload {upload_id} is being processed"
    )
