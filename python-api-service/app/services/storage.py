import boto3
from botocore.exceptions import ClientError
from datetime import datetime, timedelta
import uuid
from typing import Dict, Optional
from app.config.settings import settings


class StorageService:
    """Service for handling S3 presigned URLs and file operations"""
    
    def __init__(self):
        self.mock_mode = settings.MOCK_S3
        
        if not self.mock_mode:
            self.s3_client = boto3.client(
                's3',
                aws_access_key_id=settings.AWS_ACCESS_KEY_ID,
                aws_secret_access_key=settings.AWS_SECRET_ACCESS_KEY,
                region_name=settings.AWS_REGION
            )
        else:
            self.s3_client = None
        
        self.bucket_name = settings.S3_BUCKET_NAME
        self.expiration = settings.PRESIGNED_URL_EXPIRATION
    
    def generate_presigned_url(
        self, 
        filename: str, 
        content_type: str = "text/csv"
    ) -> Dict[str, str]:
        """
        Generate a presigned URL for file upload
        
        Args:
            filename: Original filename
            content_type: MIME type of the file
            
        Returns:
            Dict containing upload_url, file_key, expires_in, upload_id
        """
        # Generate unique file key
        upload_id = str(uuid.uuid4())
        timestamp = datetime.utcnow().strftime("%Y%m%d_%H%M%S")
        file_key = f"uploads/{timestamp}_{upload_id}_{filename}"
        
        if self.mock_mode:
            # Mock presigned URL for testing
            mock_url = f"http://localhost:8000/mock-upload/{upload_id}"
            return {
                "upload_url": mock_url,
                "file_key": file_key,
                "expires_in": self.expiration,
                "upload_id": upload_id
            }
        
        try:
            # Generate real presigned URL
            presigned_url = self.s3_client.generate_presigned_url(
                'put_object',
                Params={
                    'Bucket': self.bucket_name,
                    'Key': file_key,
                    'ContentType': content_type
                },
                ExpiresIn=self.expiration
            )
            
            return {
                "upload_url": presigned_url,
                "file_key": file_key,
                "expires_in": self.expiration,
                "upload_id": upload_id
            }
            
        except ClientError as e:
            raise Exception(f"Failed to generate presigned URL: {str(e)}")
    
    def get_file_url(self, file_key: str) -> Optional[str]:
        """Get a presigned URL for downloading a file"""
        if self.mock_mode:
            return f"http://localhost:8000/mock-download/{file_key}"
        
        try:
            url = self.s3_client.generate_presigned_url(
                'get_object',
                Params={
                    'Bucket': self.bucket_name,
                    'Key': file_key
                },
                ExpiresIn=3600
            )
            return url
        except ClientError:
            return None
    
    def delete_file(self, file_key: str) -> bool:
        """Delete a file from S3"""
        if self.mock_mode:
            return True
        
        try:
            self.s3_client.delete_object(
                Bucket=self.bucket_name,
                Key=file_key
            )
            return True
        except ClientError:
            return False


# Singleton instance
storage_service = StorageService()
