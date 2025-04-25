import { Component, inject } from '@angular/core';
import {
  FormBuilder,
  FormControl,
  FormGroup,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { Router } from '@angular/router';
import { PostService } from '../../services/post.service';
import { Post } from '../../models/post.model';
@Component({
  selector: 'app-create-post',
  imports: [ReactiveFormsModule],
  templateUrl: './create-post.component.html',
  styleUrl: './create-post.component.css',
})
export class CreatePostComponent {
  fb = inject(FormBuilder);
  router = inject(Router);
  postService = inject(PostService);

  postForm = new FormGroup({
    community: new FormControl('Nancy', [
      Validators.minLength(2),
      Validators.required,
    ]),
    title: new FormControl('Drew', [
      Validators.minLength(2),
      Validators.required,
    ]),
    content: new FormControl('Drew', [
      Validators.minLength(2),
      Validators.required,
    ]),
  });
  selectedImage: File | null = null;
  communities = ['r/Angular', 'r/WebDev', 'r/Technology', 'r/Programming'];

  ngOnInit() {
    //get post Id from url
    const url = this.router.url;
    const postId = url.split('/').pop();
    if (postId) {
      this.postService.getPostById(postId).subscribe((res: Post) => {
        this.postForm.patchValue({
          title: res.title,
          content: res.content,
          community: res.community,
        });
      });
    }
  }

  onFileSelected(event: any) {
    const file = event.target.files[0];
    if (file) {
      this.selectedImage = file;
    }

  }

  submitPost() {
    if (this.postForm.valid) {
      this.postService
        .createPost({
          title: this.postForm.value.title || '',
          content: this.postForm.value.content || '',
          image: this.selectedImage,
        })
        .subscribe(
          (response) => {
            console.log('Post created successfully:', response);
            this.router.navigate(['/']); // Redirect to feed page
          },
          (err) => {
            console.error('Error creating post:', err);
          }
        );
    }
  }
}
