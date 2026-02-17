# Pricing Data Management System

### Architect-Level Solution Design

------------------------------------------------------------------------

## 1. Overview

This system enables retail store users to:

-   Upload pricing data via CSV
-   Search pricing records
-   Update pricing records

The solution is designed using a **microservices architecture**,
leveraging: - Presigned URL uploads - Event-driven ingestion - Stateless
services - Horizontal scalability principles

The design prioritizes scalability, reliability, maintainability, and
production-readiness.

------------------------------------------------------------------------

# 2. Context Diagram

The context diagram illustrates system boundaries and external
interactions.

    +----------------------+
    |   Retail Store User  |
    +----------+-----------+
               |
               v
    +----------------------+
    |     Angular UI       |
    +----------+-----------+
               |
               v
    +----------------------+
    |  Backend Services    |
    | (API + Ingestion)    |
    +----------+-----------+
               |
        +------+-------+
        |              |
        v              v
    +--------+     +------------+
    |  RDBMS |     |  Object    |
    |Pricing |     |  Storage   |
    |Records |     |  (CSV)     |
    +--------+     +------------+

------------------------------------------------------------------------

# 3. Solution Architecture

                        +--------------------+
                        |     Angular UI     |
                        +----------+---------+
                                   |
                                   v
                        +--------------------+
                        | Python API Service |
                        |     (FastAPI)      |
                        |--------------------|
                        | - Generate         |
                        |   Presigned URL    |
                        | - Search API       |
                        | - Update API       |
                        +----------+---------+
                                   |
                                   v
                         +-------------------+
                         |   Object Storage  |
                         |   (S3 / Local)    |
                         +----------+--------+
                                    |
                                    v
                         +-------------------+
                         | Storage Event     |
                         |   Notification    |
                         +----------+--------+
                                    |
                                    v
                         +-------------------+
                         | Go Ingestion Svc  |
                         |       (Gin)       |
                         |-------------------|
                         | - Event Listener  |
                         | - CSV Streaming   |
                         | - Validation      |
                         | - DB Persistence  |
                         +----------+--------+
                                    |
                                    v
                           +----------------+
                           |  PostgreSQL    |
                           | (SQLite local) |
                           +----------------+

------------------------------------------------------------------------

# 4. Detailed Flow Design

## 4.1 CSV Upload Flow (Presigned URL + Event Driven)

1.  User requests file upload from UI.
2.  Python API generates a presigned URL.
3.  UI uploads file directly to object storage.
4.  Storage triggers an event notification.
5.  Go ingestion service consumes event.
6.  CSV is streamed and validated.
7.  Valid records inserted transactionally into DB.
8.  Errors logged for audit and traceability.

### Why This Approach?

-   Removes large file handling from API servers.
-   Reduces memory pressure.
-   Improves scalability.
-   Enables asynchronous processing.

------------------------------------------------------------------------

## 4.2 Search Flow

1.  UI calls Python API with query parameters.
2.  Indexed DB query executed.
3.  Paginated response returned.

Indexes: - store_id - sku - effective_date

------------------------------------------------------------------------

## 4.3 Update Flow

1.  UI submits update request.
2.  API validates payload.
3.  Transactional update executed.
4.  Updated record returned.

------------------------------------------------------------------------

# 5. Design Decisions

## Microservices Separation

**Reasoning:** - Ingestion is I/O-heavy and benefits from Go
concurrency. - CRUD APIs require rapid development and strong ecosystem
support (Python). - Independent scaling of services.

------------------------------------------------------------------------

## Presigned URL Upload

**Reasoning:** - Avoids loading large files in API layer. - Reduces
latency. - Enables horizontal scaling.

------------------------------------------------------------------------

## Event-Driven Processing

**Reasoning:** - Decouples upload from ingestion. - Improves fault
tolerance. - Supports retries and worker pool processing.

------------------------------------------------------------------------

## Relational Database Choice

**Reasoning:** - Structured pricing data. - ACID compliance required. -
Strong consistency preferred. - Indexed for performance.

------------------------------------------------------------------------

## Stateless Services

**Reasoning:** - Simplifies scaling. - Enables containerized
deployments. - Improves resilience.

------------------------------------------------------------------------

# 6. Non-Functional Requirements & Design Alignment

## Scalability

-   Stateless services.
-   Direct-to-storage uploads.
-   Worker pool ingestion.
-   Horizontal scaling capability.

## Performance

-   Streaming CSV parsing (constant memory).
-   Indexed queries.
-   Pagination.
-   Efficient concurrency model (Go).

## Reliability

-   Raw CSV retained before processing.
-   Transactional DB operations.
-   Idempotent ingestion logic.
-   Retry-friendly event design.

## Security

-   JWT-based authentication (conceptual).
-   Role-based access control.
-   Expiring presigned URLs.
-   Input validation & sanitization.

## Maintainability

-   Clear separation of concerns.
-   Modular services.
-   Replaceable storage abstraction.
-   Config-driven environments.

## Observability

-   Structured logging.
-   Health endpoints.
-   Metrics-ready design.
-   Error tracking during ingestion.

------------------------------------------------------------------------

# 7. Assumptions

1.  CSV schema is predefined and validated.
2.  Authentication handled by external identity provider.
3.  Storage events are reliably delivered.
4.  Strong consistency required for pricing updates.
5.  Near real-time ingestion acceptable.
6.  File sizes within acceptable operational limits.

------------------------------------------------------------------------

# 8. Source for Implementation

## Backend

-   Python (FastAPI)
-   Go (Gin)
-   PostgreSQL (SQLite for local development)
-   S3 (Simulated locally)

## Frontend

-   Angular SPA

## Local Development Notes

-   Object storage simulated using filesystem.
-   Event simulation via file watcher/background worker.
-   SQLite used for simplified setup.

------------------------------------------------------------------------

# 9. Architectural Principles

-   Separation of Concerns
-   Event-Driven Architecture
-   Stateless Design
-   Horizontal Scalability
-   Production-Oriented Thinking
-   Minimal Over-Engineering

------------------------------------------------------------------------

End of Document
