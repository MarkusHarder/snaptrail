import { ComponentFixture, TestBed } from '@angular/core/testing';

import { provideHttpClient } from '@angular/common/http';
import { MessageService } from 'primeng/api';
import { ToggleButtonModule } from 'primeng/togglebutton';
import { SessionService } from '../../services/session.service';
import { SessionsFormComponent } from '../sessions-form/sessions-form.component';
import { SessionsComponent } from './sessions.component';

describe('SessionsComponent', () => {
  let component: SessionsComponent;
  let fixture: ComponentFixture<SessionsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SessionsComponent, SessionsFormComponent, ToggleButtonModule],
      providers: [provideHttpClient(), MessageService, SessionService],
    }).compileComponents();

    fixture = TestBed.createComponent(SessionsComponent);
    component = fixture.componentInstance;
    //TODO: togglebutton currently causes errors https://github.com/primefaces/primeng/pull/18153 and textarea also has some
    // issues https://github.com/primefaces/primeng/issues/18159, will have to take a look at this later on
    // fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
