import { Component, OnInit } from '@angular/core';
import { CardModule } from 'primeng/card';
import { ButtonModule } from 'primeng/button';
import { SessionService } from '../../services/session.service';
import { Observable } from 'rxjs';
import { Session } from '../../models/session';
import { AsyncPipe, CommonModule } from '@angular/common';
import { ImageCardComponent } from '../image-card/image-card.component';
import { DialogModule } from 'primeng/dialog';

@Component({
  selector: 'app-timeline',
  imports: [
    AsyncPipe,
    CommonModule,
    CardModule,
    ButtonModule,
    ImageCardComponent,
    DialogModule,
  ],
  templateUrl: './timeline.component.html',

  styleUrl: './timeline.component.css',
})
export class TimelineComponent implements OnInit {
  sessions$: Observable<Session[]>;
  selectedSession?: Session;
  visible = false;
  constructor(private sessionService: SessionService) {
    this.sessions$ = this.sessionService.sessions$;
  }
  ngOnInit(): void {
    this.sessionService.getSessions();
  }

  showDialog(session: Session) {
    this.selectedSession = session;
    this.visible = true;
  }

  hideDialog() {
    this.visible = false;
    this.selectedSession = undefined;
  }
}
