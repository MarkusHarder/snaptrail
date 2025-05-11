import { Component, input } from '@angular/core';
import { DataViewModule } from 'primeng/dataview';
import { Session } from '../../models/session';
import { ConfirmationService, MessageService } from 'primeng/api';
import { SessionService } from '../../services/session.service';
import { ButtonModule } from 'primeng/button';
import { SessionsFormComponent } from '../sessions-form/sessions-form.component';
import { ConfirmDialog } from 'primeng/confirmdialog';
import { CardModule } from 'primeng/card';
import { ImageCardComponent } from '../image-card/image-card.component';
@Component({
  selector: 'app-sessions-display',
  imports: [
    DataViewModule,
    ButtonModule,
    SessionsFormComponent,
    ImageCardComponent,
    ConfirmDialog,
    CardModule,
  ],
  providers: [ConfirmationService, MessageService],
  templateUrl: './sessions-display.component.html',
  styleUrl: './sessions-display.component.css',
})
export class SessionsDisplayComponent {
  public = input<boolean>(true);
  sessions = input.required<Session[]>();
  visible = false;
  confirmVisible = false;
  selectedSession?: Session;

  constructor(
    private confirmationService: ConfirmationService,
    private messageService: MessageService,
    private sessionService: SessionService,
  ) {}

  showCreate() {
    this.visible = true;
    this.selectedSession = undefined;
  }

  showEdit(session: Session) {
    this.selectedSession = session;
    this.visible = true;
  }

  showDelete(event: Event, session: Session) {
    this.confirmationService.confirm({
      target: event.target as EventTarget,
      message: 'Do you want to delete this record?',
      header: 'Danger Zone',
      icon: 'pi pi-info-circle',
      rejectLabel: 'Cancel',
      rejectButtonProps: {
        label: 'Cancel',
        severity: 'secondary',
        outlined: true,
      },
      acceptButtonProps: {
        label: 'Delete',
        severity: 'danger',
      },

      accept: () => {
        if (!session.id) {
          this.messageService.add({
            severity: 'error',
            summary: 'Error',
            detail: 'There was an error',
          });
          return;
        }
        this.sessionService.deleteSession(session.id);
        this.messageService.add({
          severity: 'info',
          summary: 'Confirmed',
          detail: 'Record deleted',
        });
      },
      reject: () => {
        this.messageService.add({
          severity: 'error',
          summary: 'Rejected',
          detail: 'You have rejected',
        });
      },
    });
  }
}
