import { User } from './user.model';

export interface Post {
  ID: number;
  title: string;
  content: string;
  community: string;
  createdAt: string;
  updatedAt: string;
  imageUrl?: string;
  upvotes?: number;
  downvotes?: number;
  UserId: number;
  user?: User;
}
