/**
 * Archivo: useDashboard.ts
 * Ubicación: composables/useDashboard.ts
 * Descripción: Composable para manejar la lógica y estado del dashboard
 */

import { ref, reactive, computed } from "vue";
import { useAuthStore } from "~/stores/auth";
import type {
  Post,
  TextPost,
  MediaPost,
  Story,
  SuggestionUser,
  CurrentUser,
} from "~/types/dashboard";

/**
 * Composable para manejar la lógica y estado del dashboard
 */
export function useDashboard() {
  // Obtener store de autenticación
  const authStore = useAuthStore();

  // Estado reactivo
  const loading = ref(false);
  const page = ref(1);

  // Estado del dashboard
  const dashboardState = reactive({
    // Datos del usuario actual
    currentUser: {
      id: Number(authStore.user?.id || 1),
      username: authStore.user?.username || "usuario_actual",
      fullName: authStore.user?.fullName || "Usuario Actual",
      avatar: "assetsimagesUsuari.png",
    } as CurrentUser,

    // Historias
    stories: [
      {
        id: 1,
        username: "maria_rodriguez",
        avatar: "/assets/images/avatar1.jpg",
        hasNew: true,
      },
      {
        id: 2,
        username: "juan_perez",
        avatar: "/assets/images/avatar2.jpg",
        hasNew: true,
      },
      {
        id: 3,
        username: "ana_gomez",
        avatar: "/assets/images/avatar3.jpg",
        hasNew: false,
      },
      {
        id: 4,
        username: "carlos_lopez",
        avatar: "/assets/images/avatar4.jpg",
        hasNew: true,
      },
      {
        id: 5,
        username: "laura_martinez",
        avatar: "/assets/images/avatar5.jpg",
        hasNew: false,
      },
    ] as Story[],

    // Sugerencias de usuarios
    suggestions: [
      {
        id: 1,
        username: "pablo_sanchez",
        fullName: "Pablo Sánchez",
        avatar: "/assets/images/avatar6.jpg",
        meta: "Seguido por maria_rodriguez",
      },
      {
        id: 2,
        username: "elena_torres",
        fullName: "Elena Torres",
        avatar: "/assets/images/avatar7.jpg",
        meta: "Nuevo en SocialCue",
      },
      {
        id: 3,
        username: "miguel_fernandez",
        fullName: "Miguel Fernández",
        avatar: "/assets/images/avatar8.jpg",
        meta: "Seguido por juan_perez",
      },
    ] as SuggestionUser[],

    // Posts
    posts: [
      {
        id: 1,
        author: {
          id: 2,
          username: "maria_rodriguez",
          avatar: "/assets/images/avatar1.jpg",
        },
        mediaUrl: "/assets/images/post1.jpg",
        caption:
          "Disfrutando de un día increíble en la playa. El sol, la arena y el mar... ¡Qué más se puede pedir! #Verano #Playa #Relax",
        location: "Playa del Carmen",
        likes: 120,
        liked: false,
        comments: [
          { username: "juan_perez", text: "¡Qué envidia! Disfruta mucho 🏖️" },
          {
            username: "ana_gomez",
            text: "Hermoso lugar, las olas se ven geniales",
          },
        ],
        createdAt: "2025-05-25T14:30:00",
      } as MediaPost,
      {
        id: 2,
        author: {
          id: 3,
          username: "juan_perez",
          avatar: "/assets/images/avatar2.jpg",
        },
        mediaUrl: "/assets/images/post2.jpg",
        caption:
          "Nuevo proyecto terminado. Ha sido un desafío pero el resultado valió la pena. #Programación #Desarrollo #SocialCue",
        likes: 85,
        liked: true,
        comments: [
          {
            username: "carlos_lopez",
            text: "Excelente trabajo, ¿qué tecnologías usaste?",
          },
          {
            username: "maria_rodriguez",
            text: "Se ve increíble, felicitaciones!",
          },
          {
            username: "ana_gomez",
            text: "Quiero saber más sobre este proyecto",
          },
        ],
        createdAt: "2025-05-24T10:15:00",
      } as MediaPost,
      {
        id: 3,
        author: {
          id: 4,
          username: "ana_gomez",
          avatar: "/assets/images/avatar3.jpg",
        },
        textContent:
          "Hoy comienza una nueva etapa en mi vida profesional. Muy emocionada por los nuevos retos que vendrán. #NuevoTrabajo #Oportunidades",
        likes: 45,
        liked: false,
        comments: [
          {
            username: "laura_martinez",
            text: "¡Felicitaciones! Te mereces lo mejor 🎉",
          },
        ],
        createdAt: "2025-05-23T08:45:00",
      } as TextPost,
    ] as Post[],
  });

  /**
   * Maneja la acción de dar like a un post
   */
  const handleLike = (postId: number) => {
    console.log(`Like en post ${postId}`);

    // Encuentra el post y actualiza su estado
    const post = dashboardState.posts.find((p) => p.id === postId);
    if (post) {
      post.liked = !post.liked;
      post.likes = post.liked ? post.likes + 1 : post.likes - 1;
    }

    // En una implementación real, aquí se haría una petición a la API
  };

  /**
   * Maneja la acción de comentar en un post
   */
  const handleComment = (postId: number, comment: string) => {
    console.log(`Nuevo comentario en post ${postId}: ${comment}`);

    // Encuentra el post y añade el comentario
    const post = dashboardState.posts.find((p) => p.id === postId);
    if (post && comment.trim()) {
      post.comments.push({
        username: dashboardState.currentUser.username,
        text: comment,
      });
    }

    // En una implementación real, aquí se haría una petición a la API
  };

  /**
   * Crea un nuevo post
   */
  const createPost = (content: string, mediaFile?: File, location?: string) => {
    // Generamos un ID simple
    const newId = Math.max(...dashboardState.posts.map((p) => p.id)) + 1;

    // Si hay un archivo de media, creamos un post con imagen
    if (mediaFile) {
      // En una implementación real, aquí se subiría el archivo al servidor
      // y se obtendría la URL
      const fakeMediaUrl = "/assets/images/post1.jpg"; // URL de ejemplo

      const newPost: MediaPost = {
        id: newId,
        author: {
          id: dashboardState.currentUser.id,
          username: dashboardState.currentUser.username,
          avatar: dashboardState.currentUser.avatar,
        },
        mediaUrl: fakeMediaUrl,
        caption: content,
        location,
        likes: 0,
        liked: false,
        comments: [],
        createdAt: new Date().toISOString(),
      };

      // Añadimos el post al inicio del array
      dashboardState.posts.unshift(newPost);
    } else {
      // Si no hay archivo, creamos un post de texto
      const newPost: TextPost = {
        id: newId,
        author: {
          id: dashboardState.currentUser.id,
          username: dashboardState.currentUser.username,
          avatar: dashboardState.currentUser.avatar,
        },
        textContent: content,
        likes: 0,
        liked: false,
        comments: [],
        createdAt: new Date().toISOString(),
      };

      // Añadimos el post al inicio del array
      dashboardState.posts.unshift(newPost);
    }

    console.log("Post creado:", content);
    // En una implementación real, aquí se haría una petición a la API
  };

  /**
   * Carga más posts
   */
  const loadMorePosts = async () => {
    if (loading.value) return;

    loading.value = true;

    try {
      // Simulamos una petición a la API
      await new Promise((resolve) => setTimeout(resolve, 1500));

      // Nuevos posts de ejemplo
      const newPosts: Post[] = [
        {
          id: 4 + page.value,
          author: {
            id: 5,
            username: "carlos_lopez",
            avatar: "/assets/images/avatar4.jpg",
          },
          mediaUrl: "/assets/images/post1.jpg",
          caption:
            "Explorando nuevos lugares. La arquitectura de esta ciudad es impresionante. #Viajes #Arquitectura",
          location: "Madrid, España",
          likes: 67,
          liked: false,
          comments: [
            {
              username: "maria_rodriguez",
              text: "¡Qué hermoso lugar! Tengo que visitarlo",
            },
          ],
          createdAt: `2025-05-${20 - page.value}T16:20:00`,
        } as MediaPost,
        {
          id: 5 + page.value,
          author: {
            id: 6,
            username: "laura_martinez",
            avatar: "/assets/images/avatar5.jpg",
          },
          textContent:
            "Acabo de terminar de leer un libro increíble. Totalmente recomendado para los amantes de la ciencia ficción. #Lectura #Libros",
          likes: 32,
          liked: false,
          comments: [],
          createdAt: `2025-05-${19 - page.value}T09:10:00`,
        } as TextPost,
      ];

      // Añadimos los nuevos posts al final del array
      dashboardState.posts = [...dashboardState.posts, ...newPosts];

      // Incrementamos la página
      page.value += 1;
    } catch (error) {
      console.error("Error al cargar más posts:", error);
    } finally {
      loading.value = false;
    }
  };

  return {
    dashboardState,
    loading,
    page,
    handleLike,
    handleComment,
    createPost,
    loadMorePosts,
  };
}
