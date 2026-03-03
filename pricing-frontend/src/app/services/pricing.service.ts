import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { Observable } from 'rxjs';
import {
  PricingRecord,
  PricingSearchResponse,
  PresignedUrlResponse,
  UploadResponse
} from '../models/pricing.model';

@Injectable({
  providedIn: 'root'
})
export class PricingService {
  private pythonApiUrl = 'http://localhost:8000/api';
  private goServiceUrl = 'http://localhost:8080';

  constructor(private http: HttpClient) { }

  // Get presigned URL from Python API
  getPresignedUrl(filename: string): Observable<PresignedUrlResponse> {
    return this.http.post<PresignedUrlResponse>(
      `${this.pythonApiUrl}/pricing/upload-url`,
      { filename }
    );
  }

  // Upload file to S3 using presigned URL (or mock upload to Python)
  uploadToPresignedUrl(url: string, file: File): Observable<any> {
    return this.http.put(url, file, {
      headers: { 'Content-Type': 'text/csv' }
    });
  }

  // Upload file directly to Go service
  uploadToGoService(file: File): Observable<UploadResponse> {
    const formData = new FormData();
    formData.append('file', file);
    return this.http.post<UploadResponse>(`${this.goServiceUrl}/upload`, formData);
  }

  // Search pricing records
  searchRecords(params: {
    store_id?: string;
    sku?: string;
    product_name?: string;
    min_price?: number;
    max_price?: number;
    start_date?: string;
    end_date?: string;
    page?: number;
    page_size?: number;
  } = {}): Observable<PricingSearchResponse> {
    let httpParams = new HttpParams();

    Object.keys(params).forEach(key => {
      const value = (params as any)[key];
      if (value !== null && value !== undefined && value !== '') {
        httpParams = httpParams.set(key, value.toString());
      }
    });

    return this.http.get<PricingSearchResponse>(
      `${this.pythonApiUrl}/pricing/search`,
      { params: httpParams }
    );
  }

  // Get record by ID
  getRecordById(id: number): Observable<PricingRecord> {
    return this.http.get<PricingRecord>(`${this.pythonApiUrl}/pricing/records/${id}`);
  }

  // Update record
  updateRecord(id: number, record: Partial<PricingRecord>): Observable<PricingRecord> {
    return this.http.put<PricingRecord>(
      `${this.pythonApiUrl}/pricing/records/${id}`,
      record
    );
  }

  // Delete record
  deleteRecord(id: number): Observable<{ message: string }> {
    return this.http.delete<{ message: string }>(
      `${this.pythonApiUrl}/pricing/records/${id}`
    );
  }

  // Create record
  createRecord(record: Omit<PricingRecord, 'id'>): Observable<PricingRecord> {
    return this.http.post<PricingRecord>(
      `${this.pythonApiUrl}/pricing/records`,
      record
    );
  }
}
