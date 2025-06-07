<template>
  <div class="login-container">
    <div class="login-card">
      <h1>Recuperar contraseña</h1>
      <form @submit.prevent="requestReset" class="login-form">
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
        <button type="submit" class="login-button">Enviar código</button>
      </form>
      <p v-if="message" class="info-message">{{ message }}</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";

const email = ref("");
const message = ref("");

const requestReset = async () => {
  try {
    const nuxtApp = useNuxtApp();
    const $api = nuxtApp.$api as typeof $fetch;
    await $api("/auth/password-reset/request", {
      method: "POST",
      body: { email: email.value },
    });
    message.value = "Si el correo existe, recibirás un código";
  } catch (err) {
    message.value = "Error al enviar solicitud";
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
