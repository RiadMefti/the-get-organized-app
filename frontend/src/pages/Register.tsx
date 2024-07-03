import React, { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { useRegister } from "../hooks/useAuth";
import AuthForm from "../components/AuthForm";

const Register: React.FC = () => {
  const mutation = useRegister();
  const navigate = useNavigate();

  const handleRegister = (
    email: string,
    password: string,
    copyPassword?: string
  ) => {
    if (copyPassword) {
      mutation.mutate(
        { email, password, copyPassword },
        {
          onSuccess: () => {
            navigate("/login");
          },
          onError: () => {
            alert("Registration failed");
          },
        }
      );
    }
  };

  useEffect(() => {
    if (localStorage.getItem("token")) {
      navigate("/dashboard");
    }
  }, [navigate]); // Dependency array includes navigate to handle changes in navigate function

  return (
    <div>
      <h2>Register</h2>
      <AuthForm onSubmit={handleRegister} buttonText="Register" />
      {mutation.isPending && <p>Loading...</p>}
    </div>
  );
};

export default Register;
