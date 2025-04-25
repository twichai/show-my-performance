import { Post } from './post.model';

export interface User {
  ID: number;
  username: string;
  password: string;
  profileImageUrl: string;
  email: string;
  firstName: string;
  lastName: string;
  phoneNumber: string;
  posts?: Post[];
}
