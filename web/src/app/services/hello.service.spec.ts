import { TestBed } from '@angular/core/testing';

import { HelloService } from './hello.service';
import { provideHttpClient } from '@angular/common/http';

describe('HelloService', () => {
  let service: HelloService;

  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [provideHttpClient()],
    });
    service = TestBed.inject(HelloService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
