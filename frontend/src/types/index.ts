export type User = {
  id: number;
  name: string;
  email: string;
  created_at: string;
  updated_at: string;
};

export type Task = {
  id: number;
  user_id: number;
  title: string;
  subject: string;
  deadline: string;
  total_hours: number;
  completed_hours: number;
  daily_target: number;
  urgency: "low" | "medium" | "high";
  status: "on-track" | "behind";
  created_at: string;
  updated_at: string;
  days_left: number;
  remaining_hours: number;
  progress_pct: number;
  expected_hours: number;
  is_behind: boolean;
};

export type AuthResponse = {
  token: string;
  user: User;
};

export type TaskListResponse = {
  tasks: Task[];
};

export type SessionPayload = {
  task_id: number;
  session_date: string;
  hours_studied: number;
};

