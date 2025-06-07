<!-- components/auth/LoginForm.vue -->
<!--
  LoginForm - Componente de formulario de inicio de sesión
  
  Este componente maneja el proceso completo de autenticación de usuario,
  incluyendo la validación de campos, manejo de errores y estados de carga.
  Integra varios subcomponentes reutilizables del sistema de autenticación
  y utiliza el composable useAuth para la lógica de inicio de sesión.
-->

<template>
  <form @submit.prevent="handleSubmit">
    <!-- Campo de email -->
    <div class="form-group">
      <div class="float-label-wrapper" :class="{ 'has-value': form.email }">
        <input
          id="email"
          type="email"
          v-model="form.email"
          required
          class="input-field"
        />
        <label for="email" class="float-label">Email</label>
      </div>
    </div>

    <!-- Campo de contraseña -->
    <PasswordField v-model="form.password" label="Contraseña" />

    <!-- Opciones de cuenta -->
    <div class="account-options">
      <RememberMe v-model="remember" />
      <NuxtLink to="/forgot-password" class="forgot-link">¿Olvidaste la contraseña?</NuxtLink>
    </div>

    <!-- Mensaje de error -->
    <div v-if="error" class="error-message">{{ error }}</div>

    <!-- Botón de envío -->
    <AuthButton type="submit" :loading="loading" text="Iniciar sesión" />
  </form>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useAuth } from '~/composables/useAuth';
import PasswordField from '~/components/auth/PasswordField.vue';
import RememberMe from '~/components/auth/RememberMe.vue';
import AuthButton from '~/components/auth/AuthButton.vue';
import type { LoginCredentials } from '~/types/user';

const { login, loading, error } = useAuth()

const form = ref({
  email: '',
  password: ''
})

const remember = ref(false)

const emit = defineEmits<{
  (e: 'login-attempt', form: { email: string, password: string }): void
  (e: 'login-success'): void
  (e: 'login-error', error: unknown): void
}>()

const handleSubmit = async () => {
  const credentials: LoginCredentials = {
    email: form.value.email,
    password: form.value.password
  }

  emit('login-attempt', form.value)

  try {
    const success = await login(credentials)
    if (success) {
      emit('login-success')
    } else {
      if (error.value && error.value.includes('autenticación')) {
        emit('login-error', new Error('Credenciales inválidas'))
      } else if (error.value && error.value.includes('conexión')) {
        emit('login-error', new Error(
          'Error de conexión. Verifica tu internet e inténtalo de nuevo.'
        ))
      } else {
        emit('login-error', new Error(
          error.value || 'Error desconocido durante el inicio de sesión'
        ))
      }
    }
  } catch (e) {
    console.error('Error durante el login:', e)
    emit('login-error', e)
  }
}
</script>


<style scoped>
/* Contenedor del grupo de formulario con espaciado inferior */
.form-group {
  margin-bottom: 1rem;
  width: 100%;
}

/* Contenedor para implementar etiquetas flotantes */
.float-label-wrapper {
  position: relative;
  width: 100%;
}

/* Estilos base para los campos de entrada */
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

/* Estado de enfoque para los campos de entrada */
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

/* Contenedor para opciones adicionales de la cuenta */
.account-options {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin: 1rem 0;
}

/* Estilos para el enlace de recuperación de contraseña */
.forgot-link {
  color: #4D8CD9;
  text-decoration: none;
  font-size: 14px;
}

/* Estado hover para el enlace de recuperación */
.forgot-link:hover {
  text-decoration: underline;
}

/* Estilo para los mensajes de error */
.error-message {
  color: #FF4D4D;
  margin-bottom: 1rem;
  font-size: 14px;
}
</style>
