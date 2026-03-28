package domain

import "time"

// StudySession tracks hours studied for one task on one date.
type StudySession struct {
	ID          int64     `json:"id"`
	TaskID       int64     `json:"task_id"`
	UserID       int64     `json:"user_id"`
	SessionDate  time.Time `json:"session_date"`
	HoursStudied float64   `json:"hours_studied"`
	CreatedAt    time.Time `json:"created_at"`
}

type CreateStudySessionRequest struct {
	TaskID       int64   `json:"task_id" binding:"required"`
	SessionDate  string  `json:"session_date" binding:"required"`
	HoursStudied float64 `json:"hours_studied" binding:"required,gt=0"`
}
