import { useAuthStore } from "~/stores/auth";
import type { LoginCredentials, RegisterData } from "~/types/user";
import { useRouter } from "vue-router";
import { computed } from "vue";

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
