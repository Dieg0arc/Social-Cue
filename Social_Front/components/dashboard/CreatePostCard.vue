<!-- 
  Archivo: CreatePostCard.vue
  Ubicación: components/dashboard/CreatePostCard.vue
  Descripción: Componente que muestra la tarjeta para crear un nuevo post.
-->

<template>
  <div class="create-post-card">
    <div class="create-post-header">
      <img 
        :src="currentUser.avatar" 
        :alt="currentUser.username" 
        class="user-avatar"
      />
      <input 
        type="text" 
        placeholder="¿Qué estás pensando?" 
        class="create-post-input"
        @click="$emit('open-modal')"
      />
    </div>
    <div class="create-post-actions">
      <button class="create-action-btn" @click="$emit('open-modal', 'photo')">
        <span class="action-icon">📷</span>
        <span class="action-text">Foto</span>
      </button>
      <button class="create-action-btn" @click="$emit('open-modal', 'video')">
        <span class="action-icon">🎬</span>
        <span class="action-text">Video</span>
      </button>
      <button class="create-action-btn" @click="$emit('open-modal', 'location')">
        <span class="action-icon">📍</span>
        <span class="action-text">Ubicación</span>
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { CurrentUser } from '~/types/dashboard';

/**
 * Props del componente
 */
const props = defineProps<{
  /**
   * Información del usuario actual
   */
  currentUser: CurrentUser;
}>();

/**
 * Eventos emitidos por el componente
 */
const emit = defineEmits<{
  /**
   * Emitido cuando el usuario quiere abrir el modal de creación de post
   * @param type Tipo de contenido a crear (opcional: 'photo', 'video', 'location')
   */
  (e: 'open-modal', type?: string): void;
}>();
</script>

<style scoped>
/* Tarjeta de creación de post */
.create-post-card {
  background-color: #122237;
  border-radius: 8px;
  border: 1px solid #243447;
  padding: 1rem;
  margin-bottom: 1rem;
}

.create-post-header {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  margin-bottom: 1rem;
}

.user-avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  object-fit: cover;
}

.create-post-input {
  flex: 1;
  background-color: #243447;
  border: none;
  border-radius: 20px;
  padding: 0.75rem 1rem;
  color: #FFFFFF;
  cursor: pointer;
}

.create-post-input::placeholder {
  color: #8A97A8;
}

.create-post-actions {
  display: flex;
  justify-content: space-around;
  border-top: 1px solid #243447;
  padding-top: 1rem;
}

.create-action-btn {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  background: transparent;
  border: none;
  color: #D0D0D0;
  cursor: pointer;
  padding: 0.5rem;
  transition: color 0.2s;
}

.create-action-btn:hover {
  color: #FFFFFF;
}

.action-icon {
  font-size: 1.25rem;
}

/* Responsive */
@media (max-width: 768px) {
  .create-action-btn .action-text {
    display: none;
  }
}
</style>

