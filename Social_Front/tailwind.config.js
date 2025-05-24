/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
      "./src/**/*.{html,js,jsx,ts,tsx}",
      "./public/**/*.html",
      // Ajusta estas rutas según la estructura de tu proyecto
    ],
    theme: {
      extend: {
        colors: {
          // Puedes agregar colores personalizados para tu universidad
          'humboldt': '#00563F', // Ejemplo
        },
      },
    },
    plugins: [],
  }