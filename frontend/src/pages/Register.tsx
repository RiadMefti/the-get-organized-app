import React, { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { useCheckAuth, useRegister } from "../hooks/useAuth";
import AuthForm from "../components/AuthForm";
import { notification, Typography } from "antd";

const Register: React.FC = () => {
  const mutation = useRegister();
  const navigate = useNavigate();
  const isAuth = useCheckAuth();
  const { Title } = Typography;
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
            navigate("/");
          },
          onError: (error: any) => {
            notification.error({
              message: "Register Error",
              description:
                error.response?.data?.error ||
                "REgister failed. Please try again.",
              duration: 3,
            });
          },
        }
      );
    }
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
    <div
      style={{
        width: "100vw",
      }}
    >
      <AuthForm onSubmit={handleRegister} buttonText="Register" />
      {mutation.isPending && <p>Loading...</p>}
    </div>
  );
};

export default Register;
