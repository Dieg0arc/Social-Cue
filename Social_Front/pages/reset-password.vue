<template>
  <div class="login-container">
    <div class="login-card">
      <h1>Ingresar código</h1>
      <form @submit.prevent="confirmReset" class="login-form">
        <div class="form-group">
          <label for="email">Correo institucional</label>
          <input
            type="email"
            id="email"
            v-model="email"
            required
            class="input-field"
          />
        </div>
        <div class="form-group">
          <label for="code">Código</label>
          <input
            type="text"
            id="code"
            v-model="code"
            required
            class="input-field"
          />
        </div>
        <div class="form-group">
          <label for="password">Nueva contraseña</label>
          <input
            type="password"
            id="password"
            v-model="password"
            required
            class="input-field"
          />
        </div>
        <button type="submit" class="login-button">Restablecer</button>
      </form>
      <p v-if="message" class="info-message">{{ message }}</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";

const email = ref("");
const code = ref("");
const password = ref("");
const message = ref("");

const confirmReset = async () => {
  try {
    const nuxtApp = useNuxtApp();
    const $api = nuxtApp.$api as typeof $fetch;
    await $api("/auth/password-reset/confirm", {
      method: "POST",
      body: {
        email: email.value,
        code: code.value,
        new_password: password.value,
      },
    });
    message.value = "Contraseña actualizada";
  } catch (err) {
    message.value = "Error al actualizar contraseña";
  }
};
</script>

<style>
@import "~/assets/css/login.css";
.info-message {
  margin-top: 10px;
  text-align: center;
}
</style>
