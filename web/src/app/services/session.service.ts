import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { MessageService } from 'primeng/api';
import { BehaviorSubject } from 'rxjs';
import { environment } from '../../environments/environment';
import { Session } from '../models/session';

@Injectable({
  providedIn: 'root',
})
export class SessionService {
  private sessionsSubject = new BehaviorSubject<Session[]>([]);
  readonly sessions$ = this.sessionsSubject.asObservable();
  constructor(
    private http: HttpClient,
    private messageService: MessageService,
  ) {}

  getSessions(published = true) {
    let path = this.basePath();
    if (published) path += '/sessions';
    else path += '/admin/sessions';
    this.http.get<Session[]>(path).subscribe({
      next: (sessions) => this.sessionsSubject.next(sessions),
      error: (err) => {
        this.messageService.add({
          severity: 'error',
          summary: 'Error',
          detail: 'There was an error fetching sessions',
        });
        console.error('Failed to load sessions:', err);
      },
    });
  }

  createOrUpdateSession(sessionForm: FormData, id?: string) {
    let path = this.basePath() + '/admin/sessions';
    console.log(id);
    if (id) path += '/' + id;
    this.http.post<Session>(path, sessionForm).subscribe({
      next: () => this.getSessions(false),
      error: (err) => {
        this.messageService.add({
          severity: 'error',
          summary: 'Error',
          detail: 'There was creating or updating the session',
        });
        console.error('Failed to create or update session:', err);
      },
    });
  }

  deleteSession(id: string) {
    this.http
      .delete<void>(this.basePath() + '/admin/sessions/' + id)
      .subscribe({
        next: () => this.getSessions(false),
        error: (err) => {
          this.messageService.add({
            severity: 'error',
            summary: 'Error',
            detail: 'There was creating or updating the session',
          });
          console.error('Failed to delete session:', err);
        },
      });
  }

  basePath(): string {
    return environment.basePath;
  }
}
