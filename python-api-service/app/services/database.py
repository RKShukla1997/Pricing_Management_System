from typing import List, Optional, Tuple
from sqlalchemy import create_engine, and_, or_, desc, asc
from sqlalchemy.orm import sessionmaker, Session
from app.models.database import Base, PricingRecordDB
from app.models.schemas import PricingRecordCreate, PricingRecordUpdate, PricingSearchParams
from app.config.settings import settings
from datetime import date


# Database setup
engine = create_engine(
    settings.DATABASE_URL,
    connect_args={"check_same_thread": False} if "sqlite" in settings.DATABASE_URL else {}
)

SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)


def get_db():
    """Dependency for getting database session"""
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()


def init_db():
    """Initialize database tables"""
    Base.metadata.create_all(bind=engine)


class DatabaseService:
    """Service for database operations on pricing records"""
    
    @staticmethod
    def create_record(db: Session, record: PricingRecordCreate) -> PricingRecordDB:
        """Create a new pricing record"""
        db_record = PricingRecordDB(**record.model_dump())
        db.add(db_record)
        db.commit()
        db.refresh(db_record)
        return db_record
    
    @staticmethod
    def get_record_by_id(db: Session, record_id: int) -> Optional[PricingRecordDB]:
        """Get a pricing record by ID"""
        return db.query(PricingRecordDB).filter(PricingRecordDB.id == record_id).first()
    
    @staticmethod
    def update_record(
        db: Session, 
        record_id: int, 
        record_update: PricingRecordUpdate
    ) -> Optional[PricingRecordDB]:
        """Update a pricing record"""
        db_record = db.query(PricingRecordDB).filter(PricingRecordDB.id == record_id).first()
        
        if not db_record:
            return None
        
        # Update only provided fields
        update_data = record_update.model_dump(exclude_unset=True)
        for field, value in update_data.items():
            setattr(db_record, field, value)
        
        db.commit()
        db.refresh(db_record)
        return db_record
    
    @staticmethod
    def delete_record(db: Session, record_id: int) -> bool:
        """Delete a pricing record"""
        db_record = db.query(PricingRecordDB).filter(PricingRecordDB.id == record_id).first()
        
        if not db_record:
            return False
        
        db.delete(db_record)
        db.commit()
        return True
    
    @staticmethod
    def search_records(
        db: Session, 
        params: PricingSearchParams
    ) -> Tuple[List[PricingRecordDB], int]:
        """
        Search pricing records with filters and pagination
        
        Returns:
            Tuple of (records, total_count)
        """
        query = db.query(PricingRecordDB)
        
        # Apply filters
        filters = []
        
        if params.store_id:
            filters.append(PricingRecordDB.store_id == params.store_id)
        
        if params.sku:
            filters.append(PricingRecordDB.sku == params.sku)
        
        if params.product_name:
            filters.append(PricingRecordDB.product_name.ilike(f"%{params.product_name}%"))
        
        if params.min_price is not None:
            filters.append(PricingRecordDB.price >= params.min_price)
        
        if params.max_price is not None:
            filters.append(PricingRecordDB.price <= params.max_price)
        
        if params.date_from:
            filters.append(PricingRecordDB.date >= params.date_from)
        
        if params.date_to:
            filters.append(PricingRecordDB.date <= params.date_to)
        
        if filters:
            query = query.filter(and_(*filters))
        
        # Get total count before pagination
        total_count = query.count()
        
        # Apply sorting
        sort_column = getattr(PricingRecordDB, params.sort_by, PricingRecordDB.date)
        if params.sort_order == "desc":
            query = query.order_by(desc(sort_column))
        else:
            query = query.order_by(asc(sort_column))
        
        # Apply pagination
        offset = (params.page - 1) * params.page_size
        query = query.offset(offset).limit(params.page_size)
        
        records = query.all()
        
        return records, total_count


# Singleton instance
db_service = DatabaseService()
