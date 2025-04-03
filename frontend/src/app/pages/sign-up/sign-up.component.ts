import { Component, inject } from '@angular/core';
import {
  FormControl,
  FormGroup,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { LoginService } from '../../services/login.service';

@Component({
  selector: 'app-sign-up',
  imports: [ReactiveFormsModule],
  templateUrl: './sign-up.component.html',
  styleUrl: './sign-up.component.css',
})
export class SignUpComponent {
  loginService = inject(LoginService);
  signUpForm = new FormGroup(
    {
      username: new FormControl('testuser', [Validators.required]),
      password: new FormControl('password123', [
        Validators.required,
        Validators.minLength(6),
      ]),
      confirmPassword: new FormControl('password123', [
        Validators.required,
        Validators.minLength(6),
      ]),
      email: new FormControl('testuser@example.com', [
        Validators.required,
        Validators.email,
      ]),
      firstName: new FormControl('Test', [Validators.required]),
      lastName: new FormControl('User', [Validators.required]),
      phoneNumber: new FormControl('+1234567890', [
        Validators.required,
        Validators.pattern(/^\+?[1-9]\d{1,14}$/), // E.164 phone number format
      ]),
    },
    {
      validators: (formGroup) => {
        const password = formGroup.get('password')?.value;
        const confirmPassword = formGroup.get('confirmPassword')?.value;
        return password === confirmPassword ? null : { passwordMismatch: true };
      },
    }
  );
  errorMessage: string = '';

  onSubmit() {
    this.loginService
      .signUp({
        username: this.signUpForm.value.username!,
        password: this.signUpForm.value.password!,
        email: this.signUpForm.value.email!,
        firstName: this.signUpForm.value.firstName!,
        lastName: this.signUpForm.value.lastName!,
        phoneNumber: this.signUpForm.value.phoneNumber!,
      })
      .subscribe(
        (response) => {
          if (response.token) {
            localStorage.setItem('token', response.token);
            this.errorMessage = '';
            // Redirect to the dashboard or another page
          } else {
            this.errorMessage = 'Invalid sign-up details';
          }
        },
        (err) => {
          this.errorMessage = err.error.message;
        }
      );
  }
}
