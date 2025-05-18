import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';
import { AuthService } from '../services/auth.service';
import { Observable, Observer } from 'rxjs';

export const authGuard: CanActivateFn = () => {
  const authService = inject(AuthService);
  const router = inject(Router);

  return new Observable<boolean>((observer: Observer<boolean>) => {
    if (authService.hasToken()) {
      observer.next(true);
    } else {
      router.navigate(['/admin/login']);
      observer.next(false);
    }
    observer.complete();
  });
};
