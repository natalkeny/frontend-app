// types.ts
export enum Role {
  ADMIN = 'admin',
  MODERATOR = 'moderator',
  USER = 'user',
}

export enum SortOrder {
  ASC = 'asc',
  DESC = 'desc',
}

export enum PaginationOptions {
  LIMIT = 'limit',
  OFFSET = 'offset',
}

export type User = {
  id: number;
  username: string;
  email: string;
  role: Role;
  createdAt: Date;
  updatedAt: Date;
};