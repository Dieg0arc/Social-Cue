/**
 * useAuth.ts - Composable para gestión de autenticación
 *
 * Este composable proporciona funcionalidades de autenticación para la aplicación,
 * incluyendo login, registro y cierre de sesión. Encapsula la interacción con
 * el store de autenticación y maneja la navegación post-autenticación.
 */

import { useAuthStore } from "~/stores/auth";
import type { LoginCredentials, RegisterData } from "~/types/user";
import { useRouter } from "#app";
import { computed } from "vue";

/**
 * Hook composable que gestiona la autenticación del usuario
 *
 * @returns {Object} Objeto con métodos y propiedades reactivas para autenticación
 */
export function useAuth() {
  const authStore = useAuthStore();
  const router = useRouter();

  /**
   * Inicia sesión con las credenciales proporcionadas
   *
   * @param {LoginCredentials} credentials - Credenciales para iniciar sesión
   * @returns {Promise<boolean>} Promise que resuelve a true si el login fue exitoso
   */
  const login = async (credentials: LoginCredentials) => {
    const success = await authStore.login(credentials);
    if (success) {
      router.push("/");
    }
    return success;
  };

  /**
   * Registra un nuevo usuario con los datos proporcionados
   *
   * @param {RegisterData} data - Datos para el registro del usuario
   * @returns {Promise<boolean>} Promise que resuelve a true si el registro fue exitoso
   */
  const register = async (data: RegisterData) => {
    const success = await authStore.register(data);
    if (success) {
      router.push("/");
    }
    return success;
  };

  /**
   * Cierra la sesión del usuario actual y redirige a la página de login
   */
  const logout = () => {
    authStore.logout();
    router.push("/login");
  };

  return {
    // Estado reactivo del usuario actual
    user: computed(() => authStore.user),
    // Estado reactivo que indica si el usuario está autenticado
    isAuthenticated: computed(() => authStore.isAuthenticated),
    // Estado reactivo que indica si hay operaciones de autenticación en curso
    loading: computed(() => authStore.loading),
    // Estado reactivo que contiene errores de autenticación, si los hay
    error: computed(() => authStore.error),
    // Métodos de autenticación
    login,
    register,
    logout,
  };
}
