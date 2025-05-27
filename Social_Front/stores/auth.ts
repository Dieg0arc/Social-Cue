import { defineStore } from "pinia";
import { navigateTo } from "#app";
import type { User, LoginCredentials, RegisterData } from "~/types/user";

interface AuthResponse {
  token: string;
  user: User;
  expiresAt?: number; // Optional timestamp for token expiration
}

interface ValidationResponse {
  valid: boolean;
  user?: User;
}

interface ApiError {
  message: string;
  code?: string;
  status?: number;
}

interface AuthState {
  user: User | null;
  token: string | null;
  loading: boolean;
  error: string | null;
  isInitialized: boolean;
}

const isBrowser = () => typeof window !== "undefined";

// Helper to get API URL from runtime config or fallback to local development
const getApiBaseUrl = () => {
  const config = useRuntimeConfig();
  return config.public.apiBaseUrl || "http://localhost:3000/api";
};

export const useAuthStore = defineStore("auth", {
  state: (): AuthState => ({
    user: null,
    token: null,
    loading: false,
    error: null,
    isInitialized: false,
  }),

  getters: {
    isAuthenticated: (state) => !!state.token,
  },

  actions: {
    /**
     * Initialize auth state from localStorage and validate token if exists
     * This is called on application startup
     */
    async initialize() {
      this.loading = true;
      
      try {
        if (isBrowser()) {
          const token = localStorage.getItem("token");
          const user = localStorage.getItem("user");
          
          if (token && user) {
            // Set the initial state from localStorage
            this.token = token;
            this.user = JSON.parse(user);
            
            // Validate the token with the server
            await this.validateToken();
          }
        }
      } catch (error) {
        console.error("Error initializing auth state:", error);
        // If there's an error during initialization, clear the auth state
        this.clearAuthState();
      } finally {
        this.loading = false;
        this.isInitialized = true;
      }
    },
    
    /**
     * Validate the current token with the server
     * Returns true if the token is valid, false otherwise
     */
    async validateToken(): Promise<boolean> {
      if (!this.token) return false;
      
      // For testing, consider any token valid
      return true;
    },
    
    /**
     * Clear the authentication state
     * Used when logging out or when token validation fails
     */
    clearAuthState() {
      this.token = null;
      this.user = null;
      this.error = null;
      
      if (isBrowser()) {
        localStorage.removeItem("token");
        localStorage.removeItem("user");
      }
    },

    async login(credentials: LoginCredentials): Promise<boolean> {
      this.loading = true;
      this.error = null;

      try {
        // Hardcoded credentials for testing
        const validCredentials = {
          email: "test@example.com",
          password: "password123"
        };

        // Simulate API delay
        await new Promise(resolve => setTimeout(resolve, 500));

        if (credentials.email === validCredentials.email && 
            credentials.password === validCredentials.password) {
          
          // Simulate successful login
          const mockUser = {
            id: "1",
            email: validCredentials.email,
            name: "Usuario de Prueba",
            username: "testuser",
            role: "user",
            createdAt: new Date().toISOString()
          };

          this.token = "mock_jwt_token";
          this.user = mockUser;

          if (isBrowser()) {
            localStorage.setItem("token", this.token);
            localStorage.setItem("user", JSON.stringify(mockUser));
          }

          return true;
        }

        // Invalid credentials
        this.error = "Credenciales inválidas. Use test@example.com / password123";
        return false;

      } catch (error: any) {
        console.error("Error de autenticación:", error);
        this.error = "Error en la autenticación";
        return false;
      } finally {
        this.loading = false;
      }
    },

    /**
     * Logout the user, clear auth state and redirect to login page
     */
    async logout() {
      // Just clear the state and redirect
      this.clearAuthState();
      await navigateTo("/login");
    },

    async register(registerData: RegisterData): Promise<boolean> {
      this.loading = true;
      this.error = null;
      
      try {
        const apiBaseUrl = getApiBaseUrl();
        const { data, error } = await useFetch<AuthResponse>(`${apiBaseUrl}/auth/register`, {
          method: "POST",
          body: registerData,
          headers: {
            "Content-Type": "application/json",
          },
        });

        if (error.value) {
          // Handle different error scenarios for registration
          const status = error.value.status || 0;
          const errorData = error.value.data as ApiError | undefined;
          
          if (status === 409) {
            if (errorData?.message?.includes('email')) {
              this.error = "Este correo electrónico ya está registrado. Por favor, utiliza otro.";
            } else if (errorData?.message?.includes('username')) {
              this.error = "Este nombre de usuario ya está en uso. Por favor, elige otro.";
            } else {
              this.error = "Usuario ya existe. Por favor inicia sesión o usa credenciales diferentes.";
            }
          } else if (status === 400) {
            this.error = "Datos de registro inválidos. Por favor verifica la información ingresada.";
          } else {
            this.error = errorData?.message || "Error en el registro";
          }
          
          console.error(`Registration error (${status}):`, this.error);
          return false;
        }

        if (data.value) {
          this.token = data.value.token;
          this.user = data.value.user;

          if (isBrowser()) {
            localStorage.setItem("token", this.token);
            localStorage.setItem("user", JSON.stringify(data.value.user));
            
            // Store token expiration time if provided
            if (data.value.expiresAt) {
              localStorage.setItem("tokenExpires", data.value.expiresAt.toString());
            }
          }
          
          return true;
        } else {
          this.error = "No se recibieron datos de registro";
          return false;
        }
      } catch (error: any) {
        console.error("Error en registro:", error);
        this.error = error?.data?.message || "Error en el registro";
        return false;
      } finally {
        this.loading = false;
      }
    },
    
    /**
     * Check if the current token is expired
     * Returns true if the token is expired or not set
     */
    isTokenExpired(): boolean {
      if (!this.token) return true;
      
      if (isBrowser()) {
        const expiresAtStr = localStorage.getItem("tokenExpires");
        if (expiresAtStr) {
          const expiresAt = parseInt(expiresAtStr, 10);
          return Date.now() >= expiresAt;
        }
      }
      
      // If no expiration info, assume not expired
      return false;
    },
    
    /**
     * Get the authentication header for API requests
     * Used by API services to include the token in requests
     */
    getAuthHeader(): Record<string, string> {
      return this.token 
        ? { "Authorization": `Bearer ${this.token}` }
        : {};
    }
  },
});
