import { useMutation, UseMutationResult } from '@tanstack/react-query';
import { AuthResponse, login, register } from '../services/authService';

interface RegisterCredentials {
    email: string;
    password: string;
    copyPassword: string;
}

interface LoginCredentials {
    email: string;
    password: string;
}

export const useLogin = (): UseMutationResult<AuthResponse, Error, LoginCredentials> => {
    return useMutation<AuthResponse, Error, LoginCredentials>({
        mutationFn: login
    });
};

export const useRegister = (): UseMutationResult<AuthResponse, Error, RegisterCredentials> => {
    return useMutation<AuthResponse, Error, RegisterCredentials>({
        mutationFn: register
    });
};