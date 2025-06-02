export default defineNuxtPlugin((nuxtApp) => {
  const config = useRuntimeConfig();

  const api = $fetch.create({
    baseURL: config.public.apiBaseUrl as string,
    headers: {
      accept: "application/json",
    },
    onRequest({ options }) {
      const token = useCookie("token").value;

      // Convertimos a Headers si no lo es
      let headers: Headers;
      if (options.headers instanceof Headers) {
        headers = options.headers;
      } else {
        headers = new Headers(options.headers as HeadersInit);
      }

      if (token) {
        headers.set("Authorization", `Bearer ${token}`);
      }

      options.headers = headers;
    },
    onResponse({ response }) {
      console.log("[API Response]", response._data);
    },
    onRequestError({ error }) {
      console.error("[API Request Error]", error);
    },
    onResponseError({ response }) {
      console.error("[API Response Error]", response.status, response._data);
    },
  });

  return {
    provide: {
      api,
    },
  };
});
