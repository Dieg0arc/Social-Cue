// plugins/primevue.ts

import { defineNuxtPlugin } from "#imports";
import PrimeVue from "primevue/config";
import Button from "primevue/button";
import InputText from "primevue/inputtext";
import Card from "primevue/card";
import Message from "primevue/message";
import Toast from "primevue/toast";
import ToastService from "primevue/toastservice";
import Ripple from "primevue/ripple";

export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.vueApp.use(PrimeVue, {
    ripple: true,
    inputStyle: "filled",
  });

  // Register PrimeVue components
  nuxtApp.vueApp.component("Button", Button);
  nuxtApp.vueApp.component("InputText", InputText);
  nuxtApp.vueApp.component("Card", Card);
  nuxtApp.vueApp.component("Message", Message);
  nuxtApp.vueApp.component("Toast", Toast);

  // Register PrimeVue services
  nuxtApp.vueApp.use(ToastService);

  // Register directives
  nuxtApp.vueApp.directive("ripple", Ripple);
});
