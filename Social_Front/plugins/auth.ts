// plugins/auth.ts

import { useAuthStore } from "~/stores/auth";
import { nextTick } from "vue";

export default defineNuxtPlugin(async (nuxtApp) => {
  console.log("Configurando plugin de autenticación...");

  const initializeAuth = async () => {
    try {
      await nextTick();
      const authStore = useAuthStore();
      await authStore.initialize();
      console.log("Estado de autenticación inicializado correctamente");
    } catch (error) {
      console.error("Error al inicializar estado de autenticación:", error);
    }
  };

  if (process.client) {
    nuxtApp.hook("app:mounted", async () => {
      await initializeAuth();
    });
  }

  // forma correcta de retornar provide con typing válido
  return Promise.resolve({
    provide: {
      authInitialized: true as boolean,
    },
  }) as Promise<{ provide?: Record<string, unknown> }>;
});
