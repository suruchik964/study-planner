CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(120) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS tasks (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    subject VARCHAR(255) NOT NULL,
    deadline DATE NOT NULL,
    total_hours NUMERIC(10,2) NOT NULL CHECK (total_hours > 0),
    completed_hours NUMERIC(10,2) NOT NULL DEFAULT 0 CHECK (completed_hours >= 0),
    daily_target NUMERIC(10,2) NOT NULL DEFAULT 0,
    urgency VARCHAR(20) NOT NULL DEFAULT 'low' CHECK (urgency IN ('low', 'medium', 'high')),
    status VARCHAR(20) NOT NULL DEFAULT 'on-track' CHECK (status IN ('on-track', 'behind')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS study_sessions (
    id BIGSERIAL PRIMARY KEY,
    task_id BIGINT NOT NULL REFERENCES tasks(id) ON DELETE CASCADE,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    session_date DATE NOT NULL,
    hours_studied NUMERIC(10,2) NOT NULL CHECK (hours_studied > 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_tasks_user_id ON tasks(user_id);
CREATE INDEX IF NOT EXISTS idx_tasks_deadline ON tasks(deadline);
CREATE INDEX IF NOT EXISTS idx_study_sessions_task_id ON study_sessions(task_id);
CREATE INDEX IF NOT EXISTS idx_study_sessions_user_id ON study_sessions(user_id);
