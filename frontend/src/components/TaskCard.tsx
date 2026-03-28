import type { Task } from "../types";

type TaskCardProps = {
  task: Task;
  onLogStudy: (taskId: number, hoursStudied: number) => void;
};

export function TaskCard({ task, onLogStudy }: TaskCardProps) {
  return (
    <article className="task-card">
      <div className="task-card-header">
        <div>
          <p className="task-subject">{task.subject}</p>
          <h3>{task.title}</h3>
        </div>
        <span className={`badge badge-${task.urgency}`}>{task.urgency}</span>
      </div>

      <div className="task-grid">
        <p>Days left: {task.days_left}</p>
        <p>Daily target: {task.daily_target} hrs</p>
        <p>Completed: {task.completed_hours} hrs</p>
        <p>Remaining: {task.remaining_hours} hrs</p>
      </div>

      <div className="progress-block">
        <div className="progress-label">
          <span>Progress</span>
          <span>{task.progress_pct}%</span>
        </div>
        <div className="progress-bar">
          <div className="progress-fill" style={{ width: `${Math.min(task.progress_pct, 100)}%` }} />
        </div>
      </div>

      {task.is_behind ? (
        <p className="warning-text">
          Warning: you are behind. Expected {task.expected_hours} hours by now.
        </p>
      ) : (
        <p className="success-text">You are on track for this task.</p>
      )}

      <button className="primary-button" onClick={() => onLogStudy(task.id, task.daily_target || 1)}>
        Log today&apos;s target
      </button>
    </article>
  );
}

