import { HttpClient, HttpHandler } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Post } from '../models/post.model';

@Injectable({
  providedIn: 'root',
})
export class PostService {
  private apiUrl = 'http://localhost:3000';
  token: string = '';

  constructor(private http: HttpClient) {
    this.token = localStorage.getItem('token') || '';
  }
  createPost({
    title,
    content,
    image,
  }: {
    title: string;
    content: string;
    image: File | null;
  }) {
    const formData = new FormData();
    formData.append('title', title);
    formData.append('content', content);
    if (image) {
      formData.append('image', image, image.name);
    }

    const headers = { Authorization: `Bearer ${this.token}` };
    return this.http.post(`${this.apiUrl}/posts`, formData, { headers });
  }

  getPosts() {
    return this.http.get(`${this.apiUrl}/posts`, {
      headers: { Authorization: `Bearer ${this.token}` },
    });
  }

  getPostById(postId: string): Observable<Post> {
    return this.http.get<Post>(`${this.apiUrl}/posts/${postId}`, {
      headers: { Authorization: `Bearer ${this.token}` },
    });
  }
}
