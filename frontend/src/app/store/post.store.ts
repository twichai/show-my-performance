import { Injectable, signal, computed, effect } from '@angular/core';
import { Post } from '../models/post.model';

@Injectable({ providedIn: 'root' })
export class PostStore {
  // Private signals
  private _posts = signal<Post[]>([]);
  private _selectedPost = signal<Post | null>(null);

  // Public readonly signals
  posts = this._posts.asReadonly();
  selectedPost = this._selectedPost.asReadonly();

  // -- CRUD METHODS --

  // CREATE
  addPost(post: Post) {
    this._posts.update(posts => [...posts, post]);
  }

  // READ (by ID)
  selectPost(id: number) {
    const found = this._posts().find(p => p.ID === id) || null;
    this._selectedPost.set(found);
  }

  // UPDATE
  updatePost(updated: Post) {
    this._posts.update(posts =>
      posts.map(post => post.ID === updated.ID ? updated : post)
    );
    this._selectedPost.set(updated);
  }

  // DELETE
  deletePost(id: number) {
    this._posts.update(posts => posts.filter(post => post.ID !== id));
    if (this._selectedPost()?.ID === id) this._selectedPost.set(null);
  }

  // (Optional) Load posts from backend
  loadPosts(postsFromServer: Post[]) {
    this._posts.set(postsFromServer);
  }
}
