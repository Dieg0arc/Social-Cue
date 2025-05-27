<!-- 
  Archivo: register.vue
  Ubicación: pages/register.vue
  Descripción: Página de registro para la aplicación SocialCue.
  Esta página presenta el formulario de registro para crear una cuenta en la plataforma.
-->

<template>
  <!-- Contenedor principal con el fondo y la tarjeta de registro -->
  <div class="login-container">
    <!-- Elementos decorativos del fondo -->
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
          <img src="~/assets/images/OnlyLogo_Tp.png" alt="Logo Humboldt" class="title-logo" />
        </div>
        <!-- Subtítulo descriptivo de la aplicación -->
        <p class="app-subtitle">Crea tu cuenta de aprendizaje</p>
      </div>
      
      <!-- Contenido principal dividido en dos columnas -->
      <div class="card-content">
        <!-- Columna izquierda con imagen ilustrativa -->
        <div class="card-left">
          <img src="assets\images\login-security.png" alt="RedCue Logo" class="login-image" />
        </div>
        
        <!-- Columna derecha con el formulario de registro -->
        <div class="card-right">
          <!-- 
            Componente RegisterForm con los eventos personalizados para gestionar
            las diferentes etapas del proceso de registro
          -->
          <RegisterForm 
            @register-attempt="handleRegisterAttempt"
            @register-success="handleRegisterSuccess"
            @register-error="handleRegisterError"
          />
        </div>
      </div>

      <!-- Sección para usuarios con cuenta con enlace a login -->
      <div class="register-prompt">
        <p>¿Ya tienes una cuenta? <NuxtLink to="http://localhost:3000/login" class="register-link">Iniciar sesión</NuxtLink></p>
      </div>
    </div>
    
    <!-- Mensaje de error global -->
    <div v-if="globalError" class="error-banner">
      <p>{{ globalError }}</p>
      <button @click="globalError = ''" class="close-error">×</button>
    </div>
  </div>
</template>

<script setup lang="ts">
/**
 * Configuración de la página de registro
 * Este script maneja la lógica relacionada con el proceso de creación de cuenta
 */

// Importación del componente de formulario de registro
import RegisterForm from '~/components/auth/RegisterForm.vue';
// Importación del router para redirecciones tras el registro
// import { useRouter } from '#app';
import { useAuth } from '~/composables/useAuth';
import { ref } from 'vue';

// Inicialización del router de Nuxt
const router = useRouter();
// Obtenemos la funcionalidad de autenticación
const { isAuthenticated } = useAuth();

// Estado para mensajes de error globales
const globalError = ref('');

/**
 * Definición de la interfaz para los datos del formulario
 * Especifica la estructura que deben tener los datos recibidos
 */
interface RegisterFormData {
  fullName: string;
  username: string;
  email: string;
  password: string;
}

/**
 * Maneja el evento cuando el usuario intenta registrarse
 * @param formData - Objeto con los datos del usuario
 */
const handleRegisterAttempt = (formData: RegisterFormData) => {
  console.log('Intento de registro con:', formData);
};

/**
 * Maneja el evento cuando el registro es exitoso
 * Redirige al usuario a la página principal
 */
const handleRegisterSuccess = () => {
  console.log('Registro exitoso, redirigiendo...');
  
  // La redirección ahora es manejada por el composable useAuth
  // Verificamos que el usuario esté autenticado
  if (isAuthenticated.value) {
    console.log('Estado de autenticación confirmado');
    
    // Mostramos un mensaje de éxito usando toast si está disponible
    const toast = useToast();
    if (toast) {
      toast.add({
        severity: 'success',
        summary: 'Registro exitoso',
        detail: 'Tu cuenta ha sido creada correctamente. ¡Bienvenido a SocialCue!',
        life: 3000
      });
    }
  } else {
    console.error('Usuario no autenticado después del registro');
    globalError.value = 'La cuenta se creó correctamente, pero hubo un problema con el inicio de sesión automático.';
    
    // Redirigimos al login en caso de problemas
    setTimeout(() => {
      router.push('/login');
    }, 2000);
  }
};

/**
 * Maneja el evento cuando ocurre un error durante el registro
 * @param error - Información del error producido
 */
const handleRegisterError = (error: unknown) => {
  console.error('Error en registro:', error);
  
  // Determinar el tipo de error para mostrar un mensaje más específico
  let errorMessage = 'Error al crear la cuenta. Por favor, intenta nuevamente.';
  
  if (error instanceof Error) {
    if (error.message.includes('email') && error.message.includes('exist')) {
      errorMessage = 'Este correo electrónico ya está registrado. Por favor, utiliza otro o inicia sesión.';
    } else if (error.message.includes('username') && error.message.includes('exist')) {
      errorMessage = 'Este nombre de usuario ya está en uso. Por favor, elige otro.';
    } else if (error.message.includes('network') || error.message.includes('conexión')) {
      errorMessage = 'Error de conexión. Por favor, verifica tu conexión a internet e intenta nuevamente.';
    }
  }
  
  // Mostrar mensaje de error
  globalError.value = errorMessage;
  
  // También mostrar una notificación toast si está disponible
  const toast = useToast();
  if (toast) {
    toast.add({
      severity: 'error',
      summary: 'Error de registro',
      detail: errorMessage,
      life: 5000
    });
  }
};
</script>

<style>
/**
 * Importación de estilos específicos para la página de registro
 * Reutilizamos los estilos de login para mantener consistencia
 */
@import '~/assets/css/login.css';

/* Estilos para el banner de error */
.error-banner {
  background-color: rgba(255, 77, 77, 0.9);
  color: white;
  padding: 12px 20px;
  border-radius: 6px;
  margin-top: 20px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 90%;
  max-width: 600px;
  margin-left: auto;
  margin-right: auto;
}

.close-error {
  background: transparent;
  border: none;
  color: white;
  font-size: 20px;
  cursor: pointer;
  padding: 0 0 0 15px;
}
</style>
