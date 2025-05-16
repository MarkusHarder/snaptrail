import { Component, OnDestroy } from '@angular/core';
import {
  AbstractControl,
  FormControl,
  FormGroup,
  ReactiveFormsModule,
  ValidationErrors,
  Validators,
} from '@angular/forms';
import { ButtonModule } from 'primeng/button';
import { FloatLabel } from 'primeng/floatlabel';
import { InputTextModule } from 'primeng/inputtext';
import { Message } from 'primeng/message';
import { Password } from 'primeng/password';
import { Subscription } from 'rxjs';
import { AdminService } from '../../services/admin.service';
import { AuthService } from '../../services/auth.service';
import { DividerModule } from 'primeng/divider';

@Component({
  selector: 'app-admin-user-page',
  imports: [
    ReactiveFormsModule,
    ButtonModule,
    InputTextModule,
    Password,
    FloatLabel,
    Message,
    DividerModule,
  ],
  templateUrl: './admin-user-page.component.html',
  styleUrl: './admin-user-page.component.css',
})
export class AdminUserPageComponent implements OnDestroy {
  username$: Subscription;
  username = '';

  constructor(
    private authService: AuthService,
    private adminService: AdminService,
  ) {
    this.username$ = this.authService.username$.subscribe((username) => {
      this.username = username;
    });
  }

  form = new FormGroup(
    {
      oldPassword: new FormControl('', [Validators.required]),
      newPassword: new FormControl('', [
        Validators.required,
        Validators.minLength(8),
        this.strongPasswordValidator,
      ]),
      confirmPassword: new FormControl('', [
        Validators.required,
        Validators.minLength(8),
        this.strongPasswordValidator,
      ]),
    },
    {
      validators: this.passwordsMatchValidator,
    },
  );

  passwordMatchValidator(formGroup: FormGroup) {
    const newPassword = formGroup.get('newPassword')?.value;
    const confirmPassword = formGroup.get('confirmPassword')?.value;
    return newPassword === confirmPassword ? null : { passwordsMismatch: true };
  }

  strongPasswordValidator(control: AbstractControl): ValidationErrors | null {
    const value = control.value;
    const hasNumber = /\d/.test(value);
    const hasUpper = /[A-Z]/.test(value);
    const hasSpecial = /[!@#$%^&*]/.test(value);
    const valid = hasNumber && hasUpper && hasSpecial;
    return valid ? null : { weakPassword: true };
  }

  passwordsMatchValidator(group: AbstractControl): ValidationErrors | null {
    const newPassword = group.get('newPassword')?.value;
    const confirmPassword = group.get('confirmPassword')?.value;
    return newPassword === confirmPassword ? null : { passwordsMismatch: true };
  }

  hasError(controlName: string, error: string): boolean {
    const control = this.form.get(controlName);
    return !!(control?.hasError(error) && control?.touched);
  }

  onSubmit() {
    if (this.form.valid) {
      const pwChange = {
        username: this.username!,
        oldPassword: this.form.value.oldPassword!,
        newPassword: this.form.value.confirmPassword!,
      };
      this.adminService.changePassowrd(pwChange);
    }
  }

  ngOnDestroy(): void {
    this.username$.unsubscribe();
  }
}
