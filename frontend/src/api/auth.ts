import { apiRequest } from "./client";
import type { AuthResponse } from "../types";

export function register(payload: { name: string; email: string; password: string }) {
  return apiRequest<AuthResponse>("/auth/register", {
    method: "POST",
    body: payload,
  });
}

export function login(payload: { email: string; password: string }) {
  return apiRequest<AuthResponse>("/auth/login", {
    method: "POST",
    body: payload,
  });
}

