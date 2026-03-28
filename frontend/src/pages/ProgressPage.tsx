import { useEffect, useMemo, useState } from "react";
import { getTasks } from "../api/tasks";
import { useAuth } from "../context/AuthContext";
import type { Task } from "../types";

export function ProgressPage() {
  const { token } = useAuth();
  const [tasks, setTasks] = useState<Task[]>([]);
  const [error, setError] = useState("");

  useEffect(() => {
    async function load() {
      if (!token) {
        return;
      }

      try {
        const response = await getTasks(token);
        setTasks(response.tasks);
      } catch (loadError) {
        setError(loadError instanceof Error ? loadError.message : "Failed to load progress");
      }
    }

    load();
  }, [token]);

  const stats = useMemo(() => {
    const totalCompleted = tasks.reduce((sum, task) => sum + task.completed_hours, 0);
    const totalRemaining = tasks.reduce((sum, task) => sum + task.remaining_hours, 0);
    const totalHours = tasks.reduce((sum, task) => sum + task.total_hours, 0);
    const progressPercentage = totalHours === 0 ? 0 : (totalCompleted / totalHours) * 100;

    return {
      totalCompleted: totalCompleted.toFixed(2),
      totalRemaining: totalRemaining.toFixed(2),
      progressPercentage: progressPercentage.toFixed(0),
    };
  }, [tasks]);

  return (
    <section className="page-stack">
      <div className="page-header">
        <div>
          <p className="eyebrow">Progress</p>
          <h2>Study statistics</h2>
        </div>
      </div>

      {error && <p className="error-text">{error}</p>}

      <div className="stats-grid">
        <div className="panel">
          <span>Total completed</span>
          <strong>{stats.totalCompleted} hrs</strong>
        </div>
        <div className="panel">
          <span>Total remaining</span>
          <strong>{stats.totalRemaining} hrs</strong>
        </div>
        <div className="panel">
          <span>Overall progress</span>
          <strong>{stats.progressPercentage}%</strong>
        </div>
      </div>

      <div className="chart-grid">
        <div className="panel">
          <h3>Hours by task</h3>
          <div className="bar-chart">
            {tasks.map((task) => (
              <div className="bar-row" key={task.id}>
                <span>{task.title}</span>
                <div className="bar-track">
                  <div className="bar-fill" style={{ width: `${Math.min(task.progress_pct, 100)}%` }} />
                </div>
                <span>{task.progress_pct}%</span>
              </div>
            ))}
          </div>
        </div>

        <div className="panel">
          <h3>Completed vs remaining</h3>
          <div
            className="pie-chart"
            style={{
              background: `conic-gradient(#1f6feb 0% ${stats.progressPercentage}%, #dce6f5 ${stats.progressPercentage}% 100%)`,
            }}
          >
            <div className="pie-center">{stats.progressPercentage}%</div>
          </div>
        </div>
      </div>
    </section>
  );
}

