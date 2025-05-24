/**
 * middleware/auth.ts - Middleware de autenticación para rutas protegidas
 *
 * Este middleware controla el acceso a las rutas de la aplicación basado en
 * el estado de autenticación del usuario. Redirige automáticamente a los usuarios
 * no autenticados a la página de login cuando intentan acceder a rutas protegidas,
 * y a los usuarios autenticados al dashboard cuando intentan acceder a páginas de
 * autenticación como login o registro.
 */

import { defineNuxtRouteMiddleware, navigateTo } from "#app";
import { useAuthStore } from "~/stores/auth";

/**
 * Middleware de autenticación para aplicar a rutas Nuxt
 *
 * @param {Object} to - Objeto de ruta destino
 * @param {Object} from - Objeto de ruta origen
 * @returns {ReturnType<typeof navigateTo> | void} - Redirige o permite la navegación
 */
export default defineNuxtRouteMiddleware((to, from) => {
  // Inicializa el store de autenticación
  const authStore = useAuthStore();

  // Reglas para rutas públicas (/, /login, /register)
  if (to.path === "/" || to.name === "login" || to.name === "register") {
    // Si el usuario ya está autenticado y trata de acceder a login o register,
    // redirigirlo al dashboard para evitar autenticación redundante
    if (authStore.isAuthenticated) {
      return navigateTo("/dashboard");
    }
    // Si no está autenticado, permitir acceso a estas rutas públicas
    return;
  }

  // Para cualquier otra ruta (rutas protegidas), verificar autenticación
  if (!authStore.isAuthenticated) {
    // Si no está autenticado, redirigir a la página de login
    return navigateTo("/login");
  }

  // Si está autenticado, permitir acceso a la ruta solicitada
  return;
});
