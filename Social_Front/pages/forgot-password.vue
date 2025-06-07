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
        <button type="submit" class="login-button" :disabled="disabled">
          {{ disabled ? "Espera..." : "Enviar código" }}
        </button>
      </form>

      <p v-if="message" class="info-message">{{ message }}</p>

      <!-- Botón para ir a reset-password -->
      <NuxtLink
        v-if="showResetLink"
        :to="`/reset-password?email=${email}`"
        class="login-button"
        style="margin-top: 10px; display: block; text-align: center;"
      >
        Ya tengo el código, restablecer contraseña
      </NuxtLink>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";

const email = ref("");
const message = ref("");
const disabled = ref(false);
const showResetLink = ref(false);

const requestReset = async () => {
  try {
    const nuxtApp = useNuxtApp();
    const $api = nuxtApp.$api as typeof $fetch;

    disabled.value = true;

    await $api("/auth/password-reset/request", {
      method: "POST",
      body: { email: email.value },
    });

    message.value = "Si el correo existe, recibirás un código";
    showResetLink.value = true;

    // Desactivar por 15 segundos
    setTimeout(() => {
      disabled.value = false;
    }, 15000);
  } catch (err) {
    message.value = "Error al enviar solicitud";
    disabled.value = false;
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
