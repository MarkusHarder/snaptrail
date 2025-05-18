import { Component } from '@angular/core';
import { SessionsDisplayComponent } from '../sessions-display/sessions-display.component';
import { AsyncPipe } from '@angular/common';
import { SessionService } from '../../services/session.service';
import { Observable } from 'rxjs';
import { Session } from '../../models/session';

@Component({
  selector: 'app-public-sessions',
  imports: [SessionsDisplayComponent, AsyncPipe],
  templateUrl: './public-sessions.component.html',
  styleUrl: './public-sessions.component.css',
})
export class PublicSessionsComponent {
  sessions$: Observable<Session[]>;
  constructor(private sessionService: SessionService) {
    this.sessions$ = this.sessionService.sessions$;
    this.sessionService.getSessions()
  }
}
