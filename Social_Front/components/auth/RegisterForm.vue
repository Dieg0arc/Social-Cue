<!-- components/auth/RegisterForm.vue -->
<!--
  RegisterForm - Componente de formulario de registro de usuario
  
  Este componente maneja el proceso completo de registro de nuevos usuarios,
  incluyendo la validación de campos, manejo de errores y estados de carga.
  Integra varios subcomponentes reutilizables del sistema de autenticación
  y utiliza el composable useAuth para la lógica de registro.
-->

<template>
  <form @submit.prevent="handleSubmit">
    <!-- Campo de nombre completo con etiqueta flotante -->
    <div class="form-group">
      <div class="float-label-wrapper" :class="{ 'has-value': form.fullName, 'has-error': errors.fullName }">
        <input 
          id="fullName"
          type="text" 
          v-model="form.fullName" 
          required
          class="input-field"
          :class="{ 'error-input': errors.fullName }"
          @blur="validateField('fullName')"
        />
        <label for="fullName" class="float-label">Nombre completo</label>
        <p v-if="errors.fullName" class="field-error">{{ errors.fullName }}</p>
      </div>
    </div>

    <!-- Campo de nombre de usuario con etiqueta flotante -->
    <div class="form-group">
      <div class="float-label-wrapper" :class="{ 'has-value': form.username, 'has-error': errors.username }">
        <input 
          id="username"
          type="text" 
          v-model="form.username" 
          required
          class="input-field"
          :class="{ 'error-input': errors.username }"
          @blur="validateField('username')"
        />
        <label for="username" class="float-label">Nombre de usuario</label>
        <p v-if="errors.username" class="field-error">{{ errors.username }}</p>
      </div>
    </div>
    
    <!-- Campo de email con etiqueta flotante -->
    <div class="form-group">
      <div class="float-label-wrapper" :class="{ 'has-value': form.email, 'has-error': errors.email }">
        <input 
          id="email"
          type="email" 
          v-model="form.email" 
          required
          class="input-field"
          :class="{ 'error-input': errors.email }"
          @blur="validateField('email')"
        />
        <label for="email" class="float-label">Correo electrónico</label>
        <p v-if="errors.email" class="field-error">{{ errors.email }}</p>
      </div>
    </div>
    
    <!-- Componente de campo de contraseña con funcionalidades específicas -->
    <PasswordField 
      v-model="form.password" 
      label="Contraseña" 
      @blur="validateField('password')"
      :hasError="!!errors.password"
      :errorMessage="errors.password"
    />
    
    <!-- Mensaje de error general condicional -->
    <div v-if="error" class="error-message">
      {{ error }}
    </div>
    
    <!-- Términos y condiciones -->
    <div class="form-group">
      <div class="account-options">
        <div class="remember-container">
          <input 
            type="checkbox" 
            id="terms" 
            v-model="form.acceptTerms"
            @change="validateField('acceptTerms')"
            class="checkbox-input"
          />
          <label for="terms" class="remember-label">
            Acepto los <a href="#" class="forgot-link">términos y condiciones</a>
          </label>
        </div>
      </div>
      <p v-if="errors.acceptTerms" class="field-error">{{ errors.acceptTerms }}</p>
    </div>
    
    <!-- Botón de envío del formulario con estado de carga -->
    <AuthButton
      type="submit"
      :loading="loading"
      text="Crear cuenta"
    />
  </form>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue';
import { useAuth } from '~/composables/useAuth';
import PasswordField from '~/components/auth/PasswordField.vue';
import AuthButton from '~/components/auth/AuthButton.vue';
import type { RegisterData } from '~/types/user';

/**
 * Interfaz para definir la estructura del formulario
 */
interface RegisterFormFields {
  fullName: string;
  username: string;
  email: string;
  password: string;
  acceptTerms: boolean;
}

/**
 * Interfaz para definir la estructura de los errores de validación
 */
interface RegisterFormErrors {
  fullName: string;
  username: string;
  email: string;
  password: string;
  acceptTerms: string;
}

/**
 * Tipo para los nombres de los campos del formulario
 * Esto garantiza seguridad de tipos al acceder a propiedades por nombre
 */
type FormFieldName = keyof RegisterFormErrors;

/**
 * Obtiene las funcionalidades de autenticación del composable useAuth
 * - register: Función para registrar un nuevo usuario
 * - loading: Estado de carga durante el registro
 * - error: Mensaje de error si el registro falla
 */
const { register, loading, error } = useAuth();

/**
 * Estado reactivo del formulario
 */
const form = reactive<RegisterFormFields>({
  fullName: '',
  username: '',
  email: '',
  password: '',
  acceptTerms: false
});

/**
 * Estado reactivo para los errores de validación de cada campo
 */
const errors = reactive<RegisterFormErrors>({
  fullName: '',
  username: '',
  email: '',
  password: '',
  acceptTerms: ''
});

/**
 * Definición de eventos emitidos por el componente
 * @event register-attempt - Emitido cuando el usuario intenta registrarse
 * @event register-success - Emitido cuando el registro es exitoso
 * @event register-error - Emitido cuando ocurre un error durante el registro
 */
const emit = defineEmits<{
  (e: 'register-attempt', form: { fullName: string, username: string, email: string, password: string }): void
  (e: 'register-success'): void
  (e: 'register-error', error: unknown): void
}>();

/**
 * Valida un campo específico del formulario
 * @param field - Nombre del campo a validar
 */
const validateField = (field: FormFieldName) => {
  // Ahora TypeScript sabe que field es una clave válida de errors
  errors[field] = '';

  switch(field) {
    case 'fullName':
      if (!form.fullName) {
        errors.fullName = 'El nombre completo es obligatorio';
      } else if (form.fullName.length < 3) {
        errors.fullName = 'El nombre debe tener al menos 3 caracteres';
      }
      break;
    
    case 'username':
      if (!form.username) {
        errors.username = 'El nombre de usuario es obligatorio';
      } else if (form.username.length < 3) {
        errors.username = 'El nombre de usuario debe tener al menos 3 caracteres';
      } else if (!/^[a-zA-Z0-9_]+$/.test(form.username)) {
        errors.username = 'Solo letras, números y guiones bajos';
      }
      break;
    
    case 'email':
      if (!form.email) {
        errors.email = 'El correo electrónico es obligatorio';
      } else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(form.email)) {
        errors.email = 'Ingresa un correo electrónico válido';
      }
      break;
    
    case 'password':
      if (!form.password) {
        errors.password = 'La contraseña es obligatoria';
      } else if (form.password.length < 8) {
        errors.password = 'La contraseña debe tener al menos 8 caracteres';
      } else if (!/(?=.*[A-Z])/.test(form.password)) {
        errors.password = 'Debe incluir al menos una letra mayúscula';
      } else if (!/(?=.*[0-9])/.test(form.password)) {
        errors.password = 'Debe incluir al menos un número';
      }
      break;
    
    case 'acceptTerms':
      if (!form.acceptTerms) {
        errors.acceptTerms = 'Debes aceptar los términos y condiciones';
      }
      break;
  }
  
  return !errors[field];
};

/**
 * Valida todos los campos del formulario
 * @returns {boolean} - true si todos los campos son válidos, false en caso contrario
 */
/**
 * Valida todos los campos del formulario
 * @returns {boolean} - true si todos los campos son válidos, false en caso contrario
 */
const validateForm = () => {
  let isValid = true;
  
  // Definimos un array tipado de nombres de campos
  const fieldNames: FormFieldName[] = ['fullName', 'username', 'email', 'password', 'acceptTerms'];
  
  fieldNames.forEach(field => {
    if (!validateField(field)) {
      isValid = false;
    }
  });
  
  return isValid;
};
/**
 * Maneja el envío del formulario de registro
 * - Valida los campos del formulario
 * - Emite evento de intento de registro
 * - Intenta registrar al usuario con los datos proporcionados
 * - Emite evento de éxito o error según el resultado
 */
const handleSubmit = async () => {
  if (!validateForm()) {
    return;
  }
  
  emit('register-attempt', {
    fullName: form.fullName,
    username: form.username,
    email: form.email,
    password: form.password
  });
  
  try {
    // Creamos un objeto de registro tipado con RegisterData
    const registerData: RegisterData = {
      email: form.email,
      password: form.password,
      username: form.username,
      fullName: form.fullName
    };
    
    // Llamamos a la función register con los datos tipados y verificamos el resultado
    const success = await register(registerData);
    
    if (success) {
      emit('register-success');
    } else {
      // Si register devuelve falso, determinamos el tipo de error basado en el error del store
      if (error.value && error.value.includes('correo')) {
        emit('register-error', new Error('Este correo electrónico ya está registrado'));
      } else if (error.value && error.value.includes('usuario')) {
        emit('register-error', new Error('Este nombre de usuario ya está en uso'));
      } else if (error.value && error.value.includes('conexión')) {
        emit('register-error', new Error('Error de conexión. Verifica tu internet e inténtalo de nuevo.'));
      } else {
        emit('register-error', new Error(error.value || 'Error desconocido durante el registro'));
      }
    }
  } catch (e) {
    console.error('Error durante el registro:', e);
    emit('register-error', e);
  }
};
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

/* Estilo para campos con error */
.error-input {
  border-color: #FF4D4D;
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

/* Cuando el campo tiene error, cambia el color de la etiqueta */
.has-error .float-label {
  color: #FF4D4D;
}

/* Estilo para los mensajes de error */
.error-message {
  color: #FF4D4D;
  margin-bottom: 1rem;
  font-size: 14px;
}

/* Estilo para los mensajes de error de campo específico */
.field-error {
  color: #FF4D4D;
  font-size: 12px;
  margin-top: 4px;
  margin-bottom: 0;
}

/* Contenedor para opciones adicionales de la cuenta */
.account-options {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin: 1rem 0;
}

/* Contenedor para la opción de recordar usuario */
.remember-container {
  display: flex;
  align-items: center;
}

/* Estilos para el checkbox */
.checkbox-input {
  margin-right: 8px;
  width: 16px;
  height: 16px;
  cursor: pointer;
}

/* Estilo para la etiqueta de recordar */
.remember-label {
  display: flex;
  align-items: center;
  cursor: pointer;
  font-size: 14px;
  color: #D0D0D0;
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
</style>
