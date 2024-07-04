import axiosInstance from '../api/axiosInstance';
import { useAuthStore } from '../stores/authStore';
import { AuthResponse, LoginCredentials, RegisterCredentials } from '../types/types';



export const login = async (credentials: LoginCredentials): Promise<AuthResponse> => {
  const response = await axiosInstance.post('/login', credentials);
  useAuthStore.getState().setToken(response.data.jwt);
  return response.data;
};

export const register = async (credentials: RegisterCredentials): Promise<AuthResponse> => {
  const response = await axiosInstance.post('/register', {
    email: credentials.email,
    password: credentials.password,
    copy_password: credentials.copyPassword,
  });
  useAuthStore.getState().setToken(response.data.jwt);
  return response.data;
};

export const isUserAuthenticated = async (): Promise<boolean> => {
  const token = useAuthStore.getState().token;
  if (!token) return false;

  try {
    const response = await axiosInstance.post('/isAuthenticated', { jwt: token });
    return response.status === 200;
  } catch (error) {
    return false;
  }
};
