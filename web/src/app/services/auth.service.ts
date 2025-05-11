import { Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { BehaviorSubject } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class AuthService {
  private loggedInSubject = new BehaviorSubject<boolean>(false);
  readonly loggedIn$ = this.loggedInSubject.asObservable();

  private usernameSubject = new BehaviorSubject<string>('');
  readonly username$ = this.usernameSubject.asObservable();
  constructor(private router: Router) {
    this.hasToken();
  }

  storeToken(token: string) {
    sessionStorage.setItem('access_token', token);
    this.hasToken();
  }

  storeUsername(username: string) {
    this.usernameSubject.next(username);
  }

  getToken(): string | null {
    return sessionStorage.getItem('access_token');
  }

  clearToken() {
    sessionStorage.removeItem('access_token');
    this.hasToken();
    this.router.navigate(['/admin/login']);
  }

  hasToken(): boolean {
    const token = sessionStorage.getItem('access_token');
    this.loggedInSubject.next(!!token);
    return !!token;
  }
}
