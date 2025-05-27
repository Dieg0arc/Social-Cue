<!-- 
  Archivo: PostsList.vue
  Ubicación: components/dashboard/PostsList.vue
  Descripción: Componente que muestra la lista de posts con scroll infinito.
-->

<template>
  <div>
    <!-- Lista de posts -->
    <div class="posts-list" v-if="posts.length > 0">
      <PostCard 
        v-for="post in posts" 
        :key="post.id" 
        :post="post"
        @like="$emit('like', post.id)"
  @comment="text => handlePostComment(text, post)"
      />
      
      <!-- Indicador de carga para el scroll infinito -->
      <div v-if="loading" class="loading-indicator">
        <div class="loading-spinner"></div>
        <p>Cargando más publicaciones...</p>
      </div>
    </div>
    
    <!-- Estado vacío cuando no hay posts -->
    <div v-else class="empty-state">
      <p>No hay publicaciones para mostrar</p>
      <button class="find-friends-btn" @click="$emit('find-friends')">
        Encuentra amigos para seguir
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Post } from '~/types/dashboard';
import PostCard from '~/components/content/PostCard.vue';

/**
 * Props del componente
 */
const props = defineProps<{
  /**
   * Lista de posts a mostrar
   */
  posts: Post[];

  /**
   * Indica si se están cargando más posts
   */
  loading: boolean;
}>();

/**
 * Eventos emitidos por el componente
 */
const emit = defineEmits<{
  /**
   * Emitido cuando el usuario da like a un post
   * @param postId ID del post
   */
  (e: 'like', postId: number): void;

  /**
   * Emitido cuando el usuario comenta en un post
   * @param postId ID del post
   * @param comment Texto del comentario
   */
  (e: 'comment', postId: number, comment: string): void;

  /**
   * Emitido cuando el usuario quiere encontrar amigos
   */
  (e: 'find-friends'): void;
}>();

/**
 * Maneja el evento de comentario de un post con tipos explícitos
 * @param commentText Texto del comentario
 * @param post Objeto completo del post
 */
function handlePostComment(text: unknown, post: Post): void {
  if (typeof text === 'string') {
    emit('comment', post.id, text);
  }
}
</script>

<style scoped>
/* Lista de posts */
.posts-list {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 1rem;
}

/* Estado vacío */
.empty-state {
  background-color: #122237;
  border-radius: 8px;
  border: 1px solid #243447;
  padding: 2rem;
  text-align: center;
  margin-top: 2rem;
}

.empty-state p {
  margin-bottom: 1rem;
  color: #8A97A8;
}

.find-friends-btn {
  background-color: #0056A6;
  color: white;
  border: none;
  border-radius: 25px;
  padding: 0.75rem 1.5rem;
  font-weight: 600;
  cursor: pointer;
  transition: background-color 0.2s;
}

.find-friends-btn:hover {
  background-color: #4D8CD9;
}

/* Indicador de carga */
.loading-indicator {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin: 2rem 0;
  color: #8A97A8;
}

.loading-spinner {
  width: 30px;
  height: 30px;
  border: 3px solid rgba(77, 140, 217, 0.3);
  border-radius: 50%;
  border-top-color: #4D8CD9;
  animation: spin 1s ease-in-out infinite;
  margin-bottom: 0.5rem;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}
</style>

