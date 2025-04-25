import { Routes } from '@angular/router';
import { LoginComponent } from './pages/login/login.component';
import { SignUpComponent } from './pages/sign-up/sign-up.component';
import { HomeComponent } from './pages/home/home.component';
import { authGuard } from './auth/auth.guard';
import { FeedComponent } from './component/feed/feed.component';
import { CreatePostComponent } from './component/create-post/create-post.component';

export const routes: Routes = [
  {
    path: '',
    component: HomeComponent,
    canActivate: [authGuard],
    children: [
      { path: '', component: FeedComponent },
      { path: 'create-post', component: CreatePostComponent },
      { path: 'edit-post/:id', component: CreatePostComponent },
    ],
  },
  { path: 'login', component: LoginComponent },
  { path: 'sign-up', component: SignUpComponent },
  { path: '**', redirectTo: '' },
];
