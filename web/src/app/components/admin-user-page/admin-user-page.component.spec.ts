import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AdminUserPageComponent } from './admin-user-page.component';
import { provideHttpClient } from '@angular/common/http';
import { AdminService } from '../../services/admin.service';
import { MessageService } from 'primeng/api';

describe('AdminUserPageComponent', () => {
  let component: AdminUserPageComponent;
  let fixture: ComponentFixture<AdminUserPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [AdminUserPageComponent],
      providers: [provideHttpClient(), AdminService, MessageService],
    }).compileComponents();

    fixture = TestBed.createComponent(AdminUserPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
