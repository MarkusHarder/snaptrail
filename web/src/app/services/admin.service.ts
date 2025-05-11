import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from '../../environments/environment';
import { PasswordChange, User } from '../models/user';
import { AuthService } from './auth.service';
import { MessageService } from 'primeng/api';

interface AuthToken {
  token?: string;
}
@Injectable({
  providedIn: 'root',
})
export class AdminService {
  constructor(
    private http: HttpClient,
    private authService: AuthService,
    private messageService: MessageService,
  ) {}

  login(user: User) {
    const credentials = btoa(`${user.username}:${user.password}`); // base64 encode
    const headers = new HttpHeaders({
      Authorization: `Basic ${credentials}`,
    });

    this.http.post(this.basePath() + '/login', null, { headers }).subscribe({
      next: (res: AuthToken) => {
        if (!res.token) {
          this.messageService.add({
            severity: 'error',
            summary: 'Error',
            detail: 'There was an error logging in',
          });
          console.error('Failed to log in: invalid token');
          return;
        }
        const token = res.token;
        this.authService.storeToken(token);
        this.authService.storeUsername(user.username);
      },
      error: (err) => {
        this.messageService.add({
          severity: 'error',
          summary: 'Error',
          detail: 'There was an error logging in',
        });
        console.error('Failed to log in:', err);
      },
    });
  }

  changePassowrd(pwChange: PasswordChange) {
    this.http.post(this.basePath() + '/admin/users', pwChange).subscribe({
      next: () => {
        this.authService.clearToken();
      },
      error: (err) => {
        this.messageService.add({
          severity: 'error',
          summary: 'Error',
          detail: 'There was an updating the password',
        });
        console.error('Failed to create or update session:', err);
      },
    });
  }

  basePath(): string {
    return environment.basePath;
  }
}
