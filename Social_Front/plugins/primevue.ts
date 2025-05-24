// plugins/primevue.ts
import { defineNuxtPlugin } from '#app';
import PrimeVue from 'primevue/config';
import ToastService from 'primevue/toastservice';
import Card from 'primevue/card';
import Button from 'primevue/button';
import InputText from 'primevue/inputtext';
import Message from 'primevue/message';
import Toast from 'primevue/toast';

export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.vueApp.use(PrimeVue, { ripple: true });
  nuxtApp.vueApp.use(ToastService);

  // Register components
  nuxtApp.vueApp.component('Card', Card);
  nuxtApp.vueApp.component('Button', Button);
  nuxtApp.vueApp.component('InputText', InputText);
  nuxtApp.vueApp.component('Message', Message);
  nuxtApp.vueApp.component('Toast', Toast);
});

