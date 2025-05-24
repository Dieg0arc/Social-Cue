/**
 * Archivo: auth.ts
 * Ubicación: stores/auth.ts
 * Descripción: Store de autenticación utilizando Pinia para la gestión del estado
 * de autenticación en la aplicación SocialCue.
 */

// Importación de la función para definir stores de Pinia
import { defineStore } from "pinia";
import type { User, LoginCredentials, RegisterData, AuthState } from "~/types/user";

/**
 * Store para manejar la autenticación de usuarios
 *
 * Este store se encarga de:
 * - Iniciar sesión
 * - Registrar usuarios
 * - Cerrar sesión
 * - Mantener el estado de autenticación
 */
export const useAuthStore = defineStore('auth', {
  /**
   * Estado del store
   */
  state: (): AuthState => ({
    user: null,
    token: null,
    loading: false,
    error: null
  }),

  /**
   * Getters para obtener información del estado
   */
  getters: {
    /**
     * Verifica si el usuario está autenticado
     * @returns Boolean indicando si hay un token válido
     */
    isAuthenticated: (state: AuthState) => !!state.token,

    /**
     * Obtiene la información del usuario actual
     * @returns Objeto User o null si no hay usuario autenticado
     */
    getUser: (state: AuthState) => state.user,
  },

  /**
   * Acciones que modifican el estado
   */
  actions: {
    /**
     * Inicia sesión con credenciales de usuario
     * @param credentials Credenciales de inicio de sesión
     * @returns Promise que resuelve a true si el login es exitoso, false en caso contrario
     */
    async login(credentials: LoginCredentials) {
      this.loading = true;
      this.error = null;

      try {
        // Aquí eventualmente llamarás a tu API
        // Este es un ejemplo simulado
        // En producción, este código será reemplazado por una llamada real a la API

        // Simular un retraso de red
        await new Promise(resolve => setTimeout(resolve, 1000));

        // Verificación simulada de credenciales
        if (credentials.email !== "user@example.com" || credentials.password !== "password123") {
          throw new Error("Credenciales inválidas");
        }

        // Respuesta simulada de la API
        const response = {
          user: {
            id: 1,
            email: credentials.email,
            name: "Usuario Demo",
            fullName: "Usuario Demostración",
            username: "demo_user",
            role: "student" as const,
            createdAt: new Date()
          } as User,
          token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
        };

        // Actualiza el estado con la información de usuario y token
        this.user = response.user;
        this.token = response.token;

        // Guardar en localStorage para persistencia entre sesiones
        localStorage.setItem("auth_token", response.token);
        localStorage.setItem("auth_user", JSON.stringify(response.user));

        return true;
      } catch (error) {
        this.error = "Error de inicio de sesión. Verifica tus credenciales.";
        return false;
      } finally {
        this.loading = false;
      }
    },

    /**
     * Registra un nuevo usuario
     */
    async register(data: RegisterData) {
      this.loading = true;
      this.error = null;

      try {
        // Simular un retraso de red
        await new Promise(resolve => setTimeout(resolve, 1000));

        // Verificación simulada (por ejemplo, comprobar si el email ya existe)
        if (data.email === "user@example.com") {
          throw new Error("El email ya está registrado");
        }

        // Respuesta simulada de la API
        const response = {
          user: {
            id: 2, // ID diferente al del usuario de login
            email: data.email,
            name: data.name || "",
            fullName: data.fullName || data.name || "",
            username: data.username || "",
            role: "student" as const,
            createdAt: new Date()
          } as User,
          token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.nuevo"
        };

        // Actualiza el estado con la información de usuario y token
        this.user = response.user;
        this.token = response.token;

        // Guardar en localStorage para persistencia entre sesiones
        localStorage.setItem("auth_token", response.token);
        localStorage.setItem("auth_user", JSON.stringify(response.user));

        return true;
      } catch (error) {
        this.error = "Error en el registro. Inténtalo de nuevo.";
        return false;
      } finally {
        this.loading = false;
      }
    },

    /**
     * Cierra la sesión del usuario actual
     */
    logout() {
      this.user = null;
      this.token = null;

      // Eliminar datos de localStorage
      localStorage.removeItem("auth_token");
      localStorage.removeItem("auth_user");
    },

    /**
     * Inicializa el estado de autenticación basado en localStorage
     * Se utiliza normalmente al cargar la aplicación
     */
    initAuth() {
      const token = localStorage.getItem("auth_token");
      const userStr = localStorage.getItem("auth_user");

      if (token && userStr) {
        try {
          this.token = token;
          this.user = JSON.parse(userStr);
        } catch (e) {
          // Si hay un error al procesar los datos almacenados, cierra la sesión
          this.logout();
        }
      }
    }
  }
});
