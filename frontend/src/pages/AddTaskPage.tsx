import { useState } from "react";
import type { FormEvent } from "react";
import { useNavigate } from "react-router-dom";
import { createTask } from "../api/tasks";
import { useAuth } from "../context/AuthContext";

export function AddTaskPage() {
  const { token } = useAuth();
  const navigate = useNavigate();

  const [title, setTitle] = useState("");
  const [subject, setSubject] = useState("");
  const [deadline, setDeadline] = useState("");
  const [totalHours, setTotalHours] = useState("");
  const [error, setError] = useState("");
  const [saving, setSaving] = useState(false);

  async function handleSubmit(event: FormEvent) {
    event.preventDefault();
    if (!token) {
      return;
    }

    try {
      setSaving(true);
      setError("");
      await createTask(token, {
        title,
        subject,
        deadline,
        total_hours: Number(totalHours),
      });
      navigate("/");
    } catch (submissionError) {
      setError(submissionError instanceof Error ? submissionError.message : "Failed to create task");
    } finally {
      setSaving(false);
    }
  }

  return (
    <section className="page-stack">
      <div className="page-header">
        <div>
          <p className="eyebrow">Add Task</p>
          <h2>Create a new study goal</h2>
        </div>
      </div>

      <form className="panel form-panel" onSubmit={handleSubmit}>
        <label>
          Task title
          <input value={title} onChange={(event) => setTitle(event.target.value)} required />
        </label>

        <label>
          Subject
          <input value={subject} onChange={(event) => setSubject(event.target.value)} required />
        </label>

        <label>
          Deadline
          <input type="date" value={deadline} onChange={(event) => setDeadline(event.target.value)} required />
        </label>

        <label>
          Total hours needed
          <input
            type="number"
            min="1"
            step="0.5"
            value={totalHours}
            onChange={(event) => setTotalHours(event.target.value)}
            required
          />
        </label>

        {error && <p className="error-text">{error}</p>}

        <button className="primary-button" disabled={saving} type="submit">
          {saving ? "Saving..." : "Create task"}
        </button>
      </form>
    </section>
  );
}
