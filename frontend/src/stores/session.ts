import { writable } from "svelte/store";
import { API_URL } from "../utils/api";

interface SessionState {
  authenticated: boolean;
  username: string | null;
}

function createSessionStore() {
  const { subscribe, set, update } = writable<SessionState>({
    authenticated: false,
    username: null,
  });

  return {
    subscribe,
    setSession: (data: SessionState) => set(data),
    clearSession: () => set({ authenticated: false, username: null }),
    checkSession: async () => {
      try {
        const response = await fetch(`${API_URL}/check-login`, {
          credentials: "include",
        });
        const data = await response.json();
        if (data.statusCode === 200) {
          set({
            authenticated: true,
            username: data.username,
          });
        } else {
          set({ authenticated: false, username: null });
        }
      } catch (error) {
        console.error("Error checking session", error);
        set({ authenticated: false, username: null });
      }
    },
  };
}

export const session = createSessionStore();
