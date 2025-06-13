<template>
  <div class="app-layout">
    <aside class="sidebar" :class="{ 'sidebar-collapsed': sidebarCollapsed }">
      <div class="sidebar-header">
        <img src="~/assets/images/OnlyLogo_Tp.png" alt="SocialCue Logo" class="sidebar-logo" />
        <h1 v-if="!sidebarCollapsed" class="sidebar-title">SocialCue</h1>
      </div>

      <nav class="sidebar-nav">
        <ul class="nav-list">
          <li class="nav-item" :class="{ active: route.path === '/dashboard' }">
            <NuxtLink to="/dashboard" class="nav-link">
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
          <li class="nav-item" :class="{ active: route.path.includes('/explore') }">
            <NuxtLink to="/explore" class="nav-link">
              <i class="pi pi-search nav-icon"></i>
              <span v-if="!sidebarCollapsed" class="nav-text">Explorar</span>
            </NuxtLink>
          </li>
        </ul>
      </nav>

      <div class="sidebar-footer">
        <button @click="handleLogout" class="logout-btn">
          <i class="pi pi-sign-out nav-icon"></i>
          <h3 v-if="!sidebarCollapsed">Cerrar sesión</h3>
        </button>
      </div>
    </aside>

    <div class="main-container">
      <header class="app-header">
        <div class="header-left">
          <button @click="toggleSidebar" class="collapse-btn-header">
            <i class="pi pi-align-justify"></i>
          </button>
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
              :src="user?.profilePicture || '/assets/images/Usuari.png'" 
              alt="Foto de perfil" 
              class="user-avatar" 
            />
            <span class="user-name">{{ user?.username || 'Usuario' }}</span>
            <div v-if="userMenuOpen" class="user-menu">
              <ul>
                <li><NuxtLink to="/profile">Mi perfil</NuxtLink></li>
                <li><button @click="handleLogout">Cerrar sesión</button></li>
              </ul>
            </div>
          </div>
        </div>
      </header>

      <main class="content-area">
        <slot />
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue';
import { useAuth } from '~/composables/useAuth';

const route = useRoute();
const router = useRouter();
const { user, logout } = useAuth();
const sidebarCollapsed = ref(false);
const userMenuOpen = ref(false);

const pageTitle = computed(() => {
  switch (true) {
    case route.path === '/dashboard': return 'Inicio';
    case route.path.includes('/profile'): return 'Perfil';
    case route.path.includes('/explore'): return 'Explorar';
    default: return 'SocialCue';
  }
});

const toggleSidebar = () => sidebarCollapsed.value = !sidebarCollapsed.value;
const toggleUserMenu = () => userMenuOpen.value = !userMenuOpen.value;
const handleLogout = async () => {
  await logout();
  router.push('/login');
};
</script>

<style src="~/assets/css/dashboard.css"></style>