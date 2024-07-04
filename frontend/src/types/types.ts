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

export interface JwtToken {
    jwt: string;
}
