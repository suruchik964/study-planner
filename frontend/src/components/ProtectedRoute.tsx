import type { PropsWithChildren } from "react";
import { Navigate } from "react-router-dom";
import { useAuth } from "../context/AuthContext";

export function ProtectedRoute({ children }: PropsWithChildren) {
  const { isReady, token } = useAuth();

  if (!isReady) {
    return <p>Loading session...</p>;
  }

  if (!token) {
    return <Navigate to="/auth" replace />;
  }

  return <>{children}</>;
}
