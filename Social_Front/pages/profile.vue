<template>
  <div class="app-layout">
    <!-- Sidebar de navegación principal -->
    <aside class="sidebar">
      <div class="sidebar-header">
        <img src="/assets/images/OnlyLogo_Tp.png" alt="SocialCue Logo" class="sidebar-logo" />
        <h1 class="sidebar-title">SocialCue</h1>
      </div>
      <nav class="sidebar-nav">
        <ul class="nav-list">
          <li class="nav-item">
            <NuxtLink to="/dashboard" class="nav-link">
              <i class="pi pi-home nav-icon"></i>
              <span class="nav-text">Inicio</span>
            </NuxtLink>
          </li>
          <li class="nav-item">
            <NuxtLink to="/profile" class="nav-link">
              <i class="pi pi-user nav-icon"></i>
              <span class="nav-text">Perfil</span>
            </NuxtLink>
          </li>
          <li class="nav-item">
            <NuxtLink to="/explore" class="nav-link">
              <i class="pi pi-search nav-icon"></i>
              <span class="nav-text">Explorar</span>
            </NuxtLink>
          </li>
        </ul>
      </nav>
      <div class="sidebar-footer">
        <button @click="handleLogout" class="logout-btn">
          <i class="pi pi-sign-out nav-icon"></i>
          <span>Cerrar sesión</span>
        </button>
      </div>
    </aside>

    <!-- Contenido principal -->
    <div class="main-container">
      <header class="app-header">
        <div class="header-left">
          <button @click="toggleSidebar" class="collapse-btn-header">
            <i class="pi pi-align-justify"></i>
          </button>
          <h2 class="page-title">Perfil</h2>
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
          <div class="user-profile">
            <img :src="user?.profilePicture || '/assets/images/Usuari.png'" alt="Foto de perfil" class="user-avatar" />
            <span class="user-name">{{ user?.username || 'Usuario' }}</span>
          </div>
        </div>
      </header>

      <main class="content-area">
        <div class="profile-wrapper">
          <div class="profile-card">
            <img
              class="profile-avatar"
              src="/assets/images/Usuari.png"
              alt="Avatar"
            />
            <div class="profile-info">
              <h2 class="profile-username">{{ user?.username || 'Usuario desconocido' }}</h2>
              <p class="profile-email">Correo: {{ user?.email }}</p>
              <p class="profile-id">ID: {{ user?.id }}</p>
            </div>
          </div>
        </div>
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useAuth } from '@/composables/useAuth';
import { useRouter } from 'vue-router';

const { user, logout } = useAuth();
const router = useRouter();
const sidebarCollapsed = ref(false);

const toggleSidebar = () => {
  sidebarCollapsed.value = !sidebarCollapsed.value;
};

const handleLogout = async () => {
  await logout();
  router.push('/login');
};
</script>

  <style scoped>
  .app-layout {
    display: flex;
    background-color: #162A42; /* fondo oscuro */
    color: #fff;
    min-height: 100vh;
  }

  .sidebar {
    width: 240px;
    background-color: #0A1625;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    padding: 1rem;
  }

  .sidebar-logo {
    width: 40px;
    height: 40px;
  }

  .sidebar-title {
    font-size: 1.2rem;
    margin-left: 0.5rem;
  }

  .nav-list {
    list-style: none;
    padding: 0;
  }

  .nav-item {
    margin-bottom: 1rem;
  }

  .nav-link {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    color: #cbd5e1;
    text-decoration: none;
  }

  .logout-btn {
    background: none;
    border: none;
    color: #f87171;
    cursor: pointer;
    display: flex;
    align-items: center;
    gap: 0.5rem;
  }

  .main-container {
    flex: 1;
    display: flex;
    flex-direction: column;
  }

  .app-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 1rem;
    background-color: #0A1625;
    border-bottom: 1px solid #334155;
  }

  .header-left {
    display: flex;
    align-items: center;
    gap: 1rem;
  }

  .page-title {
    font-size: 1.25rem;
    font-weight: bold;
  }

  .header-right {
    display: flex;
    align-items: center;
    gap: 1rem;
  }

  .search-bar {
    display: flex;
    align-items: center;
    background-color: #334155;
    padding: 0.25rem 0.5rem;
    border-radius: 0.375rem;
  }

  .search-input {
    background: transparent;
    border: none;
    color: white;
    outline: none;
  }

  .user-profile {
    display: flex;
    align-items: center;
    gap: 0.5rem;
  }

  .user-avatar {
    width: 32px;
    height: 32px;
    border-radius: 9999px;
  }

  .content-area {
    flex: 1;
    padding: 2rem;
    background-color: #162A42;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .profile-card {
    background-color: #0A1625;
    padding: 2rem 80rem;
    border-radius: 1rem;
    box-shadow: 0 10px 25px rgba(0, 0, 0, 0.25);
    text-align: center;
    color: #ffffff;
    max-width: 500px;
    width: 100%;
  }

  .profile-avatar {
    width: 72px;
    height: 72px;
    border-radius: 50%;
    margin-bottom: 1rem;
    object-fit: cover;
    border: 2px solid #3b82f6;
  }

  .profile-info {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
  }

  .profile-username {
    font-size: 1.5rem;
    font-weight: bold;
  }

  .profile-email,
  .profile-id {
    font-size: 0.95rem;
    color: #cbd5e1;
  }
  </style>
