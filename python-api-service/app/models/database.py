from sqlalchemy import Column, Integer, String, DECIMAL, DateTime, Text, Date
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.sql import func
from datetime import datetime

Base = declarative_base()


class PricingRecordDB(Base):
    __tablename__ = "pricing_records"

    id = Column(Integer, primary_key=True, index=True, autoincrement=True)
    store_id = Column(String(20), nullable=False, index=True)
    sku = Column(String(50), nullable=False, index=True)
    product_name = Column(String(200), nullable=False)
    price = Column(DECIMAL(10, 2), nullable=False)
    date = Column(Date, nullable=False, index=True)
    created_at = Column(DateTime, default=datetime.utcnow, nullable=False)
    updated_at = Column(DateTime, default=datetime.utcnow, onupdate=datetime.utcnow, nullable=False)


class UploadHistoryDB(Base):
    __tablename__ = "upload_history"

    id = Column(Integer, primary_key=True, index=True, autoincrement=True)
    filename = Column(String(255), nullable=False)
    upload_date = Column(DateTime, default=datetime.utcnow, nullable=False)
    status = Column(String(50), nullable=False)  # processing, success, failed
    records_count = Column(Integer, default=0)
    error_message = Column(Text, nullable=True)


class AuditLogDB(Base):
    __tablename__ = "audit_logs"

    id = Column(Integer, primary_key=True, index=True, autoincrement=True)
    table_name = Column(String(100), nullable=False)
    record_id = Column(Integer, nullable=True)
    action = Column(String(50), nullable=False)  # INSERT, UPDATE, DELETE
    old_value = Column(Text, nullable=True)
    new_value = Column(Text, nullable=True)
    timestamp = Column(DateTime, default=datetime.utcnow, nullable=False)
