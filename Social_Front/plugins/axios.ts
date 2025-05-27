// plugins/axios.ts
import axios from "axios";

export default defineNuxtPlugin((nuxtApp) => {
  const config = useRuntimeConfig() as { public: { apiBase: string } };

  const api = axios.create({
    baseURL: config.public.apiBase, // se leerá desde .env
    timeout: 10000,
  });

  // Puedes agregar interceptores aquí si deseas
  // api.interceptors.request.use(config => {
  //   // Agregar headers por ejemplo
  //   return config
  // })

  // api.interceptors.response.use(
  //   response => response,
  //   error => {
  //     // Manejo global de errores
  //     return Promise.reject(error)
  //   }
  // )

  return {
    provide: {
      axios: api,
    },
  };
});
