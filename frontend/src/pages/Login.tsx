import React, { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { useCheckAuth, useLogin } from "../hooks/useAuth";
import AuthForm from "../components/AuthForm";
import { Typography, notification } from "antd";

const Login: React.FC = () => {
  const mutation = useLogin();
  const navigate = useNavigate();

  const isAuth = useCheckAuth();

  const handleLogin = (email: string, password: string) => {
    mutation.mutate(
      { email, password },
      {
        onSuccess: () => {
          navigate("/");
        },
        onError: (error: any) => {
          notification.error({
            message: "Login Error",
            description:
              error.response?.data?.error || "Login failed. Please try again.",
            duration: 3,
          });
        },
      }
    );
  };

  useEffect(() => {
    if (isAuth.data === true) {
      navigate("/dashboard");
    }
  }, [isAuth, navigate]);

  if (isAuth.isLoading) {
    return (
      <p
        style={{
          textAlign: "center",
          marginTop: "100px",
          fontSize: "30px",
          color: "red",
        }}
      >
        Loading...
      </p>
    );
  }

  return (
    <div>
      <AuthForm
        onSubmit={(email, password) => handleLogin(email, password)}
        buttonText="Login"
      />
      {mutation.isPending && <p>Loading...</p>}
    </div>
  );
};

export default Login;
