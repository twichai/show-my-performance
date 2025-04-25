import { Component, Inject } from '@angular/core';
import { PostService } from '../../services/post.service';
import { Post } from '../../models/post.model';
import { PostStore } from '../../store/post.store';
import { Router } from '@angular/router';

@Component({
  selector: 'app-feed',
  imports: [],
  templateUrl: './feed.component.html',
  styleUrl: './feed.component.css',
})
export class FeedComponent {

  constructor(@Inject(PostService) private postsService: PostService, public store: PostStore, private router: Router ) {}
  posts: Post[] = [];

  ngOnInit() {
    this.postsService.getPosts().subscribe((res) => {
      const post: Post[] = res as Post[];

      this.store.loadPosts(post);
      console.log(this.store.posts());
    });
  }

  deletePost(_t2: Post) {
  throw new Error('Method not implemented.');
  }

  editPost(postId: number) {
    console.log('Editing post with ID:', postId);
    this.router.navigate(['/create-post', postId]);
    this.router.navigate(['/create-post']);
  }

  upvote(post: any) {
    post.upvotes++;
  }

  downvote(post: any) {
    post.downvotes++;
  }
}
