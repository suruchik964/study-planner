import { createContext, useContext, useEffect, useState } from "react";
import type { PropsWithChildren } from "react";
import type { User } from "../types";

type AuthState = {
  isReady: boolean;
  token: string | null;
  user: User | null;
  loginUser: (token: string, user: User) => void;
  logoutUser: () => void;
};

const AuthContext = createContext<AuthState | undefined>(undefined);

const STORAGE_KEY = "smart-study-planner-auth";

export function AuthProvider({ children }: PropsWithChildren) {
  const [isReady, setIsReady] = useState(false);
  const [token, setToken] = useState<string | null>(null);
  const [user, setUser] = useState<User | null>(null);

  // Keep the user logged in after page refresh.
  useEffect(() => {
    const stored = localStorage.getItem(STORAGE_KEY);
    if (!stored) {
      return;
    }

    const parsed = JSON.parse(stored) as { token: string; user: User };
    setToken(parsed.token);
    setUser(parsed.user);
    setIsReady(true);
  }, []);

  useEffect(() => {
    if (localStorage.getItem(STORAGE_KEY) === null) {
      setIsReady(true);
    }
  }, []);

  function loginUser(nextToken: string, nextUser: User) {
    setToken(nextToken);
    setUser(nextUser);
    localStorage.setItem(STORAGE_KEY, JSON.stringify({ token: nextToken, user: nextUser }));
  }

  function logoutUser() {
    setToken(null);
    setUser(null);
    localStorage.removeItem(STORAGE_KEY);
  }

  return (
    <AuthContext.Provider value={{ isReady, token, user, loginUser, logoutUser }}>
      {children}
    </AuthContext.Provider>
  );
}

export function useAuth() {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error("useAuth must be used inside AuthProvider");
  }
  return context;
}
