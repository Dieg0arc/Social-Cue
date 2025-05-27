<!-- 
  Archivo: PostCard.vue
  Ubicación: components/content/PostCard.vue
  Descripción: Componente que muestra una publicación individual en el feed de SocialCue.
  Este componente incluye la cabecera con información del autor, el contenido del post,
  botones de interacción, contador de likes y sección de comentarios.
-->

<template>
  <article class="post-card">
    <!-- Cabecera del post con información del autor -->
    <header class="post-header">
      <div class="user-info">
        <img 
          :src="post.author.avatar || 'assets\images\Usuari.png'" 
          :alt="post.author.username" 
          class="user-avatar"
        />
        <div class="user-details">
          <h3 class="username">{{ post.author.username }}</h3>
          <span class="location" v-if="isMediaPost && mediaPost?.location">{{ mediaPost?.location }}</span>
        </div>
      </div>
      <button class="more-options-btn">⋮</button>
    </header>

    <!-- Contenido principal del post (imagen o texto) -->
    <div class="post-content">
      <img 
        v-if="isMediaPost && mediaPost?.mediaUrl" 
        :src="mediaPost?.mediaUrl" 
        :alt="mediaPost?.caption || ''" 
        class="post-image" 
        loading="lazy"
        @dblclick="handleLike"
      />
      <div v-else-if="isTextPost && textPost?.textContent" class="post-text">
        {{ textPost.textContent }}
      </div>
    </div>

    <!-- Barra de acciones (like, comentar, etc.) -->
    <div class="post-actions">
      <div class="action-buttons">
        <button 
          class="action-btn like-btn" 
          :class="{ 'liked': isLiked }"
          @click="handleLike"
        >
          <span v-if="isLiked">❤️</span>
          <span v-else>🤍</span>
        </button>
        <button class="action-btn comment-btn" @click="focusCommentInput">
          💬
        </button>
        <button class="action-btn share-btn">
          📤
        </button>
      </div>
      <button class="action-btn save-btn">
        🔖
      </button>
    </div>

    <!-- Información de likes -->
    <div class="post-likes">
      <p>
        <strong>{{ likesCount }} me gusta{{ likesCount !== 1 ? 's' : '' }}</strong>
      </p>
    </div>

    <!-- Descripción del post -->
    <div class="post-caption" v-if="isMediaPost && mediaPost?.caption">
      <p>
        <strong class="username">{{ post.author.username }}</strong>
        {{ mediaPost?.caption }}
      </p>
    </div>

    <!-- Sección de comentarios -->
    <div class="post-comments" v-if="post.comments && post.comments.length > 0">
      <button 
        v-if="post.comments.length > 2 && !showAllComments" 
        class="view-comments-btn"
        @click="showAllComments = true"
      >
        Ver los {{ post.comments.length }} comentarios
      </button>
      <ul class="comments-list">
        <li 
          v-for="(comment, index) in displayedComments" 
          :key="index"
          class="comment-item"
        >
          <strong class="username">{{ comment.username }}</strong>
          {{ comment.text }}
        </li>
      </ul>
    </div>

    <!-- Timestamp del post -->
    <div class="post-timestamp">
      <time :datetime="post.createdAt">{{ formattedTime }}</time>
    </div>

    <!-- Campo para añadir comentario -->
    <div class="add-comment">
      <input 
        type="text" 
        placeholder="Añade un comentario..." 
        v-model="newComment"
        @keyup.enter="addComment"
        ref="commentInput"
        class="comment-input"
      />
      <button 
        class="post-comment-btn"
        :disabled="!newComment.trim()"
        @click="addComment"
      >
        Publicar
      </button>
    </div>
  </article>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue';
import type { Post, MediaPost, TextPost } from '~/types/dashboard';

/**
 * Props del componente
 */
const props = defineProps<{
  /**
   * Post a mostrar
   */
  post: Post;
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
}>();

/**
 * Estado local
 */
const newComment = ref('');
const commentInput = ref<HTMLInputElement | null>(null);
const isLiked = ref(props.post.liked || false);
const likesCount = ref(props.post.likes || 0);
const showAllComments = ref(false);

/**
 * Determina si el post es un post de tipo media
 */
const isMediaPost = computed(() => {
  return 'mediaUrl' in props.post;
});

/**
 * Determina si el post es un post de tipo texto
 */
const isTextPost = computed(() => {
  return 'textContent' in props.post;
});

/**
 * Obtiene el post como MediaPost si corresponde a ese tipo
 */
const mediaPost = computed(() => {
  return isMediaPost.value ? props.post as MediaPost : null;
});

/**
 * Obtiene el post como TextPost si corresponde a ese tipo
 */
const textPost = computed(() => {
  return isTextPost.value ? props.post as TextPost : null;
});

/**
 * Comentarios a mostrar en la UI
 */
const displayedComments = computed(() => {
  if (showAllComments.value || props.post.comments.length <= 2) {
    return props.post.comments;
  } else {
    return props.post.comments.slice(0, 2);
  }
});

/**
 * Formatea la fecha del post para mostrarla de manera amigable
 */
const formattedTime = computed(() => {
  const date = new Date(props.post.createdAt);
  const now = new Date();
  const diffTime = Math.abs(now.getTime() - date.getTime());
  const diffDays = Math.floor(diffTime / (1000 * 60 * 60 * 24));
  
  if (diffDays === 0) {
    const diffHours = Math.floor(diffTime / (1000 * 60 * 60));
    if (diffHours === 0) {
      const diffMinutes = Math.floor(diffTime / (1000 * 60));
      return `HACE ${diffMinutes} MINUTOS`;
    }
    return `HACE ${diffHours} HORAS`;
  } else if (diffDays === 1) {
    return 'AYER';
  } else if (diffDays < 7) {
    return `HACE ${diffDays} DÍAS`;
  } else {
    return date.toLocaleDateString('es-ES', {
      day: 'numeric',
      month: 'long'
    }).toUpperCase();
  }
});

/**
 * Maneja la acción de dar like a un post
 */
const handleLike = () => {
  isLiked.value = !isLiked.value;
  likesCount.value = isLiked.value 
    ? likesCount.value + 1 
    : likesCount.value - 1;
  
  emit('like', props.post.id);
};

/**
 * Coloca el foco en el campo de comentario
 */
const focusCommentInput = () => {
  if (commentInput.value) {
    commentInput.value.focus();
  }
};

/**
 * Añade un nuevo comentario
 */
function addComment() {
  if (!newComment.value.trim()) return;
  
  emit('comment', props.post.id, newComment.value);
  newComment.value = '';
}
</script>

<style scoped>
/* Estilos del componente PostCard */
.post-card {
  background-color: #0A1625;
  border-radius: 8px;
  border: 1px solid #243447;
  margin-bottom: 1.5rem;
  overflow: hidden;
  max-width: 600px;
  width: 100%;
}

/* Cabecera del post */
.post-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0.75rem 1rem;
  border-bottom: 1px solid #1d2f46;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 0.75rem;
}

.user-avatar {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  object-fit: cover;
}

.user-details {
  display: flex;
  flex-direction: column;
}

.username {
  font-weight: 600;
  font-size: 0.875rem;
  color: #FFFFFF;
  margin: 0;
}

.location {
  font-size: 0.75rem;
  color: #8A97A8;
}

.more-options-btn {
  background: transparent;
  border: none;
  color: #D0D0D0;
  font-size: 1.25rem;
  cursor: pointer;
  padding: 0.25rem;
}

/* Contenido del post */
.post-content {
  width: 100%;
  background-color: #0A1625;
}

.post-image {
  width: 100%;
  max-height: 600px;
  object-fit: contain;
  display: block;
}

.post-text {
  padding: 1rem;
  font-size: 1rem;
  line-height: 1.5;
  color: #FFFFFF;
}

/* Barra de acciones */
.post-actions {
  padding: 0.75rem 1rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.action-buttons {
  display: flex;
  gap: 1rem;
}

.action-btn {
  background: transparent;
  border: none;
  color: #D0D0D0;
  font-size: 1.25rem;
  cursor: pointer;
  padding: 0.25rem;
  transition: transform 0.2s;
}

.action-btn:hover {
  transform: scale(1.1);
}

.like-btn.liked {
  color: #e74c3c;
  animation: likeAnimation 0.3s ease-in-out;
}

@keyframes likeAnimation {
  0%, 100% { transform: scale(1); }
  50% { transform: scale(1.2); }
}

/* Información de likes */
.post-likes {
  padding: 0 1rem;
}

.post-likes p {
  margin: 0.25rem 0;
  font-size: 0.875rem;
  color: #FFFFFF;
}

/* Descripción del post */
.post-caption {
  padding: 0 1rem;
  margin-bottom: 0.5rem;
}

.post-caption p {
  margin: 0.25rem 0;
  font-size: 0.875rem;
  color: #D0D0D0;
}

/* Sección de comentarios */
.post-comments {
  padding: 0 1rem;
}

.view-comments-btn {
  background: transparent;
  border: none;
  color: #8A97A8;
  font-size: 0.875rem;
  padding: 0;
  margin: 0.25rem 0;
  cursor: pointer;
  text-align: left;
  width: 100%;
}

.comments-list {
  list-style: none;
  padding: 0;
  margin: 0.5rem 0;
}

.comment-item {
  margin-bottom: 0.5rem;
  font-size: 0.875rem;
  color: #D0D0D0;
}

/* Timestamp */
.post-timestamp {
  padding: 0 1rem;
  margin: 0.5rem 0;
}

.post-timestamp time {
  font-size: 0.75rem;
  color: #8A97A8;
  text-transform: uppercase;
}

/* Campo de añadir comentario */
.add-comment {
  padding: 0.75rem 1rem;
  border-top: 1px solid #1d2f46;
  display: flex;
  align-items: center;
}

.comment-input {
  flex: 1;
  background: transparent;
  border: none;
  color: #FFFFFF;
  padding: 0.5rem 0;
  outline: none;
  font-size: 0.875rem;
}

.comment-input::placeholder {
  color: #8A97A8;
}

.post-comment-btn {
  background: transparent;
  border: none;
  color: #4D8CD9;
  font-weight: 600;
  cursor: pointer;
  font-size: 0.875rem;
  padding: 0.5rem;
}

.post-comment-btn:disabled {
  color: #8A97A8;
  cursor: default;
}

/* Responsive */
@media (max-width: 640px) {
  .post-card {
    border-radius: 0;
    border-left: none;
    border-right: none;
    margin-left: -1rem;
    margin-right: -1rem;
    width: calc(100% + 2rem);
  }
}
</style>

