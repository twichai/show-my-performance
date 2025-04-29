import { Component, Inject } from '@angular/core';
import { PostService } from '../../services/post.service';
import { Post } from '../../models/post.model';
import { PostStore } from '../../store/post.store';
import { Router } from '@angular/router';
import { User } from '../../models/user.model';

@Component({
  selector: 'app-feed',
  imports: [],
  templateUrl: './feed.component.html',
  styleUrl: './feed.component.css',
})
export class FeedComponent {

  constructor(@Inject(PostService) private postsService: PostService, public store: PostStore, private router: Router) { }
  posts: Post[] = [];
  user: User = JSON.parse(localStorage.getItem('user') || '{}');

  ngOnInit() {
    this.postsService.getPosts().subscribe((res) => {
      const post: Post[] = res as Post[];

      this.store.loadPosts(post);
      console.log(this.store.posts());
    });
  }

  deletePost(Post: Post) {
    this.postsService.deletePost(Post.ID).subscribe((res) => {
      this.store.deletePost(Post.ID);
    })
  }

  editPost(postId: number) {
    console.log('Editing post with ID:', postId);
    this.router.navigate(['/edit-post', postId]);
  }

  upvote(post: any) {
    post.upvotes++;
  }

  downvote(post: any) {
    post.downvotes++;
  }
}
