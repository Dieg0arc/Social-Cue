<!-- 
  Archivo: login.vue
  Ubicación: pages/login.vue
  Descripción: Página de inicio de sesión para la aplicación SocialCue.
  Esta página presenta el formulario de autenticación para acceder a la plataforma.
-->

<template>
  <!-- Contenedor principal con el fondo y la tarjeta de login -->
  <div class="login-container">
    <!-- Elementos decorativos del fondo -->
    <div class="background-shapes">
      <div class="shape shape-1"></div>
      <div class="shape shape-2"></div>
    </div>
    
    <!-- Tarjeta principal con el formulario de login -->
    <div class="login-card">
      <!-- Cabecera de la tarjeta con título y logo -->
      <div class="card-header">
        <div class="title-wrapper">
          <!-- Título de la aplicación -->
          <h1 class="app-title">SocialCue</h1>
          <!-- Logo de la institución Humboldt -->
          <img src="~/assets/images/OnlyLogo_Tp.png" alt="Logo Humboldt" class="title-logo" />
        </div>
        <!-- Subtítulo descriptivo de la aplicación -->
        <p class="app-subtitle">Tu red social del conocimiento</p>
      </div>
      
      <!-- Contenido principal dividido en dos columnas -->
      <div class="card-content">
        <!-- Columna izquierda con imagen ilustrativa -->
        <div class="card-left">
          <img src="~/assets/images/LoginLogo.png" alt="RedCue Logo" class="login-image" />
        </div>
        
        <!-- Columna derecha con el formulario de login -->
        <div class="card-right">
          <!-- 
            Componente LoginForm con los eventos personalizados para gestionar
            las diferentes etapas del proceso de login
          -->
          <LoginForm 
            @login-attempt="handleLoginAttempt"
            @login-success="handleLoginSuccess"
            @login-error="handleLoginError"
          />
        </div>
      </div>

      <!-- Sección para usuarios nuevos con enlace a registro -->
      <div class="register-prompt">
        <p>¿Nuevo en RedCue? <NuxtLink to="/register" class="register-link">Crear una cuenta</NuxtLink></p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
/**
 * Configuración de la página de login
 * Este script maneja la lógica relacionada con el proceso de autenticación
 */

// Importación del componente de formulario de login
import LoginForm from '~/components/auth/LoginForm.vue';
// Importación del router para redirecciones tras el login
import { useRouter } from '#app';

// Inicialización del router de Nuxt
const router = useRouter();

/**
 * Definición de la interfaz para los datos del formulario
 * Especifica la estructura que deben tener los datos recibidos
 */
interface LoginFormData {
  email: string;
  password: string;
}

/**
 * Maneja el evento cuando el usuario intenta iniciar sesión
 * @param formData - Objeto con email y password del usuario
 */
const handleLoginAttempt = (formData: LoginFormData) => {
  console.log('Intento de login con:', formData);
};

/**
 * Maneja el evento cuando el login es exitoso
 * Redirige al usuario a la página principal
 */
const handleLoginSuccess = () => {
  console.log('Login exitoso, redirigiendo...');
  router.push('/');
};

/**
 * Maneja el evento cuando ocurre un error durante el login
 * @param error - Información del error producido
 */
const handleLoginError = (error: unknown) => {
  console.error('Error en login:', error);
};
</script>

<style>
/**
 * Importación de estilos específicos para la página de login
 * Los estilos están separados en un archivo CSS externo
 */
@import '~/assets/css/login.css';
</style>