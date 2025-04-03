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

  ping(): Observable<any> {
    return this.http.get<any>(`${this.apiUrl}/ping`);
  }
}
