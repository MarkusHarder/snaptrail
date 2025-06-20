import {
  Component,
  CUSTOM_ELEMENTS_SCHEMA,
  inject,
  Input,
  model,
  NO_ERRORS_SCHEMA,
} from '@angular/core';
import { ButtonModule } from 'primeng/button';
import { DatePickerModule } from 'primeng/datepicker';
import { InputTextModule } from 'primeng/inputtext';
import { TextareaModule } from 'primeng/textarea';
import { ToggleButton } from 'primeng/togglebutton';
import { ColorPickerModule } from 'primeng/colorpicker';
import {
  FormBuilder,
  FormControl,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { Dialog } from 'primeng/dialog';
import { SessionService } from '../../services/session.service';
import { Session } from '../../models/session';
import { MessageService } from 'primeng/api';
import { FileUpload } from 'primeng/fileupload';
import { ToastModule } from 'primeng/toast';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-sessions-form',
  imports: [
    DatePickerModule,
    ToggleButton,
    InputTextModule,
    TextareaModule,
    ButtonModule,
    ColorPickerModule,
    ReactiveFormsModule,
    Dialog,
    FileUpload,
    ToastModule,
    CommonModule,
  ],
  schemas: [NO_ERRORS_SCHEMA, CUSTOM_ELEMENTS_SCHEMA],
  providers: [MessageService],
  templateUrl: './sessions-form.component.html',
  styleUrl: './sessions-form.component.css',
})
export class SessionsFormComponent {
  visible = model.required<boolean>();
  uploadedFile: File | null = null;
  filename = '';
  inputSession?: Session;

  private formBuilder = inject(FormBuilder);
  sessionForm = this.formBuilder.group({
    id: new FormControl<string>({ value: '', disabled: true }),
    sessionName: new FormControl('', Validators.required),
    subtitle: new FormControl('', Validators.required),
    description: new FormControl('', Validators.required),
    published: new FormControl<boolean>(false),
    uploadedThumbnail: new FormControl<File | null>(null, Validators.required),
    date: new FormControl<Date | null>(null, Validators.required),
  });

  @Input()
  set selectedSession(session: Session | undefined) {
    this.inputSession = session;
    let d;
    const f = session?.thumbnail
      ? new File([session.thumbnail!.rawData], session?.thumbnail?.filename, {
          type: session?.thumbnail?.mimeType,
        })
      : undefined;
    this.uploadedFile = f ?? null;
    this.filename = this.inputSession?.thumbnail?.filename ?? '';
    if (session?.date) d = new Date(session?.date);
    this.sessionForm.patchValue({
      id: session?.id,
      sessionName: session?.sessionName ?? '',
      subtitle: session?.subtitle ?? '',
      description: session?.description ?? '',
      uploadedThumbnail: f,
      published: session?.published,
      date: d,
    });
  }

  constructor(
    private sessionService: SessionService,
    private messageService: MessageService,
  ) {}

  onSubmit() {
    if (!this.sessionForm.valid) return;
    this.visible.update(() => false);
    if (this.sessionForm.valid) {
      const formData = new FormData();
      const pubVal = this.sessionForm.value.published ?? false;
      console.log(pubVal);
      Object.entries(this.sessionForm.getRawValue()).forEach(([key, value]) => {
        if (key === 'uploadedThumbnail' && value instanceof File) {
          formData.append('uploadedThumbnail', value);
        }
        if (key === 'date' && value instanceof Date) {
          const dateStr = value.toISOString();
          formData.append('date', dateStr);
        } else if (value !== null && value !== undefined) {
          formData.append(key, value.toString());
        }
      });
      formData.append('published', String(pubVal));
      this.sessionService.createOrUpdateSession(
        formData,
        this.inputSession?.id,
      );
      this.inputSession = undefined;
    }
  }

  onUpload(event: { files: File[] }) {
    const file = event.files?.slice(-1)[0];
    if (file) {
      this.uploadedFile = file;
      this.sessionForm.get('uploadedThumbnail')?.setValue(file);
      this.sessionForm.get('uploadedThumbnail')?.markAsDirty();
      this.sessionForm.get('uploadedThumbnail')?.markAsTouched();
      this.filename = file.name;
      this.messageService.add({
        severity: 'info',
        summary: 'File Uploaded',
        detail: '',
      });
    }
  }
}
