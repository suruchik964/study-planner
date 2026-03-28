import { useState } from "react";
import type { FormEvent } from "react";
import { useNavigate } from "react-router-dom";
import { login, register } from "../api/auth";
import { useAuth } from "../context/AuthContext";

export function AuthPage() {
  const navigate = useNavigate();
  const { loginUser } = useAuth();

  const [mode, setMode] = useState<"login" | "register">("login");
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);

  async function handleSubmit(event: FormEvent) {
    event.preventDefault();
    setError("");
    setLoading(true);

    try {
      const response =
        mode === "login"
          ? await login({ email, password })
          : await register({ name, email, password });

      loginUser(response.token, response.user);
      navigate("/");
    } catch (submissionError) {
      setError(submissionError instanceof Error ? submissionError.message : "Authentication failed");
    } finally {
      setLoading(false);
    }
  }

  return (
    <div className="auth-page">
      <section className="auth-hero">
        <p className="eyebrow">Smart Study Planner</p>
        <h1>Turn big study goals into daily targets.</h1>
        <p>
          Add your deadlines, track your hours, and let the planner raise the pressure only when
          you actually need it.
        </p>
      </section>

      <section className="auth-card">
        <div className="auth-tabs">
          <button
            className={mode === "login" ? "tab active-tab" : "tab"}
            onClick={() => setMode("login")}
            type="button"
          >
            Login
          </button>
          <button
            className={mode === "register" ? "tab active-tab" : "tab"}
            onClick={() => setMode("register")}
            type="button"
          >
            Register
          </button>
        </div>

        <form className="auth-form" onSubmit={handleSubmit}>
          {mode === "register" && (
            <label>
              Name
              <input value={name} onChange={(event) => setName(event.target.value)} required />
            </label>
          )}

          <label>
            Email
            <input
              type="email"
              value={email}
              onChange={(event) => setEmail(event.target.value)}
              required
            />
          </label>

          <label>
            Password
            <input
              type="password"
              value={password}
              onChange={(event) => setPassword(event.target.value)}
              minLength={6}
              required
            />
          </label>

          {error && <p className="error-text">{error}</p>}

          <button className="primary-button" disabled={loading} type="submit">
            {loading ? "Please wait..." : mode === "login" ? "Login" : "Create account"}
          </button>
        </form>
      </section>
    </div>
  );
}
