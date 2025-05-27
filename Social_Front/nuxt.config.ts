import tsconfigPaths from "vite-tsconfig-paths";

export default defineNuxtConfig({
  modules: [
    "@nuxt/ui",
    "@nuxtjs/tailwindcss",
  ],

  tailwindcss: {
    config: {
      content: [
        './components/**/*.{vue,js,ts}',
        './layouts/**/*.vue',
        './pages/**/*.vue',
        './plugins/**/*.{js,ts}',
        './app.vue',
        './assets/**/*.css'
      ],
      theme: {
        extend: {},
      },
      plugins: [],
    },
  },

  app: {
    head: {
      title: "SocialCue",
    },
  },

css: [
  'primeicons/primeicons.css',
  "~/assets/css/main.css",
  "~/assets/css/auth.css",
  "~/assets/css/register.css",
],
 

  logLevel: "silent",

  telemetry: false,

  runtimeConfig: {
    public: {
      apiBaseUrl: process.env.NUXT_PUBLIC_API_BASE || "http://localhost:3000/api",
    },
  },

  nitro: {
    devProxy: {
      "/api": {
        target: "http://localhost:3000/api",
        changeOrigin: true,
        prependPath: false,
      },
    },
    compatibilityDate: "2025-05-26",
  },

  vite: {
    plugins: [tsconfigPaths()],
  },
});
