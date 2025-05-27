/**
 * Archivo: useInfiniteScroll.ts
 * Ubicación: composables/useInfiniteScroll.ts
 * Descripción: Composable para implementar la funcionalidad de scroll infinito
 * en componentes que necesitan cargar más contenido cuando el usuario llega
 * al final de la página.
 */

import { ref, onMounted, onUnmounted } from 'vue';

/**
 * Opciones de configuración para el scroll infinito
 */
interface InfiniteScrollOptions {
  /**
   * Distancia en píxeles desde el final de la página para iniciar la carga
   * @default 200
   */
  threshold?: number;
  
  /**
   * Función a ejecutar cuando se debe cargar más contenido
   */
  onLoadMore: () => void | Promise<void>;
  
  /**
   * Elemento al que adjuntar el listener de scroll (por defecto window)
   * @default window
   */
  scrollElement?: HTMLElement | null;
}

/**
 * Composable para manejar la funcionalidad de scroll infinito
 * 
 * @param options Opciones de configuración
 * @returns Estado de carga y referencia al elemento de scroll
 */
export function useInfiniteScroll(options: InfiniteScrollOptions) {
  // Estado de carga
  const isLoading = ref(false);
  
  // Valores por defecto
  const threshold = options.threshold || 200;
  const onLoadMore = options.onLoadMore;
  
  // Elemento al que se adjunta el scroll
  const scrollTarget = ref<HTMLElement | Window>(options.scrollElement || window);
  
  /**
   * Maneja el evento de scroll
   */
  const handleScroll = async () => {
    if (isLoading.value) return;
    
    // Calcula la posición del scroll y la altura total
    let scrollHeight: number;
    let scrollPosition: number;
    let clientHeight: number;
    
    if (scrollTarget.value === window) {
      scrollHeight = document.documentElement.scrollHeight;
      scrollPosition = window.scrollY + window.innerHeight;
      clientHeight = window.innerHeight;
    } else {
      const element = scrollTarget.value as HTMLElement;
      scrollHeight = element.scrollHeight;
      scrollPosition = element.scrollTop + element.clientHeight;
      clientHeight = element.clientHeight;
    }
    
    // Si el usuario está cerca del final, carga más contenido
    if (scrollPosition >= scrollHeight - threshold) {
      isLoading.value = true;
      
      try {
        await onLoadMore();
      } catch (error) {
        console.error('Error al cargar más contenido:', error);
      } finally {
        isLoading.value = false;
      }
    }
  };
  
  /**
   * Establece el elemento de scroll
   */
  const setScrollElement = (element: HTMLElement | null) => {
    // Eliminar el listener anterior si existe
    if (scrollTarget.value) {
      (scrollTarget.value as HTMLElement | Window).removeEventListener('scroll', handleScroll);
    }
    
    // Establecer el nuevo elemento
    scrollTarget.value = element || window;
    
    // Añadir el nuevo listener
    scrollTarget.value.addEventListener('scroll', handleScroll);
  };
  
  // Ciclo de vida del componente
  onMounted(() => {
    scrollTarget.value.addEventListener('scroll', handleScroll);
  });
  
  onUnmounted(() => {
    if (scrollTarget.value) {
      scrollTarget.value.removeEventListener('scroll', handleScroll);
    }
  });
  
  return {
    isLoading,
    setScrollElement
  };
}

