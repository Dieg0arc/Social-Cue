import { defineStore } from "pinia";
import type { User, LoginCredentials, RegisterData } from "~/types/user";

interface AuthResponse {
  token: string;
  user: User;
  expiresAt?: number;
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
    async initialize() {
      this.loading = true;
      try {
        if (isBrowser()) {
          const token = localStorage.getItem("token");
          const user = localStorage.getItem("user");

          if (token && user) {
            this.token = token;
            this.user = JSON.parse(user);
            await this.validateToken();
          }
        }
      } catch (error) {
        console.error("Error initializing auth state:", error);
        this.clearAuthState();
      } finally {
        this.loading = false;
        this.isInitialized = true;
      }
    },

    async validateToken(): Promise<boolean> {
      if (!this.token) return false;
      return true;
    },

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
        const nuxtApp = useNuxtApp();
        const $api = nuxtApp.$api as typeof $fetch;

        const response: AuthResponse = await $api("/auth/login", {
          method: "POST",
          body: credentials,
        });

        this.token = response.token;
        this.user = response.user;

        if (isBrowser()) {
          localStorage.setItem("token", this.token);
          localStorage.setItem("user", JSON.stringify(this.user));
        }

        return true;
      } catch (error: any) {
        console.error("Error en login:", error);
        this.error = error?.data?.message || "Error en la autenticación";
        return false;
      } finally {
        this.loading = false;
      }
    },

    async register(registerData: RegisterData): Promise<boolean> {
      this.loading = true;
      this.error = null;

      try {
        const nuxtApp = useNuxtApp();
        const $api = nuxtApp.$api as typeof $fetch;

        const response: AuthResponse = await $api("/auth/register", {
          method: "POST",
          body: registerData,
        });

        this.token = response.token;
        this.user = response.user;

        if (isBrowser()) {
          localStorage.setItem("token", this.token);
          localStorage.setItem("user", JSON.stringify(this.user));

          if (response.expiresAt) {
            localStorage.setItem("tokenExpires", response.expiresAt.toString());
          }
        }

        return true;
      } catch (error: any) {
        console.error("Error en registro:", error);
        this.error = error?.data?.message || "Error en el registro";
        return false;
      } finally {
        this.loading = false;
      }
    },

    async logout() {
      this.clearAuthState();
      await navigateTo("/login");
    },

    async requestPasswordReset(email: string): Promise<boolean> {
      this.loading = true;
      try {
        const nuxtApp = useNuxtApp();
        const $api = nuxtApp.$api as typeof $fetch;
        await $api("/auth/password-reset/request", {
          method: "POST",
          body: { email },
        });
        return true;
      } catch (error) {
        console.error("Error solicitando reset:", error);
        return false;
      } finally {
        this.loading = false;
      }
    },

    async confirmPasswordReset(
      email: string,
      code: string,
      newPassword: string
    ): Promise<boolean> {
      this.loading = true;
      try {
        const nuxtApp = useNuxtApp();
        const $api = nuxtApp.$api as typeof $fetch;
        await $api("/auth/password-reset/confirm", {
          method: "POST",
          body: { email, code, new_password: newPassword },
        });
        return true;
      } catch (error) {
        console.error("Error confirmando reset:", error);
        return false;
      } finally {
        this.loading = false;
      }
    },

    isTokenExpired(): boolean {
      if (!this.token) return true;
      if (isBrowser()) {
        const expiresAtStr = localStorage.getItem("tokenExpires");
        if (expiresAtStr) {
          const expiresAt = parseInt(expiresAtStr, 10);
          return Date.now() >= expiresAt;
        }
      }
      return false;
    },

    getAuthHeader(): Record<string, string> {
      return this.token ? { Authorization: `Bearer ${this.token}` } : {};
    },
  },
});
