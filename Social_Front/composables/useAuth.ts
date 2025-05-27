/**
 * Este composable proporciona funcionalidades de autenticación para la aplicación,
 * incluyendo login, registro y cierre de sesión. Encapsula la interacción con
 * el store de autenticación y maneja la navegación post-autenticación.
 */

import { useAuthStore } from "~/stores/auth";
import type { LoginCredentials, RegisterData } from "~/types/user";
import { useRouter } from "#imports";
import { computed } from "vue";

/**
 * Hook composable que gestiona la autenticación del usuario
 *
 * @returns {Object} Objeto con métodos y propiedades reactivas para autenticación
 */
export function useAuth() {
  const authStore = useAuthStore();
  const router = useRouter();

  return {
    ...authStore,
    router,
    isAuthenticated: computed(() => authStore.isAuthenticated),
    isInitialized: computed(() => authStore.isInitialized),
  };
}
