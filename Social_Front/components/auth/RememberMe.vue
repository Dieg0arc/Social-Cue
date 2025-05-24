<!-- components/auth/RememberMe.vue -->
<!--
  RememberMe - Componente de casilla de verificación para recordar sesión
  
  Este componente simple proporciona una casilla de verificación que permite
  al usuario optar por mantener su sesión activa en el dispositivo actual.
  Implementa la funcionalidad de v-model para una integración bidireccional
  con componentes padres.
-->

<template>
  <label class="remember-label">
    <input 
      type="checkbox" 
      :checked="modelValue"
      @change="e => $emit('update:modelValue', (e.target as HTMLInputElement).checked)" 
    />
    <span>Recordar este dispositivo</span>
  </label>
</template>
  
<script setup lang="ts">
/**
 * Props del componente
 * @prop {Boolean} modelValue - Estado actual de la casilla de verificación (requerido para v-model)
 */
defineProps({
  modelValue: {
    type: Boolean,
    required: true
  }
});

/**
 * Evento para actualizar el valor del modelo (v-model)
 * Se emite cuando cambia el estado de la casilla de verificación
 */
defineEmits(['update:modelValue']);
</script>

<style scoped>
/* Puedes agregar estilos personalizados para la casilla de verificación aquí */
.remember-label {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 14px;
  color: #8A97A8;
  cursor: pointer;
}

/* Estilo para que el cursor cambie al pasar sobre el texto también */
.remember-label span {
  cursor: pointer;
}
</style>