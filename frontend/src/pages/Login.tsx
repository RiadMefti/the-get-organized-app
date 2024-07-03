import React, { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { useLogin } from "../hooks/useAuth";
import AuthForm from "../components/AuthForm";

const Login: React.FC = () => {
  const mutation = useLogin();
  const navigate = useNavigate();

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
    if (localStorage.getItem("token")) {
      navigate("/dashboard");
    }
  }, [navigate]); // Dependency array includes navigate to handle changes in navigate function

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
