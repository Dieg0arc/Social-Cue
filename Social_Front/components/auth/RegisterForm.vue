<template>
  <form @submit.prevent="handleSubmit">
    <!-- Campo de nombre completo con etiqueta flotante -->
    <div class="form-group">
      <div class="float-label-wrapper" :class="{ 'has-value': form.fullName }">
        <input
          id="fullName"
          type="text"
          v-model="form.fullName"
          required
          class="input-field"
        />
        <label for="fullName" class="float-label">Nombre completo</label>
      </div>
    </div>

    <!-- Campo de nombre de usuario con etiqueta flotante -->
    <div class="form-group">
      <div class="float-label-wrapper" :class="{ 'has-value': form.username }">
        <input
          id="username"
          type="text"
          v-model="form.username"
          required
          class="input-field"
        />
        <label for="username" class="float-label">Nombre de usuario</label>
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
        />
        <label for="email" class="float-label">Correo electrónico</label>
        <p v-if="errors.email" class="field-error">{{ errors.email }}</p>
      </div>
    </div>

    <!-- Campo de contraseña con componente reutilizable -->
    <PasswordField
      v-model="form.password"
      label="Contraseña"
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
            class="checkbox-input"
          />
          <label for="terms" class="remember-label">
            Acepto los <a href="#" class="forgot-link">términos y condiciones</a>
          </label>
        </div>
      </div>
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
import { reactive } from 'vue';
import { useAuth } from '~/composables/useAuth';
import PasswordField from '~/components/auth/PasswordField.vue';
import AuthButton from '~/components/auth/AuthButton.vue';
import type { RegisterData } from '~/types/user';

const { register, loading, error } = useAuth();

const form = reactive({
  fullName: '',
  username: '',
  email: '',
  password: '',
  acceptTerms: false
});

const errors = reactive({
  email: ''
});

const validateEmail = (email: string): boolean => {
  return /@(?:cue\.edu\.co|unihumboldt\.edu\.co)$/.test(email);
};

const handleSubmit = async () => {
  errors.email = '';

  if (!form.acceptTerms) {
    alert('Debes aceptar los términos y condiciones.');
    return;
  }

  if (!validateEmail(form.email)) {
    errors.email = 'El correo debe terminar en @cue.edu.co o @unihumboldt.edu.co';
    return;
  }

  const registerData: RegisterData = {
    fullName: form.fullName,
    username: form.username,
    email: form.email,
    password: form.password
  };

  const success = await register(registerData);
  if (success) {
    navigateTo("/dashboard");
  }
};
</script>

<style scoped>
.form-group {
  margin-bottom: 1rem;
  width: 100%;
}

.float-label-wrapper {
  position: relative;
  width: 100%;
}

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

.input-field:focus {
  outline: none;
  border-color: #4D8CD9;
}

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

.input-field:focus ~ .float-label,
.float-label-wrapper.has-value .float-label {
  top: 25%;
  font-size: 12px;
  color: #4D8CD9;
}

.float-label-wrapper.has-error .float-label {
  color: #FF4D4D;
}

.input-field.error-input,
.float-label-wrapper.has-error .input-field {
  border-color: #FF4D4D;
}

.field-error {
  color: #FF4D4D;
  font-size: 12px;
  margin-top: 4px;
  margin-bottom: 0;
}

.error-message {
  color: #FF4D4D;
  margin-bottom: 1rem;
  font-size: 14px;
}

.account-options {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin: 1rem 0;
}

.remember-container {
  display: flex;
  align-items: center;
}

.checkbox-input {
  margin-right: 8px;
  width: 16px;
  height: 16px;
  cursor: pointer;
}

.remember-label {
  display: flex;
  align-items: center;
  cursor: pointer;
  font-size: 14px;
  color: #D0D0D0;
}

.forgot-link {
  color: #4D8CD9;
  text-decoration: none;
  font-size: 14px;
}

.forgot-link:hover {
  text-decoration: underline;
}
</style>
