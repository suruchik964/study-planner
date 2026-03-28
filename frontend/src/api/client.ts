const API_URL = import.meta.env.VITE_API_URL || "http://localhost:8080";

type RequestOptions = {
  method?: string;
  token?: string | null;
  body?: unknown;
};

// This helper keeps all fetch logic in one place.
export async function apiRequest<T>(path: string, options: RequestOptions = {}): Promise<T> {
  const response = await fetch(`${API_URL}${path}`, {
    method: options.method ?? "GET",
    headers: {
      "Content-Type": "application/json",
      ...(options.token ? { Authorization: `Bearer ${options.token}` } : {}),
    },
    body: options.body ? JSON.stringify(options.body) : undefined,
  });

  if (!response.ok) {
    const data = await response.json().catch(() => ({ error: "Request failed" }));
    throw new Error(data.error || "Request failed");
  }

  return response.json() as Promise<T>;
}

