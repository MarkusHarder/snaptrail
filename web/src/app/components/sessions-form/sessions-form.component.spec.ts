import { ComponentFixture, TestBed } from '@angular/core/testing';

import { provideHttpClient } from '@angular/common/http';
import { MessageService } from 'primeng/api';
import { ToggleButton } from 'primeng/togglebutton';
import { AdminService } from '../../services/admin.service';
import { SessionsFormComponent } from './sessions-form.component';
import { CUSTOM_ELEMENTS_SCHEMA, NO_ERRORS_SCHEMA } from '@angular/core';

describe('SessionsFormComponent', () => {
  let component: SessionsFormComponent;
  let fixture: ComponentFixture<SessionsFormComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SessionsFormComponent, ToggleButton],
      schemas: [NO_ERRORS_SCHEMA, CUSTOM_ELEMENTS_SCHEMA],
      providers: [provideHttpClient(), MessageService, AdminService],
    }).compileComponents();

    fixture = TestBed.createComponent(SessionsFormComponent);
    component = fixture.componentInstance;
    fixture.componentRef.setInput('visible', false);
    //TODO: togglebutton currently causes errors https://github.com/primefaces/primeng/pull/18153 and textarea also has some
    // issues https://github.com/primefaces/primeng/issues/18159, will have to take a look at this later on
    // fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
