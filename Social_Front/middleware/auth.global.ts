import { defineNuxtRouteMiddleware, navigateTo } from "nuxt/app";
import { useAuthStore } from "@/stores/auth";
import type { RouteLocationNormalized } from "vue-router";

const publicRoutes = ["/", "/login", "/register"];

export default defineNuxtRouteMiddleware((to: RouteLocationNormalized) => {
  try {
    const auth = useAuthStore();

    if (publicRoutes.includes(to.path)) {
      if (auth.isAuthenticated && ["/login", "/register"].includes(to.path)) {
        return navigateTo("/dashboard");
      }
      return;
    }

    if (!auth.isAuthenticated) {
      return navigateTo("/login");
    }
  } catch (error) {
    // If Pinia is not ready or there's an error accessing the store,
    // allow access to public routes and redirect to login for private routes
    if (publicRoutes.includes(to.path)) {
      return;
    }
    return navigateTo("/login");
  }
});
