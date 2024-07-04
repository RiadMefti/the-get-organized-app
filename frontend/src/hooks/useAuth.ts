import { useMutation, UseMutationResult, useQuery } from '@tanstack/react-query';
import { login, register, isUserAuthenticated } from '../services/authService';
import { AuthResponse, LoginCredentials, RegisterCredentials } from '../types/types';

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

export const useCheckAuth = () => {
    return useQuery({
        queryKey: ['isAuthenticated'],
        queryFn: isUserAuthenticated
    });
};