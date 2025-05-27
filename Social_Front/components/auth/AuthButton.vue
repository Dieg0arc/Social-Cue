<template>
  <button
    class="login-button"
    :type="type"
    :disabled="loading"
  >
    <span v-if="loading" class="loading-text">
      <span class="loading-indicator"></span>
      Cargando...
    </span>
    <span v-else>{{ text }}</span>
  </button>
</template>

<script setup lang="ts">
/**
 * AuthButton - Botón para acciones de autenticación
 * 
 * Componente reutilizable para acciones de autenticación como login y registro.
 * Incluye estados de carga y estilos consistentes con el diseño de la aplicación.
 */

defineProps<{
  loading: boolean;
  text: string;
  type?: 'submit' | 'button';
}>();
</script>

<style scoped>
/* Estilo base del botón de acción */
.login-button {
  width: 100%;
  padding: 14px;
  background-color: #0056A6; /* Color primario: azul institucional */
  color: white;
  border: none;
  border-radius: 25px;
  font-weight: 600;
  font-size: 16px;
  cursor: pointer;
  transition: background-color 0.2s;
  text-align: center;
  position: relative;
}

/* Estado hover del botón */
.login-button:hover {
  background-color: #4D8CD9; /* Color primario hover: azul más brillante */
}

/* Estado deshabilitado del botón */
.login-button:disabled {
  background-color: #4D8CD9; /* Versión más oscura cuando desactivado */
  opacity: 0.7;
  cursor: not-allowed;
}

/* Contenedor del texto de carga */
.loading-text {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
}

/* Animación del indicador de carga */
.loading-indicator {
  display: inline-block;
  width: 16px;
  height: 16px;
  border: 2px solid rgba(255, 255, 255, 0.3);
  border-radius: 50%;
  border-top-color: #fff;
  animation: spin 1s ease-in-out infinite;
}

/* Keyframes para la animación de giro */
@keyframes spin {
  to { transform: rotate(360deg); }
}
</style>
