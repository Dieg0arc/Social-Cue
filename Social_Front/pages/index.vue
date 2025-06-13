<template>
  <div class="login-container">
    <div class="login-card home-card">
      <div class="card-header">
        <h1 class="app-title">Social-Cue</h1>
        <p class="app-subtitle">La red social del conocimiento</p>
      </div>
      <div class="card-content home-content">
        <div class="card-right home-content-center">
          <NuxtLink to="http://localhost:3000/login" class="login-button home-login-button">
            Iniciar sesión
          </NuxtLink>
          <!-- <p class="home-description">
            Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque nec est at libero ultrices suscipit.
          </p> -->
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
// No requiere lógica adicional
</script>

<style scoped>
/* Reutilizando tu sistema de diseño */
.login-container {
  display: flex;
  width: 100vw;
  height: 100vh;
  justify-content: center;
  align-items: center;
  background-color: #0A1625;
  color: #FFFFFF;
  position: relative;
  overflow: hidden;
}

.login-card.home-card {
  width: 90%;
  max-width: 600px;
  background-color: #0A1625;
  border-radius: 20px;
  display: flex;
  flex-direction: column;
  padding: 40px 30px;
  align-items: center;
  justify-content: center;
  box-shadow: 0 8px 30px rgba(0, 86, 166, 0.4);
}

.card-header {
  text-align: center;
  margin-bottom: 30px;
}

.app-title {
  font-size: 2.5rem;
  color: #FFFFFF;
  margin: 0;
  font-weight: 600;
}

.app-subtitle {
  color: #D0D0D0;
  font-size: 1.2rem;
  margin-top: 10px;
}

.card-content.home-content {
  justify-content: center;
}

.card-right.home-content-center {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 100%;
}

.login-button.home-login-button {
  width: 100%;
  max-width: 300px;
  padding: 16px;
  background-color: #0056A6;
  color: white;
  border: none;
  border-radius: 25px;
  font-weight: 600;
  font-size: 18px;
  cursor: pointer;
  transition: background-color 0.2s;
  text-align: center;
  margin-bottom: 24px;
}

.login-button.home-login-button:hover {
  background-color: #4D8CD9;
}

.home-description {
  text-align: center;
  max-width: 450px;
  color: #D0D0D0;
  font-size: 1rem;
  line-height: 1.5;
}
</style>
