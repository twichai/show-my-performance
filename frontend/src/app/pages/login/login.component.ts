import { Component, inject } from '@angular/core';
import {
  FormGroup,
  FormControl,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { LoginService } from '../../services/login.service';
import { Token } from '../../models/token.model';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  imports: [ReactiveFormsModule],
  templateUrl: './login.component.html',
  styleUrl: './login.component.css',
})
export class LoginComponent {
  private loginService = inject(LoginService);
  private route = inject(Router);
  loginForm = new FormGroup({
    email: new FormControl('johndoe@example.com', [
      Validators.required,
      Validators.email,
    ]),
    password: new FormControl('securepassword123', [
      Validators.required,
      Validators.minLength(6),
    ]),
  });
  errorMessage: string = '';

  onSubmit() {
    if (this.loginForm.valid) {
      this.loginService
        .login({
          email: this.loginForm.value.email!,
          password: this.loginForm.value.password!,
        })
        .subscribe(
          (response: Token) => {
            if (response.token) {
              localStorage.setItem('token', response.token);
              localStorage.setItem('user', JSON.stringify(response.user));
              this.errorMessage = '';
              this.route.navigate(['/']);
              // Redirect to the dashboard or another page
            } else {
              this.errorMessage = 'Invalid email or password';
            }
          },
          (err) => {
            this.errorMessage = err.error.message;
            console.log(this.errorMessage);
          }
        );
    }
  }
}
