import { Component, inject } from '@angular/core';
import { CardModule } from 'primeng/card';
import { InputTextModule } from 'primeng/inputtext';
import { FloatLabel } from 'primeng/floatlabel';
import {
  FormBuilder,
  FormControl,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { ButtonModule } from 'primeng/button';
import { PasswordModule } from 'primeng/password';
import { AdminService } from '../../services/admin.service';
import { User } from '../../models/user';
import { AuthService } from '../../services/auth.service';
import { Observable } from 'rxjs';
import { AsyncPipe } from '@angular/common';

@Component({
  selector: 'app-admin-login',
  imports: [
    CardModule,
    AsyncPipe,
    CardModule,
    ButtonModule,
    InputTextModule,
    FloatLabel,
    ReactiveFormsModule,
    PasswordModule,
  ],
  templateUrl: './admin-login.component.html',
  styleUrl: './admin-login.component.css',
})
export class AdminLoginComponent {
  private formBuilder = inject(FormBuilder);
  loggedIn$: Observable<boolean>;

  adminForm = this.formBuilder.group({
    username: new FormControl('', Validators.required),
    password: new FormControl('', Validators.required),
  });

  constructor(
    private adminService: AdminService,
    private authService: AuthService,
  ) {
    this.loggedIn$ = this.authService.loggedIn$;
  }

  onSubmit() {
    const admin = {
      username: this.adminForm.value.username,
      password: this.adminForm.value.password,
    } as User;
    this.adminService.login(admin);
  }
}
