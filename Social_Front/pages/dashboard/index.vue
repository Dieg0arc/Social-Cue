<!-- 
  Archivo: dashboard/index.vue
  Ubicación: pages/dashboard/index.vue
  Descripción: Página principal del dashboard para SocialCue.
  Esta página muestra el feed de publicaciones similar a Instagram.
-->

<template>
  <div class="dashboard-container">
    <!-- Contenedor principal del feed de posts -->
    <div class="feed-container">
      <!-- Sección de historias (simulado) -->
<div class="stories-container">
  <div class="stories-header">
    <h3>Historias</h3>
    <button class="view-all-btn">Ver todas</button>
  </div>
  <div class="stories-list">
    <div 
      v-for="story in stories" 
      :key="story.id" 
      class="story-item"
      :class="{ 'has-new': story.hasNew }"
    >
      <div class="story-avatar-container">
        <img :src="story.avatar" :alt="story.username" class="story-avatar" />
      </div>
      <span class="story-username">{{ story.username }}</span>
    </div>
  </div>
</div>


      <!-- Tarjeta de creación de post -->
      <div class="create-post-card">
        <div class="create-post-header">
          <img 
            :src="currentUser.avatar" 
            :alt="currentUser.username" 
            class="user-avatar"
          />
          <input 
            type="text" 
            placeholder="¿Qué estás pensando?" 
            class="create-post-input"
            @click="showCreatePostModal = true"
          />
        </div>
        <div class="create-post-actions">
          <button class="create-action-btn">
            <span class="action-icon">📷</span>
            <span class="action-text">Foto</span>
          </button>
          <button class="create-action-btn">
            <span class="action-icon">🎬</span>
            <span class="action-text">Video</span>
          </button>
          <button class="create-action-btn">
            <span class="action-icon">📍</span>
            <span class="action-text">Ubicación</span>
          </button>
        </div>
      </div>

      <!-- Lista de posts -->
      <div class="posts-list" v-if="posts.length > 0">
        <PostCard 
          v-for="post in posts" 
          :key="post.id" 
          :post="post"
          @like="handleLike"
          @comment="handleComment"
        />
        
        <!-- Indicador de carga para el scroll infinito -->
        <div v-if="loading" class="loading-indicator">
          <div class="loading-spinner"></div>
          <p>Cargando más publicaciones...</p>
        </div>
      </div>
      <div v-else class="empty-state">
        <p>No hay publicaciones para mostrar</p>
        <button class="find-friends-btn">Encuentra amigos para seguir</button>
      </div>
    </div>

    <!-- Sidebar de sugerencias (solo visible en pantallas grandes) -->
    <div class="suggestions-sidebar">
      <!-- Perfil del usuario -->
      <div class="user-profile-card">
        <img :src="currentUser.avatar" :alt="currentUser.username" class="profile-avatar" />
        <div class="profile-info">
          <h3 class="profile-username">{{ currentUser.username }}</h3>
          <p class="profile-name">{{ currentUser.fullName }}</p>
        </div>
        <button class="switch-btn">Cambiar</button>
      </div>

      <!-- Sugerencias de amistad -->
      <div class="suggestions-container">
        <div class="suggestions-header">
          <h4>Sugerencias para ti</h4>
          <button class="view-all-btn">Ver todo</button>
        </div>
        <ul class="suggestions-list">
          <li 
            v-for="suggestion in suggestions" 
            :key="suggestion.id" 
            class="suggestion-item"
          >
            <div class="suggestion-user">
              <img :src="suggestion.avatar" :alt="suggestion.username" class="suggestion-avatar" />
              <div class="suggestion-info">
                <p class="suggestion-username">{{ suggestion.username }}</p>
                <p class="suggestion-meta">{{ suggestion.meta }}</p>
              </div>
            </div>
            <button class="follow-btn">Seguir</button>
          </li>
        </ul>
      </div>

      <!-- Footer con enlaces -->
      <div class="sidebar-footer">
        <nav class="footer-links">
          <a href="#" class="footer-link">Acerca de</a>
          <a href="#" class="footer-link">Ayuda</a>
          <a href="#" class="footer-link">Privacidad</a>
          <a href="#" class="footer-link">Términos</a>
        </nav>
        <p class="copyright">© 2025 SocialCue</p>
      </div>
    </div>

    <!-- Modal para crear post (simulado) -->
    <div v-if="showCreatePostModal" class="create-post-modal">
      <div class="modal-overlay" @click="showCreatePostModal = false"></div>
      <div class="modal-content">
        <div class="modal-header">
          <h3>Crear publicación</h3>
          <button class="close-modal-btn" @click="showCreatePostModal = false">✕</button>
        </div>
        <div class="modal-body">
          <div class="post-form">
            <textarea
              placeholder="¿Qué quieres compartir?"
              class="post-textarea"
              v-model="newPostText"
            ></textarea>
            <div class="post-form-actions">
              <button class="attach-btn">Adjuntar foto</button>
              <button 
                class="publish-btn"
                :disabled="!newPostText.trim()"
                @click="createPost"
              >
                Publicar
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue';
import PostCard from '~/components/content/PostCard.vue';

// Define el layout autenticado para esta página
definePageMeta({
  layout: 'authenticated'
});

// Estado para controlar la creación de posts
const showCreatePostModal = ref(false);
const newPostText = ref('');

// Estado para el scroll infinito
const loading = ref(false);
const page = ref(1);

// Simula el usuario actual
const currentUser = {
  id: 1,
  username: 'usuario_actual',
  fullName: 'Usuario Actual',
  avatar: 'assets\images\Usuari.png'
};

// Datos de ejemplo para historias
const stories = ref([
  { id: 1, username: 'maria_rodriguez', avatar: '/assets/images/avatar1.jpg', hasNew: true },
  { id: 2, username: 'juan_perez', avatar: '/assets/images/avatar2.jpg', hasNew: true },
  { id: 3, username: 'ana_gomez', avatar: '/assets/images/avatar3.jpg', hasNew: false },
  { id: 4, username: 'carlos_lopez', avatar: '/assets/images/avatar4.jpg', hasNew: true },
  { id: 5, username: 'laura_martinez', avatar: '/assets/images/avatar5.jpg', hasNew: false }
]);

// Datos de ejemplo para sugerencias
const suggestions = ref([
  { id: 1, username: 'pablo_sanchez', fullName: 'Pablo Sánchez', avatar: '/assets/images/avatar6.jpg', meta: 'Seguido por maria_rodriguez' },
  { id: 2, username: 'elena_torres', fullName: 'Elena Torres', avatar: '/assets/images/avatar7.jpg', meta: 'Nuevo en SocialCue' },
  { id: 3, username: 'miguel_fernandez', fullName: 'Miguel Fernández', avatar: '/assets/images/avatar8.jpg', meta: 'Seguido por juan_perez' }
]);

// Datos de ejemplo para posts
const posts = ref([
  {
    id: 1,
    author: {
      id: 2,
      username: 'maria_rodriguez',
      avatar: '/assets/images/avatar1.jpg'
    },
    mediaUrl: '/assets/images/post1.jpg',
    caption: 'Disfrutando de un día increíble en la playa. El sol, la arena y el mar... ¡Qué más se puede pedir! #Verano #Playa #Relax',
    location: 'Playa del Carmen',
    likes: 120,
    liked: false,
    comments: [
      { username: 'juan_perez', text: '¡Qué envidia! Disfruta mucho 🏖️' },
      { username: 'ana_gomez', text: 'Hermoso lugar, las olas se ven geniales' }
    ],
    createdAt: '2025-05-25T14:30:00'
  },
  {
    id: 2,
    author: {
      id: 3,
      username: 'juan_perez',
      avatar: '/assets/images/avatar2.jpg'
    },
    mediaUrl: '/assets/images/post2.jpg',
    caption: 'Nuevo proyecto terminado. Ha sido un desafío pero el resultado valió la pena. #Programación #Desarrollo #SocialCue',
    likes: 85,
    liked: true,
    comments: [
      { username: 'carlos_lopez', text: 'Excelente trabajo, ¿qué tecnologías usaste?' },
      { username: 'maria_rodriguez', text: 'Se ve increíble, felicitaciones!' },
      { username: 'ana_gomez', text: 'Quiero saber más sobre este proyecto' }
    ],
    createdAt: '2025-05-24T10:15:00'
  },
  {
    id: 3,
    author: {
      id: 4,
      username: 'ana_gomez',
      avatar: '/assets/images/avatar3.jpg'
    },
    textContent: 'Hoy comienza una nueva etapa en mi vida profesional. Muy emocionada por los nuevos retos que vendrán. #NuevoTrabajo #Oportunidades',
    likes: 45,
    liked: false,
    comments: [
      { username: 'laura_martinez', text: '¡Felicitaciones! Te mereces lo mejor 🎉' }
    ],
    createdAt: '2025-05-23T08:45:00'
  }
]);

/**
 * Maneja la acción de dar like a un post
 */
const handleLike = (postId: string | number) => {
  console.log(`Like en post ${postId}`);
  // En una implementación real, aquí se haría una petición a la API
};

/**
 * Maneja la acción de comentar en un post
 */
const handleComment = (postId: string | number, comment: string) => {
  console.log(`Nuevo comentario en post ${postId}: ${comment}`);
  // En una implementación real, aquí se haría una petición a la API
};

/**
 * Crea un nuevo post (simulado)
 */
const createPost = () => {
  if (!newPostText.value.trim()) return;
  
  // Simulamos la creación de un post
  const newPost = {
    id: Date.now(),
    author: {
      id: currentUser.id,
      username: currentUser.username,
      avatar: currentUser.avatar
    },
    textContent: newPostText.value,
    likes: 0,
    liked: false,
    comments: [],
    createdAt: new Date().toISOString()
  };
  
  // Añadimos el post al inicio del array
  posts.value.unshift(newPost);
  
  // Cerramos el modal y limpiamos el texto
  showCreatePostModal.value = false;
  newPostText.value = '';
};

/**
 * Función para cargar más posts (simulado)
 */
const loadMorePosts = () => {
  if (loading.value) return;
  
  loading.value = true;
  
  // Simulamos una petición a la API
  setTimeout(() => {
    // Nuevos posts de ejemplo
    const newPosts = [
      {
        id: 4 + page.value,
        author: {
          id: 5,
          username: 'carlos_lopez',
          avatar: '/assets/images/avatar4.jpg'
        },
        mediaUrl: '/assets/images/post3.jpg',
        caption: 'Explorando nuevos lugares. La arquitectura de esta ciudad es impresionante. #Viajes #Arquitectura',
        location: 'Madrid, España',
        likes: 67,
        liked: false,
        comments: [
          { username: 'maria_rodriguez', text: '¡Qué hermoso lugar! Tengo que visitarlo' }
        ],
        createdAt: `2025-05-${20 - page.value}T16:20:00`
      },
      {
        id: 5 + page.value,
        author: {
          id: 6,
          username: 'laura_martinez',
          avatar: '/assets/images/avatar5.jpg'
        },
        textContent: 'Acabo de terminar de leer un libro increíble. Totalmente recomendado para los amantes de la ciencia ficción. #Lectura #Libros',
        likes: 32,
        liked: false,
        comments: [],
        createdAt: `2025-05-${19 - page.value}T09:10:00`
      }
    ];
    
    // Añadimos los nuevos posts al final del array
    posts.value = [...posts.value, ...newPosts];
    
    // Actualizamos el estado
    loading.value = false;
    page.value += 1;
  }, 1500);
};

/**
 * Función para detectar cuando el usuario llega al final de la página
 */
const handleScroll = () => {
  const scrollPosition = window.innerHeight + window.scrollY;
  const documentHeight = document.body.offsetHeight;
  
  // Si el usuario está cerca del final de la página, cargamos más posts
  if (scrollPosition >= documentHeight - 500 && !loading.value) {
    loadMorePosts();
  }
};

// Agregamos el evento de scroll al montar el componente
onMounted(() => {
  window.addEventListener('scroll', handleScroll);
});

// Eliminamos el evento al desmontar el componente
onUnmounted(() => {
  window.removeEventListener('scroll', handleScroll);
});
</script>

<style scoped>
/* Contenedor principal del dashboard */
.dashboard-container {
  display: flex;
  width: 100%;
  gap: 2rem;
}

/* Contenedor del feed */
.feed-container {
  flex: 1;
  max-width: 600px;
  margin: 0 auto;
}

/* Contenedor de historias */
.stories-container {
  background-color: #0A1625;
  border-radius: 8px;
  border: 1px solid #243447;
  padding: 1rem;
  margin-bottom: 1rem;
  overflow: hidden;
}

.stories-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1rem;
}

.stories-header h3 {
  margin: 0;
  font-size: 1rem;
  font-weight: 600;
}

.view-all-btn {
  background: transparent;
  border: none;
  color: #4D8CD9;
  cursor: pointer;
  font-size: 0.875rem;
}

.stories-list {
  display: flex;
  gap: 1rem; /* Espacio entre historias */
  overflow-x: auto;
  padding-bottom: 0.5rem;
}

.story-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: flex-start;
  min-width: 70px; /* Aumentado para evitar colisiones */
  max-width: 70px;
  text-align: center;
  flex: 0 0 auto;
}

.story-avatar-container {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  padding: 2px;
  background: linear-gradient(45deg, #FCAF45, #FD1D1D, #833AB4);
  margin-bottom: 0.5rem;
  display: flex;
  align-items: center;
  justify-content: center;
}

.story-item:not(.has-new) .story-avatar-container {
  background: #243447;
}

.story-avatar {
  width: 100%;
  height: 100%;
  border-radius: 50%;
  object-fit: cover;
  border: 2px solid #0A1625;
}

.story-username {
  font-size: 0.75rem;
  color: #D0D0D0;
  max-width: 100%;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

/* Tarjeta de creación de post */
.create-post-card {
  background-color: #0A1625;
  border-radius: 8px;
  border: 1px solid #243447;
  padding: 1rem;
  margin-bottom: 1rem;
}

.create-post-header {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  margin-bottom: 1rem;
}

.user-avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  object-fit: cover;
}

.create-post-input {
  flex: 1;
  background-color: #243447;
  border: none;
  border-radius: 20px;
  padding: 0.75rem 1rem;
  color: #FFFFFF;
  cursor: pointer;
}

.create-post-input::placeholder {
  color: #8A97A8;
}

.create-post-actions {
  display: flex;
  justify-content: space-around;
  border-top: 1px solid #243447;
  padding-top: 1rem;
}

.create-action-btn {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  background: transparent;
  border: none;
  color: #D0D0D0;
  cursor: pointer;
  padding: 0.5rem;
  transition: color 0.2s;
}

.create-action-btn:hover {
  color: #FFFFFF;
}

.action-icon {
  font-size: 1.25rem;
}

/* Lista de posts */
.posts-list {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 1rem;
}

/* Estado vacío */
.empty-state {
  background-color: #122237;
  border-radius: 8px;
  border: 1px solid #243447;
  padding: 2rem;
  text-align: center;
  margin-top: 2rem;
}

.empty-state p {
  margin-bottom: 1rem;
  color: #8A97A8;
}

.find-friends-btn {
  background-color: #0056A6;
  color: white;
  border: none;
  border-radius: 25px;
  padding: 0.75rem 1.5rem;
  font-weight: 600;
  cursor: pointer;
  transition: background-color 0.2s;
}

.find-friends-btn:hover {
  background-color: #4D8CD9;
}

/* Indicador de carga */
.loading-indicator {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin: 2rem 0;
  color: #8A97A8;
}

.loading-spinner {
  width: 30px;
  height: 30px;
  border: 3px solid rgba(77, 140, 217, 0.3);
  border-radius: 50%;
  border-top-color: #4D8CD9;
  animation: spin 1s ease-in-out infinite;
  margin-bottom: 0.5rem;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

/* Sidebar de sugerencias */
.suggestions-sidebar {
  width: 300px;
  display: none; /* Oculto por defecto en móviles */
}

.user-profile-card {
  display: flex;
  align-items: center;
  margin-bottom: 1.5rem;
}

.profile-avatar {
  width: 56px;
  height: 56px;
  border-radius: 50%;
  object-fit: cover;
  margin-right: 1rem;
}

.profile-info {
  flex: 1;
}

.profile-username {
  margin: 0;
  font-size: 0.875rem;
  font-weight: 600;
  color: #FFFFFF;
}

.profile-name {
  margin: 0;
  font-size: 0.75rem;
  color: #8A97A8;
}

.switch-btn {
  background: transparent;
  border: none;
  color: #4D8CD9;
  font-weight: 600;
  cursor: pointer;
  font-size: 0.75rem;
}

/* Sugerencias */
.suggestions-container {
  background-color: #0A1625;
  border-radius: 8px;
  border: 1px solid #243447;
  padding: 1rem;
  margin-bottom: 1.5rem;
}

.suggestions-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1rem;
}

.suggestions-header h4 {
  margin: 0;
  font-size: 0.875rem;
  font-weight: 600;
  color: #8A97A8;
}

.suggestions-list {
  list-style: none;
  padding: 0;
  margin: 0;
}

.suggestion-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 0.75rem;
}

.suggestion-user {
  display: flex;
  align-items: center;
  gap: 0.75rem;
}

.suggestion-avatar {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  object-fit: cover;
}

.suggestion-info p {
  margin: 0;
}

.suggestion-username {
  font-size: 0.75rem;
  font-weight: 600;
  color: #FFFFFF;
}

.suggestion-meta {
  font-size: 0.6875rem;
  color: #8A97A8;
}

.follow-btn {
  background: transparent;
  border: none;
  color: #4D8CD9;
  font-weight: 600;
  cursor: pointer;
  font-size: 0.75rem;
}

/* Footer del sidebar */
.sidebar-footer {
  margin-top: 1.5rem;
  color: #8A97A8;
  font-size: 0.75rem;
}

.footer-links {
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
  margin-bottom: 0.5rem;
}

.footer-link {
  color: #8A97A8;
  text-decoration: none;
}

.footer-link:hover {
  text-decoration: underline;
}

.copyright {
  margin: 0;
}

/* Modal de creación de post */
.create-post-modal {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.modal-overlay {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.8);
}

.modal-content {
  position: relative;
  width: 90%;
  max-width: 600px;
  background-color: #122237;
  border-radius: 8px;
  overflow: hidden;
  z-index: 1001;
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem;
  border-bottom: 1px solid #243447;
}

.modal-header h3 {
  margin: 0;
  font-size: 1.25rem;
}

.close-modal-btn {
  background: transparent;
  border: none;
  color: #D0D0D0;
  font-size: 1.5rem;
  cursor: pointer;
}

.modal-body {
  padding: 1rem;
}

.post-form {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.post-textarea {
  width: 100%;
  min-height: 120px;
  background-color: #243447;
  border: none;
  border-radius: 8px;
  padding: 1rem;
  color: #FFFFFF;
  resize: none;
}

.post-form-actions {
  display: flex;
  justify-content: space-between;
}

.attach-btn {
  background-color: transparent;
  border: 1px solid #243447;
  border-radius: 25px;
  padding: 0.5rem 1rem;
  color: #D0D0D0;
  cursor: pointer;
}

.publish-btn {
  background-color: #0056A6;
  border: none;
  border-radius: 25px;
  padding: 0.5rem 1.5rem;
  color: white;
  font-weight: 600;
  cursor: pointer;
}

.publish-btn:disabled {
  background-color: #243447;
  color: #8A97A8;
  cursor: not-allowed;
}

/* Responsive */
@media (min-width: 1024px) {
  .suggestions-sidebar {
    display: block;
  }
}

@media (max-width: 768px) {
  .create-action-btn .action-text {
    display: none;
  }
}
</style>

