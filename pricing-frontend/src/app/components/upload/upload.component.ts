import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { PricingService } from '../../services/pricing.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-upload',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './upload.component.html',
  styleUrls: ['./upload.component.css']
})
export class UploadComponent {
  selectedFile: File | null = null;
  uploading = false;
  uploadSuccess = false;
  uploadError = false;
  errorMessage = '';
  successMessage = '';
  uploadMethod: 'presigned' | 'direct' = 'direct';

  constructor(
    private pricingService: PricingService,
    private router: Router
  ) {}

  onFileSelected(event: Event): void {
    const input = event.target as HTMLInputElement;
    if (input.files && input.files.length > 0) {
      this.selectedFile = input.files[0];
      this.uploadSuccess = false;
      this.uploadError = false;
      this.errorMessage = '';
      this.successMessage = '';
    }
  }

  uploadFile(): void {
    if (!this.selectedFile) {
      this.showError('Please select a file first');
      return;
    }

    this.uploading = true;
    this.uploadSuccess = false;
    this.uploadError = false;

    if (this.uploadMethod === 'presigned') {
      this.uploadViaPresignedUrl();
    } else {
      this.uploadDirectly();
    }
  }

  private uploadViaPresignedUrl(): void {
    if (!this.selectedFile) return;

    // Step 1: Get presigned URL
    this.pricingService.getPresignedUrl(this.selectedFile.name).subscribe({
      next: (response) => {
        // Step 2: Upload to presigned URL
        if (this.selectedFile) {
          this.pricingService.uploadToPresignedUrl(response.upload_url, this.selectedFile).subscribe({
            next: () => {
              this.uploading = false;
              this.showSuccess('File uploaded successfully via presigned URL');
            },
            error: (error) => {
              this.uploading = false;
              this.showError(`Upload failed: ${error.error?.detail || error.message}`);
            }
          });
        }
      },
      error: (error) => {
        this.uploading = false;
        this.showError(`Failed to get presigned URL: ${error.error?.detail || error.message}`);
      }
    });
  }

  private uploadDirectly(): void {
    if (!this.selectedFile) return;

    // Upload directly to Go service
    this.pricingService.uploadToGoService(this.selectedFile).subscribe({
      next: (response) => {
        this.uploading = false;
        this.showSuccess(`File uploaded successfully! ${response.records_processed || ''} records processed`);
      },
      error: (error) => {
        this.uploading = false;
        this.showError(`Upload failed: ${error.error?.error || error.message}`);
      }
    });
  }

  private showSuccess(message: string): void {
    this.uploadSuccess = true;
    this.successMessage = message;
    this.selectedFile = null;

    // Clear file input
    const fileInput = document.getElementById('fileInput') as HTMLInputElement;
    if (fileInput) fileInput.value = '';
  }

  private showError(message: string): void {
    this.uploadError = true;
    this.errorMessage = message;
  }

  viewRecords(): void {
    this.router.navigate(['/records']);
  }

  setUploadMethod(method: 'presigned' | 'direct'): void {
    this.uploadMethod = method;
    this.uploadSuccess = false;
    this.uploadError = false;
  }
}
