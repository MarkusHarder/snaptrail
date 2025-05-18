import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Session } from '../models/session';
import { BehaviorSubject, forkJoin, map, Observable } from 'rxjs';
import { environment } from '../../environments/environment';
import { MessageService } from 'primeng/api';

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
      next: (sessions) => {
        const thumbnailRequests = sessions.map((session) =>
          this.getThumbnailBlobs(
            session.id!,
            session.thumbnail!.id!,
            published,
          ).pipe(
            map((blob) => {
              session.thumbnail!.data = URL.createObjectURL(blob);
              session.thumbnail!.rawData = blob;
              return session;
            }),
          ),
        );
        if (thumbnailRequests.length === 0) {
          this.sessionsSubject.next(sessions);
        } else
          forkJoin(thumbnailRequests).subscribe({
            next: (updatedSessions) => {
              this.sessionsSubject.next(updatedSessions);
            },
            error: (err) => {
              console.error('Failed to fetch thumbnails:', err);
              this.messageService.add({
                severity: 'error',
                summary: 'Error',
                detail: 'There was an error fetching thumbnails',
              });
            },
          });
      },
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

  createOrUpdateSession(sessionForm: FormData, id?: number) {
    let path = this.basePath() + '/admin/sessions';
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

  deleteSession(id: number) {
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

  getThumbnailBlobs(
    sessionId: number,
    thumbnailId: number,
    published = true,
  ): Observable<Blob> {
    let path = this.basePath();
    if (!published) path += '/admin';
    path += `/sessions/${sessionId}/thumbnails/${thumbnailId}`;
    return this.http.get(path, {
      responseType: 'blob',
    });
  }
  basePath(): string {
    return environment.basePath;
  }
}
