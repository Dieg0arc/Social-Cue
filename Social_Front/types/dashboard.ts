/**
 * Archivo: dashboard.ts
 * Ubicación: types/dashboard.ts
 * Descripción: Definición de tipos para el dashboard de SocialCue.
 */

/**
 * Representa un autor de un post o historia
 */
export interface Author {
  id: number;
  username: string;
  avatar: string;
}

/**
 * Representa un comentario en un post
 */
export interface Comment {
  username: string;
  text: string;
}

/**
 * Estructura base para todos los tipos de posts
 */
export interface BasePost {
  id: number;
  author: Author;
  likes: number;
  liked: boolean;
  comments: Comment[];
  createdAt: string;
}

/**
 * Post que solo contiene texto
 */
export interface TextPost extends BasePost {
  textContent: string;
}

/**
 * Post que contiene contenido multimedia (imagen o video)
 */
export interface MediaPost extends BasePost {
  mediaUrl: string;
  caption: string;
  location?: string;
}

/**
 * Tipo unión que representa cualquier tipo de post
 */
export type Post = TextPost | MediaPost;

/**
 * Representa una historia de usuario
 */
export interface Story {
  id: number;
  username: string;
  avatar: string;
  hasNew: boolean;
}

/**
 * Representa un usuario sugerido para seguir
 */
export interface SuggestionUser {
  id: number;
  username: string;
  fullName: string;
  avatar: string;
  meta: string;
}

/**
 * Representa el usuario actual
 */
export interface CurrentUser {
  id: number;
  username: string;
  fullName: string;
  avatar: string;
}

