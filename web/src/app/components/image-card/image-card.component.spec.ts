import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ImageCardComponent } from './image-card.component';
import { Thumbnail } from '../../models/session';

describe('ImageCardComponent', () => {
  let component: ImageCardComponent;
  let fixture: ComponentFixture<ImageCardComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ImageCardComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(ImageCardComponent);
    component = fixture.componentInstance;
    const t = {
      id: 'abc-123',
      filename: 'image123.jpg',
      mimeType: 'image/jpeg',
      imageSrc: 'some/image/url/img.png',
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
