import { Component } from '@angular/core';

@Component({
  selector: 'app-feed',
  imports: [],
  templateUrl: './feed.component.html',
  styleUrl: './feed.component.css',
})
export class FeedComponent {
  posts = [
    {
      id: 1,
      username: 'JohnDoe',
      time: '2 hours ago',
      title: 'My first Angular post!',
      content: 'Just started learning Angular, and I love it! 🚀',
      upvotes: 120,
      downvotes: 3,
      comments: ['Great!', 'Keep going!'],
      image:
        'https://plus.unsplash.com/premium_photo-1681412205359-a803b2649d57?q=80&w=3087&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D',
    },
    {
      id: 2,
      username: 'JaneSmith',
      time: '5 hours ago',
      title: 'Why DaisyUI is amazing',
      content: 'DaisyUI makes TailwindCSS even better. What do you guys think?',
      upvotes: 250,
      downvotes: 10,
      comments: ['Agreed!', 'Not a fan.'],
      image: '',
    },
  ];

  upvote(post: any) {
    post.upvotes++;
  }

  downvote(post: any) {
    post.downvotes++;
  }
}
