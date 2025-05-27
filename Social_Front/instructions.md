# Instrucciones para usar y corregir el sistema

## Credenciales de acceso

Para acceder al sistema, debes usar las siguientes credenciales de prueba:

Email: user@example.com
Password: password123

Estas son las credenciales que están configuradas en el store de autenticación. Una vez que inicies sesión correctamente, serás redirigido al dashboard.

## Problemas a solucionar

### 1. Rutas de imágenes

Los errores en consola indican que varias imágenes no se encuentran. Asegúrate de que las siguientes imágenes existan en los directorios correspondientes:

- `~/assets/images/OnlyLogo_Tp.png`
- `~/assets/images/LoginLogo.png`
- `assets\images\Usuari.png`
- `~/assets/images/avatar1.jpg` hasta `avatar8.jpg`
- `~/assets/images/post1.jpg` y `post2.jpg`

Puedes crear estas imágenes o ajustar las rutas en el código para que apunten a imágenes existentes.

### 2. Warnings de plugins duplicados

Para solucionar los warnings de "Plugin has already been applied to target app" y componentes duplicados, necesitas modificar la configuración de Nuxt:

1. Abre o crea el archivo `nuxt.config.ts` en la raíz del proyecto
2. Asegúrate de que los plugins se configuren con la opción `mode: 'client'` si están destinados a ejecutarse solo en el cliente

Por ejemplo:

```typescript
// nuxt.config.ts
export default defineNuxtConfig({
  plugins: [
    { src: "~/plugins/primevue.ts", mode: "client" },
    // otros plugins...
  ],
  // otras configuraciones...
});
```

### 3. Errores de TypeScript corregidos

Los errores de TypeScript en los componentes ya han sido corregidos:

- Se solucionó el problema con `mediaPreviewUrl.value = null` en CreatePostModal.vue
- Se corrigió el manejo de tipos en el evento de comentarios en PostsList.vue
- Se actualizaron las importaciones de tipos en SuggestionsSidebar.vue

### 4. Próximos pasos

Una vez corregidos estos problemas, el sistema debería funcionar correctamente. Para futuras mejoras, considera:

- Implementar la autenticación real con una API en lugar de usar credenciales hardcoded
- Optimizar la carga de imágenes con técnicas como lazy loading
- Revisar la configuración de Nuxt para evitar problemas de registro múltiple de componentes
