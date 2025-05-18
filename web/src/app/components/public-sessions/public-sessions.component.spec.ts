import { ComponentFixture, TestBed } from '@angular/core/testing';

import { provideHttpClient } from '@angular/common/http';
import { MessageService } from 'primeng/api';
import { ToggleButtonModule } from 'primeng/togglebutton';
import { PublicSessionsComponent } from './public-sessions.component';
import { SessionsFormComponent } from '../sessions-form/sessions-form.component';
import { ReactiveFormsModule } from '@angular/forms';

describe('PublicSessionsComponent', () => {
  let component: PublicSessionsComponent;
  let fixture: ComponentFixture<PublicSessionsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        PublicSessionsComponent,
        SessionsFormComponent,
        ReactiveFormsModule,
      ],
      providers: [provideHttpClient(), MessageService, ToggleButtonModule],
    }).compileComponents();

    fixture = TestBed.createComponent(PublicSessionsComponent);
    component = fixture.componentInstance;
    //TODO: togglebutton currently causes errors https://github.com/primefaces/primeng/pull/18153 and textarea also has some
    // issues https://github.com/primefaces/primeng/issues/18159, will have to take a look at this later on
    // fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
