import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SessionsDisplayComponent } from './sessions-display.component';

describe('SessionsDisplayComponent', () => {
  let component: SessionsDisplayComponent;
  let fixture: ComponentFixture<SessionsDisplayComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SessionsDisplayComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SessionsDisplayComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
