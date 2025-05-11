import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PublicSessionsComponent } from './public-sessions.component';

describe('PublicSessionsComponent', () => {
  let component: PublicSessionsComponent;
  let fixture: ComponentFixture<PublicSessionsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [PublicSessionsComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PublicSessionsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
