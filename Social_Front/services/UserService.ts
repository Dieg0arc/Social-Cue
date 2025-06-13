// services/userService.ts

import axios from "axios";

export async function getUserById(userId: string) {
  try {
    const response = await axios.get(
      `http://localhost:8082/api/users/${userId}`
    );
    return response.data;
  } catch (error) {
    console.error("Error al obtener el usuario:", error);
    throw error;
  }
}
