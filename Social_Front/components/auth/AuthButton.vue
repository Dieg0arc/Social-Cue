<!-- components/auth/AuthButton.vue -->
<!-- 
  AuthButton - Componente botón para acciones de autenticación
  
  Este componente renderiza un botón personalizado para acciones de autenticación
  con estados de carga y deshabilitado. Cambia dinámicamente su texto según 
  el estado actual y mantiene un estilo consistente con los demás elementos
  de autenticación.
-->

<template>
  <button 
    :type="type as 'button' | 'submit' | 'reset'" 
    class="auth-button" 
    :disabled="loading || disabled"
  >
    <span v-if="loading">{{ loadingText }}</span>
    <span v-else>{{ defaultText }}</span>
  </button>
</template>

<script setup lang="ts">
/**
 * Propiedades del componente AuthButton
 * 
 * @prop {('button'|'submit'|'reset')} type - Tipo de botón HTML, por defecto 'button'
 * @prop {Boolean} loading - Indica si el botón está en estado de carga
 * @prop {Boolean} disabled - Indica si el botón está deshabilitado
 * @prop {String} loadingText - Texto a mostrar cuando loading=true, por defecto 'Cargando...'
 * @prop {String} defaultText - Texto principal del botón (requerido)
 */
defineProps({
  type: {
    type: String as () => 'button' | 'submit' | 'reset',
    default: 'button'
  },
  loading: {
    type: Boolean,
    default: false
  },
  disabled: {
    type: Boolean,
    default: false
  },
  loadingText: {
    type: String,
    default: 'Cargando...'
  },
  defaultText: {
    type: String,
    required: true
  }
});
</script>
<style>
/* Importa los estilos específicos para componentes de autenticación */
@import '~/assets/css/auth.css';
</style>