/**
 * services/posts.ts
 * 
 * Service for handling all social post related API calls
 * This file contains the interfaces and methods needed to interact with the
 * posts endpoints of the API, including fetching posts, creating posts,
 * liking/unliking posts, and managing comments.
 */

import { useAuthStore } from '~/stores/auth';
import { User } from '~/types/user';

/**
 * Interface for comment data
 */
export interface Comment {
  id: string;
  postId: string;
  author: User;
  content: string;
  createdAt: string;
  updatedAt?: string;
}

/**
 * Interface for post data
 */
export interface Post {
  id: string;
  author: User;
  content: string;
  likesCount: number;
  commentsCount: number;
  isLiked: boolean;
  createdAt: string;
  updatedAt?: string;
  media?: string[];
  comments?: Comment[];
}

/**
 * Interface for creating a new post
 */
export interface CreatePostRequest {
  content: string;
  media?: string[];
}

/**
 * Interface for creating a new comment
 */
export interface CreateCommentRequest {
  content: string;
}

/**
 * Interface for API pagination
 */
export interface Pagination {
  page: number;
  limit: number;
  total: number;
  totalPages: number;
}

/**
 * Interface for post response from API
 */
export interface PostResponse {
  success: boolean;
  data: Post;
}

/**
 * Interface for posts list response from API
 */
export interface PostsResponse {
  success: boolean;
  data: Post[];
  pagination?: Pagination;
}

/**
 * Interface for comments list response from API
 */
export interface CommentsResponse {
  success: boolean;
  data: Comment[];
  pagination?: Pagination;
}

/**
 * Interface for error response from API
 */
interface ApiError {
  message: string;
  code?: string;
  errors?: Record<string, string[]>;
}

/**
 * Helper to get API URL from runtime config
 */
const getApiBaseUrl = () => {
  const config = useRuntimeConfig();
  return config.public.apiBaseUrl || 'http://localhost:3000';
};

/**
 * Helper to handle API errors
 */
const handleApiError = (error: any): never => {
  const errorMessage = error?.data?.message || 'Error desconocido';
  console.error('API Error:', error);
  throw new Error(errorMessage);
};

/**
 * Posts service with methods for interacting with the API
 */
export const postsService = {
  /**
   * Get posts feed - returns all posts or posts by user if userId is provided
   * 
   * @param page Page number for pagination
   * @param limit Number of posts per page
   * @param userId Optional user ID to filter posts by user
   * @returns Promise with posts data
   */
  async getFeed(page = 1, limit = 10, userId?: string): Promise<PostsResponse> {
    const authStore = useAuthStore();
    const apiBaseUrl = getApiBaseUrl();
    
    try {
      // Build query params
      const queryParams = new URLSearchParams();
      queryParams.append('page', page.toString());
      queryParams.append('limit', limit.toString());
      if (userId) {
        queryParams.append('userId', userId);
      }
      
      const { data, error } = await useFetch<PostsResponse>(`${apiBaseUrl}/api/posts?${queryParams.toString()}`, {
        method: 'GET',
        headers: {
          ...authStore.getAuthHeader(),
          'Content-Type': 'application/json',
        },
      });
      
      if (error.value) {
        return handleApiError(error.value);
      }
      
      return data.value as PostsResponse;
    } catch (error) {
      return handleApiError(error);
    }
  },
  
  /**
   * Create a new post
   * 
   * @param postData Post content and optional media
   * @returns Promise with created post data
   */
  async createPost(postData: CreatePostRequest): Promise<PostResponse> {
    const authStore = useAuthStore();
    const apiBaseUrl = getApiBaseUrl();
    
    try {
      const { data, error } = await useFetch<PostResponse>(`${apiBaseUrl}/api/posts`, {
        method: 'POST',
        body: postData,
        headers: {
          ...authStore.getAuthHeader(),
          'Content-Type': 'application/json',
        },
      });
      
      if (error.value) {
        return handleApiError(error.value);
      }
      
      return data.value as PostResponse;
    } catch (error) {
      return handleApiError(error);
    }
  },
  
  /**
   * Like a post
   * 
   * @param postId ID of the post to like
   * @returns Promise with updated post data
   */
  async likePost(postId: string): Promise<PostResponse> {
    const authStore = useAuthStore();
    const apiBaseUrl = getApiBaseUrl();
    
    try {
      const { data, error } = await useFetch<PostResponse>(`${apiBaseUrl}/api/posts/${postId}/like`, {
        method: 'POST',
        headers: {
          ...authStore.getAuthHeader(),
          'Content-Type': 'application/json',
        },
      });
      
      if (error.value) {
        return handleApiError(error.value);
      }
      
      return data.value as PostResponse;
    } catch (error) {
      return handleApiError(error);
    }
  },
  
  /**
   * Unlike a post
   * 
   * @param postId ID of the post to unlike
   * @returns Promise with updated post data
   */
  async unlikePost(postId: string): Promise<PostResponse> {
    const authStore = useAuthStore();
    const apiBaseUrl = getApiBaseUrl();
    
    try {
      const { data, error } = await useFetch<PostResponse>(`${apiBaseUrl}/api/posts/${postId}/unlike`, {
        method: 'POST',
        headers: {
          ...authStore.getAuthHeader(),
          'Content-Type': 'application/json',
        },
      });
      
      if (error.value) {
        return handleApiError(error.value);
      }
      
      return data.value as PostResponse;
    } catch (error) {
      return handleApiError(error);
    }
  },
  
  /**
   * Create a new comment on a post
   * 
   * @param postId ID of the post to comment on
   * @param commentData Comment content
   * @returns Promise with created comment data
   */
  async createComment(postId: string, commentData: CreateCommentRequest): Promise<Comment> {
    const authStore = useAuthStore();
    const apiBaseUrl = getApiBaseUrl();
    
    try {
      const { data, error } = await useFetch<{ success: boolean, data: Comment }>(`${apiBaseUrl}/api/posts/${postId}/comments`, {
        method: 'POST',
        body: commentData,
        headers: {
          ...authStore.getAuthHeader(),
          'Content-Type': 'application/json',
        },
      });
      
      if (error.value) {
        return handleApiError(error.value);
      }
      
      return data.value?.data as Comment;
    } catch (error) {
      return handleApiError(error);
    }
  },
  
  /**
   * Get comments for a post
   * 
   * @param postId ID of the post to get comments for
   * @param page Page number for pagination
   * @param limit Number of comments per page
   * @returns Promise with comments data
   */
  async getComments(postId: string, page = 1, limit = 10): Promise<CommentsResponse> {
    const authStore = useAuthStore();
    const apiBaseUrl = getApiBaseUrl();
    
    try {
      // Build query params
      const queryParams = new URLSearchParams();
      queryParams.append('page', page.toString());
      queryParams.append('limit', limit.toString());
      
      const { data, error } = await useFetch<CommentsResponse>(`${apiBaseUrl}/api/posts/${postId}/comments?${queryParams.toString()}`, {
        method: 'GET',
        headers: {
          ...authStore.getAuthHeader(),
          'Content-Type': 'application/json',
        },
      });
      
      if (error.value) {
        return handleApiError(error.value);
      }
      
      return data.value as CommentsResponse;
    } catch (error) {
      return handleApiError(error);
    }
  },
  
  /**
   * Get a single post by ID
   * 
   * @param postId ID of the post to retrieve
   * @returns Promise with post data
   */
  async getPost(postId: string): Promise<PostResponse> {
    const authStore = useAuthStore();
    const apiBaseUrl = getApiBaseUrl();
    
    try {
      const { data, error } = await useFetch<PostResponse>(`${apiBaseUrl}/api/posts/${postId}`, {
        method: 'GET',
        headers: {
          ...authStore.getAuthHeader(),
          'Content-Type': 'application/json',
        },
      });
      
      if (error.value) {
        return handleApiError(error.value);
      }
      
      return data.value as PostResponse;
    } catch (error) {
      return handleApiError(error);
    }
  },
  
  /**
   * Delete a post
   * 
   * @param postId ID of the post to delete
   * @returns Promise with success status
   */
  async deletePost(postId: string): Promise<{ success: boolean }> {
    const authStore = useAuthStore();
    const apiBaseUrl = getApiBaseUrl();
    
    try {
      const { data, error } = await useFetch<{ success: boolean }>(`${apiBaseUrl}/api/posts/${postId}`, {
        method: 'DELETE',
        headers: {
          ...authStore.getAuthHeader(),
          'Content-Type': 'application/json',
        },
      });
      
      if (error.value) {
        return handleApiError(error.value);
      }
      
      return data.value as { success: boolean };
    } catch (error) {
      return handleApiError(error);
    }
  }
};

