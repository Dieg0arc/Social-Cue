<!-- 
  Archivo: authenticated.vue
  Ubicación: layouts/authenticated.vue
  Descripción: Layout principal para usuarios autenticados en SocialCue.
  Este layout contiene la estructura de navegación lateral, header y área de contenido.
-->

<template>
  <div class="app-layout">
    <!-- Sidebar de navegación principal -->
    <aside class="sidebar" :class="{ 'sidebar-collapsed': sidebarCollapsed }">
      <!-- Logo y nombre de la aplicación -->
      <div class="sidebar-header">
        <img src="~/assets/images/OnlyLogo_Tp.png" alt="SocialCue Logo" class="sidebar-logo" />
        <h1 v-if="!sidebarCollapsed" class="sidebar-title">SocialCue</h1>
      </div>

      <!-- Navegación principal -->
      <nav class="sidebar-nav">
        <ul class="nav-list">
          <li class="nav-item" :class="{ active: route.path === '/dashboard' }">
            <NuxtLink to="http://localhost:3000/dashboard" class="nav-link">
              <i class="pi pi-home nav-icon"></i>
              <span v-if="!sidebarCollapsed" class="nav-text">Inicio</span>
            </NuxtLink>
          </li>
          <li class="nav-item" :class="{ active: route.path.includes('/profile') }">
            <NuxtLink to="/profile" class="nav-link">
              <i class="pi pi-user nav-icon"></i>
              <span v-if="!sidebarCollapsed" class="nav-text">Perfil</span>
            </NuxtLink>
          </li>
          <li class="nav-item" :class="{ active: route.path.includes('/messages') }">
            <NuxtLink to="http://localhost:3000/messages" class="nav-link">
              <i class="pi pi-comments nav-icon"></i>
              <span v-if="!sidebarCollapsed" class="nav-text">Mensajes</span>
            </NuxtLink>
          </li>
          <li class="nav-item" :class="{ active: route.path.includes('/events') }">
            <NuxtLink to="http://localhost:3000/events" class="nav-link">
              <i class="pi pi-calendar nav-icon"></i>
              <span v-if="!sidebarCollapsed" class="nav-text">Eventos</span>
            </NuxtLink>
          </li>
          <li class="nav-item" :class="{ active: route.path.includes('/explore') }">
            <NuxtLink to="http://localhost:3000/explore" class="nav-link">
              <i class="pi pi-search nav-icon"></i>
              <span v-if="!sidebarCollapsed" class="nav-text">Explorar</span>
            </NuxtLink>
          </li>
        </ul>
      </nav>

      <!-- Sección inferior del sidebar -->
      <div class="sidebar-footer">
        <button @click="handleLogout" class="logout-btn">
          <i class="pi pi-sign-out nav-icon" style="font-size: 1.2rem;"></i>
          <h3 v-if="!sidebarCollapsed">Cerrar sesión</h3>
        </button>
      </div>
    </aside>

    <!-- Contenido principal -->
    <div class="main-container">
      <!-- Header con información del usuario y acciones -->
      <header class="app-header">
        <div class="header-left">
          <!-- Botón para colapsar el menú -->
          <button @click="toggleSidebar" class="collapse-btn-header">
            <i class="pi pi-align-justify"></i>
          </button>
          <!-- Título de la página -->
          <h2 class="page-title">{{ pageTitle }}</h2>
        </div>

        <div class="header-right">
          <div class="search-bar">
            <input type="text" placeholder="Buscar..." class="search-input" />
            <button class="search-btn">
              <i class="pi pi-search"></i>
            </button>
          </div>
          <div class="notifications">
            <button class="notification-btn">
              <i class="pi pi-bell"></i>
            </button>
          </div>
          <div class="user-profile" @click="toggleUserMenu">
            <img 
              :src="user?.profilePicture || 'assets/images/Usuari.png'" 
              alt="Foto de perfil" 
              class="user-avatar" 
            />
            <span class="user-name">{{ user?.fullName || 'Usuario' }}</span>
            <div v-if="userMenuOpen" class="user-menu">
              <ul>
                <li><NuxtLink to="/profile">Mi perfil</NuxtLink></li>
                <li><NuxtLink to="/settings">Configuración</NuxtLink></li>
                <li><button @click="handleLogout">Cerrar sesión</button></li>
              </ul>
            </div>
          </div>
        </div>
      </header>

      <!-- Contenido de la página actual -->
      <main class="content-area">
        <slot />
      </main>
    </div>
  </div>
</template>



<script setup lang="ts">
import { ref, computed } from 'vue';
// import { useRoute, useRouter } from '#app';
import { useAuth } from '~/composables/useAuth';

// Obtener instancias de route y router
const route = useRoute();
const router = useRouter();

// Obtener la información de autenticación
const { user, logout } = useAuth();

// Estado para controlar el colapso del sidebar
const sidebarCollapsed = ref(false);

// Estado para controlar el menú de usuario
const userMenuOpen = ref(false);

// Determinar el título de la página basado en la ruta actual
const pageTitle = computed(() => {
  switch (true) {
    case route.path === '/dashboard':
      return 'Inicio';
    case route.path.includes('/profile'):
      return 'Perfil';
    case route.path.includes('/messages'):
      return 'Mensajes';
    case route.path.includes('/events'):
      return 'Eventos';
    case route.path.includes('/explore'):
      return 'Explorar';
    default:
      return 'SocialCue';
  }
});

/**
 * Alterna el estado de colapso del sidebar
 */
const toggleSidebar = () => {
  sidebarCollapsed.value = !sidebarCollapsed.value;
};

/**
 * Alterna la visibilidad del menú de usuario
 */
const toggleUserMenu = () => {
  userMenuOpen.value = !userMenuOpen.value;
};

/**
 * Maneja el cierre de sesión del usuario
 */
const handleLogout = async () => {
  await logout();
  router.push('/login');
};
</script>

<style src="~/assets/css/dashboard.css"></style>
