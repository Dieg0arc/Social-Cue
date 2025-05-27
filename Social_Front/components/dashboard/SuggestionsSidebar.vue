<!-- 
  Archivo: SuggestionsSidebar.vue
  Ubicación: components/dashboard/SuggestionsSidebar.vue
  Descripción: Componente de barra lateral que muestra el perfil del usuario actual
  y sugerencias de personas para seguir. Solo visible en pantallas grandes.
-->

<template>
  <div class="suggestions-sidebar">
    <!-- Perfil del usuario -->
    <div class="user-profile-card">
      <img :src="currentUser.avatar" :alt="currentUser.username" class="profile-avatar" />
      <div class="profile-info">
        <h3 class="profile-username">{{ currentUser.username }}</h3>
        <p class="profile-name">{{ currentUser.fullName }}</p>
      </div>
      <button class="switch-btn" @click="$emit('switch-account')">Cambiar</button>
    </div>

    <!-- Sugerencias de amistad -->
    <div class="suggestions-container">
      <div class="suggestions-header">
        <h4>Sugerencias para ti</h4>
        <button class="view-all-btn" @click="$emit('view-all-suggestions')">Ver todo</button>
      </div>
      <ul class="suggestions-list">
        <li 
          v-for="suggestion in suggestions" 
          :key="suggestion.id" 
          class="suggestion-item"
        >
          <div class="suggestion-user">
            <img :src="suggestion.avatar" :alt="suggestion.username" class="suggestion-avatar" />
            <div class="suggestion-info">
              <p class="suggestion-username">{{ suggestion.username }}</p>
              <p class="suggestion-meta">{{ suggestion.meta }}</p>
            </div>
          </div>
          <button 
            class="follow-btn" 
            @click="$emit('follow', suggestion.id)"
          >
            Seguir
          </button>
        </li>
      </ul>
    </div>

    <!-- Footer con enlaces -->
    <div class="sidebar-footer">
      <nav class="footer-links">
        <a href="#" class="footer-link">Acerca de</a>
        <a href="#" class="footer-link">Ayuda</a>
        <a href="#" class="footer-link">Privacidad</a>
        <a href="#" class="footer-link">Términos</a>
      </nav>
      <p class="copyright">© 2025 SocialCue</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { CurrentUser, SuggestionUser } from '~/types/dashboard';

/**
 * Props del componente
 */
const props = defineProps<{
  /**
   * Información del usuario actual
   */
  currentUser: CurrentUser;

  /**
   * Lista de sugerencias de usuarios para seguir
   */
  suggestions: SuggestionUser[];
}>();

/**
 * Eventos emitidos por el componente
 */
const emit = defineEmits<{
  /**
   * Emitido cuando el usuario quiere cambiar de cuenta
   */
  (e: 'switch-account'): void;

  /**
   * Emitido cuando el usuario quiere ver todas las sugerencias
   */
  (e: 'view-all-suggestions'): void;

  /**
   * Emitido cuando el usuario quiere seguir a alguien
   * @param userId ID del usuario a seguir
   */
  (e: 'follow', userId: number): void;
}>();
</script>

<style scoped>
/* Sidebar de sugerencias */
.suggestions-sidebar {
  width: 300px;
  display: none; /* Oculto por defecto en móviles */
}

/* Tarjeta de perfil de usuario */
.user-profile-card {
  display: flex;
  align-items: center;
  margin-bottom: 1.5rem;
}

.profile-avatar {
  width: 56px;
  height: 56px;
  border-radius: 50%;
  object-fit: cover;
  margin-right: 1rem;
}

.profile-info {
  flex: 1;
}

.profile-username {
  margin: 0;
  font-size: 0.875rem;
  font-weight: 600;
  color: #FFFFFF;
}

.profile-name {
  margin: 0;
  font-size: 0.75rem;
  color: #8A97A8;
}

.switch-btn {
  background: transparent;
  border: none;
  color: #4D8CD9;
  font-weight: 600;
  cursor: pointer;
  font-size: 0.75rem;
}

/* Sugerencias */
.suggestions-container {
  background-color: #122237;
  border-radius: 8px;
  border: 1px solid #243447;
  padding: 1rem;
  margin-bottom: 1.5rem;
}

.suggestions-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1rem;
}

.suggestions-header h4 {
  margin: 0;
  font-size: 0.875rem;
  font-weight: 600;
  color: #8A97A8;
}

.view-all-btn {
  background: transparent;
  border: none;
  color: #4D8CD9;
  cursor: pointer;
  font-size: 0.875rem;
}

.suggestions-list {
  list-style: none;
  padding: 0;
  margin: 0;
}

.suggestion-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 0.75rem;
}

.suggestion-user {
  display: flex;
  align-items: center;
  gap: 0.75rem;
}

.suggestion-avatar {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  object-fit: cover;
}

.suggestion-info p {
  margin: 0;
}

.suggestion-username {
  font-size: 0.75rem;
  font-weight: 600;
  color: #FFFFFF;
}

.suggestion-meta {
  font-size: 0.6875rem;
  color: #8A97A8;
}

.follow-btn {
  background: transparent;
  border: none;
  color: #4D8CD9;
  font-weight: 600;
  cursor: pointer;
  font-size: 0.75rem;
}

/* Footer del sidebar */
.sidebar-footer {
  margin-top: 1.5rem;
  color: #8A97A8;
  font-size: 0.75rem;
}

.footer-links {
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
  margin-bottom: 0.5rem;
}

.footer-link {
  color: #8A97A8;
  text-decoration: none;
}

.footer-link:hover {
  text-decoration: underline;
}

.copyright {
  margin: 0;
}

/* Responsive */
@media (min-width: 1024px) {
  .suggestions-sidebar {
    display: block;
  }
}
</style>

