<!-- 
  Archivo: CreatePostModal.vue
  Ubicación: components/dashboard/CreatePostModal.vue
  Descripción: Componente de modal para la creación de nuevos posts.
  Permite al usuario crear publicaciones de texto o con imágenes.
-->

<template>
  <div class="create-post-modal">
    <!-- Overlay de fondo para cerrar el modal al hacer clic fuera -->
    <div class="modal-overlay" @click="$emit('close')"></div>
    
    <!-- Contenido principal del modal -->
    <div class="modal-content">
      <!-- Cabecera del modal -->
      <div class="modal-header">
        <h3>Crear publicación</h3>
        <button class="close-modal-btn" @click="$emit('close')">✕</button>
      </div>
      
      <!-- Cuerpo del modal -->
      <div class="modal-body">
        <!-- Información del usuario -->
        <div class="user-info" v-if="currentUser">
            <img :src="mediaPreviewUrl || ''" alt="Vista previa" class="preview-image" />
          <div class="user-details">
            <p class="username">{{ currentUser.username }}</p>
            <select v-if="showVisibilityOptions" v-model="visibility" class="visibility-select">
              <option value="public">Público</option>
              <option value="friends">Amigos</option>
              <option value="private">Solo yo</option>
            </select>
          </div>
        </div>
        
        <!-- Formulario de creación de post -->
        <div class="post-form">
          <!-- Textarea para el contenido del post -->
          <textarea
            ref="postTextarea"
            v-model="postContent"
            placeholder="¿Qué quieres compartir?"
            class="post-textarea"
            :class="{ 'with-media': selectedMedia }"
            @keydown="handleKeydown"
          ></textarea>
          
          <!-- Vista previa de la imagen seleccionada -->
          <div v-if="selectedMedia" class="media-preview">
            <button class="remove-media-btn" @click="removeMedia">✕</button>
            <img :src="mediaPreviewUrl" alt="Vista previa" class="preview-image" />
          </div>
          
          <!-- Campo de ubicación opcional -->
          <div v-if="showLocationField" class="location-field">
            <input 
              type="text" 
              v-model="location" 
              placeholder="Añade una ubicación" 
              class="location-input"
            />
          </div>
          
          <!-- Mensajes de error -->
          <div v-if="error" class="error-message">
            {{ error }}
          </div>
          
          <!-- Acciones del formulario -->
          <div class="post-form-actions">
            <!-- Opciones de adjuntar contenido -->
            <div class="attach-options">
              <label class="attach-btn">
                <input 
                  type="file" 
                  accept="image/*" 
                  class="file-input" 
                  @change="handleFileSelection"
                />
                <span class="attach-icon">📷</span>
              </label>
              <button 
                class="attach-btn"
                @click.prevent="toggleLocationField"
              >
                <span class="attach-icon">📍</span>
              </button>
            </div>
            
            <!-- Botón de publicar -->
            <button 
              class="publish-btn"
              :disabled="!isValidPost"
              @click="createPost"
            >
              Publicar
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch, nextTick } from 'vue';
import type { PropType } from 'vue';
import type { CurrentUser } from '~/types/dashboard';

/**
 * Props del componente
 */
const props = defineProps({
  /**
   * Valor del texto del post (two-way binding)
   */
  modelValue: {
    type: String,
    default: ''
  },
  
  /**
   * Información del usuario actual
   */
  currentUser: {
    type: Object as PropType<CurrentUser>,
    required: true
  },
  
  /**
   * Tipo inicial de post a crear
   */
  initialType: {
    type: String,
    default: 'text'
  }
});

/**
 * Eventos emitidos por el componente
 */
const emit = defineEmits<{
  /**
   * Actualización del valor del modelo (v-model)
   */
  (e: 'update:modelValue', value: string): void;
  
  /**
   * Emitido cuando el usuario quiere cerrar el modal
   */
  (e: 'close'): void;
  
  /**
   * Emitido cuando el usuario crea un nuevo post
   * @param content Contenido del post
   * @param mediaFile Archivo de imagen seleccionado (si existe)
   * @param location Ubicación (si se especificó)
   */
  (e: 'create', content: string, mediaFile?: File, location?: string): void;
}>();

// Estado reactivo del formulario
const postContent = ref(props.modelValue);
const selectedMedia = ref<File | null>(null);
const mediaPreviewUrl = ref<string | undefined>(undefined);
const location = ref('');
const visibility = ref('public');
const error = ref('');
const showLocationField = ref(false);
const showVisibilityOptions = ref(true);
const postTextarea = ref<HTMLTextAreaElement | null>(null);

// Observamos cambios en el contenido para emitir actualizaciones
watch(postContent, (newValue) => {
  emit('update:modelValue', newValue);
});

// Observamos cambios en el prop modelValue para actualizar el estado local
watch(() => props.modelValue, (newValue) => {
  postContent.value = newValue;
});

// Computamos si el post es válido para habilitar/deshabilitar el botón de publicar
const isValidPost = computed(() => {
  return postContent.value.trim().length > 0 || selectedMedia.value !== null;
});

/**
 * Maneja la selección de archivos
 */
function handleFileSelection(event: Event) {
  const input = event.target as HTMLInputElement;
  if (!input.files || input.files.length === 0) {
    return;
  }
  
  const file = input.files[0];
  
  // Validar que sea una imagen
  if (!file.type.startsWith('image/')) {
    error.value = 'Solo se permiten imágenes';
    return;
  }
  
  // Validar tamaño (máximo 5MB)
  if (file.size > 5 * 1024 * 1024) {
    error.value = 'La imagen no debe superar 5MB';
    return;
  }
  
  selectedMedia.value = file;
  
  // Crear URL para la vista previa
  if (mediaPreviewUrl.value) {
    URL.revokeObjectURL(mediaPreviewUrl.value);
  }
  
  mediaPreviewUrl.value = undefined;
  error.value = '';
}

/**
 * Elimina la imagen seleccionada
 */
function removeMedia() {
  if (mediaPreviewUrl.value) {
    URL.revokeObjectURL(mediaPreviewUrl.value);
  }
  
  selectedMedia.value = null;
  mediaPreviewUrl.value = undefined;
}

/**
 * Activa/desactiva el campo de ubicación
 */
function toggleLocationField() {
  showLocationField.value = !showLocationField.value;
  
  // Si se oculta el campo, limpiar su valor
  if (!showLocationField.value) {
    location.value = '';
  } else {
    // Enfocar el campo cuando se muestra
    nextTick(() => {
      const locationInput = document.querySelector('.location-input') as HTMLInputElement;
      if (locationInput) {
        locationInput.focus();
      }
    });
  }
}

/**
 * Crea un nuevo post
 */
function createPost() {
  if (!isValidPost.value) return;
  
  error.value = '';
  
  try {
    // Emitir evento con la información del post
    emit('create', postContent.value, selectedMedia.value || undefined, location.value || undefined);
    
    // Limpiar el formulario
    resetForm();
  } catch (err) {
    error.value = 'Error al crear la publicación. Inténtalo de nuevo.';
    console.error('Error al crear post:', err);
  }
}

/**
 * Resetea el formulario a su estado inicial
 */
function resetForm() {
  postContent.value = '';
  removeMedia();
  location.value = '';
  showLocationField.value = false;
  error.value = '';
}

/**
 * Maneja el evento keydown en el textarea para enviar el formulario con Ctrl+Enter
 */
function handleKeydown(event: KeyboardEvent) {
  // Ctrl+Enter para enviar
  if (event.ctrlKey && event.key === 'Enter') {
    event.preventDefault();
    createPost();
  }
}

// Al montar el componente, enfocar el textarea
nextTick(() => {
  if (postTextarea.value) {
    postTextarea.value.focus();
  }
});
</script>

<style scoped>
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
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.3);
}

/* Cabecera del modal */
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
  color: #FFFFFF;
  font-weight: 600;
}

.close-modal-btn {
  background: transparent;
  border: none;
  color: #D0D0D0;
  font-size: 1.5rem;
  cursor: pointer;
  padding: 0.25rem;
  line-height: 1;
}

/* Cuerpo del modal */
.modal-body {
  padding: 1rem;
}

/* Información del usuario */
.user-info {
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

.user-details {
  display: flex;
  flex-direction: column;
}

.username {
  margin: 0;
  font-size: 0.875rem;
  font-weight: 600;
  color: #FFFFFF;
}

.visibility-select {
  background-color: #243447;
  border: none;
  border-radius: 12px;
  color: #8A97A8;
  font-size: 0.75rem;
  padding: 0.25rem 0.5rem;
  margin-top: 0.25rem;
}

/* Formulario de post */
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
  font-family: inherit;
  font-size: 1rem;
}

.post-textarea.with-media {
  min-height: 80px;
}

.post-textarea:focus {
  outline: none;
}

/* Vista previa de la imagen */
.media-preview {
  position: relative;
  max-height: 300px;
  border-radius: 8px;
  overflow: hidden;
}

.preview-image {
  width: 100%;
  max-height: 300px;
  object-fit: contain;
  background-color: #0A1625;
}

.remove-media-btn {
  position: absolute;
  top: 0.5rem;
  right: 0.5rem;
  background-color: rgba(0, 0, 0, 0.6);
  color: white;
  border: none;
  border-radius: 50%;
  width: 24px;
  height: 24px;
  font-size: 0.75rem;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
}

/* Campo de ubicación */
.location-field {
  padding-top: 0.5rem;
}

.location-input {
  width: 100%;
  background-color: #243447;
  border: none;
  border-radius: 20px;
  padding: 0.5rem 1rem;
  color: #FFFFFF;
  font-size: 0.875rem;
}

.location-input:focus {
  outline: none;
}

/* Mensaje de error */
.error-message {
  color: #FF4D4D;
  font-size: 0.875rem;
  margin-top: 0.5rem;
}

/* Acciones del formulario */
.post-form-actions {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-top: 0.5rem;
}

.attach-options {
  display: flex;
  gap: 0.75rem;
}

.attach-btn {
  background-color: transparent;
  border: none;
  color: #D0D0D0;
  font-size: 1.25rem;
  cursor: pointer;
  padding: 0.5rem;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.attach-btn:hover {
  background-color: #243447;
}

.file-input {
  display: none;
}

.publish-btn {
  background-color: #0056A6;
  border: none;
  border-radius: 25px;
  padding: 0.5rem 1.5rem;
  color: white;
  font-weight: 600;
  cursor: pointer;
  transition: background-color 0.2s;
}

.publish-btn:hover:not(:disabled) {
  background-color: #4D8CD9;
}

.publish-btn:disabled {
  background-color: #243447;
  color: #8A97A8;
  cursor: not-allowed;
}

/* Responsive */
@media (max-width: 480px) {
  .modal-content {
    width: 100%;
    height: 100%;
    max-width: none;
    border-radius: 0;
  }
  
  .modal-body {
    height: calc(100% - 60px);
    display: flex;
    flex-direction: column;
  }
  
  .post-form {
    flex: 1;
    display: flex;
    flex-direction: column;
  }
  
  .post-textarea {
    flex: 1;
    min-height: 120px;
  }
}
</style>

