export default defineNuxtConfig({
  modules: [
    "@nuxt/image",
    "@nuxt/ui",
    "@primevue/nuxt-module",
    "@nuxtjs/tailwindcss",
    "@pinia/nuxt",
  ],

  primevue: {
    components: {
      include: ["Button", "InputText", "Card", "Message", "Toast"]
    },
    directives: {
      include: ["ripple"]
    },
    options: {
      ripple: true
    }
  },

  css: ["~/assets/css/main.css"],

  routeRules: {
    "/": { redirect: "/login" },
  },

  typescript: {
    strict: true,
    typeCheck: true,
    shim: false,
  },

  nitro: {
    compressPublicAssets: true,
    compatibilityDate: "2025-05-01",
  },

  imports: {
    dirs: ["composables", "stores"],
  },

  compatibilityDate: "2025-05-01",
});
