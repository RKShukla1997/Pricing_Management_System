import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { Router } from '@angular/router';
import { PricingService } from '../../services/pricing.service';
import { PricingRecord } from '../../models/pricing.model';

@Component({
  selector: 'app-records',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './records.component.html',
  styleUrls: ['./records.component.css']
})
export class RecordsComponent implements OnInit {
  records: PricingRecord[] = [];
  loading = false;
  error = '';

  // Pagination
  currentPage = 1;
  pageSize = 10;
  totalRecords = 0;
  totalPages = 0;

  // Filters
  filters = {
    store_id: '',
    sku: '',
    product_name: '',
    min_price: null as number | null,
    max_price: null as number | null,
    start_date: '',
    end_date: ''
  };

  // Edit mode
  editingRecord: PricingRecord | null = null;
  editForm: Partial<PricingRecord> = {};

  constructor(
    private pricingService: PricingService,
    private router: Router
  ) {}

  ngOnInit(): void {
    this.loadRecords();
  }

  loadRecords(): void {
    this.loading = true;
    this.error = '';

    const params: any = {
      ...this.filters,
      page: this.currentPage,
      page_size: this.pageSize
    };

    // Remove null values
    Object.keys(params).forEach(key => {
      if (params[key] === null) {
        delete params[key];
      }
    });

    this.pricingService.searchRecords(params).subscribe({
      next: (response) => {
        this.records = response.items;
        this.totalRecords = response.total;
        this.totalPages = response.total_pages;
        this.currentPage = response.page;
        this.loading = false;
      },
      error: (error) => {
        this.error = `Failed to load records: ${error.error?.detail || error.message}`;
        this.loading = false;
      }
    });
  }

  applyFilters(): void {
    this.currentPage = 1;
    this.loadRecords();
  }

  clearFilters(): void {
    this.filters = {
      store_id: '',
      sku: '',
      product_name: '',
      min_price: null,
      max_price: null,
      start_date: '',
      end_date: ''
    };
    this.currentPage = 1;
    this.loadRecords();
  }

  nextPage(): void {
    if (this.currentPage < this.totalPages) {
      this.currentPage++;
      this.loadRecords();
    }
  }

  previousPage(): void {
    if (this.currentPage > 1) {
      this.currentPage--;
      this.loadRecords();
    }
  }

  goToPage(page: number): void {
    if (page >= 1 && page <= this.totalPages) {
      this.currentPage = page;
      this.loadRecords();
    }
  }

  editRecord(record: PricingRecord): void {
    this.editingRecord = record;
    this.editForm = { ...record };
  }

  cancelEdit(): void {
    this.editingRecord = null;
    this.editForm = {};
  }

  saveEdit(): void {
    if (!this.editingRecord || !this.editingRecord.id) return;

    this.loading = true;
    this.pricingService.updateRecord(this.editingRecord.id, this.editForm).subscribe({
      next: (updated) => {
        // Update the record in the list
        const index = this.records.findIndex(r => r.id === updated.id);
        if (index !== -1) {
          this.records[index] = updated;
        }
        this.cancelEdit();
        this.loading = false;
      },
      error: (error) => {
        this.error = `Failed to update record: ${error.error?.detail || error.message}`;
        this.loading = false;
      }
    });
  }

  deleteRecord(record: PricingRecord): void {
    if (!record.id) return;

    if (!confirm(`Are you sure you want to delete this record?\n\nStore: ${record.store_id}\nSKU: ${record.sku}`)) {
      return;
    }

    this.loading = true;
    this.pricingService.deleteRecord(record.id).subscribe({
      next: () => {
        this.loadRecords();
      },
      error: (error) => {
        this.error = `Failed to delete record: ${error.error?.detail || error.message}`;
        this.loading = false;
      }
    });
  }

  goToUpload(): void {
    this.router.navigate(['/upload']);
  }

  getPageNumbers(): number[] {
    const pages: number[] = [];
    const start = Math.max(1, this.currentPage - 2);
    const end = Math.min(this.totalPages, this.currentPage + 2);

    for (let i = start; i <= end; i++) {
      pages.push(i);
    }

    return pages;
  }
}
