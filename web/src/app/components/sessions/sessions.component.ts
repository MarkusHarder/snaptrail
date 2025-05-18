import { Component } from '@angular/core';
import { ButtonModule } from 'primeng/button';
import { DataViewModule } from 'primeng/dataview';
import { ToastModule } from 'primeng/toast';
import { SessionService } from '../../services/session.service';
import { Observable } from 'rxjs';
import { Session } from '../../models/session';
import { AsyncPipe } from '@angular/common';
import { ConfirmationService, MessageService } from 'primeng/api';
import { SessionsDisplayComponent } from '../sessions-display/sessions-display.component';

@Component({
  selector: 'app-sessions',
  imports: [
    ButtonModule,
    DataViewModule,
    AsyncPipe,
    ToastModule,
    SessionsDisplayComponent,
  ],
  providers: [ConfirmationService, MessageService],
  templateUrl: './sessions.component.html',
  styleUrl: './sessions.component.css',
})
export class SessionsComponent {
  sessions$: Observable<Session[]>;

  constructor(private sessionService: SessionService) {
    this.sessions$ = this.sessionService.sessions$;
    this.sessionService.getSessions(false);
  }
}
