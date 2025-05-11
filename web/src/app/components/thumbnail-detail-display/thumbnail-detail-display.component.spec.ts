import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ThumbnailDetailDisplayComponent } from './thumbnail-detail-display.component';

describe('ThumbnailDetailDisplayComponent', () => {
  let component: ThumbnailDetailDisplayComponent;
  let fixture: ComponentFixture<ThumbnailDetailDisplayComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ThumbnailDetailDisplayComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ThumbnailDetailDisplayComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
