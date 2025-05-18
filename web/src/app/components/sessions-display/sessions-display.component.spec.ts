import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SessionsDisplayComponent } from './sessions-display.component';
import { provideHttpClient } from '@angular/common/http';
import { SessionService } from '../../services/session.service';
import { MessageService } from 'primeng/api';
import { Session } from '../../models/session';

describe('SessionsDisplayComponent', () => {
  let component: SessionsDisplayComponent;
  let fixture: ComponentFixture<SessionsDisplayComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SessionsDisplayComponent],
      providers: [provideHttpClient(), SessionService, MessageService],
    }).compileComponents();

    fixture = TestBed.createComponent(SessionsDisplayComponent);
    component = fixture.componentInstance;
    const mockSessions: Session[] = [
      {
        id: 1,
        sessionName: 'Sunset at the Beach',
        subtitle: 'Golden hour by the sea',
        description:
          'Captured during a calm summer evening with perfect lighting.',
        published: true,
        date: new Date('2024-08-01T18:30:00'),
        thumbnail: {
          id: 101,
          filename: 'sunset_beach.jpg',
          mimeType: 'image/jpeg',
          data: 'base64EncodedData1',
          rawData: new Blob(['image data 1'], { type: 'image/jpeg' }),
          cameraModel: 'Sony Alpha 7 III',
          make: 'Sony',
          lensModel: 'Sony FE 24-70mm f/2.8 GM',
          exposure: '1/320s',
          dateTime: '2024-08-01T18:30:00',
          aperture: 2.8,
          iso: 200,
          fc: 35,
        },
      },
      {
        id: 2,
        sessionName: 'Mountain Hike',
        subtitle: 'Morning in the Alps',
        description: 'A fresh early morning hike with misty mountain views.',
        published: false,
        date: new Date('2024-09-15T07:00:00'),
        thumbnail: {
          id: 102,
          filename: 'alps_morning.jpg',
          mimeType: 'image/jpeg',
          data: 'base64EncodedData2',
          rawData: new Blob(['image data 2'], { type: 'image/jpeg' }),
          cameraModel: 'Canon EOS R5',
          make: 'Canon',
          lensModel: 'Canon RF 15-35mm f/2.8L IS USM',
          exposure: '1/125s',
          dateTime: '2024-09-15T07:00:00',
          aperture: 4.0,
          iso: 100,
          fc: 20,
        },
      },
    ];
    fixture.componentRef.setInput('sessions', mockSessions);
    //TODO: togglebutton currently causes errors https://github.com/primefaces/primeng/pull/18153 and textarea also has some
    // issues https://github.com/primefaces/primeng/issues/18159, will have to take a look at this later on
    // fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
