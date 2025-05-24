<template>
  <!-- Contenedor principal que abarca toda la página de inicio -->
  <div class="home-page">
    <!-- Contenedor responsivo con margen automático y espaciado interno -->
    <div class="container mx-auto px-4 py-12">
      <!-- Título principal de la plataforma -->
      <h1 class="text-3xl font-bold mb-6">Red Social del Conocimiento Humboldt</h1>
      
      <!-- Sección visible solo para usuarios autenticados -->
      <div v-if="isAuthenticated">
        <!-- Mensaje de bienvenida personalizado con el nombre del usuario -->
        <p class="mb-4">Bienvenido, {{ user?.fullName || 'Usuario' }}!</p>
        
        <!-- Cuadrícula de características principales (1 columna en móvil, 3 en escritorio) -->
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mt-8">
          <!-- Tarjeta de característica: Explorar recursos -->
          <div class="feature-card">
            <h2 class="text-xl font-semibold mb-2">Explorar Recursos</h2>
            <p>Descubre tutoriales, videos y materiales compartidos por la comunidad universitaria.</p>
            <button class="btn-primary mt-4">Explorar</button>
          </div>
          
          <!-- Tarjeta de característica: Compartir conocimiento -->
          <div class="feature-card">
            <h2 class="text-xl font-semibold mb-2">Compartir Conocimiento</h2>
            <p>Comparte tus propios recursos con otros estudiantes y profesores.</p>
            <button class="btn-primary mt-4">Publicar</button>
          </div>
          
          <!-- Tarjeta de característica: Eventos académicos -->
          <div class="feature-card">
            <h2 class="text-xl font-semibold mb-2">Eventos Académicos</h2>
            <p>Descubre y participa en eventos organizados por la comunidad.</p>
            <button class="btn-primary mt-4">Ver Eventos</button>
          </div>
        </div>
      </div>
      
      <!-- Sección visible para usuarios no autenticados -->
      <div v-else class="text-center">
        <!-- Descripción breve de la plataforma -->
        <p class="mb-6">Comparte y descubre conocimiento en la comunidad Alexander Von Humboldt</p>
        
        <!-- Botones de acceso y registro -->
        <div class="flex justify-center gap-4">
          <!-- Enlace para iniciar sesión usando componente NuxtLink -->
          <NuxtLink to="/login" class="btn-primary">Iniciar Sesión</NuxtLink>
          
          <!-- Enlace para registrarse usando componente NuxtLink -->
          <NuxtLink to="/register" class="btn-secondary">Registrarse</NuxtLink>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
/**
 * Script de configuración para la página principal
 * 
 * Usa composition API de Vue con TypeScript para gestionar
 * el estado de autenticación de los usuarios.
 */

// Importación del composable de autenticación personalizado
import { useAuth } from '~/composables/useAuth';

// Desestructuración para obtener el usuario actual y su estado de autenticación
const { user, isAuthenticated } = useAuth();
</script>

<style scoped>
/**
 * Estilos específicos para la página principal
 * Los estilos están encapsulados con 'scoped' para evitar afectar otros componentes
 */

/* Establece altura mínima para la página principal */
.home-page {
  min-height: 80vh;
}

/* Estilo para las tarjetas de características */
.feature-card {
  padding: 1.5rem;
  border-radius: 0.5rem;
  background-color: white;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  transition: transform 0.2s, box-shadow 0.2s;
}

/* Efecto hover para las tarjetas de características */
.feature-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 6px 12px rgba(0, 0, 0, 0.15);
}

/* Estilo para el botón secundario */
.btn-secondary {
  background-color: white;
  color: var(--primary-color);
  border: 1px solid var(--primary-color);
  padding: 0.75rem 1rem;
  border-radius: 0.25rem;
  font-weight: bold;
  cursor: pointer;
  transition: background-color 0.2s;
}

/* Efecto hover para el botón secundario */
.btn-secondary:hover {
  background-color: #f3f4f6;
}

/* Nota: El estilo .btn-primary se asume que está definido globalmente */
</style>