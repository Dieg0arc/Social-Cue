<!-- components/auth/PasswordField.vue -->
<!--
  PasswordField - Componente de campo de contraseña con visibilidad togglable
  
  Este componente proporciona un campo de entrada para contraseñas con la capacidad
  de mostrar u ocultar el texto ingresado. Implementa etiquetas flotantes para mejorar
  la experiencia de usuario y se integra con v-model para un manejo de datos bidireccional.
-->

<template>
  <div class="form-group password-field">
    <div class="float-label-wrapper" :class="{ 'has-value': modelValue }">
      <div class="password-input-wrapper">
        <input 
          :id="id"
          :type="showPassword ? 'text' : 'password'" 
          :value="modelValue"
          @input="handleInput"
          required
          class="input-field"
        />
        <label :for="id" class="float-label">{{ label }}</label>
        <button 
          type="button" 
          class="toggle-password" 
          @click="showPassword = !showPassword"
          tabindex="-1"
          aria-label="Toggle password visibility"
        >
          <img 
            :src="showPassword ? eyeHideIcon : eyeViewIcon" 
            :alt="showPassword ? 'Ocultar contraseña' : 'Mostrar contraseña'"
            class="eye-icon"
          />
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';

/**
 * Importación de iconos para el toggle de visibilidad
 * Nota: Ajustar las rutas según la estructura del proyecto
 */
import eyeViewIcon from '@/assets/images/Eye_view.png';
import eyeHideIcon from '@/assets/images/Eye_hide.png';

/**
 * Props del componente
 * @prop {String} modelValue - Valor actual del campo (requerido para v-model)
 * @prop {String} label - Texto de la etiqueta del campo
 * @prop {String} placeholder - Texto de placeholder (no visible con etiqueta flotante)
 * @prop {String} id - ID único para el campo y la etiqueta
 */
const props = defineProps({
  modelValue: {
    type: String,
    required: true
  },
  label: {
    type: String,
    default: 'Contraseña'
  },
  placeholder: {
    type: String,
    default: 'Ingresa tu contraseña'
  },
  id: {
    type: String,
    default: 'password'
  }
});

/**
 * Evento para actualizar el valor del modelo (v-model)
 */
const emit = defineEmits(['update:modelValue']);

/**
 * Maneja los cambios en el campo de entrada
 * Utiliza currentTarget para mayor seguridad en lugar de target
 * 
 * @param {Event} event - Evento de input del navegador
 */
const handleInput = (event: Event) => {
  const input = event.currentTarget as HTMLInputElement;
  emit('update:modelValue', input.value);
};

/**
 * Estado reactivo que controla la visibilidad de la contraseña
 */
const showPassword = ref(false);
</script>

<style scoped>
/* Contenedor principal del campo de contraseña */
.password-field {
  width: 100%;
}

/* Contenedor para implementar etiquetas flotantes */
.float-label-wrapper {
  position: relative;
  width: 100%;
}

/* Contenedor para el input y el botón de toggle */
.password-input-wrapper {
  position: relative;
  display: flex;
  width: 100%;
}

/* Estilos base para el campo de entrada */
.input-field {
  width: 100%;
  padding: 15px;
  padding-top: 25px;
  padding-bottom: 5px;
  background-color: #243447;
  border: 1px solid #3A4A5C;
  border-radius: 10px;
  color: #FFFFFF;
  font-size: 16px;
  transition: border-color 0.2s;
  box-sizing: border-box;
}

/* Estado de enfoque para el campo de entrada */
.input-field:focus {
  outline: none;
  border-color: #4D8CD9;
}

/* Etiqueta flotante que se posiciona sobre el campo */
.float-label {
  position: absolute;
  top: 50%;
  left: 15px;
  transform: translateY(-50%);
  color: #8A97A8;
  pointer-events: none;
  transition: all 0.2s ease;
  font-size: 16px;
}

/* Animación de la etiqueta cuando el campo tiene enfoque o contiene valor */
.input-field:focus ~ .float-label,
.has-value .float-label {
  top: 25%;
  font-size: 12px;
  color: #4D8CD9;
}

/* Botón para alternar la visibilidad de la contraseña */
.toggle-password {
  position: absolute;
  right: 10px;
  top: 50%;
  transform: translateY(-50%);
  background: none;
  border: none;
  cursor: pointer;
  padding: 0;
  display: flex;
  align-items: center;
  justify-content: center;
}

/* Icono del botón de visibilidad */
.eye-icon {
  width: 20px;
  height: 20px;
}
</style>