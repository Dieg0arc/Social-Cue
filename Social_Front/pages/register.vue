<!-- 
  Archivo: register.vue
  Ubicación: pages/register.vue
  Descripción: Página de registro para nuevos usuarios en la aplicación SocialCue.
  Esta página permite a los usuarios crear una nueva cuenta en la plataforma.
-->

<template>
  <!-- Contenedor principal con el mismo diseño que la página de login -->
  <div class="login-container">
    <!-- Elementos decorativos de fondo -->
    <div class="background-shapes">
      <div class="shape shape-1"></div>
      <div class="shape shape-2"></div>
    </div>
    
    <!-- Tarjeta principal con el formulario de registro -->
    <div class="login-card">
      <!-- Cabecera de la tarjeta con título y logo -->
      <div class="card-header">
        <div class="title-wrapper">
          <!-- Título de la aplicación -->
          <h1 class="app-title">SocialCue</h1>
          <!-- Logo de la institución Humboldt -->
          <img src="assets\images\OnlyLogo_Tp.png" alt="Logo Humboldt" class="title-logo" />
        </div>
        <!-- Subtítulo específico para la página de registro -->
        <p class="app-subtitle">Crea tu cuenta</p>
      </div>
      
      <!-- Contenido principal dividido en dos columnas -->
      <div class="card-content">
        <!-- Columna izquierda con imagen ilustrativa -->
        <div class="card-left">
          <img src="~/assets/images/login-security.png" alt="RedCue Logo" class="login-image" />
        </div>
        
        <!-- Columna derecha con el formulario de registro -->
        <div class="card-right">
          <!-- Formulario de registro con prevención de envío por defecto -->
          <form @submit.prevent="handleSubmit">
            <!-- Campo para nombre completo -->
            <div class="form-group">
              <label for="fullName">Nombre completo</label>
              <input 
                id="fullName"
                type="text" 
                v-model="form.fullName" 
                placeholder="Ingresa tu nombre completo" 
                required
                class="input-field"
              />
            </div>
            
            <!-- Campo para nombre de usuario -->
            <div class="form-group">
              <label for="username">Nombre de usuario</label>
              <input 
                id="username"
                type="text" 
                v-model="form.username" 
                placeholder="Ingresa tu nombre de usuario" 
                required
                class="input-field"
              />
            </div>
            
            <!-- Campo para correo electrónico -->
            <div class="form-group">
              <label for="email">Correo electrónico</label>
              <input 
                id="email"
                type="email" 
                v-model="form.email" 
                placeholder="Ingresa tu@email.com" 
                required
                class="input-field"
              />
            </div>
            
            <!-- Campo para contraseña con opción de mostrar/ocultar -->
            <div class="form-group">
              <label for="password">Contraseña</label>
              <input 
                id="password"
                :type="showPassword ? 'text' : 'password'" 
                v-model="form.password" 
                placeholder="Ingresa tu contraseña" 
                required
                class="input-field"
              />
            </div>
            
            <!-- Mensaje de error condicional -->
            <div v-if="error" class="error-message">
              {{ error }}
            </div>
            
            <!-- Botón de envío con estado de carga -->
            <button type="submit" class="login-button" :disabled="loading">
              <span v-if="loading">Registrando...</span>
              <span v-else>Crear cuenta</span>
            </button>
          </form>
        </div>
      </div>

      <!-- Enlace para usuarios que ya tienen cuenta -->
      <div class="register-prompt">
        <p>¿Ya tienes una cuenta? <NuxtLink to="/login" class="register-link">Iniciar sesión</NuxtLink></p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
/**
 * Configuración de la página de registro
 * Este script maneja la lógica para crear nuevas cuentas de usuario
 */

// Importación de ref para variables reactivas
import { ref } from 'vue';
// Importación del composable de autenticación personalizado
import { useAuth } from '~/composables/useAuth';

// Obtención de funciones y estados del composable de autenticación
const { register, loading, error } = useAuth();

/**
 * Estado reactivo del formulario de registro
 * Contiene todos los campos necesarios para crear una cuenta
 */
const form = ref({
  fullName: '',   // Nombre completo del usuario
  username: '',   // Nombre de usuario único
  email: '',      // Correo electrónico del usuario
  password: ''    // Contraseña elegida
});

// Estado para controlar la visibilidad de la contraseña
const showPassword = ref(false);

/**
 * Maneja el envío del formulario de registro
 * Llama a la función register del composable de autenticación
 * con los datos introducidos por el usuario
 */
const handleSubmit = async () => {
  await register({
    fullName: form.value.fullName,
    username: form.value.username,
    email: form.value.email,
    password: form.value.password
  });
};
</script>

<style>
/**
 * Importación de estilos específicos para la página de registro
 * Los estilos están separados en un archivo CSS externo
 */
@import '~/assets/css/register.css';
</style>