<template>
  <div class="login-container">
    <div class="login-card">
      <h1>Restablecer contraseña</h1>
      <form @submit.prevent="confirmReset" class="login-form">
        <div class="form-group">
          <label for="email">Correo institucional</label>
          <input type="email" id="email" v-model="email" required class="input-field" />
        </div>
        <div class="form-group">
          <label for="code">Código de verificación</label>
          <input type="text" id="code" v-model="code" required class="input-field" maxlength="6" />
        </div>
        <div class="form-group">
          <label for="password">Nueva contraseña</label>
          <input type="password" id="password" v-model="password" required class="input-field" />
        </div>
      <div class="button-row">
        <button type="submit" class="login-button">Restablecer</button>
        <NuxtLink to="/login" class="login-button">
          Volver
        </NuxtLink>
      </div>
      </form>


      <p v-if="message" :class="{'info-message': true, 'error': isError}">
        {{ message }}
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useRoute, useRouter } from "vue-router";

const route = useRoute();
const router = useRouter();

const email = ref("");
const code = ref("");
const password = ref("");
const message = ref("");
const isError = ref(false);

// Prellenar email si viene como query param
onMounted(() => {
  const queryEmail = route.query.email as string;
  if (queryEmail) {
    email.value = queryEmail;
  }
});

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

    message.value = "¡Contraseña actualizada correctamente!";
    isError.value = false;

    setTimeout(() => {
      router.push("/login");
    }, 2000);
  } catch (err: any) {
    message.value = err?.data?.error || "Error al actualizar la contraseña";
    isError.value = true;
  }
};

const goToLogin = () => {
  router.push("/login");
};
</script>

<style>
@import "~/assets/css/login.css";

.info-message {
  margin-top: 10px;
  text-align: center;
  font-weight: bold;
}

.info-message.error {
  color: red;
}

.info-message:not(.error) {
  color: green;
}

/* Puedes personalizar este estilo si tienes un sistema de clases para botones secundarios */
.secondary-button {
  margin-top: 10px;
  background-color: transparent;
  color: #007bff;
  border: 1px solid #007bff;
}

.button-row {
  display: flex;
  gap: 10px;
  margin-top: 15px;
}

.button-row .login-button {
  flex: 1;
}
</style>
