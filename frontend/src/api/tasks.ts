import { apiRequest } from "./client";
import type { SessionPayload, Task, TaskListResponse } from "../types";

export function getTasks(token: string) {
  return apiRequest<TaskListResponse>("/tasks", { token });
}

export function getTask(token: string, taskId: number) {
  return apiRequest<Task>(`/tasks/${taskId}`, { token });
}

export function createTask(
  token: string,
  payload: { title: string; subject: string; deadline: string; total_hours: number }
) {
  return apiRequest<Task>("/tasks", {
    method: "POST",
    token,
    body: payload,
  });
}

export function updateTask(
  token: string,
  taskId: number,
  payload: Partial<{ title: string; subject: string; deadline: string; total_hours: number; completed_hours: number }>
) {
  return apiRequest<Task>(`/tasks/${taskId}`, {
    method: "PUT",
    token,
    body: payload,
  });
}

export function createSession(token: string, payload: SessionPayload) {
  return apiRequest<Task>("/sessions", {
    method: "POST",
    token,
    body: payload,
  });
}

