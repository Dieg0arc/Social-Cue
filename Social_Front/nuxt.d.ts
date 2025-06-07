// Social_Front/nuxt.d.ts
import type { $Fetch } from "ofetch"; // $fetch viene de ofetch (lo trae Nuxt)

declare module "#app" {
  interface NuxtApp {
    $api: $Fetch;
  }
}

declare module "@vue/runtime-core" {
  interface ComponentCustomProperties {
    $api: $Fetch;
  }
}

export {};
