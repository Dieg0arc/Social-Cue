<template>
  <div class="login-container">
    <div class="login-card profile-card">
      <div class="card-header">
        <h1 class="app-title">Perfil de Usuario</h1>
        <p class="app-subtitle">Consulta y edita tu información personal</p>
      </div>

      <div class="card-content profile-content">
        <!-- Columna izquierda: avatar -->
        <div class="card-left profile-avatar-container">
          <img
            src="assets\images\Usuari.png"
            alt="Avatar del usuario"
            class="profile-avatar"
          />
        </div>

        <!-- Columna derecha: datos del usuario -->
        <div class="card-right profile-info">
          <div class="form-group">
            <label for="name">Nombre completo</label>
            <input type="text" id="name" value="Juan Pérez" readonly />
          </div>

          <div class="form-group">
            <label for="email">Correo electrónico</label>
            <input type="email" id="email" value="juan.perez@correo.com" readonly />
          </div>

          <div class="form-group">
            <label for="rol">Rol</label>
            <input type="text" id="rol" value="Estudiante" readonly />
          </div>

          <div class="form-group">
            <label for="codigo">Código institucional</label>
            <input type="text" id="codigo" value="202312345" readonly />
          </div>

          <NuxtLink to="/dashboard" class="login-button profile-back-button">
            Volver al Dashboard
          </NuxtLink>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
// No requiere lógica adicional aún
</script>

<style scoped>
.profile-card {
  max-width: 900px;
  padding: 40px 30px;
}

.profile-content {
  display: flex;
  flex-direction: row;
  gap: 40px;
}

.profile-avatar-container {
  justify-content: center;
}

.profile-avatar {
  width: 200px;
  height: 200px;
  border-radius: 100%;
  border: 4px solid #4D8CD9;
  object-fit: cover;
}

.profile-info {
  flex: 1;
}

.profile-back-button {
  margin-top: 20px;
}

@media (max-width: 768px) {
  .profile-content {
    flex-direction: column;
    align-items: center;
  }

  .profile-avatar {
    width: 150px;
    height: 150px;
  }

  .profile-info {
    width: 100%;
  }
}
</style>
