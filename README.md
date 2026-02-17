# Retail Pricing Management System# Retail Pricing Management System# Retail Pricing Management System# Retail Pricing Management System# Retail Pricing Management System# Retail Pricing Management System# Product Service - CSV Upload API



A multi-service web application for managing pricing feeds from 3000+ retail stores. Built with Angular, Go, and Python.



---A multi-service web application for managing pricing feeds from 3000+ retail stores. Built with Angular, Go, and Python.



## Context Diagram



The system allows retail store users to upload CSV files, search pricing data, and edit records through a modern web interface.---A scalable multi-service platform for managing pricing data from 3000+ retail stores across multiple countries. Built with Angular, Go, and Python.



```

                  ┌─────────────────────┐

                  │  Retail Store User  │## Context Diagram

                  └──────────┬──────────┘

                             │

                             ▼

                  ┌─────────────────────┐The system allows retail store users to upload CSV files, search pricing data, and edit records through a modern web interface.---A scalable, multi-service web application for managing and analyzing pricing feeds from retail stores across multiple countries. The system enables upload, persistence, search, and modification of pricing data from 3000+ retail stores.

                  │   Angular Web UI    │

                  │  - Upload CSV       │

                  │  - Search & Filter  │

                  │  - Edit Records     │```

                  └──────────┬──────────┘

                             │ REST APIs                  ┌─────────────────────┐

                             ▼

                  ┌─────────────────────┐                  │  Retail Store User  │## 📋 Table of Contents

                  │  Backend Services   │

                  │                     │                  └──────────┬──────────┘

                  │  ┌────────────────┐ │

                  │  │ Python API     │ │                             │

                  │  │ (FastAPI)      │ │

                  │  └────────────────┘ │                             ▼

                  │                     │

                  │  ┌────────────────┐ │                  ┌─────────────────────┐- [Overview](#overview)---A scalable, multi-service web application for managing and analyzing pricing feeds from retail stores across multiple countries. The system enables upload, persistence, search, and modification of pricing data from 3000+ retail stores.

                  │  │ Go Ingestion   │ │

                  │  │ (Gin)          │ │                  │   Angular Web UI    │

                  │  └────────────────┘ │

                  └──────────┬──────────┘                  │  - Upload CSV       │- [Context Diagram](#context-diagram)

                             │

           ┌─────────────────┴─────────────────┐                  │  - Search & Filter  │

           │                                   │

           ▼                                   ▼                  │  - Edit Records     │- [Solution Architecture](#solution-architecture)

   ┌───────────────┐                  ┌───────────────┐

   │  PostgreSQL   │                  │ S3 / Storage  │                  └──────────┬──────────┘

   │  - Records    │                  │ - CSV Files   │

   │  - Audit Logs │                  │ - Archives    │                             │ REST APIs- [Design Decisions](#design-decisions)

   └───────────────┘                  └───────────────┘

```                             ▼



**Key Components:**                  ┌─────────────────────┐- [Non-Functional Requirements](#non-functional-requirements)## 📋 Table of Contents

- **Users:** Store managers, pricing analysts, administrators

- **Frontend:** Single Page Application for all interactions                  │  Backend Services   │

- **Backend:** Separate services for file processing (Go) and data operations (Python)

- **Storage:** Relational database for records, object storage for files                  │                     │- [Assumptions](#assumptions)



---                  │  ┌────────────────┐ │



## Solution Architecture                  │  │ Python API     │ │- [Getting Started](#getting-started)



The architecture uses an event-driven approach where file uploads are processed asynchronously.                  │  │ (FastAPI)      │ │



```                  │  └────────────────┘ │

                  ┌──────────────────┐

                  │   Angular UI     │                  │                     │

                  │  (Port: 4200)    │

                  └────────┬─────────┘                  │  ┌────────────────┐ │---- [Overview](#overview)---A scalable, multi-service web application for managing and analyzing pricing feeds from retail stores across multiple countries. The system enables upload, persistence, search, and modification of pricing data from 3000+ retail stores.A microservice built with Go and the Gin web framework that provides an endpoint for uploading CSV files.

                           │

          ┌────────────────┴────────────────┐                  │  │ Go Ingestion   │ │

          │                                 │

          ▼                                 ▼                  │  │ (Gin)          │ │

   ┌─────────────────┐             ┌──────────────────┐

   │  Python API     │             │  Go Ingestion    │                  │  └────────────────┘ │

   │  (Port: 8000)   │             │  (Port: 8080)    │

   │                 │             │                  │                  └──────────┬──────────┘## 🎯 Overview- [Context Diagram](#context-diagram)

   │ • Presigned URL │             │ • Event Listener │

   │ • Search API    │             │ • CSV Validator  │                             │

   │ • CRUD Ops      │             │ • Stream Parser  │

   │ • Auth          │             │ • Batch Insert   │           ┌─────────────────┴─────────────────┐

   └────────┬────────┘             └────────┬─────────┘

            │                               │           │                                   │

            │        ┌──────────────────────┘

            │        │  Storage Event           ▼                                   ▼**Retail Pricing Management System** enables retail chains to:- [Solution Architecture](#solution-architecture)

            │        │

   ┌────────┴────────┴────────────────────────────────┐   ┌───────────────┐                  ┌───────────────┐

   │                  DATA LAYER                       │

   │                                                   │   │  PostgreSQL   │                  │ S3 / Storage  │- Upload CSV pricing feeds (Store ID, SKU, Product Name, Price, Date)

   │  ┌────────────────┐         ┌─────────────────┐  │

   │  │ S3 / Storage   │         │   PostgreSQL    │  │   │  - Records    │                  │ - CSV Files   │

   │  │ • CSV Upload   │         │ • Records       │  │

   │  │ • Event Trigger│         │ • Upload History│  │   │  - Audit Logs │                  │ - Archives    │- Search and filter pricing records- [Technology Stack](#technology-stack)

   │  └────────────────┘         └─────────────────┘  │

   │                                                   │   └───────────────┘                  └───────────────┘

   │  ┌─────────────────┐                             │

   │  │ Redis (Cache)   │                             │```- Edit and update pricing data in real-time

   │  │ • Search Cache  │                             │

   │  └─────────────────┘                             │

   └───────────────────────────────────────────────────┘

```**Key Components:**- [Design Decisions](#design-decisions)## 📋 Table of Contents



**Architecture Highlights:**- **Users:** Store managers, pricing analysts, administrators



1. **Python API Service (FastAPI)**- **Frontend:** Single Page Application for all interactions**Technology Stack:**

   - Generates presigned URLs for direct S3 uploads

   - Handles search queries with Redis caching- **Backend:** Separate services for file processing (Go) and data operations (Python)

   - Manages CRUD operations with PostgreSQL

   - Provides authentication and authorization- **Storage:** Relational database for records, object storage for files- **Frontend:** Angular 17+- [Non-Functional Requirements](#non-functional-requirements)



2. **Go Ingestion Service (Gin)**

   - Listens to S3 storage events

   - Validates and parses CSV files using streaming---- **Upload Service:** Go + Gin (This Repository)

   - Performs batch inserts to PostgreSQL

   - Efficient concurrent processing with goroutines



3. **Data Flow**## Solution Architecture- **Data API:** Python + FastAPI- [Assumptions](#assumptions)

   - **Upload:** User → Angular → Python API (Presigned URL) → S3 → Storage Event → Go Service → PostgreSQL

   - **Search:** User → Angular → Python API → Redis (Cache) → PostgreSQL

   - **Edit:** User → Angular → Python API → PostgreSQL (Transaction) → Audit Log

The architecture uses an event-driven approach where file uploads are processed asynchronously.- **Database:** PostgreSQL + Redis

**Key Design Decisions:**

- **Event-Driven:** Decouples upload from processing for better scalability

- **Direct S3 Upload:** Reduces backend load and improves upload performance

- **Streaming Parser:** Handles large CSV files efficiently without loading into memory```- **Storage:** S3 / Local FileSystem- [Getting Started](#getting-started)

- **Batch Inserts:** Optimizes database performance (1000 records per batch)

- **Microservices:** Go excels at I/O operations; Python excels at data manipulation                  ┌──────────────────┐


                  │   Angular UI     │

                  │  (Port: 4200)    │

                  └────────┬─────────┘---- [Source Implementation](#source-implementation)- [Overview](#overview)---## Features

                           │

          ┌────────────────┴────────────────┐

          │                                 │

          ▼                                 ▼## 🗺️ Context Diagram

   ┌─────────────────┐             ┌──────────────────┐

   │  Python API     │             │  Go Ingestion    │

   │  (Port: 8000)   │             │  (Port: 8080)    │

   │                 │             │                  │```---- [Context Diagram](#context-diagram)

   │ • Presigned URL │             │ • Event Listener │

   │ • Search API    │             │ • CSV Validator  │                  ┌─────────────────────┐

   │ • CRUD Ops      │             │ • Stream Parser  │

   │ • Auth          │             │ • Batch Insert   │                  │  Retail Store User  │

   └────────┬────────┘             └────────┬─────────┘

            │                               │                  └──────────┬──────────┘

            │        ┌──────────────────────┘

            │        │  Storage Event                             │## 🎯 Overview- [Solution Architecture](#solution-architecture)

            │        │

   ┌────────┴────────┴────────────────────────────────┐                             ▼

   │                  DATA LAYER                       │

   │                                                   │                  ┌─────────────────────┐

   │  ┌────────────────┐         ┌─────────────────┐  │

   │  │ S3 / Storage   │         │   PostgreSQL    │  │                  │   Angular Web UI    │

   │  │ • CSV Upload   │         │ • Records       │  │

   │  │ • Event Trigger│         │ • Upload History│  │                  │  - Upload CSV       │This solution provides a comprehensive pricing management platform for retail chains operating 3000+ stores across multiple countries.- [Technology Stack](#technology-stack)

   │  └────────────────┘         └─────────────────┘  │

   │                                                   │                  │  - Search & Filter  │

   │  ┌─────────────────┐                             │

   │  │ Redis (Cache)   │                             │                  │  - Edit Records     │

   │  │ • Search Cache  │                             │

   │  └─────────────────┘                             │                  └──────────┬──────────┘

   └───────────────────────────────────────────────────┘

```                             │**Key Capabilities:**- [Functional Requirements](#functional-requirements)## 📋 Table of Contents- CSV file upload via HTTP POST



**Architecture Highlights:**                             ▼



1. **Python API Service (FastAPI)**                  ┌─────────────────────┐- CSV pricing feed uploads from retail stores

   - Generates presigned URLs for direct S3 uploads

   - Handles search queries with Redis caching                  │  Backend Services   │

   - Manages CRUD operations with PostgreSQL

   - Provides authentication and authorization                  │                     │- Persistent storage of pricing records (Store ID, SKU, Product Name, Price, Date)- [Non-Functional Requirements](#non-functional-requirements)



2. **Go Ingestion Service (Gin)**                  │  ┌────────────────┐ │

   - Listens to S3 storage events

   - Validates and parses CSV files using streaming                  │  │ Python API     │ │- Advanced search and filter capabilities

   - Performs batch inserts to PostgreSQL

   - Efficient concurrent processing with goroutines                  │  │ (FastAPI)      │ │



3. **Data Flow**                  │  └────────────────┘ │- Real-time editing and updates to pricing records- [Design Decisions](#design-decisions)- File size validation (max 10MB)

   - **Upload:** User → Angular → Python API (Presigned URL) → S3 → Storage Event → Go Service → PostgreSQL

   - **Search:** User → Angular → Python API → Redis (Cache) → PostgreSQL                  │                     │

   - **Edit:** User → Angular → Python API → PostgreSQL (Transaction) → Audit Log

                  │  ┌────────────────┐ │- Scalable architecture supporting multi-country operations

**Key Design Decisions:**

- **Event-Driven:** Decouples upload from processing for better scalability                  │  │ Go Ingestion   │ │

- **Direct S3 Upload:** Reduces backend load and improves upload performance

- **Streaming Parser:** Handles large CSV files efficiently without loading into memory                  │  │ (Gin)          │ │- [Assumptions](#assumptions)

- **Batch Inserts:** Optimizes database performance (1000 records per batch)

- **Microservices:** Go excels at I/O operations; Python excels at data manipulation                  │  └────────────────┘ │



---                  └──────────┬──────────┘**Technology Stack:**



## Quick Start                             │



```powershell           ┌─────────────────┴─────────────────┐- **Frontend:** Angular (Single Page Application)- [Project Structure](#project-structure)- [Overview](#overview)- CSV format validation

# Start all services with Docker Compose

docker-compose up -d           │                                   │



# Access applications           ▼                                   ▼- **Backend:** Go (Upload/Ingestion Service) + Python (Data API Service)

# - Frontend: http://localhost:4200

# - Python API: http://localhost:8000   ┌───────────────┐                  ┌───────────────┐

# - Go Service: http://localhost:8080

```   │  PostgreSQL   │                  │ S3 / Storage  │- **Database:** PostgreSQL (Primary) + Redis (Cache)- [Getting Started](#getting-started)



---   │  - Records    │                  │ - CSV Files   │



## Repository   │  - Audit Logs │                  │ - Archives    │- **Storage:** S3/Local FileSystem (CSV Files)



**Project:** [Pricing_Management_System](https://github.com/RKShukla1997/Pricing_Management_System)     └───────────────┘                  └───────────────┘

**Last Updated:** February 17, 2026

```- [API Documentation](#api-documentation)- [Context Diagram](#context-diagram)- File type validation (only .csv files)



------



## 🏗️ Solution Architecture- [Deployment](#deployment)



```## 🗺️ Context Diagram

┌────────────────────────────────────────────────────────────┐

│                     PRESENTATION LAYER                      │- [Solution Architecture](#solution-architecture)- Automatic file storage with timestamps

│                                                             │

│                  ┌──────────────────┐                       │### High-Level System Context

│                  │   Angular UI     │                       │

│                  │  (Port: 4200)    │                       │---

│                  └────────┬─────────┘                       │

└─────────────────────────┬─────────────────────────────────-┘```

                          │ HTTP/REST

┌─────────────────────────┴──────────────────────────────────┐                ┌────────────────────────┐- [Technology Stack](#technology-stack)- Health check endpoint

│                   APPLICATION LAYER                         │

│                                                             │                │   Retail Store User    │

│    ┌─────────────────┐              ┌──────────────────┐   │

│    │  Python API     │              │  Go Ingestion    │   │                │  (Store Managers,      │## 🎯 Overview

│    │  (Port: 8000)   │              │  (Port: 8080)    │   │

│    │                 │              │                  │   │                │   Pricing Analysts,    │

│    │ • Presigned URL │              │ • Event Listener │   │

│    │ • Search API    │              │ • CSV Validator  │   │                │   Administrators)      │- [Functional Requirements](#functional-requirements)- JSON responses with proper error handling

│    │ • CRUD Ops      │              │ • Stream Parser  │   │

│    │ • Auth          │              │ • Batch Insert   │   │                └────────────┬───────────┘

│    └────────┬────────┘              └────────┬─────────┘   │

└─────────────┼──────────────────────────────┬─────────────-┘                             │This solution provides a comprehensive pricing management platform for retail chains operating 3000+ stores across multiple countries. The system handles:

              │                               │

              │        ┌──────────────────────┘                             │ HTTPS

              │        │  Storage Event

              │        │                             ▼- [Non-Functional Requirements](#non-functional-requirements)- Built with Gin framework for high performance

┌─────────────┴────────┴────────────────────────────────────┐

│                       DATA LAYER                           │                ┌────────────────────────┐

│                                                            │

│  ┌────────────────────┐         ┌────────────────────┐    │                │     Angular Web UI     │- **CSV pricing feed uploads** from retail stores

│  │   S3 / Storage     │         │    PostgreSQL      │    │

│  │                    │         │                    │    │                │   (Single Page App)    │

│  │  • CSV Upload      │         │  • pricing_records │    │

│  │  • Event Trigger   │         │  • upload_history  │    │                │                        │- **Persistent storage** of pricing records (Store ID, SKU, Product Name, Price, Date)- [Design Decisions](#design-decisions)

│  │  • File Archive    │         │  • audit_logs      │    │

│  └────────────────────┘         └────────────────────┘    │                │ - Upload Interface     │

│                                                            │

│  ┌────────────────────┐                                    │                │ - Search & Filter      │- **Advanced search capabilities** with multiple criteria

│  │   Redis (Cache)    │                                    │

│  │  • Session         │                                    │                │ - Edit Data Grid       │

│  │  • Search Cache    │                                    │

│  └────────────────────┘                                    │                └────────────┬───────────┘- **Real-time editing** and updates to pricing records- [Assumptions](#assumptions)## Prerequisites

└────────────────────────────────────────────────────────────┘

```                             │



### Data Flow                             │ REST APIs- **Scalable architecture** supporting multi-country operations



**Upload Flow:**                             │ (JSON)

```

User → Angular → Python API (Presigned URL) → S3 Direct Upload                             ▼- [Project Structure](#project-structure)

→ Storage Event → Go Service → Validate & Parse → PostgreSQL

```                ┌────────────────────────┐



**Search Flow:**                │   Backend Services     │---

```

User → Angular → Python API → Redis (Check Cache) → PostgreSQL                │                        │

→ Cache Results → Return to UI

```                │ ┌────────────────────┐ │- [Getting Started](#getting-started)- Go 1.25.1 or higher



**Edit Flow:**                │ │  Python API Svc    │ │

```

User → Angular → Python API → PostgreSQL (Transaction: Update + Audit)                │ │  (FastAPI)         │ │## 🗺️ Context Diagram

→ Invalidate Cache → Return Updated Record

```                │ └────────────────────┘ │



---                │                        │- [API Documentation](#api-documentation)- Git



## 🎨 Design Decisions                │ ┌────────────────────┐ │



### 1. Microservices with Language Optimization                │ │ Go Ingestion Svc   │ │### High-Level System Context



**Go for Upload/Ingestion:**                │ │ (Gin Framework)    │ │

- Excellent I/O performance

- Efficient memory management                │ └────────────────────┘ │- [Deployment](#deployment)

- Native concurrency (goroutines)

- Fast CSV streaming                └────────────┬───────────┘



**Python for Data API:**                             │```

- Rich data manipulation (Pandas, SQLAlchemy)

- Rapid API development (FastAPI)          ┌──────────────────┴──────────────────┐

- Better for business logic

          │                                     │                ┌────────────────────────┐## Installation

### 2. Event-Driven Upload

          ▼                                     ▼

**Presigned URL + Storage Events:**

- Frontend uploads directly to S3 (no backend bottleneck) ┌────────────────────┐               ┌────────────────────┐                │   Retail Store User    │

- S3 event triggers async processing

- Reduced bandwidth costs │ Relational Database│               │ Object Storage     │

- Better scalability

 │   (PostgreSQL)     │               │   (S3 / Local)     │                │  (Store Managers,      │---

### 3. Streaming CSV Parser

 │                    │               │                    │

**Why:** Memory-efficient processing of large files

```go │ - Pricing Records  │               │ - CSV Files        │                │   Pricing Analysts,    │

reader := csv.NewReader(file)

for { │ - Audit Logs       │               │ - File Archive     │

    record, err := reader.Read()

    if err == io.EOF { break } │ - User Sessions    │               │ - Backups          │                │   Administrators)      │1. Clone the repository

    processBatch(record)

} └────────────────────┘               └────────────────────┘

```

```                └────────────┬───────────┘

### 4. PostgreSQL Database



**Why:** ACID compliance, complex queries, data integrity, proven at scale

---                             │## 🎯 Overview2. Install dependencies:

### 5. Batch Inserts



**Why:** 10-100x faster than individual inserts

```go## 🏗️ Solution Architecture                             │ HTTPS

const batchSize = 1000

for i := 0; i < len(records); i += batchSize {

    batch := records[i:min(i+batchSize, len(records))]

    db.BatchInsert(batch)### Detailed Architecture with Data Flow                             ▼```bash

}

```



### 6. Optimistic Locking```                ┌────────────────────────┐



**Why:** Better concurrency, no database locks┌─────────────────────────────────────────────────────────────────┐

```python

UPDATE pricing_records│                         PRESENTATION LAYER                       │                │     Angular Web UI     │This solution provides a comprehensive pricing management platform for retail chains operating 3000+ stores across multiple countries. The system handles:go mod download

SET price = $1, version = version + 1

WHERE id = $2 AND version = $3│                                                                  │

```

│                    ┌────────────────────┐                        │                │   (Single Page App)    │

---

│                    │     Angular UI     │                        │

## 🚀 Non-Functional Requirements

│                    │  (Port: 4200)      │                        │                │                        │```

### Performance

- **Search Response:** < 2 seconds│                    │                    │                        │

- **Upload Speed:** 10MB file in < 5 seconds

- **Concurrent Users:** 1000+│                    │ - Material Design  │                        │                │ - Upload Interface     │

- **Throughput:** 100+ uploads/minute

│                    │ - State Management │                        │

**Approach:** DB indexing, Redis caching, connection pooling, async processing

│                    │ - Form Validation  │                        │                │ - Search & Filter      │- **CSV pricing feed uploads** from retail stores

### Scalability

- **Data Volume:** 100M+ records│                    └─────────┬──────────┘                        │

- **Store Growth:** 3000 → 10,000 stores

- **Multi-Region:** Geographic distribution└──────────────────────────────┼───────────────────────────────────┘                │ - Edit Data Grid       │



**Approach:** Microservices, DB sharding, read replicas, auto-scaling, message queues                               │



### Availability                               │ HTTP/REST                └────────────┬───────────┘- **Persistent storage** of pricing records (Store ID, SKU, Product Name, Price, Date)## Running the Service

- **Uptime:** 99.9% (< 8.76 hours downtime/year)

- **Recovery:** RTO < 4 hours, RPO < 1 hour                               │



**Approach:** Multi-AZ deployment, DB replication, health checks, daily backups┌──────────────────────────────┼───────────────────────────────────┐                             │



### Security│                       APPLICATION LAYER                          │

- **Auth:** OAuth 2.0 / JWT

- **Encryption:** TLS 1.3, AES-256│                               │                                  │                             │ REST APIs- **Advanced search capabilities** with multiple criteria

- **Compliance:** GDPR, SOC 2

│     ┌─────────────────────────┴─────────────────────────┐       │

**Approach:** API gateway, JWT tokens, encrypted connections, input validation, audit logs

│     │                                                     │       │                             │ (JSON)

### Maintainability

- **Test Coverage:** > 80%│     ▼                                                     ▼       │

- **Documentation:** API docs, architecture diagrams

│ ┌────────────────────┐                       ┌────────────────┐ │                             ▼- **Real-time editing** and updates to pricing records```bash

**Approach:** Unit/integration tests, OpenAPI/Swagger, CI/CD pipelines

│ │  Python API Service│                       │ Go Ingestion   │ │

---

│ │  (FastAPI)         │                       │ Service (Gin)  │ │                ┌────────────────────────┐

## 📝 Assumptions

│ │  Port: 8000        │                       │ Port: 8080     │ │

### Business

- 3000 stores → 10,000 in 3 years│ │                    │                       │                │ │                │   Backend Services     │- **Scalable architecture** supporting multi-country operationsgo run main.go

- Daily uploads per store

- Files: 1K-10K records (< 10MB)│ │ ┌────────────────┐ │                       │ ┌────────────┐ │ │

- 2-year data retention

- 500 concurrent users│ │ │ Generate       │ │                       │ │ Event      │ │ │                │                        │



### Technical│ │ │ Presigned URL  │ │                       │ │ Listener   │ │ │

- 1 Mbps minimum upload speed

- Modern browsers (last 2 versions)│ │ └────────────────┘ │                       │ └────────────┘ │ │                │ ┌────────────────────┐ │```

- Database: 50GB → 500GB

- Cloud-hosted (AWS/Azure/GCP)│ │                    │                       │                │ │



### Data│ │ ┌────────────────┐ │                       │ ┌────────────┐ │ │                │ │  Python API Svc    │ │

- CSV format: `Store ID, SKU, Product Name, Price, Date`

- UTF-8 encoding│ │ │ Search API     │ │                       │ │ CSV        │ │ │

- All fields required

- Price: 0.01 - 999,999.99│ │ │ (Multi-Criteria)│ │                      │ │ Validation │ │ │                │ │  (FastAPI)         │ │---

- Date: YYYY-MM-DD format

│ │ └────────────────┘ │                       │ └────────────┘ │ │

### Security

- Corporate SSO (OAuth 2.0)│ │                    │                       │                │ │                │ └────────────────────┘ │

- Three roles: Store Manager, Analyst, Admin

- VPN or whitelisted IPs│ │ ┌────────────────┐ │                       │ ┌────────────┐ │ │

- PCI DSS, GDPR compliance

│ │ │ Update/Edit    │ │                       │ │ Streaming  │ │ │                │                        │The server will start on port 8080 by default. You can change the port by setting the `PORT` environment variable:

### Operational

- 24/7 monitoring│ │ │ API (CRUD)     │ │                       │ │ Parser     │ │ │

- Daily backups (30-day retention)

- Monthly maintenance windows│ │ └────────────────┘ │                       │ └────────────┘ │ │                │ ┌────────────────────┐ │

- 99.9% SLA

│ │                    │                       │                │ │

---

│ │ ┌────────────────┐ │                       │ ┌────────────┐ │ │                │ │ Go Ingestion Svc   │ │## 🗺️ Context Diagram

## 🚀 Getting Started

│ │ │ Authentication │ │                       │ │ DB Batch   │ │ │

### Quick Start (Docker Compose)

│ │ │ & Authorization│ │                       │ │ Insert     │ │ │                │ │ (Gin Framework)    │ │

```powershell

# Start all services│ │ └────────────────┘ │                       │ └────────────┘ │ │

docker-compose up -d

│ └─────────┬──────────┘                       └────────┬───────┘ │                │ └────────────────────┘ │```bash

# View logs

docker-compose logs -f│           │                                           │         │



# Stop services└───────────┼───────────────────────────────────────────┼─────────┘                └────────────┬───────────┘

docker-compose down

```            │                                           │



**Services:**            │              ┌────────────────────────────┘                             │```# Windows PowerShell

- Angular: http://localhost:4200

- Python API: http://localhost:8000            │              │

- Go Ingestion: http://localhost:8080

            │              │     ┌──────────────────────┐          ┌──────────────────┴──────────────────┐

### Manual Setup - Go Service

            │              │     │ Storage Event Trigger│

```powershell

# Clone repository            │              │     │ (S3 Event / Watcher) │          │                                     │┌──────────────────────────────────────────────────────────────────┐$env:PORT="3000"; go run main.go

git clone https://github.com/RKShukla1997/Pricing_Management_System.git

cd golang-project-product-service            │              │     └──────────┬───────────┘



# Install dependencies            │              │                │          ▼                                     ▼

go mod download

┌───────────┼──────────────┼────────────────┼───────────────────────┐

# Configure environment

cp .env.example .env│           │       DATA LAYER              │                       │ ┌────────────────────┐               ┌────────────────────┐│                     External Actors                               │



# Run service│           │              │                │                       │

go run cmd/server/main.go

```│           ▼              ▼                ▼                       │ │ Relational Database│               │ Object Storage     │



### Test Upload│  ┌─────────────────────────────────────────────────┐             │



```powershell│  │         Object Storage (S3 / Local FS)          │             │ │   (PostgreSQL)     │               │   (S3 / Local)     │├──────────────────────────────────────────────────────────────────┤# Linux/Mac

# Create test CSV

@"│  │                                                  │             │

Store ID,SKU,Product Name,Price,Date

ST001,SKU12345,Laptop,999.99,2026-02-17│  │  1. User uploads CSV via UI                     │             │ │                    │               │                    │

ST001,SKU12346,Mouse,29.99,2026-02-17

"@ | Out-File -FilePath test.csv -Encoding utf8│  │  2. Python API generates presigned URL          │             │



# Upload│  │  3. Direct upload to S3 from browser            │             │ │ - Pricing Records  │               │ - CSV Files        ││  👤 Store Managers  │  👤 Pricing Analysts  │  👤 Administrators │PORT=3000 go run main.go

Invoke-WebRequest -Uri "http://localhost:8080/upload" `

    -Method Post -Form @{file = Get-Item -Path "test.csv"}│  │  4. S3 event triggers Go Ingestion Service      │             │

```

│  │  5. File stored with metadata                   │             │ │ - Audit Logs       │               │ - File Archive     │

**Response:**

```json│  │                                                  │             │

{

  "message": "File uploaded successfully",│  │  Structure: /uploads/YYYY/MM/DD/filename.csv    │             │ │ - User Sessions    │               │ - Backups          │└──────────────┬───────────────────┬─────────────────┬─────────────┘```

  "filename": "1708171234_test.csv",

  "size": 128,│  └──────────────────────────────────────────────────┘             │

  "rows": 3,

  "columns": 5│                                                                   │ └────────────────────┘               └────────────────────┘

}

```│  ┌──────────────────────────────────────────────────┐            │



---│  │      Relational Database (PostgreSQL)            │            │```               │                   │                 │



## 📦 Repository Structure│  │                                                   │            │



```│  │  Tables:                                          │            │

golang-project-product-service/

├── cmd/server/main.go           # Entry point│  │  ├─ pricing_records                              │            │

├── internal/

│   ├── handlers/                # HTTP handlers│  │  │  ├─ id (PK)                                   │            │### System Boundaries               └───────────────────┼─────────────────┘## API Endpoints

│   ├── middleware/              # Auth, CORS, logging

│   ├── models/                  # Data models│  │  │  ├─ store_id (Indexed)                        │            │

│   ├── services/                # Business logic

│   └── config/                  # Configuration│  │  │  ├─ sku (Indexed)                             │            │

├── pkg/utils/                   # Utilities

├── uploads/                     # Temporary storage│  │  │  ├─ product_name                              │            │

├── tests/                       # Unit & integration tests

├── go.mod│  │  │  ├─ price                                     │            │- **External Actors:** Store managers, pricing analysts, administrators                                   │

├── Dockerfile

└── README.md│  │  │  ├─ date (Indexed)                            │            │

```

│  │  │  ├─ created_at                                │            │- **System Boundary:** Web UI + Backend Services + Data Stores

---

│  │  │  └─ updated_at                                │            │

## 🧪 Testing

│  │  │                                                │            │- **External Systems:** Corporate SSO, Monitoring Systems, Backup Services                                   ▼### 1. Upload CSV File

```powershell

# Run all tests│  │  ├─ upload_history                               │            │

go test ./... -v

│  │  │  ├─ id (PK)                                   │            │

# With coverage

go test ./... -v -cover -coverprofile=coverage.out│  │  │  ├─ filename                                  │            │



# View coverage│  │  │  ├─ upload_date                               │            │---        ┌──────────────────────────────────────────────────┐

go tool cover -html=coverage.out

```│  │  │  ├─ status                                    │            │



---│  │  │  ├─ records_count                             │            │



## 🚢 Deployment│  │  │  └─ user_id                                   │            │



### Docker│  │  │                                                │            │## 🏗️ Solution Architecture        │        Angular Frontend (SPA)                     │**Endpoint:** `POST /upload`

```powershell

docker build -t go-ingestion:latest .│  │  └─ audit_logs                                   │            │

docker run -d -p 8080:8080 go-ingestion:latest

```│  │     ├─ id (PK)                                   │            │



### Kubernetes│  │     ├─ table_name                                │            │

```powershell

kubectl apply -f k8s/deployment.yaml│  │     ├─ record_id                                 │            │### Detailed Architecture with Data Flow        │  - File Upload Interface                          │

kubectl scale deployment go-ingestion --replicas=3

```│  │     ├─ action (INSERT/UPDATE/DELETE)             │            │



---│  │     ├─ old_value                                 │            │



## 🔗 Related Repositories│  │     ├─ new_value                                 │            │



- [Angular Frontend](https://github.com/RKShukla1997/frontend)│  │     ├─ user_id                                   │            │```        │  - Search & Filter UI                             │**Content-Type:** `multipart/form-data`

- [Python Data Service](https://github.com/RKShukla1997/data-service)

- [Infrastructure (K8s/Terraform)](https://github.com/RKShukla1997/infrastructure)│  │     └─ timestamp                                 │            │



---│  └──────────────────────────────────────────────────┘            │┌─────────────────────────────────────────────────────────────────┐



## 📄 License│                                                                   │



MIT License│  ┌──────────────────────────────────────────────────┐            ││                         PRESENTATION LAYER                       │        │  - Data Grid with Edit Capabilities               │



---│  │           Cache Layer (Redis) - Optional         │            │



**Repository:** [Pricing_Management_System](https://github.com/RKShukla1997/Pricing_Management_System)  │  │                                                   │            ││                                                                  │

**Last Updated:** February 17, 2026

│  │  - Session Storage                                │            │

│  │  - Search Results Cache                           │            ││                    ┌────────────────────┐                        │        └────────────────┬──────────────┬──────────────────┘**Parameters:**

│  │  - Rate Limiting                                  │            │

│  └──────────────────────────────────────────────────┘            ││                    │     Angular UI     │                        │

└───────────────────────────────────────────────────────────────────┘

```│                    │  (Port: 4200)      │                        │                         │              │- `file`: CSV file (required)



### Data Flow Sequences│                    │                    │                        │



#### Upload Flow│                    │ - Material Design  │                        │                ┌────────┘              └────────┐

```

1. User → Angular UI: Select CSV file│                    │ - State Management │                        │

2. Angular → Python API: Request presigned URL

3. Python API → S3: Generate presigned URL│                    │ - Form Validation  │                        │                │                                 │**Example using curl:**

4. Python API → Angular: Return presigned URL

5. Angular → S3: Direct upload using presigned URL│                    └─────────┬──────────┘                        │

6. S3 → Event Trigger: File upload complete event

7. Event Trigger → Go Ingestion Svc: Trigger processing└──────────────────────────────┼───────────────────────────────────┘                ▼                                 ▼```bash

8. Go Ingestion Svc → S3: Download and stream file

9. Go Ingestion Svc: Validate CSV structure                               │

10. Go Ingestion Svc: Parse CSV (streaming)

11. Go Ingestion Svc → PostgreSQL: Batch insert records                               │ HTTP/REST    ┌───────────────────────┐       ┌─────────────────────────┐curl -X POST http://localhost:8080/upload -F "file=@data.csv"

12. Go Ingestion Svc → PostgreSQL: Update upload_history

13. PostgreSQL → Python API: Notify completion                               │

14. Python API → Angular: Push notification (WebSocket)

15. Angular: Display success message┌──────────────────────────────┼───────────────────────────────────┐    │  Go Upload Service    │       │  Python Data Service    │```

```

│                       APPLICATION LAYER                          │

#### Search Flow

```│                               │                                  │    │  (Port: 8080)         │       │  (Port: 8000)           │

1. User → Angular UI: Enter search criteria

2. Angular → Python API: POST /api/v1/pricing/search│     ┌─────────────────────────┴─────────────────────────┐       │

3. Python API → Redis: Check cache

4. Redis → Python API: Cache miss│     │                                                     │       │    │  - CSV Upload         │       │  - Search/Query         │**Example using PowerShell:**

5. Python API → PostgreSQL: Execute search query

6. PostgreSQL → Python API: Return results│     ▼                                                     ▼       │

7. Python API → Redis: Store in cache

8. Python API → Angular: Return JSON response│ ┌────────────────────┐                       ┌────────────────┐ │    │  - Validation         │       │  - CRUD Operations      │```powershell

9. Angular: Display results in data grid

```│ │  Python API Service│                       │ Go Ingestion   │ │



#### Edit Flow│ │  (FastAPI)         │                       │ Service (Gin)  │ │    │  - File Persistence   │       │  - Data Processing      │$uri = "http://localhost:8080/upload"

```

1. User → Angular UI: Edit pricing record inline│ │  Port: 8000        │                       │ Port: 8080     │ │

2. Angular → Python API: PUT /api/v1/pricing/{id}

3. Python API: Validate changes│ │                    │                       │                │ │    └──────────┬────────────┘       └───────────┬─────────────┘$filePath = "C:\path\to\your\file.csv"

4. Python API → PostgreSQL: Begin transaction

5. Python API → PostgreSQL: Update pricing_record│ │ ┌────────────────┐ │                       │ ┌────────────┐ │ │

6. Python API → PostgreSQL: Insert audit_log

7. Python API → PostgreSQL: Commit transaction│ │ │ Generate       │ │                       │ │ Event      │ │ │               │                                 │$form = @{

8. PostgreSQL → Python API: Confirm update

9. Python API → Redis: Invalidate cache│ │ │ Presigned URL  │ │                       │ │ Listener   │ │ │

10. Python API → Angular: Return updated record

11. Angular: Update UI with new values│ │ └────────────────┘ │                       │ └────────────┘ │ │               └────────────┬────────────────────┘    file = Get-Item -Path $filePath

```

│ │                    │                       │                │ │

---

│ │ ┌────────────────┐ │                       │ ┌────────────┐ │ │                            │}

## 🛠️ Technology Stack

│ │ │ Search API     │ │                       │ │ CSV        │ │ │

### Frontend

- **Framework:** Angular 17+│ │ │ (Multi-Criteria)│ │                      │ │ Validation │ │ │                            ▼Invoke-RestMethod -Uri $uri -Method Post -Form $form

- **UI Library:** Angular Material / PrimeNG

- **State Management:** NgRx / Akita│ │ └────────────────┘ │                       │ └────────────┘ │ │

- **HTTP Client:** Angular HttpClient

- **Form Validation:** Reactive Forms│ │                    │                       │                │ │              ┌──────────────────────────┐```



### Backend Services│ │ ┌────────────────┐ │                       │ ┌────────────┐ │ │



#### Python API Service (Data Operations)│ │ │ Update/Edit    │ │                       │ │ Streaming  │ │ │              │   Database Layer         │

- **Language:** Python 3.11+

- **Framework:** FastAPI│ │ │ API (CRUD)     │ │                       │ │ Parser     │ │ │

- **ORM:** SQLAlchemy

- **Data Processing:** Pandas│ │ └────────────────┘ │                       │ └────────────┘ │ │              │   - PostgreSQL (Primary) │**Success Response (200 OK):**

- **Validation:** Pydantic

- **Responsibilities:**│ │                    │                       │                │ │

  - Generate presigned URLs for S3 uploads

  - Search and filter pricing records│ │ ┌────────────────┐ │                       │ ┌────────────┐ │ │              │   - Redis (Cache)        │```json

  - CRUD operations on pricing data

  - User authentication and authorization│ │ │ Authentication │ │                       │ │ DB Batch   │ │ │



#### Go Ingestion Service (File Processing)│ │ │ & Authorization│ │                       │ │ Insert     │ │ │              └──────────────────────────┘{

- **Language:** Go 1.25+

- **Framework:** Gin Web Framework│ │ └────────────────┘ │                       │ └────────────┘ │ │

- **CSV Processing:** encoding/csv (streaming)

- **Concurrency:** Goroutines│ └─────────┬──────────┘                       └────────┬───────┘ │                            │  "message": "File uploaded successfully",

- **Responsibilities:**

  - Listen to S3 storage events│           │                                           │         │

  - Validate CSV file structure

  - Stream and parse large CSV files└───────────┼───────────────────────────────────────────┼─────────┘                            ▼  "filename": "1739151234_data.csv",

  - Batch insert records into database

            │                                           │

### Data Layer

- **Primary Database:** PostgreSQL 15+            │              ┌────────────────────────────┘              ┌──────────────────────────┐  "size": 1024,

- **Cache:** Redis 7+ (Optional)

- **Object Storage:** AWS S3 / Azure Blob / Local FileSystem            │              │



### DevOps            │              │     ┌──────────────────────┐              │   Storage Layer          │  "rows": 100,

- **Containerization:** Docker

- **Orchestration:** Kubernetes / Docker Compose            │              │     │ Storage Event Trigger│

- **CI/CD:** GitHub Actions / GitLab CI

- **Monitoring:** Prometheus + Grafana            │              │     │ (S3 Event / Watcher) │              │   - S3/Blob Storage      │  "columns": 5

- **Logging:** ELK Stack

            │              │     └──────────┬───────────┘

---

            │              │                │              │   - File Archive         │}

## 🎨 Design Decisions

┌───────────┼──────────────┼────────────────┼───────────────────────┐

### 1. Microservices Architecture with Language Optimization

│           │       DATA LAYER              │                       │              └──────────────────────────┘```

**Decision:** Use Go for ingestion and Python for API operations

│           │              │                │                       │

**Rationale:**

- **Go for File Processing:**│           ▼              ▼                ▼                       │```

  - Superior performance for I/O operations

  - Efficient memory management for large files│  ┌─────────────────────────────────────────────────┐             │

  - Native goroutines for concurrent processing

  - Fast CSV parsing with streaming│  │         Object Storage (S3 / Local FS)          │             │**Error Responses:**

  - Single binary deployment

  │  │                                                  │             │

- **Python for Data Operations:**

  - Rich ecosystem for data manipulation (Pandas)│  │  1. User uploads CSV via UI                     │             │---- `400 Bad Request`: File too large, invalid CSV format, or not a CSV file

  - FastAPI for rapid API development

  - SQLAlchemy for complex queries│  │  2. Python API generates presigned URL          │             │

  - Better for business logic

│  │  3. Direct upload to S3 from browser            │             │- `500 Internal Server Error`: Server error while saving file

**Trade-offs:**

- Need to maintain two codebases│  │  4. S3 event triggers Go Ingestion Service      │             │

- Different deployment pipelines

- Team needs expertise in both languages│  │  5. File stored with metadata                   │             │## 🏗️ Solution Architecture



### 2. Event-Driven Upload Architecture│  │                                                  │             │



**Decision:** Use presigned URLs and storage events instead of direct upload to backend│  │  Structure: /uploads/YYYY/MM/DD/filename.csv    │             │### 2. Health Check



**Rationale:**│  └──────────────────────────────────────────────────┘             │

- **Scalability:** Direct S3 upload bypasses backend bottleneck

- **Performance:** No file data through backend servers│                                                                   │### High-Level Architecture

- **Bandwidth:** Reduced backend bandwidth costs

- **Security:** Presigned URLs with expiration (5 minutes)│  ┌──────────────────────────────────────────────────┐            │

- **Decoupling:** Upload and processing are independent

│  │      Relational Database (PostgreSQL)            │            │**Endpoint:** `GET /health`

**Flow:**

1. Frontend requests presigned URL from Python API│  │                                                   │            │

2. Frontend uploads directly to S3

3. S3 event triggers Go ingestion service│  │  Tables:                                          │            │```

4. Go service processes asynchronously

│  │  ├─ pricing_records                              │            │

**Trade-offs:**

- More complex upload flow│  │  │  ├─ id (PK)                                   │            │┌─────────────────────────────────────────────────────────────────┐**Example:**

- Requires S3 event configuration

- Need to handle failed uploads│  │  │  ├─ store_id (Indexed)                        │            │



### 3. Streaming CSV Processing│  │  │  ├─ sku (Indexed)                             │            ││                         Load Balancer / API Gateway             │```bash



**Decision:** Use streaming parser instead of loading entire file into memory│  │  │  ├─ product_name                              │            │



**Rationale:**│  │  │  ├─ price                                     │            ││                    (Rate Limiting, Auth, CORS)                  │curl http://localhost:8080/health

- **Memory Efficiency:** Process 10MB+ files with constant memory

- **Performance:** Start processing immediately│  │  │  ├─ date (Indexed)                            │            │

- **Scalability:** Handle unlimited file sizes

- **Error Recovery:** Stop on first error│  │  │  ├─ created_at                                │            │└────────────────────────────┬────────────────────────────────────┘```



**Implementation:**│  │  │  └─ updated_at                                │            │

```go

// Stream CSV line by line│  │  │                                                │            │                             │

reader := csv.NewReader(file)

for {│  │  ├─ upload_history                               │            │

    record, err := reader.Read()

    if err == io.EOF {│  │  │  ├─ id (PK)                                   │            │        ┌────────────────────┼────────────────────┐**Response (200 OK):**

        break

    }│  │  │  ├─ filename                                  │            │

    processBatch(record)

}│  │  │  ├─ upload_date                               │            │        │                    │                    │```json

```

│  │  │  ├─ status                                    │            │

### 4. PostgreSQL Over NoSQL

│  │  │  ├─ records_count                             │            │        ▼                    ▼                    ▼{

**Decision:** Use PostgreSQL as primary database

│  │  │  └─ user_id                                   │            │

**Rationale:**

- **ACID Compliance:** Critical for financial pricing data│  │  │                                                │            │┌──────────────┐    ┌──────────────┐    ┌──────────────┐  "status": "healthy",

- **Complex Queries:** Support for JOINs, aggregations, window functions

- **Data Integrity:** Foreign keys, constraints, triggers│  │  └─ audit_logs                                   │            │

- **Full-Text Search:** Built-in search capabilities

- **Proven at Scale:** Battle-tested with large datasets│  │     ├─ id (PK)                                   │            ││ Angular SPA  │    │   Go Upload  │    │   Python     │  "timestamp": "2026-02-10T12:00:00Z"



| Feature | PostgreSQL | MongoDB |│  │     ├─ table_name                                │            │

|---------|-----------|---------|

| ACID | ✅ Strong | ⚠️ Weak |│  │     ├─ record_id                                 │            ││   (Static)   │    │   Service    │    │ Data Service │}

| Joins | ✅ Excellent | ❌ Limited |

| Schema | ✅ Enforced | ⚠️ Flexible |│  │     ├─ action (INSERT/UPDATE/DELETE)             │            │

| Scale | ✅ Vertical + Horizontal | ✅ Horizontal |

│  │     ├─ old_value                                 │            ││              │    │              │    │              │```

### 5. Batch Insert for Database Operations

│  │     ├─ new_value                                 │            │

**Decision:** Batch insert records instead of individual inserts

│  │     ├─ user_id                                   │            ││ - Angular 17 │    │ - Gin HTTP   │    │ - FastAPI    │

**Rationale:**

- **Performance:** 10-100x faster than single inserts│  │     └─ timestamp                                 │            │

- **Network:** Reduced round trips to database

- **Transactions:** Single transaction for batch│  └──────────────────────────────────────────────────┘            ││ - Material   │    │ - CSV Parse  │    │ - SQLAlchemy │## File Storage

- **Error Handling:** Rollback entire batch on error

│                                                                   │

**Implementation:**

```go│  ┌──────────────────────────────────────────────────┐            ││ - State Mgmt │    │ - Validation │    │ - Pandas     │

// Batch size: 1000 records

const batchSize = 1000│  │           Cache Layer (Redis) - Optional         │            │

for i := 0; i < len(records); i += batchSize {

    batch := records[i:min(i+batchSize, len(records))]│  │                                                   │            │└──────────────┘    └──────┬───────┘    └──────┬───────┘Uploaded files are stored in the `./uploads` directory with a timestamp prefix to ensure uniqueness.

    db.BatchInsert(batch)

}│  │  - Session Storage                                │            │

```

│  │  - Search Results Cache                           │            │                           │                    │

### 6. Optimistic Locking for Concurrent Edits

│  │  - Rate Limiting                                  │            │

**Decision:** Use version column for optimistic locking

│  └──────────────────────────────────────────────────┘            │                           └──────────┬─────────┘## Testing

**Rationale:**

- **Performance:** No locks, better concurrency└───────────────────────────────────────────────────────────────────┘

- **User Experience:** Users notified only on actual conflicts

- **Scalability:** Works in distributed systems```                                      │



**Implementation:**

```python

# Check version before update### Data Flow Sequence                    ┌─────────────────┼─────────────────┐### Create a test CSV file:

UPDATE pricing_records

SET price = $1, version = version + 1

WHERE id = $2 AND version = $3

# If affected rows = 0, conflict detected#### Upload Flow                    │                 │                 │

```

```

### 7. Direct S3 Upload from Frontend

1. User → Angular UI: Select CSV file                    ▼                 ▼                 ▼**PowerShell:**

**Decision:** Upload files directly from browser to S3

2. Angular → Python API: Request presigned URL

**Rationale:**

- **Performance:** No backend involvement in file transfer3. Python API → S3: Generate presigned URL            ┌──────────────┐  ┌─────────────┐  ┌──────────────┐```powershell

- **Scalability:** Backend not bottleneck for uploads

- **Cost:** Reduced bandwidth costs4. Python API → Angular: Return presigned URL

- **User Experience:** Progress bar works directly

5. Angular → S3: Direct upload using presigned URL            │  PostgreSQL  │  │    Redis    │  │  S3 Storage  │@"

**Security:**

- Presigned URLs with 5-minute expiration6. S3 → Event Trigger: File upload complete event

- Limit file size at S3 level

- Restrict to specific content type7. Event Trigger → Go Ingestion Svc: Trigger processing            │   Database   │  │    Cache    │  │  (CSV Files) │name,email,age



---8. Go Ingestion Svc → S3: Download and stream file



## 🚀 Non-Functional Requirements9. Go Ingestion Svc: Validate CSV structure            │              │  │             │  │              │John Doe,john@example.com,30



### 1. Performance10. Go Ingestion Svc: Parse CSV (streaming)



**Requirements:**11. Go Ingestion Svc → PostgreSQL: Batch insert records            │ - Pricing    │  │ - Sessions  │  │ - Archive    │Jane Smith,jane@example.com,25

- Response Time: < 2 seconds for search queries

- Upload Speed: Support 10MB files in < 5 seconds12. Go Ingestion Svc → PostgreSQL: Update upload_history

- Concurrent Users: Support 1000+ concurrent users

- Throughput: Process 100+ CSV uploads per minute13. PostgreSQL → Python API: Notify completion            │   Records    │  │ - Search    │  │ - Backup     │"@ | Out-File -FilePath test.csv -Encoding utf8

- Database Queries: < 100ms for indexed lookups

14. Python API → Angular: Push notification (WebSocket)

**Design Approach:**

- Database indexing on store_id, sku, date columns15. Angular: Display success message            │ - Indexes    │  │   Results   │  │              │```

- Redis caching for frequent searches (TTL: 5 minutes)

- Connection pooling (50-100 connections)```

- Asynchronous processing for large files

- CDN for static assets            └──────────────┘  └─────────────┘  └──────────────┘



### 2. Scalability#### Search Flow



**Requirements:**``````### Upload the test file:

- Horizontal Scaling: Stateless microservices

- Data Volume: Handle 100M+ pricing records1. User → Angular UI: Enter search criteria

- Store Growth: Support 3000+ stores, expandable to 10,000+

- Multi-Region: Deploy across multiple geographic regions2. Angular → Python API: POST /api/v1/pricing/search```powershell



**Design Approach:**3. Python API → Redis: Check cache

- Microservices architecture

- Database sharding by region/country4. Redis → Python API: Cache miss### Microservices ArchitectureInvoke-RestMethod -Uri "http://localhost:8080/upload" -Method Post -Form @{file = Get-Item -Path "test.csv"}

- Read replicas for query distribution (3 replicas)

- Auto-scaling based on CPU/memory (50-80% threshold)5. Python API → PostgreSQL: Execute search query

- Message queues for async tasks (RabbitMQ/SQS)

6. PostgreSQL → Python API: Return results```

### 3. Availability & Reliability

7. Python API → Redis: Store in cache

**Requirements:**

- Uptime: 99.9% availability (< 8.76 hours downtime/year)8. Python API → Angular: Return JSON response1. **Frontend Service (Angular)**

- Disaster Recovery: RTO < 4 hours, RPO < 1 hour

- Fault Tolerance: No single point of failure9. Angular: Display results in data grid

- Data Integrity: Zero data loss

```   - Single Page Application## Building for Production

**Design Approach:**

- Multi-AZ deployment (3 availability zones)

- Database replication (primary + 2 replicas)

- Service redundancy (min 3 instances per service)#### Edit Flow   - Responsive UI with Material Design

- Health checks and circuit breakers

- Daily automated backups with 30-day retention```



### 4. Security1. User → Angular UI: Edit pricing record inline   - Client-side validation```bash



**Requirements:**2. Angular → Python API: PUT /api/v1/pricing/{id}

- Authentication: OAuth 2.0 / JWT tokens

- Authorization: Role-based access control (RBAC)3. Python API: Validate changes   - Real-time updatesgo build -o product-service.exe main.go

- Data Encryption: TLS 1.3 in transit, AES-256 at rest

- Audit Logging: All actions logged and immutable4. Python API → PostgreSQL: Begin transaction

- Compliance: GDPR, SOC 2 compliance

5. Python API → PostgreSQL: Update pricing_record```

**Design Approach:**

- API Gateway with authentication6. Python API → PostgreSQL: Insert audit_log

- JWT tokens (1-hour expiration)

- Encrypted database connections7. Python API → PostgreSQL: Commit transaction2. **Upload Service (Go + Gin)**

- Secrets management (AWS Secrets Manager/Vault)

- Input validation and sanitization8. PostgreSQL → Python API: Confirm update

- Parameterized queries (prevent SQL injection)

9. Python API → Redis: Invalidate cache   - High-performance file upload handlingRun the executable:

### 5. Maintainability

10. Python API → Angular: Return updated record

**Requirements:**

- Code Quality: > 80% test coverage11. Angular: Update UI with new values   - CSV validation and parsing```bash

- Documentation: API docs, architecture diagrams

- Monitoring: Real-time metrics and alerts```

- Logging: Centralized structured logging

   - Concurrent processing./product-service.exe

**Design Approach:**

- Comprehensive unit and integration tests---

- OpenAPI/Swagger documentation

- Prometheus metrics + Grafana dashboards   - File storage management```

- ELK stack for log aggregation

- CI/CD pipelines with automated testing## 🛠️ Technology Stack



### 6. Usability



**Requirements:**### Frontend

- Responsive Design: Mobile, tablet, desktop support

- Accessibility: WCAG 2.1 Level AA compliance- **Framework:** Angular 17+3. **Data Service (Python + FastAPI)**## Configuration

- Internationalization: Multi-language support

- User Experience: Intuitive interface, < 3 clicks to any feature- **UI Library:** Angular Material / PrimeNG



**Design Approach:**- **State Management:** NgRx / Akita   - RESTful API for CRUD operations

- Angular Material responsive components

- ARIA labels and keyboard navigation- **HTTP Client:** Angular HttpClient

- i18n with Angular's built-in support

- Loading indicators and progress bars- **Form Validation:** Reactive Forms   - Complex search queries- **MAX_UPLOAD_SIZE**: Default is 10MB (can be modified in `main.go`)

- Error messages with actionable guidance

- **File Upload:** ng-file-upload / custom directive

### 7. Data Consistency

   - Data transformation- **UPLOAD_PATH**: Default is `./uploads` (can be modified in `main.go`)

**Requirements:**

- ACID Compliance: Transactional integrity### Backend Services

- Eventual Consistency: For cross-region sync

- Data Validation: Strong type checking   - Business logic processing- **PORT**: Set via environment variable (default: 8080)



**Design Approach:**#### 1. Python API Service (Data Operations)

- PostgreSQL transactions

- Optimistic locking for concurrent edits- **Language:** Python 3.11+

- Data validation at multiple layers (frontend, backend, database)

- Conflict resolution strategies- **Framework:** FastAPI



---- **ORM:** SQLAlchemy---## Project Structure



## 📝 Assumptions- **Data Processing:** Pandas



### Business Assumptions- **Validation:** Pydantic

1. **Store Count:** Starting with 3000 stores, growing to 10,000 in 3 years

2. **Upload Frequency:** Each store uploads pricing data daily- **Authentication:** JWT / OAuth2

3. **File Size:** Average CSV file contains 1,000-10,000 records (< 10MB)

4. **Data Retention:** Pricing history retained for 2 years minimum- **Key Responsibilities:**## 🛠️ Technology Stack```

5. **Peak Load:** 20% of stores upload simultaneously during business hours

6. **User Base:** 500 concurrent users (store managers, analysts, admins)  - Generate presigned URLs for S3 uploads

7. **Pricing Changes:** 10% of records updated monthly via edit feature

  - Search and filter pricing records.

### Technical Assumptions

1. **Network:** Minimum 1 Mbps upload speed at store locations  - CRUD operations on pricing data

2. **Browser Support:** Modern browsers (Chrome, Firefox, Safari, Edge - last 2 versions)

3. **Database Size:** Initial ~50GB, growing to 500GB in 3 years  - User authentication and authorization### Frontend├── main.go           # Main application file with HTTP handlers

4. **API Calls:** Average 1000 API requests per minute during peak hours

5. **Search Queries:** 70% use 1-2 criteria, 30% use 3+ criteria  - Real-time notifications (WebSocket)

6. **Edit Operations:** 80% single record edits, 20% bulk updates

7. **Infrastructure:** Cloud-hosted (AWS/Azure/GCP)- **Framework:** Angular 17+├── go.mod            # Go module definition



### Data Assumptions#### 2. Go Ingestion Service (File Processing)

1. **CSV Format:** Consistent column order: `Store ID, SKU, Product Name, Price, Date`

2. **Store ID:** Alphanumeric, max 20 characters- **Language:** Go 1.25+- **UI Library:** Angular Material / PrimeNG├── go.sum            # Go module checksums

3. **SKU:** Alphanumeric, max 50 characters

4. **Product Name:** UTF-8 string, max 200 characters- **Framework:** Gin Web Framework

5. **Price:** Decimal with 2 decimal places, positive values, range: 0.01 - 999,999.99

6. **Date:** ISO 8601 format (YYYY-MM-DD)- **CSV Processing:** encoding/csv (streaming)- **State Management:** NgRx / Akita├── README.md         # This file

7. **Character Encoding:** UTF-8

8. **Delimiter:** Comma (,)- **Concurrency:** Goroutines for parallel processing

9. **Header Row:** First row contains column names

10. **No Empty Fields:** All fields are required- **Validation:** Custom validators- **HTTP Client:** Angular HttpClient├── .gitignore        # Git ignore file



### Security Assumptions- **Key Responsibilities:**

1. **Authentication:** Corporate SSO integration (SAML/OAuth 2.0)

2. **Authorization:** Three roles:  - Listen to S3 storage events- **Form Validation:** Reactive Forms└── uploads/          # Directory for uploaded files (created automatically)

   - **Store Manager:** Upload files, view own store data

   - **Analyst:** Search, view, edit all data  - Validate CSV file structure

   - **Admin:** All permissions + user management

3. **Network:** Application accessible via corporate VPN or whitelisted IPs  - Stream and parse large CSV files```

4. **Compliance:** PCI DSS for pricing data, GDPR for EU stores

  - Batch insert records into database

### Operational Assumptions

1. **Monitoring:** 24/7 monitoring with on-call support  - Error handling and retry logic### Backend Services

2. **Backup:** Daily automated backups at 2 AM (off-peak)

3. **Backup Retention:** 30 days online, 2 years archived

4. **Updates:** Monthly maintenance windows (4 hours max, Saturday 2-6 AM)

5. **Support:** Email and ticket-based support (24-hour response time)### Data Layer#### Upload Service (Go)

6. **SLA:** 99.9% uptime

- **Primary Database:** PostgreSQL 15+- **Language:** Go 1.25+

---

  - ACID compliance for pricing data- **Framework:** Gin Web Framework

## 🚀 Getting Started

  - Full-text search capabilities- **CSV Processing:** encoding/csv

### Prerequisites

  - Partitioning for large datasets- **File Handling:** io, os packages

- **Go** 1.25+

- **Python** 3.11+  - Replication for high availability- **Validation:** Custom middleware

- **PostgreSQL** 15+

- **Redis** 7+ (optional)

- **Docker** and Docker Compose (recommended)

- **Cache:** Redis 7+ (Optional)#### Data Service (Python)

### Quick Start with Docker Compose

  - Session storage- **Language:** Python 3.11+

```powershell

# Clone repository  - Search results caching- **Framework:** FastAPI / Flask

git clone https://github.com/yourusername/retail-pricing-system.git

cd retail-pricing-system  - Rate limiting- **ORM:** SQLAlchemy



# Start all services- **Data Processing:** Pandas

docker-compose up -d

- **Object Storage:** AWS S3 / Azure Blob / Local FileSystem- **Validation:** Pydantic

# View logs

docker-compose logs -f  - CSV file storage



# Stop services  - File archival### Data Layer

docker-compose down

```  - Event notifications- **Primary Database:** PostgreSQL 15+



**Services available at:**- **Cache:** Redis 7+

- Frontend: http://localhost:4200

- Python API: http://localhost:8000### DevOps & Infrastructure- **Object Storage:** AWS S3 / Azure Blob Storage

- Go Ingestion: http://localhost:8080

- PostgreSQL: localhost:5432- **Containerization:** Docker- **Message Queue:** RabbitMQ / AWS SQS (for async processing)

- Redis: localhost:6379

- **Orchestration:** Kubernetes / Docker Compose

### Manual Setup - Go Ingestion Service

- **CI/CD:** GitHub Actions / GitLab CI### DevOps & Infrastructure

```powershell

# Clone this repository- **Monitoring:** Prometheus + Grafana- **Containerization:** Docker

git clone https://github.com/yourusername/golang-project-product-service.git

cd golang-project-product-service- **Logging:** ELK Stack (Elasticsearch, Logstash, Kibana)- **Orchestration:** Kubernetes



# Install dependencies- **API Gateway:** Kong / Nginx (Optional)- **CI/CD:** GitHub Actions / GitLab CI

go mod download

- **Monitoring:** Prometheus + Grafana

# Configure environment

cp .env.example .env---- **Logging:** ELK Stack (Elasticsearch, Logstash, Kibana)

notepad .env

- **API Gateway:** Kong / AWS API Gateway

# Run the service

go run cmd/server/main.go## ✅ Functional Requirements

```

---

**Service runs on:** http://localhost:8080

### 1. CSV Upload and Persistence

### Environment Variables

- ✅ Upload CSV files containing pricing data## ✅ Functional Requirements

```env

# Server- ✅ CSV Structure: `Store ID, SKU, Product Name, Price, Date`

PORT=8080

GIN_MODE=release- ✅ File validation (format, size, structure)### 1. CSV Upload and Persistence



# File Upload- ✅ Automatic parsing and data extraction- ✅ Upload CSV files containing pricing data

MAX_UPLOAD_SIZE=10485760

UPLOAD_PATH=./uploads- ✅ Persistent storage in database- ✅ CSV Structure: `Store ID, SKU, Product Name, Price, Date`



# Database- ✅ File archival in object storage- ✅ File validation (format, size, structure)

DATABASE_URL=postgresql://user:pass@localhost:5432/pricing_db

DB_MAX_CONNECTIONS=50- ✅ Upload history and audit trail- ✅ Automatic parsing and data extraction



# S3 (Optional)- ✅ Support for large files (streaming processing)- ✅ Persistent storage in database

USE_S3=false

AWS_REGION=us-east-1- ✅ Duplicate detection and handling- ✅ File archival in object storage

S3_BUCKET=retail-pricing-uploads

- ✅ Error reporting with line numbers- ✅ Upload history and audit trail

# Security

JWT_SECRET=your-secret-key



# Logging**Implementation:**### 2. Search and Filter Capabilities

LOG_LEVEL=info

- Frontend: Angular file upload with drag-and-drop- ✅ Search by Store ID

# Processing

BATCH_SIZE=1000- Backend: Go service for efficient CSV processing- ✅ Search by SKU

```

- Storage: S3 for files, PostgreSQL for records- ✅ Search by Product Name (partial match)

### Testing the Service

- ✅ Filter by Price range

#### Health Check

```powershell### 2. Search and Filter Capabilities- ✅ Filter by Date range

Invoke-RestMethod -Uri "http://localhost:8080/health" -Method Get

```- ✅ Search by Store ID (exact match)- ✅ Combined search criteria



#### Create Test CSV- ✅ Search by SKU (exact match)- ✅ Pagination and sorting

```powershell

@"- ✅ Search by Product Name (partial match, case-insensitive)- ✅ Export search results

Store ID,SKU,Product Name,Price,Date

ST001,SKU12345,Laptop Computer,999.99,2026-02-17- ✅ Filter by Price range (min/max)

ST001,SKU12346,Wireless Mouse,29.99,2026-02-17

ST002,SKU12347,USB Keyboard,49.99,2026-02-17- ✅ Filter by Date range (start/end)### 3. Edit and Update Records

"@ | Out-File -FilePath pricing_test.csv -Encoding utf8

```- ✅ Combined search criteria (AND logic)- ✅ Inline editing in data grid



#### Upload CSV- ✅ Pagination (configurable page size)- ✅ Bulk update capabilities

```powershell

$response = Invoke-WebRequest -Uri "http://localhost:8080/upload" `- ✅ Sorting (ascending/descending)- ✅ Data validation on updates

    -Method Post `

    -Form @{file = Get-Item -Path "pricing_test.csv"}- ✅ Export search results to CSV- ✅ Change tracking and versioning



$response.Content | ConvertFrom-Json | ConvertTo-Json- ✅ Save search filters as presets- ✅ Rollback functionality

```

- ✅ Audit log of all changes

**Expected Response:**

```json**Implementation:**

{

  "message": "File uploaded successfully",- Frontend: Advanced search form with reactive forms---

  "filename": "1708171234_pricing_test.csv",

  "size": 256,- Backend: Python FastAPI with SQLAlchemy queries

  "rows": 4,

  "columns": 5- Database: Indexed columns for fast lookups## 🚀 Non-Functional Requirements

}

```- Cache: Redis for frequently searched queries



---### 1. Performance



## 📦 Source Implementation### 3. Edit and Update Records- **Response Time:** < 2 seconds for search queries



### Repository Structure- ✅ Inline editing in data grid- **Upload Speed:** Support 10MB files in < 5 seconds



```- ✅ Single record updates- **Concurrent Users:** Support 1000+ concurrent users

golang-project-product-service/        # THIS REPOSITORY

│- ✅ Bulk update capabilities- **Throughput:** Process 100+ CSV uploads per minute

├── cmd/

│   └── server/- ✅ Data validation on updates (type, range, format)- **Database Queries:** < 100ms for indexed lookups

│       └── main.go                    # Application entry point

│- ✅ Change tracking and versioning

├── internal/

│   ├── handlers/- ✅ Audit log of all changes (who, what, when)**Design Approach:**

│   │   ├── upload.go                  # File upload handler

│   │   ├── health.go                  # Health check- ✅ Rollback functionality- Database indexing on Store ID, SKU, and Date

│   │   └── event.go                   # S3 event handler

│   ├── middleware/- ✅ Optimistic locking for concurrent edits- Redis caching for frequent searches

│   │   ├── auth.go                    # Authentication

│   │   ├── cors.go                    # CORS- ✅ Confirmation dialog for destructive operations- Connection pooling

│   │   └── logging.go                 # Request logging

│   ├── models/- Asynchronous processing for large files

│   │   ├── pricing.go                 # Pricing record model

│   │   └── response.go                # API responses**Implementation:**- CDN for static assets

│   ├── services/

│   │   ├── validator.go               # CSV validation- Frontend: Editable data grid with Material Table

│   │   ├── parser.go                  # CSV parsing (streaming)

│   │   ├── storage.go                 # File storage (S3/local)- Backend: Python FastAPI with transaction support### 2. Scalability

│   │   └── database.go                # Database operations

│   └── config/- Database: Audit logs table with triggers- **Horizontal Scaling:** Stateless microservices

│       └── config.go                  # Configuration

│- Validation: Server-side validation with Pydantic- **Data Volume:** Handle 100M+ pricing records

├── pkg/

│   └── utils/- **Store Growth:** Support 3000+ stores, expandable to 10,000+

│       ├── csv.go                     # CSV utilities

│       └── s3.go                      # S3 utilities---- **Multi-Region:** Deploy across multiple geographic regions

│

├── uploads/                           # Temporary storage

│

├── tests/## 🚀 Non-Functional Requirements**Design Approach:**

│   ├── integration/

│   └── unit/- Microservices architecture

│

├── go.mod                             # Go module### 1. Performance- Database sharding by region/country

├── go.sum                             # Dependencies

├── Dockerfile                         # Container image**Requirements:**- Read replicas for query distribution

├── .env.example                       # Environment template

└── README.md                          # This file- Response Time: < 2 seconds for search queries- Auto-scaling based on load

```

- Upload Speed: Support 10MB files in < 5 seconds- Message queues for async tasks

### Related Repositories

- Concurrent Users: Support 1000+ concurrent users

- **Frontend:** [Angular Web Application](https://github.com/yourusername/frontend)

- **Python API Service:** [FastAPI Data Service](https://github.com/yourusername/data-service)- Throughput: Process 100+ CSV uploads per minute### 3. Availability & Reliability

- **Infrastructure:** [Kubernetes & Terraform](https://github.com/yourusername/infrastructure)

- **Database:** [Schema & Migrations](https://github.com/yourusername/database)- Database Queries: < 100ms for indexed lookups- **Uptime:** 99.9% availability (< 8.76 hours downtime/year)



### API Endpoints- **Disaster Recovery:** RTO < 4 hours, RPO < 1 hour



#### Go Ingestion Service (Port 8080)**Design Approach:**- **Fault Tolerance:** No single point of failure



**POST /upload**- **Database Optimization:**- **Data Integrity:** Zero data loss

- Upload CSV file with pricing data

- Max file size: 10MB  - B-tree indexes on store_id, sku, date columns

- Format: multipart/form-data

- Returns: Upload metadata (filename, size, rows, columns)  - Composite indexes for common query patterns**Design Approach:**



**GET /health**  - Query optimization with EXPLAIN ANALYZE- Multi-AZ deployment

- Health check endpoint

- Returns: Service status and timestamp  - Connection pooling (50-100 connections)- Database replication (primary-replica)



#### Python Data Service (Port 8000)  - Regular automated backups



**POST /api/v1/upload/presigned-url**- **Caching Strategy:**- Health checks and circuit breakers

- Generate presigned URL for S3 upload

  - Redis for search results (TTL: 5 minutes)- Retry mechanisms with exponential backoff

**GET /api/v1/pricing**

- Search pricing records with filters  - Browser caching for static assets



**PUT /api/v1/pricing/{id}**  - API response caching with ETag### 4. Security

- Update pricing record

  - **Authentication:** OAuth 2.0 / JWT tokens

**DELETE /api/v1/pricing/{id}**

- Delete pricing record- **Concurrent Processing:**- **Authorization:** Role-based access control (RBAC)



**Full API documentation:** http://localhost:8000/docs  - Go goroutines for parallel CSV processing- **Data Encryption:** TLS 1.3 in transit, AES-256 at rest



---  - Batch inserts (1000 records per batch)- **Audit Logging:** All actions logged and immutable



## 📊 Deployment  - Async processing for large files- **Compliance:** GDPR, SOC 2 compliance



### Docker Build  



```powershell- **CDN:****Design Approach:**

# Build image

docker build -t go-ingestion-service:latest .  - CloudFront/CloudFlare for static assets- API Gateway with authentication



# Run container  - Geographic distribution- Encrypted database connections

docker run -d -p 8080:8080 --name go-ingestion go-ingestion-service:latest

```- Secrets management (Vault/AWS Secrets Manager)



### Kubernetes Deployment### 2. Scalability- Input validation and sanitization



```powershell**Requirements:**- Regular security audits

# Deploy to Kubernetes

kubectl apply -f k8s/deployment.yaml- Horizontal Scaling: Stateless microservices

kubectl apply -f k8s/service.yaml

- Data Volume: Handle 100M+ pricing records### 5. Maintainability

# Scale deployment

kubectl scale deployment go-ingestion --replicas=3- Store Growth: Support 3000+ stores, expandable to 10,000+- **Code Quality:** > 80% test coverage



# Auto-scaling- Multi-Region: Deploy across multiple geographic regions- **Documentation:** API docs, architecture diagrams

kubectl autoscale deployment go-ingestion --cpu-percent=70 --min=2 --max=10

```- Storage: Accommodate 10TB+ of CSV files- **Monitoring:** Real-time metrics and alerts



---- **Logging:** Centralized structured logging



## 🧪 Testing**Design Approach:**



```powershell- **Microservices:****Design Approach:**

# Run all tests

go test ./... -v  - Independent scaling of upload and API services- Comprehensive unit and integration tests



# Run with coverage  - Load balancing with round-robin- OpenAPI/Swagger documentation

go test ./... -v -cover -coverprofile=coverage.out

  - Auto-scaling based on CPU/memory (50-80% threshold)- Prometheus metrics + Grafana dashboards

# View coverage report

go tool cover -html=coverage.out  - ELK stack for log aggregation

```

- **Database Scaling:**- CI/CD pipelines

---

  - Table partitioning by date (monthly partitions)

## 📄 License

  - Read replicas for query distribution (3 replicas)### 6. Usability

This project is licensed under the MIT License.

  - Database sharding by region/country- **Responsive Design:** Mobile, tablet, desktop support

---

  - **Accessibility:** WCAG 2.1 Level AA compliance

## 📧 Contact

- **Object Storage:**- **Internationalization:** Multi-language support

**Repository:** https://github.com/yourusername/golang-project-product-service

  - S3 with lifecycle policies- **User Experience:** Intuitive interface, < 3 clicks to any feature

**Project:** Retail Pricing Management System

  - Automatic archival to Glacier after 90 days

**Last Updated:** February 17, 2026

  - Multi-region replication**Design Approach:**

  - Angular Material responsive components

- **Message Queues:**- ARIA labels and keyboard navigation

  - RabbitMQ/SQS for async processing- i18n with Angular's built-in support

  - Dead letter queues for failed uploads- User testing and feedback loops



### 3. Availability & Reliability### 7. Data Consistency

**Requirements:**- **ACID Compliance:** Transactional integrity

- Uptime: 99.9% availability (< 8.76 hours downtime/year)- **Eventual Consistency:** For cross-region sync

- Disaster Recovery: RTO < 4 hours, RPO < 1 hour- **Data Validation:** Strong type checking

- Fault Tolerance: No single point of failure

- Data Integrity: Zero data loss**Design Approach:**

- PostgreSQL transactions

**Design Approach:**- Optimistic locking for concurrent edits

- **High Availability:**- Data validation at multiple layers

  - Multi-AZ deployment (3 availability zones)- Conflict resolution strategies

  - Database replication (primary + 2 replicas)

  - Service redundancy (min 3 instances per service)---

  

- **Health Checks:**## 🎨 Design Decisions

  - Liveness probes (every 10 seconds)

  - Readiness probes before routing traffic### 1. Microservices Architecture

  - Circuit breakers for external dependencies**Decision:** Split responsibilities into specialized services (Upload, Data)

  

- **Backup & Recovery:****Rationale:**

  - Daily automated database backups- **Separation of Concerns:** Upload service focuses on file handling; data service handles business logic

  - Point-in-time recovery (PITR) enabled- **Technology Optimization:** Go excels at concurrent I/O (file uploads); Python excels at data manipulation

  - S3 versioning for file recovery- **Independent Scaling:** Scale upload service during peak upload times, scale data service during high query load

  - Backup retention: 30 days- **Fault Isolation:** Failure in one service doesn't affect others

  - **Team Autonomy:** Different teams can own different services

- **Monitoring & Alerts:**

  - Real-time monitoring with Prometheus### 2. Go for Upload Service

  - Alert on error rate > 5%**Decision:** Use Go with Gin framework for file upload service

  - Alert on response time > 3s

  - 24/7 on-call rotation**Rationale:**

- **Performance:** Go's goroutines handle concurrent uploads efficiently

### 4. Security- **Low Memory Footprint:** Critical for processing large files

**Requirements:**- **Fast Compilation:** Quick iteration during development

- Authentication: OAuth 2.0 / JWT tokens- **Built-in Concurrency:** Native support for parallel CSV parsing

- Authorization: Role-based access control (RBAC)- **Static Binary:** Easy deployment without dependencies

- Data Encryption: TLS 1.3 in transit, AES-256 at rest

- Audit Logging: All actions logged and immutable### 3. Python for Data Service

- Compliance: GDPR, SOC 2 compliance**Decision:** Use Python with FastAPI/Flask for data operations



**Design Approach:****Rationale:**

- **Authentication & Authorization:**- **Data Processing:** Pandas library for complex data transformations

  - JWT tokens with 1-hour expiration- **ORM Maturity:** SQLAlchemy for robust database interactions

  - Refresh tokens with 7-day expiration- **Rich Ecosystem:** Libraries for validation, serialization, testing

  - Role-based permissions (Store Manager, Analyst, Admin)- **Rapid Development:** Quick implementation of business logic

  - API key authentication for service-to-service- **Team Expertise:** Wider talent pool familiar with Python

  

- **Encryption:**### 4. PostgreSQL as Primary Database

  - TLS 1.3 for all API communication**Decision:** Use PostgreSQL over MongoDB or MySQL

  - Database encryption at rest (AES-256)

  - S3 server-side encryption (SSE-S3)**Rationale:**

  - Secrets management with AWS Secrets Manager- **ACID Compliance:** Critical for financial pricing data

  - **Complex Queries:** Support for JOINs, aggregations, window functions

- **Input Validation:**- **JSON Support:** Flexibility for schema evolution

  - Server-side validation for all inputs- **Full-Text Search:** Built-in search capabilities

  - Parameterized queries (prevent SQL injection)- **Proven at Scale:** Battle-tested with large datasets

  - File type validation (CSV only)- **Open Source:** No licensing costs

  - File size limits (10MB max)

  - CSV content sanitization### 5. Redis for Caching

  **Decision:** Implement Redis for caching layer

- **Audit & Compliance:**

  - Immutable audit logs**Rationale:**

  - Log retention: 2 years- **Speed:** In-memory storage for sub-millisecond response times

  - GDPR right to erasure implementation- **Common Queries:** Cache frequent searches (e.g., by Store ID)

  - Data anonymization for non-production environments- **Session Storage:** User sessions and temporary data

- **Rate Limiting:** API throttling implementation

### 5. Maintainability- **Reduced Database Load:** Offload read operations

**Requirements:**

- Code Quality: > 80% test coverage### 6. S3 for File Storage

- Documentation: API docs, architecture diagrams**Decision:** Store uploaded CSV files in object storage

- Monitoring: Real-time metrics and alerts

- Logging: Centralized structured logging**Rationale:**

- **Durability:** 99.999999999% durability (11 nines)

**Design Approach:**- **Scalability:** Unlimited storage capacity

- **Testing:**- **Cost-Effective:** Pay-per-use, cheaper than database storage

  - Unit tests (80% coverage)- **Audit Trail:** Keep original files for compliance

  - Integration tests for critical paths- **Lifecycle Policies:** Automatic archival and deletion

  - End-to-end tests for user workflows

  - Load testing with k6/JMeter### 7. Single Page Application (Angular)

  **Decision:** Build SPA instead of server-rendered pages

- **Documentation:**

  - OpenAPI/Swagger for APIs**Rationale:**

  - Architecture Decision Records (ADR)- **User Experience:** Instant navigation, no page reloads

  - README for each service- **Responsive:** Better mobile and tablet experience

  - Inline code comments- **API-First:** Clean separation between frontend and backend

  - **Offline Capability:** Service workers for offline functionality

- **Observability:**- **Rich Interactions:** Real-time updates, drag-drop uploads

  - Structured JSON logging

  - Correlation IDs for request tracing### 8. API Gateway Pattern

  - Distributed tracing with Jaeger**Decision:** Use API Gateway as single entry point

  - Application Performance Monitoring (APM)

  **Rationale:**

- **CI/CD:**- **Centralized Auth:** Single authentication point

  - Automated testing on every commit- **Rate Limiting:** Protect backend services from abuse

  - Automated deployment to staging- **Request Routing:** Dynamic routing to services

  - Blue-green deployment for production- **CORS Handling:** Centralized CORS configuration

  - Rollback capabilities- **Monitoring:** Single point for metrics and logging



### 6. Usability### 9. Event-Driven for Large Files

**Requirements:****Decision:** Use message queues for processing large CSV files

- Responsive Design: Mobile, tablet, desktop support

- Accessibility: WCAG 2.1 Level AA compliance**Rationale:**

- Internationalization: Multi-language support- **Async Processing:** Don't block user while processing

- User Experience: Intuitive interface, < 3 clicks to any feature- **Retry Logic:** Automatic retry on failures

- **Load Leveling:** Prevent system overload

**Design Approach:**- **Auditability:** Track processing status

- **UI/UX:**- **Scalability:** Add workers as needed

  - Material Design principles

  - Consistent color scheme and typography### 10. Multi-Region Deployment

  - Loading indicators for async operations**Decision:** Deploy across multiple geographic regions

  - Error messages with actionable guidance

  - Keyboard shortcuts for power users**Rationale:**

  - **Latency:** Reduce latency for global users

- **Accessibility:**- **Compliance:** Data residency requirements (GDPR)

  - ARIA labels for screen readers- **Disaster Recovery:** Geographic redundancy

  - Keyboard navigation support- **Load Distribution:** Distribute traffic globally

  - High contrast mode

  - Font size adjustment---

  

- **Internationalization:**## 📝 Assumptions

  - Angular i18n support

  - Separate language files (en, es, fr, de)### Business Assumptions

  - Date/time localization1. **Store Count:** Starting with 3000 stores, growing to 10,000 in 3 years

  - Currency formatting2. **Upload Frequency:** Each store uploads pricing data daily

  3. **File Size:** Average CSV file contains 1,000-10,000 records (< 10MB)

- **Performance:**4. **Data Retention:** Pricing history retained for 2 years minimum

  - Lazy loading for large datasets5. **Peak Load:** 20% of stores upload simultaneously during business hours

  - Virtual scrolling for data grids6. **User Base:** 500 concurrent users (store managers, analysts, admins)

  - Debounced search inputs7. **Pricing Changes:** 10% of records updated monthly via edit feature

  - Progressive web app (PWA) capabilities

### Technical Assumptions

### 7. Data Consistency1. **Network:** Minimum 1 Mbps upload speed at store locations

**Requirements:**2. **Browser Support:** Modern browsers (Chrome, Firefox, Safari, Edge - last 2 versions)

- ACID Compliance: Transactional integrity3. **Database Size:** Initial database size ~50GB, growing to 500GB in 3 years

- Eventual Consistency: For cross-region sync4. **API Calls:** Average 1000 API requests per minute during peak hours

- Data Validation: Strong type checking5. **Search Queries:** 70% of searches use 1-2 criteria, 30% use 3+ criteria

6. **Edit Operations:** 80% single record edits, 20% bulk updates

**Design Approach:**7. **Infrastructure:** Cloud-hosted (AWS/Azure/GCP), not on-premises

- **Database Transactions:**

  - Use PostgreSQL transactions for updates### Data Assumptions

  - Isolation level: READ COMMITTED1. **CSV Format:** Consistent column order: Store ID, SKU, Product Name, Price, Date

  - Optimistic locking with version column2. **Store ID:** Alphanumeric, max 20 characters, unique per country

  3. **SKU:** Alphanumeric, max 50 characters, unique per product

- **Data Validation:**4. **Product Name:** UTF-8 string, max 200 characters

  - Schema validation at API level (Pydantic)5. **Price:** Decimal with 2 decimal places, positive values only

  - Database constraints (NOT NULL, CHECK, FK)6. **Date:** ISO 8601 format (YYYY-MM-DD)

  - Application-level validation7. **Character Encoding:** UTF-8 for international character support

  

- **Conflict Resolution:**### Security Assumptions

  - Last-write-wins for concurrent edits1. **Authentication:** Corporate SSO integration (SAML/OAuth)

  - User notification on conflicts2. **Authorization:** Three roles: Store Manager (upload only), Analyst (search/edit), Admin (all)

  - Merge strategies for bulk updates3. **Network:** Application accessible only via corporate VPN or whitelisted IPs

4. **Compliance:** Subject to PCI DSS for pricing data, GDPR for EU stores

---

### Operational Assumptions

## 🎨 Design Decisions1. **Monitoring:** 24/7 monitoring with on-call support

2. **Backup:** Daily automated backups with 30-day retention

### 1. Microservices Architecture with Language Optimization3. **Updates:** Monthly maintenance windows for updates (4 hours max)

**Decision:** Use Go for ingestion and Python for API operations4. **Support:** Email and ticket-based support (no phone support initially)



**Rationale:**---

- **Go for File Processing:**

  - Superior performance for I/O operations## 📁 Project Structure

  - Efficient memory management for large files

  - Native goroutines for concurrent processing```

  - Fast CSV parsing with streamingretail-pricing-system/

  - Single binary deployment│

  ├── frontend/                          # Angular Frontend

- **Python for Data Operations:**│   ├── src/

  - Rich ecosystem for data manipulation (Pandas)│   │   ├── app/

  - FastAPI for rapid API development│   │   │   ├── components/

  - SQLAlchemy for complex queries│   │   │   │   ├── upload/           # File upload component

  - Easy integration with data science tools│   │   │   │   ├── search/           # Search interface

  - Better for business logic│   │   │   │   └── data-grid/        # Editable data grid

│   │   │   ├── services/

**Trade-offs:**│   │   │   │   ├── upload.service.ts

- Need to maintain two codebases│   │   │   │   ├── pricing.service.ts

- Different deployment pipelines│   │   │   │   └── auth.service.ts

- Team needs expertise in both languages│   │   │   ├── models/

│   │   │   └── guards/

### 2. Event-Driven Upload Architecture│   │   ├── assets/

**Decision:** Use presigned URLs and storage events instead of direct upload to backend│   │   └── environments/

│   ├── angular.json

**Rationale:**│   ├── package.json

- **Scalability:** Direct S3 upload bypasses backend bottleneck│   └── README.md

- **Performance:** No file data through backend servers│

- **Bandwidth:** Reduced backend bandwidth costs├── upload-service/                    # Go Upload Microservice (THIS REPO)

- **Security:** Presigned URLs with expiration (5 minutes)│   ├── cmd/

- **Decoupling:** Upload and processing are independent│   │   └── server/

│   │       └── main.go               # Entry point

**Flow:**│   ├── internal/

1. Frontend requests presigned URL from Python API│   │   ├── handlers/                 # HTTP handlers

2. Frontend uploads directly to S3│   │   ├── middleware/               # Custom middleware

3. S3 event triggers Go ingestion service│   │   ├── models/                   # Data models

4. Go service processes asynchronously│   │   ├── services/                 # Business logic

│   │   └── config/                   # Configuration

**Trade-offs:**│   ├── pkg/                          # Shared packages

- More complex upload flow│   ├── uploads/                      # Temporary upload storage

- Requires S3 event configuration│   ├── go.mod

- Need to handle failed uploads│   ├── go.sum

│   ├── Dockerfile

### 3. Streaming CSV Processing│   └── README.md

**Decision:** Use streaming parser instead of loading entire file into memory│

├── data-service/                      # Python Data Microservice

**Rationale:**│   ├── app/

- **Memory Efficiency:** Process 10MB+ files with constant memory│   │   ├── main.py                   # FastAPI entry point

- **Performance:** Start processing immediately│   │   ├── routers/

- **Scalability:** Handle unlimited file sizes│   │   │   ├── pricing.py            # CRUD endpoints

- **Error Recovery:** Stop on first error, don't process entire file│   │   │   ├── search.py             # Search endpoints

│   │   │   └── health.py

**Implementation:**│   │   ├── models/

```go│   │   │   ├── pricing.py            # SQLAlchemy models

// Stream CSV line by line│   │   │   └── schemas.py            # Pydantic schemas

reader := csv.NewReader(file)│   │   ├── services/

for {│   │   ├── database/

    record, err := reader.Read()│   │   ├── middleware/

    if err == io.EOF {│   │   └── config/

        break│   ├── tests/

    }│   ├── requirements.txt

    // Process record immediately│   ├── Dockerfile

    processBatch(record)│   └── README.md

}│

```├── infrastructure/                    # Infrastructure as Code

│   ├── kubernetes/

### 4. PostgreSQL Over NoSQL│   ├── terraform/

**Decision:** Use PostgreSQL as primary database│   └── docker-compose.yml

│

**Rationale:**├── database/

- **ACID Compliance:** Critical for financial pricing data│   ├── migrations/

- **Complex Queries:** Support for JOINs, aggregations, window functions│   ├── seeds/

- **Data Integrity:** Foreign keys, constraints, triggers│   └── schema.sql

- **Full-Text Search:** Built-in search capabilities│

- **Proven at Scale:** Battle-tested with large datasets└── docs/

- **Cost:** Open source, no licensing fees    ├── architecture/

    ├── api/

**Comparison:**    └── deployment/

| Feature | PostgreSQL | MongoDB |```

|---------|-----------|---------|

| ACID | ✅ Strong | ⚠️ Weak |---

| Joins | ✅ Excellent | ❌ Limited |

| Schema | ✅ Enforced | ⚠️ Flexible |## 🚀 Getting Started

| Scale | ✅ Vertical + Horizontal | ✅ Horizontal |

| Query Language | SQL | MongoDB Query |### Upload Service (Go + Gin) - This Repository



### 5. Redis for Caching (Optional)#### Prerequisites

**Decision:** Implement caching layer with Redis- **Go** 1.25+

- **PostgreSQL** 15+ (optional for local dev)

**Rationale:**- **Redis** 7+ (optional for caching)

- **Performance:** 100x faster than database for common queries

- **Reduced Load:** Offload read operations from PostgreSQL#### Installation

- **Session Storage:** Store user sessions in memory

- **Rate Limiting:** Implement API throttling1. **Clone the repository**

```powershell

**Caching Strategy:**git clone https://github.com/yourusername/golang-project-product-service.git

- Cache search results (TTL: 5 minutes)cd golang-project-product-service

- Cache user sessions (TTL: 1 hour)```

- Invalidate cache on data updates

- Use cache-aside pattern2. **Install dependencies**

```powershell

### 6. Batch Insert for Database Operationsgo mod download

**Decision:** Batch insert records instead of individual inserts```



**Rationale:**3. **Configure environment variables** (optional)

- **Performance:** 10-100x faster than single inserts```powershell

- **Network:** Reduced round trips to database# Create .env file

- **Transactions:** Single transaction for batch@"

- **Error Handling:** Rollback entire batch on errorPORT=8080

MAX_UPLOAD_SIZE=10485760

**Implementation:**UPLOAD_PATH=./uploads

```goLOG_LEVEL=info

// Batch size: 1000 records"@ | Out-File -FilePath .env -Encoding utf8

const batchSize = 1000```

for i := 0; i < len(records); i += batchSize {

    batch := records[i:min(i+batchSize, len(records))]4. **Run the service**

    db.BatchInsert(batch)```powershell

}go run cmd/server/main.go

``````



### 7. Audit Logging with Database TriggersThe service will start on http://localhost:8080

**Decision:** Use database triggers for audit logging

#### Testing the Upload Endpoint

**Rationale:**

- **Consistency:** Can't bypass audit logging**Create a test CSV file:**

- **Performance:** Happens at database level```powershell

- **Completeness:** Captures all changes@"

- **Immutable:** Logs stored separatelyStore ID,SKU,Product Name,Price,Date

ST001,SKU12345,Laptop Computer,999.99,2026-02-17

**Trigger Example:**ST001,SKU12346,Wireless Mouse,29.99,2026-02-17

```sqlST002,SKU12347,USB Keyboard,49.99,2026-02-17

CREATE TRIGGER audit_pricing_update"@ | Out-File -FilePath pricing_test.csv -Encoding utf8

AFTER UPDATE ON pricing_records```

FOR EACH ROW

INSERT INTO audit_logs (table_name, record_id, action, old_value, new_value, user_id, timestamp)**Upload using PowerShell:**

VALUES ('pricing_records', NEW.id, 'UPDATE', row_to_json(OLD), row_to_json(NEW), current_user, NOW());```powershell

```$headers = @{

    "Accept" = "application/json"

### 8. Optimistic Locking for Concurrent Edits}

**Decision:** Use version column for optimistic locking$response = Invoke-WebRequest -Uri "http://localhost:8080/upload" `

    -Method Post `

**Rationale:**    -Headers $headers `

- **Performance:** No locks, better concurrency    -Form @{file = Get-Item -Path "pricing_test.csv"}

- **User Experience:** Users notified only on actual conflicts

- **Scalability:** Works in distributed systems$response.Content | ConvertFrom-Json | ConvertTo-Json

```

**Implementation:**

```python**Expected Response:**

# Check version before update```json

UPDATE pricing_records{

SET price = $1, version = version + 1  "message": "File uploaded successfully",

WHERE id = $2 AND version = $3  "filename": "1739151234_pricing_test.csv",

# If affected rows = 0, conflict detected  "size": 256,

```  "rows": 4,

  "columns": 5

### 9. Direct S3 Upload from Frontend}

**Decision:** Upload files directly from browser to S3```



**Rationale:**#### Health Check

- **Performance:** No backend involvement in file transfer```powershell

- **Scalability:** Backend not bottleneck for uploadsInvoke-RestMethod -Uri "http://localhost:8080/health" -Method Get

- **Cost:** Reduced bandwidth costs```

- **User Experience:** Progress bar works directly

---

**Security:**

- Presigned URLs with 5-minute expiration## 📚 API Documentation

- Limit file size at S3 level

- Restrict to specific content type### Upload Service Endpoints



### 10. Separation of Concerns (SoC)#### POST /upload

**Decision:** Separate file upload, processing, and data operationsUpload a CSV file with pricing data.



**Rationale:****Request:**

- **Maintainability:** Each service has single responsibility- **Method:** POST

- **Scalability:** Scale services independently- **Content-Type:** multipart/form-data

- **Fault Isolation:** Failure in one doesn't affect others- **Parameters:**

- **Team Organization:** Different teams own different services  - `file` (required): CSV file containing pricing data



**Service Boundaries:****CSV Format:**

- **Frontend:** User interface, validation, display```csv

- **Python API:** Data operations, search, authenticationStore ID,SKU,Product Name,Price,Date

- **Go Ingestion:** File processing, validation, database insertST001,SKU12345,Laptop Computer,999.99,2026-02-17

- **PostgreSQL:** Data persistence, queries```

- **S3:** File storage

**Success Response (200 OK):**

---```json

{

## 📝 Assumptions  "message": "File uploaded successfully",

  "filename": "1739151234_pricing.csv",

### Business Assumptions  "size": 524288,

1. **Store Count:** Starting with 3000 stores, growing to 10,000 in 3 years  "rows": 5000,

2. **Upload Frequency:** Each store uploads pricing data daily (once per day)  "columns": 5

3. **File Size:** Average CSV file contains 1,000-10,000 records (< 10MB)}

4. **Data Retention:** Pricing history retained for 2 years minimum```

5. **Peak Load:** 20% of stores upload simultaneously during business hours (8-10 AM)

6. **User Base:** 500 concurrent users (store managers, analysts, admins)**Error Responses:**

7. **Pricing Changes:** 10% of records updated monthly via edit feature- **400 Bad Request:**

8. **Growth Rate:** 15% annual growth in stores and data volume  ```json

  {

### Technical Assumptions    "error": "File too large. Maximum size is 10MB"

1. **Network:** Minimum 1 Mbps upload speed at store locations  }

2. **Browser Support:** Modern browsers (Chrome, Firefox, Safari, Edge - last 2 versions)  ```

3. **Database Size:** Initial database size ~50GB, growing to 500GB in 3 years  ```json

4. **API Calls:** Average 1000 API requests per minute during peak hours  {

5. **Search Queries:** 70% of searches use 1-2 criteria, 30% use 3+ criteria    "error": "Only CSV files are allowed"

6. **Edit Operations:** 80% single record edits, 20% bulk updates  }

7. **Infrastructure:** Cloud-hosted (AWS/Azure/GCP), not on-premises  ```

8. **Internet Connectivity:** 99.5% uptime at store locations  ```json

  {

### Data Assumptions    "error": "Invalid CSV file format"

1. **CSV Format:** Consistent column order: `Store ID, SKU, Product Name, Price, Date`  }

2. **Store ID:** Alphanumeric, max 20 characters, unique per country  ```

3. **SKU:** Alphanumeric, max 50 characters, unique per product

4. **Product Name:** UTF-8 string, max 200 characters, may contain special characters- **500 Internal Server Error:**

5. **Price:** Decimal with 2 decimal places, positive values only, range: 0.01 - 999,999.99  ```json

6. **Date:** ISO 8601 format (YYYY-MM-DD), range: 2020-01-01 to current date  {

7. **Character Encoding:** UTF-8 for international character support    "error": "Error saving file"

8. **Delimiter:** Comma (,) as CSV delimiter  }

9. **Header Row:** First row contains column names  ```

10. **No Empty Fields:** All fields are required

#### GET /health

### Security AssumptionsHealth check endpoint.

1. **Authentication:** Corporate SSO integration (SAML/OAuth 2.0)

2. **Authorization:** Three roles:**Response (200 OK):**

   - Store Manager: Upload files, view own store data```json

   - Analyst: Search, view, edit all data{

   - Admin: All permissions + user management  "status": "healthy",

3. **Network:** Application accessible only via:  "timestamp": "2026-02-17T10:30:00Z"

   - Corporate VPN}

   - Whitelisted IP ranges```

   - Authenticated users only

4. **Compliance:** Subject to:---

   - PCI DSS for pricing data

   - GDPR for EU stores## 🚢 Deployment

   - SOC 2 Type II

### Docker Deployment

### Operational Assumptions

1. **Monitoring:** 24/7 monitoring with on-call support**Build Docker image:**

2. **Backup:** Daily automated backups at 2 AM (off-peak hours)```powershell

3. **Backup Retention:** 30 days online, 2 years archiveddocker build -t upload-service:latest .

4. **Updates:** Monthly maintenance windows for updates (4 hours max, Saturday 2-6 AM)```

5. **Support:** 

   - Email support (24-hour response time)**Run container:**

   - Ticket-based support system```powershell

   - No phone support initiallydocker run -d `

6. **SLA:** 99.9% uptime (8.76 hours downtime per year allowed)  -p 8080:8080 `

7. **Disaster Recovery Testing:** Quarterly DR drills  -e PORT=8080 `

  -e MAX_UPLOAD_SIZE=10485760 `

### Integration Assumptions  --name upload-service `

1. **SSO:** Integration with existing identity provider (Okta/Azure AD)  upload-service:latest

2. **Monitoring:** Integration with existing monitoring tools (Datadog/New Relic)```

3. **Logging:** Logs shipped to central logging system (Splunk/ELK)

4. **Notifications:** Email notifications via SMTP server### Kubernetes Deployment

5. **No POS Integration:** Initially standalone system, future integration possible

**Create deployment:**

---```yaml

apiVersion: apps/v1

## 📁 Project Structurekind: Deployment

metadata:

### Multi-Repository Structure (Recommended)  name: upload-service

spec:

```  replicas: 3

retail-pricing-system/  selector:

│    matchLabels:

├── frontend/                          # Angular Frontend Repository      app: upload-service

│   ├── src/  template:

│   │   ├── app/    metadata:

│   │   │   ├── components/      labels:

│   │   │   │   ├── upload/           # File upload component        app: upload-service

│   │   │   │   │   ├── upload.component.ts    spec:

│   │   │   │   │   ├── upload.component.html      containers:

│   │   │   │   │   └── upload.component.scss      - name: upload-service

│   │   │   │   ├── search/           # Search interface        image: upload-service:latest

│   │   │   │   │   ├── search.component.ts        ports:

│   │   │   │   │   ├── search.component.html        - containerPort: 8080

│   │   │   │   │   └── search.component.scss        env:

│   │   │   │   └── data-grid/        # Editable data grid        - name: PORT

│   │   │   │       ├── data-grid.component.ts          value: "8080"

│   │   │   │       ├── data-grid.component.html        - name: MAX_UPLOAD_SIZE

│   │   │   │       └── data-grid.component.scss          value: "10485760"

│   │   │   ├── services/```

│   │   │   │   ├── upload.service.ts

│   │   │   │   ├── pricing.service.ts**Deploy:**

│   │   │   │   └── auth.service.ts```powershell

│   │   │   ├── models/kubectl apply -f k8s/deployment.yaml

│   │   │   │   ├── pricing-record.model.tskubectl apply -f k8s/service.yaml

│   │   │   │   ├── search-criteria.model.ts```

│   │   │   │   └── user.model.ts

│   │   │   ├── guards/---

│   │   │   │   ├── auth.guard.ts

│   │   │   │   └── role.guard.ts## 🧪 Testing

│   │   │   └── interceptors/

│   │   │       ├── auth.interceptor.ts### Unit Tests

│   │   │       └── error.interceptor.ts```powershell

│   │   ├── assets/go test ./... -v -cover

│   │   └── environments/```

│   ├── angular.json

│   ├── package.json### Integration Tests

│   ├── tsconfig.json```powershell

│   └── README.mdgo test ./... -tags=integration -v

│```

├── golang-project-product-service/    # Go Ingestion Service (THIS REPO)

│   ├── cmd/### Load Testing

│   │   └── server/```powershell

│   │       └── main.go               # Entry point# Using Apache Bench

│   ├── internal/ab -n 1000 -c 10 -T "multipart/form-data" -p pricing_test.csv http://localhost:8080/upload

│   │   ├── handlers/                 # HTTP handlers```

│   │   │   ├── upload.go

│   │   │   ├── health.go---

│   │   │   └── event.go              # S3 event handler

│   │   ├── middleware/               # Custom middleware## 📊 Monitoring

│   │   │   ├── auth.go

│   │   │   ├── cors.go### Metrics Exposed

│   │   │   └── logging.go- Total uploads count

│   │   ├── models/                   # Data models- Upload success/failure rate

│   │   │   ├── pricing.go- Upload duration histogram

│   │   │   └── upload_history.go- File size histogram

│   │   ├── services/                 # Business logic- Active connections

│   │   │   ├── validator.go

│   │   │   ├── parser.go### Health Checks

│   │   │   ├── storage.go- `/health` endpoint for liveness probe

│   │   │   └── database.go- Database connectivity check (if implemented)

│   │   └── config/                   # Configuration

│   │       └── config.go---

│   ├── pkg/                          # Shared packages

│   │   └── utils/## 🔒 Security Features

│   │       ├── csv.go

│   │       └── s3.go- File size validation (max 10MB)

│   ├── uploads/                      # Temporary upload storage- File type validation (CSV only)

│   ├── tests/- CSV structure validation

│   │   ├── integration/- Input sanitization

│   │   └── unit/- Error handling without sensitive info disclosure

│   ├── go.mod

│   ├── go.sum---

│   ├── Dockerfile

│   ├── .env.example## 🤝 Contributing

│   └── README.md

│1. Fork the repository

├── data-service/                      # Python Data Microservice2. Create a feature branch (`git checkout -b feature/amazing-feature`)

│   ├── app/3. Commit your changes (`git commit -m 'Add amazing feature'`)

│   │   ├── main.py                   # FastAPI entry point4. Push to the branch (`git push origin feature/amazing-feature`)

│   │   ├── routers/5. Open a Pull Request

│   │   │   ├── __init__.py

│   │   │   ├── pricing.py            # CRUD endpoints---

│   │   │   ├── search.py             # Search endpoints

│   │   │   ├── upload.py             # Presigned URL generation## 📄 License

│   │   │   ├── auth.py               # Authentication

│   │   │   └── health.py             # Health checkThis project is licensed under the MIT License - see the LICENSE file for details.

│   │   ├── models/

│   │   │   ├── __init__.py---

│   │   │   ├── pricing.py            # SQLAlchemy models

│   │   │   └── upload_history.py## 📧 Contact & Support

│   │   ├── schemas/

│   │   │   ├── __init__.py**Project Link:** https://github.com/yourusername/golang-project-product-service

│   │   │   ├── pricing.py            # Pydantic schemas

│   │   │   ├── search.py**Related Services:**

│   │   │   └── user.py- Frontend: [Angular Frontend Repository]

│   │   ├── services/- Data Service: [Python Data Service Repository]

│   │   │   ├── __init__.py

│   │   │   ├── pricing_service.py---

│   │   │   ├── search_service.py

│   │   │   ├── upload_service.py**Last Updated:** February 17, 2026

│   │   │   └── cache_service.py
│   │   ├── database/
│   │   │   ├── __init__.py
│   │   │   ├── connection.py
│   │   │   └── session.py
│   │   ├── middleware/
│   │   │   ├── __init__.py
│   │   │   ├── auth.py
│   │   │   └── error_handler.py
│   │   └── config/
│   │       ├── __init__.py
│   │       └── settings.py
│   ├── migrations/                   # Alembic migrations
│   │   └── versions/
│   ├── tests/
│   │   ├── __init__.py
│   │   ├── test_pricing.py
│   │   └── test_search.py
│   ├── requirements.txt
│   ├── requirements-dev.txt
│   ├── Dockerfile
│   ├── .env.example
│   └── README.md
│
├── infrastructure/                    # Infrastructure as Code
│   ├── kubernetes/
│   │   ├── deployments/
│   │   │   ├── frontend.yaml
│   │   │   ├── go-ingestion.yaml
│   │   │   └── python-api.yaml
│   │   ├── services/
│   │   │   ├── frontend-service.yaml
│   │   │   ├── go-service.yaml
│   │   │   └── python-service.yaml
│   │   ├── ingress/
│   │   │   └── ingress.yaml
│   │   ├── configmaps/
│   │   │   └── app-config.yaml
│   │   └── secrets/
│   │       └── app-secrets.yaml
│   ├── terraform/
│   │   ├── modules/
│   │   │   ├── eks/
│   │   │   ├── rds/
│   │   │   ├── s3/
│   │   │   └── vpc/
│   │   ├── environments/
│   │   │   ├── dev/
│   │   │   ├── staging/
│   │   │   └── prod/
│   │   └── main.tf
│   └── docker-compose.yml            # Local development
│
├── database/
│   ├── migrations/                   # Database migrations
│   │   ├── 001_initial_schema.sql
│   │   ├── 002_add_audit_logs.sql
│   │   └── 003_add_indexes.sql
│   ├── seeds/                        # Test data
│   │   └── sample_data.sql
│   └── schema.sql                    # Complete schema
│
├── docs/
│   ├── architecture/
│   │   ├── context-diagram.md
│   │   ├── solution-architecture.md
│   │   ├── component-diagram.md
│   │   └── adr/                      # Architecture Decision Records
│   │       ├── 001-microservices.md
│   │       ├── 002-go-for-ingestion.md
│   │       └── 003-event-driven-upload.md
│   ├── api/
│   │   ├── upload-service-api.md
│   │   └── data-service-api.md
│   ├── deployment/
│   │   └── deployment-guide.md
│   └── user-guides/
│       ├── upload-guide.md
│       ├── search-guide.md
│       └── edit-guide.md
│
├── scripts/
│   ├── setup-dev.sh
│   ├── run-tests.sh
│   ├── deploy.sh
│   └── seed-database.sh
│
├── .github/
│   └── workflows/
│       ├── ci-go-service.yml
│       ├── ci-python-service.yml
│       ├── ci-frontend.yml
│       └── deploy-prod.yml
│
├── docker-compose.yml                # Full stack local development
├── .gitignore
├── LICENSE
└── README.md                         # Main project README
```

### Current Repository (Go Ingestion Service)

```
golang-project-product-service/
│
├── cmd/
│   └── server/
│       └── main.go                   # Application entry point
│
├── internal/
│   ├── handlers/
│   │   ├── upload.go                 # File upload handler
│   │   ├── health.go                 # Health check handler
│   │   └── event.go                  # S3 event handler (future)
│   ├── middleware/
│   │   ├── auth.go                   # Authentication middleware
│   │   ├── cors.go                   # CORS middleware
│   │   └── logging.go                # Request logging
│   ├── models/
│   │   ├── pricing.go                # Pricing record model
│   │   └── response.go               # API response models
│   ├── services/
│   │   ├── validator.go              # CSV validation
│   │   ├── parser.go                 # CSV parsing (streaming)
│   │   ├── storage.go                # File storage (S3/local)
│   │   └── database.go               # Database operations
│   └── config/
│       └── config.go                 # Configuration management
│
├── pkg/
│   └── utils/
│       ├── csv.go                    # CSV utilities
│       └── s3.go                     # S3 utilities
│
├── uploads/                          # Temporary file storage
│
├── tests/
│   ├── integration/
│   │   └── upload_test.go
│   └── unit/
│       ├── validator_test.go
│       └── parser_test.go
│
├── go.mod                            # Go module definition
├── go.sum                            # Dependency checksums
├── Dockerfile                        # Container image
├── .env.example                      # Environment variables template
├── .gitignore
└── README.md                         # This file
```

---

## 🚀 Getting Started

### Prerequisites

- **Go** 1.25+
- **Node.js** 18+ and npm (for frontend)
- **Python** 3.11+ (for data service)
- **PostgreSQL** 15+
- **Redis** 7+ (optional, for caching)
- **Docker** and Docker Compose (recommended for local development)
- **AWS CLI** (if using S3)

### Local Development Setup

#### Option 1: Docker Compose (Recommended)

This starts all services with a single command:

```powershell
# Clone the main repository
git clone https://github.com/yourusername/retail-pricing-system.git
cd retail-pricing-system

# Start all services
docker-compose up -d

# View logs
docker-compose logs -f

# Stop all services
docker-compose down
```

**Services will be available at:**
- Frontend: http://localhost:4200
- Python API: http://localhost:8000
- Go Ingestion: http://localhost:8080
- PostgreSQL: localhost:5432
- Redis: localhost:6379

#### Option 2: Manual Setup (Individual Services)

##### 1. Go Ingestion Service (This Repository)

```powershell
# Clone repository
git clone https://github.com/yourusername/golang-project-product-service.git
cd golang-project-product-service

# Install dependencies
go mod download

# Copy environment variables
cp .env.example .env

# Edit .env with your configuration
notepad .env

# Run the service
go run cmd/server/main.go
```

**Service runs on:** http://localhost:8080

##### 2. Database Setup (PostgreSQL)

```powershell
# Using Docker
docker run -d `
  --name pricing-postgres `
  -e POSTGRES_DB=pricing_db `
  -e POSTGRES_USER=pricing_user `
  -e POSTGRES_PASSWORD=your_password `
  -p 5432:5432 `
  postgres:15

# Run migrations
psql -h localhost -U pricing_user -d pricing_db -f database/schema.sql
```

##### 3. Python Data Service

```powershell
# Clone repository
git clone https://github.com/yourusername/data-service.git
cd data-service

# Create virtual environment
python -m venv venv
.\venv\Scripts\Activate.ps1

# Install dependencies
pip install -r requirements.txt

# Copy environment variables
cp .env.example .env

# Run migrations
alembic upgrade head

# Start service
uvicorn app.main:app --reload --port 8000
```

**Service runs on:** http://localhost:8000

##### 4. Angular Frontend

```powershell
# Clone repository
git clone https://github.com/yourusername/frontend.git
cd frontend

# Install dependencies
npm install

# Update environment configuration
notepad src/environments/environment.ts

# Start development server
ng serve
```

**Application runs on:** http://localhost:4200

### Testing the Go Ingestion Service

#### 1. Health Check

```powershell
Invoke-RestMethod -Uri "http://localhost:8080/health" -Method Get
```

**Expected Response:**
```json
{
  "status": "healthy",
  "timestamp": "2026-02-17T10:30:00Z"
}
```

#### 2. Create Test CSV File

```powershell
@"
Store ID,SKU,Product Name,Price,Date
ST001,SKU12345,Laptop Computer,999.99,2026-02-17
ST001,SKU12346,Wireless Mouse,29.99,2026-02-17
ST002,SKU12347,USB Keyboard,49.99,2026-02-17
ST002,SKU12348,Monitor 24 inch,299.99,2026-02-17
ST003,SKU12349,USB Cable,9.99,2026-02-17
"@ | Out-File -FilePath pricing_test.csv -Encoding utf8
```

#### 3. Upload CSV File

```powershell
$response = Invoke-WebRequest -Uri "http://localhost:8080/upload" `
    -Method Post `
    -Form @{file = Get-Item -Path "pricing_test.csv"}

$response.Content | ConvertFrom-Json | ConvertTo-Json -Depth 10
```

**Expected Response:**
```json
{
  "message": "File uploaded successfully",
  "filename": "1708171234_pricing_test.csv",
  "size": 256,
  "rows": 6,
  "columns": 5
}
```

#### 4. Verify Upload

```powershell
# Check uploaded file
Get-ChildItem .\uploads\
```

### Environment Variables

#### Go Ingestion Service (.env)

```env
# Server Configuration
PORT=8080
GIN_MODE=release

# File Upload
MAX_UPLOAD_SIZE=10485760
UPLOAD_PATH=./uploads

# Database
DATABASE_URL=postgresql://pricing_user:your_password@localhost:5432/pricing_db
DB_MAX_CONNECTIONS=50
DB_MAX_IDLE_CONNECTIONS=10

# S3 Configuration (Optional)
USE_S3=false
AWS_REGION=us-east-1
S3_BUCKET=retail-pricing-uploads
AWS_ACCESS_KEY_ID=your_access_key
AWS_SECRET_ACCESS_KEY=your_secret_key

# Redis (Optional)
REDIS_URL=redis://localhost:6379/0

# Security
JWT_SECRET=your-secret-key-change-in-production

# Logging
LOG_LEVEL=info
LOG_FORMAT=json

# Processing
BATCH_SIZE=1000
MAX_GOROUTINES=10
```

#### Python Data Service (.env)

```env
# Server Configuration
PORT=8000
WORKERS=4

# Database
DATABASE_URL=postgresql://pricing_user:your_password@localhost:5432/pricing_db

# Redis
REDIS_URL=redis://localhost:6379/0
CACHE_TTL=300

# Security
SECRET_KEY=your-secret-key-change-in-production
JWT_ALGORITHM=HS256
ACCESS_TOKEN_EXPIRE_MINUTES=60

# CORS
CORS_ORIGINS=http://localhost:4200,https://yourdomain.com

# S3
AWS_REGION=us-east-1
S3_BUCKET=retail-pricing-uploads
PRESIGNED_URL_EXPIRATION=300

# Logging
LOG_LEVEL=info
```

---

## 📚 API Documentation

### Go Ingestion Service (Port 8080)

#### POST /upload
Upload a CSV file with pricing data.

**Request:**
```http
POST /upload HTTP/1.1
Host: localhost:8080
Content-Type: multipart/form-data

file: <binary-csv-file>
```

**CSV Format Requirements:**
- **Header Row:** `Store ID,SKU,Product Name,Price,Date`
- **Delimiter:** Comma (,)
- **Encoding:** UTF-8
- **Max Size:** 10MB
- **File Extension:** .csv

**Example CSV:**
```csv
Store ID,SKU,Product Name,Price,Date
ST001,SKU12345,Laptop Computer,999.99,2026-02-17
ST002,SKU12346,Wireless Mouse,29.99,2026-02-17
```

**Success Response (200 OK):**
```json
{
  "message": "File uploaded successfully",
  "filename": "1708171234_pricing.csv",
  "size": 524288,
  "rows": 5000,
  "columns": 5
}
```

**Error Responses:**

**400 Bad Request - File too large:**
```json
{
  "error": "File too large. Maximum size is 10MB"
}
```

**400 Bad Request - Invalid file type:**
```json
{
  "error": "Only CSV files are allowed"
}
```

**400 Bad Request - Invalid CSV format:**
```json
{
  "error": "Invalid CSV file format"
}
```

**400 Bad Request - Empty file:**
```json
{
  "error": "CSV file is empty"
}
```

**500 Internal Server Error:**
```json
{
  "error": "Error saving file"
}
```

#### GET /health
Health check endpoint for monitoring.

**Request:**
```http
GET /health HTTP/1.1
Host: localhost:8080
```

**Response (200 OK):**
```json
{
  "status": "healthy",
  "timestamp": "2026-02-17T10:30:00Z"
}
```

### Python Data Service (Port 8000)

Full API documentation available at: http://localhost:8000/docs (Swagger UI)

#### POST /api/v1/upload/presigned-url
Generate presigned URL for direct S3 upload.

**Request:**
```json
{
  "filename": "pricing_data.csv",
  "content_type": "text/csv"
}
```

**Response (200 OK):**
```json
{
  "upload_url": "https://s3.amazonaws.com/bucket/...",
  "file_key": "uploads/2026/02/17/1708171234_pricing_data.csv",
  "expires_in": 300
}
```

#### GET /api/v1/pricing
Search and retrieve pricing records.

**Query Parameters:**
- `store_id` (optional): Filter by store ID (exact match)
- `sku` (optional): Filter by SKU (exact match)
- `product_name` (optional): Filter by product name (partial match, case-insensitive)
- `min_price` (optional): Minimum price (inclusive)
- `max_price` (optional): Maximum price (inclusive)
- `start_date` (optional): Start date (YYYY-MM-DD, inclusive)
- `end_date` (optional): End date (YYYY-MM-DD, inclusive)
- `page` (optional): Page number (default: 1)
- `limit` (optional): Records per page (default: 50, max: 1000)
- `sort_by` (optional): Sort field (default: date)
- `sort_order` (optional): asc or desc (default: desc)

**Example Request:**
```http
GET /api/v1/pricing?store_id=ST001&start_date=2026-01-01&page=1&limit=50
```

**Response (200 OK):**
```json
{
  "data": [
    {
      "id": 1,
      "store_id": "ST001",
      "sku": "SKU12345",
      "product_name": "Laptop Computer",
      "price": 999.99,
      "date": "2026-02-17",
      "created_at": "2026-02-17T10:00:00Z",
      "updated_at": "2026-02-17T10:00:00Z",
      "version": 1
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 50,
    "total": 5000,
    "pages": 100
  }
}
```

#### GET /api/v1/pricing/{id}
Get a single pricing record by ID.

**Response (200 OK):**
```json
{
  "id": 1,
  "store_id": "ST001",
  "sku": "SKU12345",
  "product_name": "Laptop Computer",
  "price": 999.99,
  "date": "2026-02-17",
  "created_at": "2026-02-17T10:00:00Z",
  "updated_at": "2026-02-17T10:00:00Z",
  "version": 1
}
```

#### PUT /api/v1/pricing/{id}
Update a pricing record.

**Request:**
```json
{
  "price": 1099.99,
  "product_name": "Laptop Computer - Updated",
  "version": 1
}
```

**Response (200 OK):**
```json
{
  "id": 1,
  "store_id": "ST001",
  "sku": "SKU12345",
  "product_name": "Laptop Computer - Updated",
  "price": 1099.99,
  "date": "2026-02-17",
  "created_at": "2026-02-17T10:00:00Z",
  "updated_at": "2026-02-17T11:00:00Z",
  "version": 2
}
```

**Error (409 Conflict - Concurrent Edit):**
```json
{
  "error": "Concurrent modification detected. Please refresh and try again.",
  "code": "CONCURRENT_MODIFICATION"
}
```

#### DELETE /api/v1/pricing/{id}
Delete a pricing record.

**Response (200 OK):**
```json
{
  "message": "Record deleted successfully"
}
```

#### GET /api/v1/health
Health check endpoint.

**Response (200 OK):**
```json
{
  "status": "healthy",
  "timestamp": "2026-02-17T10:30:00Z",
  "database": "connected",
  "redis": "connected"
}
```

---

## 🚢 Deployment

### Docker Deployment

#### Build Docker Images

```powershell
# Go Ingestion Service
docker build -t go-ingestion-service:latest -f Dockerfile .

# Python Data Service
cd ../data-service
docker build -t python-data-service:latest -f Dockerfile .

# Angular Frontend
cd ../frontend
docker build -t angular-frontend:latest -f Dockerfile .
```

#### Run with Docker Compose

```powershell
# Start all services
docker-compose up -d

# View logs
docker-compose logs -f go-ingestion
docker-compose logs -f python-api
docker-compose logs -f frontend

# Stop services
docker-compose down

# Stop and remove volumes
docker-compose down -v
```

### Kubernetes Deployment

#### Prerequisites
- Kubernetes cluster (EKS, AKS, GKE, or local minikube)
- kubectl configured
- Docker images pushed to container registry

#### Deploy Services

```powershell
# Create namespace
kubectl create namespace retail-pricing

# Deploy PostgreSQL
kubectl apply -f infrastructure/kubernetes/deployments/postgres.yaml

# Deploy Redis
kubectl apply -f infrastructure/kubernetes/deployments/redis.yaml

# Deploy Go Ingestion Service
kubectl apply -f infrastructure/kubernetes/deployments/go-ingestion.yaml

# Deploy Python Data Service
kubectl apply -f infrastructure/kubernetes/deployments/python-api.yaml

# Deploy Frontend
kubectl apply -f infrastructure/kubernetes/deployments/frontend.yaml

# Deploy Ingress
kubectl apply -f infrastructure/kubernetes/ingress/ingress.yaml

# Check deployment status
kubectl get pods -n retail-pricing
kubectl get services -n retail-pricing
```

#### Scale Services

```powershell
# Scale Go Ingestion Service
kubectl scale deployment go-ingestion --replicas=3 -n retail-pricing

# Scale Python API Service
kubectl scale deployment python-api --replicas=5 -n retail-pricing

# Auto-scaling
kubectl autoscale deployment go-ingestion --cpu-percent=70 --min=2 --max=10 -n retail-pricing
```

### Production Deployment Checklist

- [ ] Environment variables configured
- [ ] Database migrations applied
- [ ] Database backups configured
- [ ] SSL/TLS certificates installed
- [ ] Monitoring and alerting set up
- [ ] Logging configured
- [ ] Rate limiting enabled
- [ ] CORS origins configured
- [ ] Authentication/authorization enabled
- [ ] Health checks configured
- [ ] Auto-scaling policies set
- [ ] Disaster recovery plan documented
- [ ] Load testing completed
- [ ] Security audit completed

---

## 🧪 Testing

### Unit Tests

#### Go Service
```powershell
# Run all tests
go test ./... -v

# Run with coverage
go test ./... -v -cover -coverprofile=coverage.out

# View coverage report
go tool cover -html=coverage.out

# Run specific test
go test ./internal/services -v -run TestValidateCSV
```

#### Python Service
```powershell
# Activate virtual environment
.\venv\Scripts\Activate.ps1

# Run all tests
pytest tests/ -v

# Run with coverage
pytest tests/ -v --cov=app --cov-report=html

# Run specific test
pytest tests/test_pricing.py::test_search_pricing -v
```

### Integration Tests

```powershell
# Start test environment
docker-compose -f docker-compose.test.yml up -d

# Run integration tests
go test ./tests/integration -v

# Cleanup
docker-compose -f docker-compose.test.yml down -v
```

### Load Testing

#### Using Apache Bench
```powershell
# Test health endpoint
ab -n 10000 -c 100 http://localhost:8080/health

# Test upload endpoint (requires proper multipart data)
ab -n 1000 -c 10 -T "multipart/form-data" -p test.csv http://localhost:8080/upload
```

#### Using k6
```powershell
# Install k6
choco install k6

# Run load test
k6 run scripts/load-tests/upload-test.js

# Run with custom VUs and duration
k6 run --vus 100 --duration 5m scripts/load-tests/search-test.js
```

### End-to-End Tests

```powershell
# Frontend E2E tests (Cypress)
cd frontend
npm run e2e

# Run headless
npm run e2e:headless
```

---

## 📊 Monitoring & Observability

### Metrics (Prometheus)

**Key Metrics:**
- `http_requests_total` - Total HTTP requests
- `http_request_duration_seconds` - Request latency
- `http_errors_total` - Error count
- `csv_files_processed_total` - Files processed
- `csv_records_inserted_total` - Records inserted
- `database_connections_active` - Active DB connections
- `cache_hit_rate` - Redis cache hit rate

**Accessing Metrics:**
- Go Service: http://localhost:8080/metrics
- Python Service: http://localhost:8000/metrics

### Dashboards (Grafana)

**Pre-built Dashboards:**
1. System Overview
2. Upload Service Metrics
3. Data Service Metrics
4. Database Performance
5. User Activity

**Access:** http://localhost:3000 (default credentials: admin/admin)

### Logging (ELK Stack)

**Log Levels:**
- DEBUG: Detailed diagnostic information
- INFO: General informational messages
- WARN: Warning messages
- ERROR: Error messages
- FATAL: Critical errors

**Log Format (JSON):**
```json
{
  "timestamp": "2026-02-17T10:30:00Z",
  "level": "INFO",
  "service": "go-ingestion",
  "message": "File uploaded successfully",
  "correlation_id": "abc123",
  "user_id": "user@example.com",
  "filename": "pricing.csv",
  "records": 5000
}
```

### Alerts

**Critical Alerts:**
- Service downtime (> 1 minute)
- Error rate > 5%
- Response time > 3 seconds
- Database connection failures
- Disk space < 10%

**Warning Alerts:**
- CPU usage > 80%
- Memory usage > 85%
- Error rate > 2%
- Response time > 2 seconds

---

## 🔒 Security

### Authentication Flow
1. User logs in via SSO (OAuth 2.0)
2. Backend issues JWT token (1-hour expiration)
3. Frontend stores token in memory
4. Token included in Authorization header for API requests
5. Backend validates token on each request

### Authorization (RBAC)

**Roles:**
- **Store Manager:**
  - Upload files
  - View own store data
  
- **Analyst:**
  - Upload files
  - Search all data
  - Edit all data
  - Export data
  
- **Administrator:**
  - All analyst permissions
  - User management
  - System configuration
  - View audit logs

### Data Protection

**In Transit:**
- TLS 1.3 for all API communication
- Certificate pinning for mobile apps

**At Rest:**
- Database encryption (AES-256)
- S3 server-side encryption (SSE-S3)
- Encrypted backups

**Input Validation:**
- Server-side validation for all inputs
- Parameterized SQL queries
- File type validation
- File size limits
- Content sanitization

---

## 🤝 Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.

### Development Workflow
1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Write tests for your changes
4. Make your changes
5. Run tests and linters
6. Commit your changes (`git commit -m 'Add amazing feature'`)
7. Push to the branch (`git push origin feature/amazing-feature`)
8. Open a Pull Request

### Code Standards

**Go:**
- Follow [Effective Go](https://golang.org/doc/effective_go.html)
- Use `gofmt` for formatting
- Run `golangci-lint` for linting
- Minimum 80% test coverage

**Python:**
- Follow PEP 8
- Use `black` for formatting
- Use `flake8` for linting
- Use type hints
- Minimum 80% test coverage

**Angular:**
- Follow [Angular Style Guide](https://angular.io/guide/styleguide)
- Use `prettier` for formatting
- Use `eslint` for linting
- Use TypeScript strict mode

**Commit Messages:**
- Follow [Conventional Commits](https://www.conventionalcommits.org/)
- Format: `type(scope): description`
- Example: `feat(upload): add progress bar for large files`

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 👥 Team & Support

**Project Team:**
- Architecture Lead: [Name]
- Backend Lead (Go): [Name]
- Backend Lead (Python): [Name]
- Frontend Lead (Angular): [Name]
- DevOps Lead: [Name]
- QA Lead: [Name]

**Contact:**
- Email: pricing-system-support@company.com
- Slack: #retail-pricing-system
- Issue Tracker: GitHub Issues

---

## 🗺️ Roadmap

### Phase 1 (Q1 2026) - ✅ Completed
- [x] Basic upload functionality
- [x] Search and filter
- [x] Edit capabilities
- [x] Authentication
- [x] PostgreSQL integration
- [x] Docker deployment

### Phase 2 (Q2 2026) - 🚧 In Progress
- [x] S3 integration with presigned URLs
- [x] Event-driven architecture
- [ ] Redis caching
- [ ] Real-time notifications (WebSocket)
- [ ] Bulk operations
- [ ] Advanced analytics dashboard

### Phase 3 (Q3 2026) - 📋 Planned
- [ ] Scheduled imports (FTP/SFTP)
- [ ] Price comparison across stores
- [ ] Mobile app (iOS/Android)
- [ ] Multi-language support
- [ ] Advanced reporting
- [ ] Data export (Excel, PDF)

### Phase 4 (Q4 2026) - 💡 Future
- [ ] Machine learning price predictions
- [ ] Automated price optimization
- [ ] Competitive price monitoring
- [ ] AI-powered insights
- [ ] Integration with POS systems
- [ ] Real-time price updates

---

## 📖 Additional Resources

### Documentation
- [Architecture Decision Records (ADR)](docs/architecture/adr/)
- [API Documentation](docs/api/)
- [User Guides](docs/user-guides/)
- [Deployment Guide](docs/deployment/deployment-guide.md)
- [Troubleshooting Guide](docs/troubleshooting.md)
- [Database Schema](database/schema.sql)

### Related Repositories
- [Frontend Repository](https://github.com/yourusername/frontend)
- [Python Data Service](https://github.com/yourusername/data-service)
- [Infrastructure](https://github.com/yourusername/infrastructure)

### External Resources
- [Go Documentation](https://golang.org/doc/)
- [Gin Framework](https://gin-gonic.com/)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)
- [FastAPI Documentation](https://fastapi.tiangolo.com/)
- [Angular Documentation](https://angular.io/docs)

---

## 🙏 Acknowledgments

- Gin framework team for excellent Go HTTP framework
- FastAPI team for modern Python web framework
- Angular team for robust frontend framework
- PostgreSQL community for reliable database
- Open source community for amazing tools and libraries

---

**Repository:** https://github.com/yourusername/golang-project-product-service

**Last Updated:** February 17, 2026

**Version:** 1.0.0
