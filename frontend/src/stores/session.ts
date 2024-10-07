import { writable } from "svelte/store";

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
        const response = await fetch("http://localhost:8080/api/check-login", {
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
