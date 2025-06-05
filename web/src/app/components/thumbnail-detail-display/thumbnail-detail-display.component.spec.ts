import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ThumbnailDetailDisplayComponent } from './thumbnail-detail-display.component';
import { provideHttpClientTesting } from '@angular/common/http/testing';
import { Thumbnail } from '../../models/session';

describe('ThumbnailDetailDisplayComponent', () => {
  let component: ThumbnailDetailDisplayComponent;
  let fixture: ComponentFixture<ThumbnailDetailDisplayComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ThumbnailDetailDisplayComponent],
      providers: [provideHttpClientTesting],
    }).compileComponents();

    fixture = TestBed.createComponent(ThumbnailDetailDisplayComponent);
    component = fixture.componentInstance;
    const t = {
      id: 'abc-123',
      filename: 'image123.jpg',
      mimeType: 'image/jpeg',
      cameraModel: 'Canon EOS 5D Mark IV',
      make: 'Canon',
      lensModel: 'Canon EF 24-70mm f/2.8L II USM',
      exposure: '1/200s',
      dateTime: '2024-05-11T14:00:00',
      aperture: 2.8,
      iso: 400,
      fc: 35,
    } as Thumbnail;
    fixture.componentRef.setInput('thumbnail', t);

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
