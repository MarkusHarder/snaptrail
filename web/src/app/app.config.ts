import { ApplicationConfig, provideZoneChangeDetection } from '@angular/core';
import { provideAnimationsAsync } from '@angular/platform-browser/animations/async';
import { provideRouter } from '@angular/router';
import { provideHttpClient, withInterceptors } from '@angular/common/http';
import Material from '@primeng/themes/material';

import { routes } from './app.routes';
import { providePrimeNG } from 'primeng/config';
import {
  authHeaderInterceptor,
  tokenValidatorInterceptor,
} from './interceptors/auth';
import { MessageService } from 'primeng/api';

export const appConfig: ApplicationConfig = {
  providers: [
    MessageService,
    provideZoneChangeDetection({ eventCoalescing: true }),
    provideRouter(routes),
    provideHttpClient(
      withInterceptors([authHeaderInterceptor, tokenValidatorInterceptor]),
    ),
    provideAnimationsAsync(),
    providePrimeNG({
      theme: {
        preset: Material,
        options: {
          cssLayer: {
            name: 'primeng',
            order: 'tailwind, primeng',
          },
        },
      },
    }),
  ],
};
