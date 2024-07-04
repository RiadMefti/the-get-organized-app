import React, { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { useCheckAuth, useLogin } from "../hooks/useAuth";
import AuthForm from "../components/AuthForm";

const Login: React.FC = () => {
  const mutation = useLogin();
  const navigate = useNavigate();

  const isAuth = useCheckAuth();

  const handleLogin = (email: string, password: string) => {
    mutation.mutate(
      { email, password },
      {
        onSuccess: () => {
          navigate("/dashboard");
        },
        onError: () => {
          alert("Login failed");
        },
      }
    );
  };

  useEffect(() => {
    if (isAuth.data === true) {
      navigate("/dashboard");
    }
  }, [isAuth]);
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
      <h2>Login</h2>
      <AuthForm
        onSubmit={(email, password) => handleLogin(email, password)}
        buttonText="Login"
      />
      {mutation.isPending && <p>Loading...</p>}
    </div>
  );
};

export default Login;
