export interface PricingRecord {
  id?: number;
  store_id: string;
  sku: string;
  product_name: string;
  price: string | number;
  date: string;
  created_at?: string;
  updated_at?: string;
}

export interface PricingSearchResponse {
  items: PricingRecord[];
  total: number;
  page: number;
  page_size: number;
  total_pages: number;
}

export interface PresignedUrlResponse {
  upload_url: string;
  file_key: string;
  expires_in: number;
}

export interface UploadResponse {
  message: string;
  filename: string;
  size: number;
  records_processed?: number;
}
