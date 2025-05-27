export default defineNuxtRouteMiddleware((to) => {
  const authStore = useAuthStore();

  if (!authStore.isInitialized) {
    // Verifica si implementaste el método initAuth en el store
    if (typeof authStore.initAuth === "function") {
      authStore.initAuth();
    } else {
      console.warn("initAuth no está definido en el authStore");
    }
  }

  if (to.path !== "/login" && !authStore.isAuthenticated) {
    return navigateTo("/login");
  }
});
