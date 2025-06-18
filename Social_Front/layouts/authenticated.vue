<!-- Layout principal de la aplicación SocialCue -->
<template>
  <!-- Contenedor general del layout -->
  <div class="app-layout">

    <!-- Barra lateral de navegación -->
    <aside class="sidebar" :class="{ 'sidebar-collapsed': sidebarCollapsed }">
      
      <!-- Encabezado con logo y título -->
      <div class="sidebar-header">
        <img src="~/assets/images/OnlyLogo_Tp.png" alt="SocialCue Logo" class="sidebar-logo" />
        <h1 v-if="!sidebarCollapsed" class="sidebar-title">SocialCue</h1>
      </div>

      <!-- Navegación del sidebar -->
      <nav class="sidebar-nav">
        <ul class="nav-list">
          <!-- Opción: Inicio -->
          <li class="nav-item" :class="{ active: route.path === '/dashboard' }">
            <NuxtLink to="/dashboard" class="nav-link">
              <i class="pi pi-home nav-icon"></i>
              <span v-if="!sidebarCollapsed" class="nav-text">Inicio</span>
            </NuxtLink>
          </li>

          <!-- Opción: Perfil -->
          <li class="nav-item" :class="{ active: route.path.includes('/profile') }">
            <NuxtLink to="/profile" class="nav-link">
              <i class="pi pi-user nav-icon"></i>
              <span v-if="!sidebarCollapsed" class="nav-text">Perfil</span>
            </NuxtLink>
          </li>

          <!-- Opción: Explorar -->
          <!-- <li class="nav-item" :class="{ active: route.path.includes('/explore') }">
            <NuxtLink to="/explore" class="nav-link">
              <i class="pi pi-search nav-icon"></i>
              <span v-if="!sidebarCollapsed" class="nav-text">Explorar</span>
            </NuxtLink>
          </li> -->
        </ul>
      </nav>

      <!-- Botón de cerrar sesión -->
      <div class="sidebar-footer">
        <button @click="handleLogout" class="logout-btn">
          <i class="pi pi-sign-out nav-icon"></i>
          <h3 v-if="!sidebarCollapsed">Cerrar sesión</h3>
        </button>
      </div>
    </aside>

    <!-- Contenedor principal de contenido -->
    <div class="main-container">
      
      <!-- Encabezado principal -->
      <header class="app-header">
        <div class="header-left">
          <!-- Botón para colapsar/expandir sidebar -->
          <button @click="toggleSidebar" class="collapse-btn-header">
            <i class="pi pi-align-justify"></i>
          </button>
          <!-- Título dinámico según ruta -->
          <h2 class="page-title">{{ pageTitle }}</h2>
        </div>

        <!-- Panel derecho del header: búsqueda, notificaciones y usuario -->
        <div class="header-right">

          <!-- Barra de búsqueda (sin funcionalidad aún) -->
          <div class="search-bar">
            <input type="text" placeholder="Buscar..." class="search-input" />
            <button class="search-btn">
              <i class="pi pi-search"></i>
            </button>
          </div>

          <!-- Icono de notificaciones -->
          <div class="notifications">
            <button class="notification-btn">
              <i class="pi pi-bell"></i>
            </button>
          </div>

          <!-- Perfil de usuario y menú desplegable -->
          <div class="user-profile" @click="toggleUserMenu">
            <img 
              :src="user?.profilePicture || '/assets/images/Usuari.png'" 
              alt="Foto de perfil" 
              class="user-avatar" 
            />
            <span class="user-name">{{ user?.username || 'Usuario' }}</span>

            <!-- Menú del usuario -->
            <div v-if="userMenuOpen" class="user-menu">
              <ul>
                <li><NuxtLink to="/profile">Mi perfil</NuxtLink></li>
                <li><button @click="handleLogout">Cerrar sesión</button></li>
              </ul>
            </div>
          </div>
        </div>
      </header>

      <!-- Área donde se renderizan las vistas hijas -->
      <main class="content-area">
        <slot />
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
/**
 * Script principal para el layout `authenticated`.
 * Se encarga de manejar la lógica de la barra lateral,
 * sesión de usuario y encabezado de la vista.
 */

import { ref, computed } from 'vue';
import { useAuth } from '~/composables/useAuth';

// Rutas actuales del sistema
const route = useRoute();
const router = useRouter();

// Estado global de usuario y función de logout
const { user, logout } = useAuth();

// Sidebar colapsado o expandido
const sidebarCollapsed = ref(false);

// Estado de apertura del menú de usuario
const userMenuOpen = ref(false);

// Título dinámico de la página basado en la ruta actual
const pageTitle = computed(() => {
  switch (true) {
    case route.path === '/dashboard': return 'Inicio';
    case route.path.includes('/profile'): return 'Perfil';
    case route.path.includes('/explore'): return 'Explorar';
    default: return 'SocialCue';
  }
});

// Alternar visibilidad del sidebar
const toggleSidebar = () => sidebarCollapsed.value = !sidebarCollapsed.value;

// Alternar visibilidad del menú del usuario
const toggleUserMenu = () => userMenuOpen.value = !userMenuOpen.value;

// Cierra sesión y redirige al login
const handleLogout = async () => {
  await logout();
  router.push('/login');
};
</script>

<!-- Estilo base del dashboard -->
<style src="~/assets/css/dashboard.css"></style>
