import type { PropsWithChildren } from "react";
import { NavLink } from "react-router-dom";
import { useAuth } from "../context/AuthContext";

export function Layout({ children }: PropsWithChildren) {
  const { user, logoutUser } = useAuth();

  return (
    <div className="app-shell">
      <aside className="sidebar">
        <div>
          <p className="eyebrow">Smart Study Planner</p>
          <h1>Study with a real plan</h1>
          <p className="sidebar-text">
            Build daily momentum instead of guessing what to study next.
          </p>
        </div>

        <nav className="nav-links">
          <NavLink className={({ isActive }) => (isActive ? "active" : "")} to="/">
            Dashboard
          </NavLink>
          <NavLink className={({ isActive }) => (isActive ? "active" : "")} to="/add-task">
            Add Task
          </NavLink>
          <NavLink className={({ isActive }) => (isActive ? "active" : "")} to="/progress">
            Progress
          </NavLink>
          <NavLink className={({ isActive }) => (isActive ? "active" : "")} to="/timetable">
            Timetable
          </NavLink>
        </nav>

        <div className="user-panel">
          <p>Signed in as</p>
          <strong>{user?.name}</strong>
          <button className="secondary-button" onClick={logoutUser}>
            Logout
          </button>
        </div>
      </aside>

      <main className="content">{children}</main>
    </div>
  );
}
