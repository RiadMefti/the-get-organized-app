import axiosInstance from '../api/axiosInstance';

export interface AuthResponse {
    jwt: string;
}

export interface RegisterCredentials {
    email: string;
    password: string;
    copyPassword: string;
}

export interface LoginCredentials {
    email: string;
    password: string;
}

export const login = async (credentials: LoginCredentials): Promise<AuthResponse> => {
    const response = await axiosInstance.post('/login', credentials);
    localStorage.setItem('token', response.data.jwt);
    return response.data;
};

export const register = async (credentials: RegisterCredentials): Promise<AuthResponse> => {
    const response = await axiosInstance.post('/register', {
        email: credentials.email,
        password: credentials.password,
        copy_password: credentials.copyPassword,
    });
    localStorage.setItem('token', response.data.jwt);
    return response.data;
};
