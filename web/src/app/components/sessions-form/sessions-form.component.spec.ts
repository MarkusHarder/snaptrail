import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SessionsFormComponent } from './sessions-form.component';

describe('SessionsFormComponent', () => {
  let component: SessionsFormComponent;
  let fixture: ComponentFixture<SessionsFormComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SessionsFormComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SessionsFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
