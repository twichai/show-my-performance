import { User } from './user.model';

export interface Post {
  id: number;
  title: string;
  content: string;
  community: string;
  createdAt: string;
  updatedAt: string;
  imageUrl?: string;
  userId: number;
  user?: User;
}
