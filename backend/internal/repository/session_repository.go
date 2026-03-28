package repository

import (
	"database/sql"

	"smart-study-planner/backend/internal/domain"
)

type SessionRepository struct {
	db *sql.DB
}

func NewSessionRepository(db *sql.DB) *SessionRepository {
	return &SessionRepository{db: db}
}

func (r *SessionRepository) Create(session *domain.StudySession) error {
	query := `
		INSERT INTO study_sessions (task_id, user_id, session_date, hours_studied)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at
	`

	return r.db.QueryRow(
		query,
		session.TaskID,
		session.UserID,
		session.SessionDate,
		session.HoursStudied,
	).Scan(&session.ID, &session.CreatedAt)
}

