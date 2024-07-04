import { create } from 'zustand';

interface AuthState {
    token: string | null;
    isAuthenticated: boolean;
    setToken: (token: string | null) => void;

    setIsAuthenticated: (isAuthenticated: boolean) => void;

}

export const useAuthStore = create<AuthState>((set) => ({
    token: localStorage.getItem('token'),
    isAuthenticated: false,
    setToken: (token: string | null) => {
        if (token) {
            localStorage.setItem('token', token);
        } else {
            localStorage.removeItem('token');
        }
        set({ token, isAuthenticated: !!token });
    },
    setIsAuthenticated: (isAuthenticated: boolean) => {
        set({ isAuthenticated });
    },

}));
