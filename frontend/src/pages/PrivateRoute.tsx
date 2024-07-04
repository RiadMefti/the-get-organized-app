import React, { useEffect } from "react";
import { Navigate, Outlet, useNavigate } from "react-router-dom";
import { useCheckAuth } from "../hooks/useAuth";

interface PrivateRouteProps {
  redirectPath?: string;
  children?: React.ReactNode;
}

const PrivateRoute: React.FC<PrivateRouteProps> = ({
  redirectPath = "/login",
  children,
}) => {
  const navigate = useNavigate();

  const isAuth = useCheckAuth();
  useEffect(() => {
    if (isAuth.data === false) {
      navigate(redirectPath);
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
  return children ? <>{children}</> : <Outlet />;
};

export default PrivateRoute;
