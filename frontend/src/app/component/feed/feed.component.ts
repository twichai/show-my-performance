import { Component, Inject } from '@angular/core';
import { PostService } from '../../services/post.service';
import { Post } from '../../models/post.model';

@Component({
  selector: 'app-feed',
  imports: [],
  templateUrl: './feed.component.html',
  styleUrl: './feed.component.css',
})
export class FeedComponent {
  constructor(@Inject(PostService) private postsService: PostService) {}
  posts: Post[] = [];

  ngOnInit() {
    this.postsService.getPosts().subscribe((res) => {
      this.posts = res as Post[];
    });
  }

  upvote(post: any) {
    post.upvotes++;
  }

  downvote(post: any) {
    post.downvotes++;
  }
}
