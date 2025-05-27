// stores/dashboard.ts
import { defineStore } from "pinia";

interface Post {
  id: string;
  content: string;
  authorId: string;
  // Agrega más propiedades si es necesario
}

interface Story {
  id: string;
  imageUrl: string;
  userId: string;
  // Agrega más propiedades si es necesario
}

interface User {
  id: string;
  name: string;
  avatarUrl: string;
  // Agrega más propiedades si es necesario
}

export const useDashboardStore = defineStore("dashboard", {
  state: () => ({
    posts: [] as Post[],
    stories: [] as Story[],
    currentUser: null as User | null,
  }),
  actions: {
    setPosts(posts: Post[]) {
      this.posts = posts;
    },
    setStories(stories: Story[]) {
      this.stories = stories;
    },
    setCurrentUser(user: User) {
      this.currentUser = user;
    },
  },
});
