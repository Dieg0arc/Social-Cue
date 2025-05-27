<!-- 
  Archivo: StorySection.vue
  Ubicación: components/dashboard/StorySection.vue
  Descripción: Componente que muestra la sección de historias de usuarios.
-->

<template>
  <div class="stories-container">
    <div class="stories-header">
      <h3>Historias</h3>
      <button class="view-all-btn" @click="$emit('view-all')">Ver todas</button>
    </div>
    <div class="stories-list">
      <div 
        v-for="story in stories" 
        :key="story.id" 
        class="story-item"
        :class="{ 'has-new': story.hasNew }"
        @click="$emit('view-story', story.id)"
      >
        <div class="story-avatar-container">
          <img :src="story.avatar" :alt="story.username" class="story-avatar" />
        </div>
        <span class="story-username">{{ story.username }}</span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Story } from '~/types/dashboard';

/**
 * Props del componente
 */
const props = defineProps<{
  /**
   * Lista de historias a mostrar
   */
  stories: Story[];
}>();

/**
 * Eventos emitidos por el componente
 */
const emit = defineEmits<{
  /**
   * Emitido cuando el usuario quiere ver todas las historias
   */
  (e: 'view-all'): void;

  /**
   * Emitido cuando el usuario hace clic en una historia específica
   * @param storyId ID de la historia seleccionada
   */
  (e: 'view-story', storyId: number): void;
}>();
</script>

<style scoped>
/* Contenedor de historias */
.stories-container {
  background-color: #122237;
  border-radius: 8px;
  border: 1px solid #243447;
  padding: 1rem;
  margin-bottom: 1rem;
  overflow: hidden;
}

.stories-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1rem;
}

.stories-header h3 {
  margin: 0;
  font-size: 1rem;
  font-weight: 600;
}

.view-all-btn {
  background: transparent;
  border: none;
  color: #4D8CD9;
  cursor: pointer;
  font-size: 0.875rem;
}

.stories-list {
  display: flex;
  gap: 1rem;
  overflow-x: auto;
  padding-bottom: 0.5rem;
}

.stories-list::-webkit-scrollbar {
  height: 4px;
}

.stories-list::-webkit-scrollbar-track {
  background: #0A1625;
}

.stories-list::-webkit-scrollbar-thumb {
  background: #243447;
  border-radius: 4px;
}

.story-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  cursor: pointer;
  min-width: 65px;
}

.story-avatar-container {
  width: 56px;
  height: 56px;
  border-radius: 50%;
  padding: 2px;
  background: linear-gradient(45deg, #FCAF45, #FD1D1D, #833AB4);
  margin-bottom: 0.5rem;
}

.story-item:not(.has-new) .story-avatar-container {
  background: #243447;
}

.story-avatar {
  width: 100%;
  height: 100%;
  border-radius: 50%;
  object-fit: cover;
  border: 2px solid #0A1625;
}

.story-username {
  font-size: 0.75rem;
  color: #D0D0D0;
  max-width: 65px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  text-align: center;
}
</style>

