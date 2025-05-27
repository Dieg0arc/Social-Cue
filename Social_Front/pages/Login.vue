<template>
  <div class="login-container">
    <div class="background-shapes">
      <div class="shape shape-1"></div>
      <div class="shape shape-2"></div>
    </div>
    
    <div class="login-card">
      <div class="card-header">
        <div class="title-wrapper">
          <h1 class="app-title">SocialCue</h1>
          <img src="~/assets/images/OnlyLogo_Tp.png" alt="Logo Humboldt" class="title-logo" />
        </div>
        <p class="app-subtitle">Tu red social del conocimiento</p>
      </div>
      
<div class="card-content">
    <div class="card-left">
      <img src="~/assets/images/LoginLogo.png" alt="RedCue Logo" class="login-image" />
    </div>

    <div class="card-right">
      <form @submit.prevent="handleLogin" class="login-form">
        <div class="form-group">
          <label for="email">Correo institucional</label>
          <input 
            type="email" 
            id="email" 
            v-model="email" 
            required 
            placeholder="ejemplo@cue.edu.co"
            class="input-field"
          />
        </div>

        <div class="form-group">
          <div class="password-header">
            <label for="password">Contraseña</label>
            <a href="#" class="forgot-link">¿Olvidaste tu contraseña?</a>
          </div>

          <div class="password-input-wrapper">
            <input 
              :type="showPassword ? 'text' : 'password'" 
              id="password" 
              v-model="password" 
              required 
              placeholder="Ingresa tu contraseña"
              class="input-field"
            />
            <i
              :class="showPassword ? 'pi pi-eye-slash' : 'pi pi-eye'"
              class="eye-icon"
              @click="showPassword = !showPassword"
            ></i>
          </div>
        </div>

        <div class="remember-container">
          <label class="remember-label">
            <input type="checkbox" v-model="rememberMe" />
            Recordar sesión
          </label>
        </div>

        <button type="submit" :disabled="authStore.loading" class="login-button">
          {{ authStore.loading ? 'Iniciando sesión...' : 'Iniciar sesión' }}
        </button>

        <div
          v-if="authStore.error"
          class="error-message"
          style="color: #ff4d4d; margin-top: 10px; text-align: center;"
        >
          {{ authStore.error }}
        </div>
      </form>
    </div>
  </div>

      <div class="register-prompt">
        <p>¿Nuevo en SocialCue? <NuxtLink to="http://localhost:3000/register" class="register-link">Crear una cuenta</NuxtLink></p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useAuthStore } from '~/stores/auth';
// import { navigateTo } from '#app';

const email = ref('');
const password = ref('');
const rememberMe = ref(false);
const authStore = useAuthStore();
const showPassword = ref(false)

const handleLogin = async () => {
  try {
    const success = await authStore.login({
      email: email.value,
      password: password.value
    });

    if (success) {
      await navigateTo('/dashboard');
    }
  } catch (error) {
    console.error('Error during login:', error);
  }
};
</script>

<style>
@import '~/assets/css/login.css';
</style>
