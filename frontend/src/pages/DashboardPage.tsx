import { useEffect, useMemo, useState } from "react";
import { createSession, getTasks } from "../api/tasks";
import { TaskCard } from "../components/TaskCard";
import { useAuth } from "../context/AuthContext";
import type { Task } from "../types";

export function DashboardPage() {
  const { token } = useAuth();
  const [tasks, setTasks] = useState<Task[]>([]);
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(true);

  async function loadTasks() {
    if (!token) {
      return;
    }

    try {
      setLoading(true);
      const response = await getTasks(token);
      setTasks(response.tasks);
    } catch (loadError) {
      setError(loadError instanceof Error ? loadError.message : "Failed to load tasks");
    } finally {
      setLoading(false);
    }
  }

  useEffect(() => {
    loadTasks();
  }, [token]);

  const todaysStudyTarget = useMemo(
    () => tasks.reduce((total, task) => total + Number(task.daily_target || 0), 0).toFixed(2),
    [tasks]
  );

  async function handleLogStudy(taskId: number, hoursStudied: number) {
    if (!token) {
      return;
    }

    const now = new Date();
    const sessionDate = `${now.getFullYear()}-${String(now.getMonth() + 1).padStart(2, "0")}-${String(
      now.getDate()
    ).padStart(2, "0")}`;

    try {
      await createSession(token, {
        task_id: taskId,
        session_date: sessionDate,
        hours_studied: Number(hoursStudied.toFixed(2)),
      });
      await loadTasks();
    } catch (sessionError) {
      setError(sessionError instanceof Error ? sessionError.message : "Failed to save session");
    }
  }

  return (
    <section className="page-stack">
      <div className="page-header">
        <div>
          <p className="eyebrow">Dashboard</p>
          <h2>Your study plan</h2>
        </div>
        <div className="target-panel">
          <span>Today&apos;s Study Target</span>
          <strong>{todaysStudyTarget} hrs</strong>
        </div>
      </div>

      {error && <p className="error-text">{error}</p>}
      {loading ? <p>Loading tasks...</p> : null}

      <div className="task-list">
        {tasks.map((task) => (
          <TaskCard key={task.id} task={task} onLogStudy={handleLogStudy} />
        ))}
      </div>

      {!loading && tasks.length === 0 ? (
        <div className="empty-panel">
          <h3>No tasks yet</h3>
          <p>Create your first task to start generating a study plan.</p>
        </div>
      ) : null}
    </section>
  );
}
