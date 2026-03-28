import { useEffect, useState } from "react";
import { getTasks } from "../api/tasks";
import { useAuth } from "../context/AuthContext";
import type { Task } from "../types";

export function TimetablePage() {
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
        setError(loadError instanceof Error ? loadError.message : "Failed to load timetable");
      }
    }

    load();
  }, [token]);

  return (
    <section className="page-stack">
      <div className="page-header">
        <div>
          <p className="eyebrow">Timetable</p>
          <h2>Daily study plan</h2>
        </div>
      </div>

      {error && <p className="error-text">{error}</p>}

      <div className="panel timetable-list">
        {tasks.map((task) => (
          <div className="timetable-row" key={task.id}>
            <div>
              <h3>{task.title}</h3>
              <p>
                {task.subject} • deadline {String(task.deadline).slice(0, 10)}
              </p>
            </div>
            <div className="timetable-hours">
              <strong>{task.daily_target} hrs/day</strong>
              <span className={`badge badge-${task.urgency}`}>{task.urgency}</span>
            </div>
          </div>
        ))}

        {tasks.length === 0 ? <p>No timetable yet. Add a task to generate one.</p> : null}
      </div>
    </section>
  );
}
