<script setup lang="ts">
definePageMeta({
  layout: 'authenticated',
});

import { useAuth } from '~/composables/useAuth';
import { ref, onMounted } from 'vue';

const { user } = useAuth();
const { $api, $authApi } = useNuxtApp();

const seguidores = ref([]);
const seguidos = ref([]);
const posts = ref([]);

const obtenerRelaciones = async () => {
  if (!user.value || !user.value._id) return;

  try {
    const [followsSeguidores, followsSeguidos] = await Promise.all([
      $api(`/users/${user.value._id}/followers`),
      $api(`/users/${user.value._id}/following`),
    ]);

    const detalleSeguidores = await Promise.all(
      followsSeguidores.map(f => $authApi(`/users/${f.seguidorId}`))
    );

    const detalleSeguidos = await Promise.all(
      followsSeguidos.map(f => $authApi(`/users/${f.seguidoId}`))
    );

    seguidores.value = detalleSeguidores;
    seguidos.value = detalleSeguidos;
  } catch (error) {
    console.error("Error al cargar seguidores/seguidos:", error);
  }
};

const obtenerPostsUsuario = async () => {
  if (!user.value || !user.value._id) return;

  try {
    const res = await $api(`/posts/usuario/${user.value._id}`);
    posts.value = res;
  } catch (error) {
    console.error("Error al obtener posts:", error);
  }
};

onMounted(() => {
  obtenerRelaciones();
  obtenerPostsUsuario();
});
</script>

<template>
  <div class="profile-page">
    <!-- PERFIL -->
    <div class="profile-card">
      <div class="left-section">
        <div class="avatar-placeholder">
          <img
            :src="user?.profilePicture || '/assets/images/Usuari.png'"
            alt="Avatar"
            class="avatar-img"
          />
        </div>
        <div class="user-info">
          <h2 class="username">{{ user?.username || 'Usuario' }}</h2>
          <p class="email">{{ user?.email || 'Correo no disponible' }}</p>
          <span class="role">{{ user?.role || 'Estudiante' }}</span>
        </div>
      </div>
      <div class="right-section">
        <button class="edit-button">Editar Perfil</button>
      </div>
    </div>

    <!-- RELACIONES -->
    <div class="relations-container">
      <div class="relation-card">
        <h3 class="relation-title">Seguidores ({{ seguidores.length }})</h3>
        <ul class="relation-list">
          <li
            v-for="seguidor in seguidores"
            :key="seguidor._id"
            class="relation-item"
          >
            <img
              :src="seguidor.profilePicture || '/assets/images/Usuari.png'"
              class="mini-avatar"
            />
            <span class="relation-name">{{ seguidor.username }}</span>
          </li>
        </ul>
      </div>

      <div class="relation-card">
        <h3 class="relation-title">Siguiendo ({{ seguidos.length }})</h3>
        <ul class="relation-list">
          <li
            v-for="seguido in seguidos"
            :key="seguido._id"
            class="relation-item"
          >
            <img
              :src="seguido.profilePicture || '/assets/images/Usuari.png'"
              class="mini-avatar"
            />
            <span class="relation-name">{{ seguido.username }}</span>
          </li>
        </ul>
      </div>
    </div>

    <!-- POSTS -->
    <div class="user-posts-container">
      <h3 class="section-title">Mis Publicaciones ({{ posts.length }})</h3>
      <div v-if="posts.length === 0" class="no-posts">No has publicado nada aún.</div>
      <ul class="posts-list">
        <li v-for="post in posts" :key="post._id" class="post-item">
          <h4>{{ post.titulo }}</h4>
          <p>{{ post.contenido }}</p>
          <span class="post-fecha">{{ new Date(post.fecha).toLocaleString() }}</span>
        </li>
      </ul>
    </div>
  </div>
</template>

<style scoped>
.profile-page {
  padding: 2rem;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.profile-card {
  width: 100%;
  max-width: 900px;
  background-color: #101d2e;
  border-radius: 0.75rem;
  padding: 1.5rem 2rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
  box-shadow: 0 0 12px rgba(0, 0, 0, 0.4);
}

.left-section {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.avatar-placeholder {
  width: 64px;
  height: 64px;
  border-radius: 50%;
  overflow: hidden;
  background-color: #999;
  display: flex;
  align-items: center;
  justify-content: center;
}

.avatar-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.user-info {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.username {
  color: white;
  font-size: 1.2rem;
  margin: 0;
}

.email {
  font-size: 0.9rem;
  color: #ccc;
  margin: 0;
}

.role {
  font-size: 0.8rem;
  background-color: #045a8d;
  color: white;
  padding: 2px 8px;
  border-radius: 12px;
  display: inline-block;
  width: fit-content;
}

.right-section {
  display: flex;
  align-items: center;
}

.edit-button {
  background-color: #00aaff;
  color: white;
  border: none;
  padding: 0.5rem 1rem;
  border-radius: 0.5rem;
  font-weight: 600;
  cursor: pointer;
  transition: background-color 0.3s;
  box-shadow: 0 0 10px #00aaff80;
}

.edit-button:hover {
  background-color: #008fcc;
}

.relations-container {
  display: flex;
  gap: 2rem;
  margin-top: 2rem;
  justify-content: center;
  flex-wrap: wrap;
  width: 100%;
  max-width: 900px;
}

.relation-card {
  background-color: #0d1b2a;
  padding: 1.2rem;
  border-radius: 0.75rem;
  width: 100%;
  max-width: 400px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.4);
}

.relation-title {
  color: #fff;
  margin-bottom: 0.8rem;
  font-size: 1.1rem;
  font-weight: 600;
}

.relation-list {
  list-style: none;
  padding: 0;
  margin: 0;
}

.relation-item {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  padding: 0.4rem 0;
  color: #ddd;
}

.mini-avatar {
  width: 36px;
  height: 36px;
  border-radius: 50%;
  object-fit: cover;
}

.relation-name {
  font-size: 0.95rem;
}

.user-posts-container {
  margin-top: 2rem;
  width: 100%;
  max-width: 900px;
  background-color: #0f2238;
  padding: 1.5rem;
  border-radius: 0.75rem;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.3);
}

.section-title {
  font-size: 1.2rem;
  font-weight: 600;
  color: white;
  margin-bottom: 1rem;
}

.no-posts {
  color: #ccc;
  font-style: italic;
}

.posts-list {
  list-style: none;
  padding: 0;
}

.post-item {
  background-color: #142b43;
  padding: 1rem;
  margin-bottom: 1rem;
  border-radius: 0.5rem;
  color: white;
}

.post-item h4 {
  margin: 0 0 0.5rem 0;
}

.post-item p {
  margin: 0 0 0.5rem 0;
}

.post-fecha {
  font-size: 0.75rem;
  color: #aaa;
}
</style>
