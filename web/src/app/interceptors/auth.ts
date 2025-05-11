import { inject } from '@angular/core';
import { AuthService } from '../services/auth.service';
import {
  HttpErrorResponse,
  HttpEvent,
  HttpHandlerFn,
  HttpRequest,
  HttpStatusCode,
} from '@angular/common/http';
import { catchError, Observable, throwError } from 'rxjs';

export function authHeaderInterceptor(
  req: HttpRequest<unknown>,
  next: HttpHandlerFn,
): Observable<HttpEvent<unknown>> {
  if (!req.url.includes('/admin')) {
    return next(req);
  }
  const token = inject(AuthService).getToken();

  if (token) {
    const cloned = req.clone({
      setHeaders: {
        Authorization: `Bearer ${token}`,
      },
    });
    return next(cloned);
  }

  return next(req);
}

export function tokenValidatorInterceptor(
  req: HttpRequest<unknown>,
  next: HttpHandlerFn,
): Observable<HttpEvent<unknown>> {
  const authService = inject(AuthService);
  return next(req).pipe(
    catchError((error: HttpErrorResponse) => {
      console.log('intercepting error: ', error);
      if (error.status === HttpStatusCode.Unauthorized) {
        authService.clearToken();
      }
      return throwError(() => error);
    }),
  );
}
