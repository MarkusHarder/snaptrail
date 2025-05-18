import { TestBed } from '@angular/core/testing';

import { provideHttpClient } from '@angular/common/http';
import { AdminService } from './admin.service';
import { MessageService } from 'primeng/api';

describe('AdminService', () => {
  let service: AdminService;

  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [provideHttpClient(), MessageService],
    });
    service = TestBed.inject(AdminService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
