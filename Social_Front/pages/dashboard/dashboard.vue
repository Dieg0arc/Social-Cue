<template>
  <div class="dashboard-container">
    <!-- Header -->
    <!-- <header class="dashboard-header">
      <div class="header-content">
        <div class="header-left">
          <h1 class="app-title">SocialCue</h1>
          <img src="~/assets/images/OnlyLogo_Tp.png" alt="Logo" class="header-logo" />
        </div>
        <div class="header-right">
          <div class="user-info">
            <span class="user-name">{{ authStore.user?.name || 'Usuario' }}</span>
            <img src="assets\images\Usuari.png" alt="Avatar" class="user-avatar" />
          </div>
          <button @click="handleLogout" class="logout-button">
            Cerrar sesión
          </button>
        </div>
      </div>
    </header> -->

    <!-- Main Content -->
    <main class="main-content">
      <div class="welcome-card">
        <h2 class="welcome-title">¡Bienvenido a SocialCue!</h2>
        <p class="welcome-subtitle">Has iniciado sesión correctamente</p>
        
        <div class="user-details">
          <div class="detail-item">
            <span class="detail-label">Usuario:</span>
            <span class="detail-value">{{ authStore.user?.username }}</span>
          </div>
          <div class="detail-item">
            <span class="detail-label">Email:</span>
            <span class="detail-value">{{ authStore.user?.email }}</span>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { useAuthStore } from '~/stores/auth';

const authStore = useAuthStore();

const handleLogout = async () => {
  await authStore.logout();
};
</script>

<style scoped>
.dashboard-container {
  min-height: 100vh;
  background-color: #0A1625;
  color: #FFFFFF;
}

.dashboard-header {
  background-color: #1A2535;
  padding: 1rem 2rem;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.2);
  position: sticky;
  top: 0;
  z-index: 100;
}

.header-content {
  max-width: 1200px;
  margin: 0 auto;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.header-left {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.app-title {
  font-size: 1.8rem;
  font-weight: 600;
  color: #4D8CD9;
  margin: 0;
}

.header-logo {
  height: 40px;
  width: auto;
}

.header-right {
  display: flex;
  align-items: center;
  gap: 1.5rem;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 0.8rem;
}

.user-name {
  color: #D0D0D0;
  font-size: 0.9rem;
}

.user-avatar {
  width: 35px;
  height: 35px;
  border-radius: 50%;
  object-fit: cover;
}

.logout-button {
  background-color: #0056A6;
  color: white;
  border: none;
  padding: 0.6rem 1.2rem;
  border-radius: 8px;
  font-size: 0.9rem;
  cursor: pointer;
  transition: background-color 0.2s;
}

.logout-button:hover {
  background-color: #4D8CD9;
}

.main-content {
  max-width: 1200px;
  margin: 2rem auto;
  padding: 0 2rem;
}

.welcome-card {
  background-color: #1A2535;
  border-radius: 15px;
  padding: 2rem;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.2);
}

.welcome-title {
  color: #4D8CD9;
  font-size: 1.8rem;
  margin: 0 0 0.5rem 0;
}

.welcome-subtitle {
  color: #D0D0D0;
  margin-bottom: 2rem;
}

.user-details {
  background-color: #243447;
  border-radius: 10px;
  padding: 1.5rem;
}

.detail-item {
  display: flex;
  margin-bottom: 1rem;
}

.detail-item:last-child {
  margin-bottom: 0;
}

.detail-label {
  font-weight: 600;
  color: #4D8CD9;
  width: 100px;
}

.detail-value {
  color: #D0D0D0;
}

@media (max-width: 768px) {
  .dashboard-header {
    padding: 1rem;
  }

  .header-content {
    flex-direction: column;
    gap: 1rem;
  }

  .header-right {
    width: 100%;
    justify-content: space-between;
  }

  .main-content {
    padding: 0 1rem;
    margin: 1rem auto;
  }

  .welcome-card {
    padding: 1.5rem;
  }
}
</style>