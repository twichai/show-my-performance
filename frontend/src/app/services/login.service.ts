import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { Token } from '../model/token.model';

@Injectable({
  providedIn: 'root',
})
export class LoginService {
  private apiUrl = 'http://localhost:3000';

  private http = inject(HttpClient);

  login({
    email,
    password,
  }: {
    email: string;
    password: string;
  }): Observable<Token> {
    const body = { email, password };
    return this.http.post<Token>(`${this.apiUrl}/login`, body);
  }
  signUp({
    username,
    password,
    email,
    firstName,
    lastName,
    phoneNumber,
  }: {
    username: string;
    password: string;
    email: string;
    firstName: string;
    lastName: string;
    phoneNumber: string;
  }): Observable<Token> {
    const body = {
      username,
      password,
      email,
      firstName,
      lastName,
      phoneNumber,
    };
    return this.http.post<Token>(`${this.apiUrl}/signup`, body);
  }
}
