package domain

import "time"

// Task stores one study goal created by a user.
type Task struct {
	ID             int64     `json:"id"`
	UserID         int64     `json:"user_id"`
	Title          string    `json:"title"`
	Subject        string    `json:"subject"`
	Deadline       time.Time `json:"deadline"`
	TotalHours     float64   `json:"total_hours"`
	CompletedHours float64   `json:"completed_hours"`
	DailyTarget    float64   `json:"daily_target"`
	Urgency        string    `json:"urgency"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// TaskView includes calculated values used directly by the frontend.
type TaskView struct {
	Task
	DaysLeft       int     `json:"days_left"`
	RemainingHours float64 `json:"remaining_hours"`
	ProgressPct    float64 `json:"progress_pct"`
	ExpectedHours  float64 `json:"expected_hours"`
	IsBehind       bool    `json:"is_behind"`
}

type CreateTaskRequest struct {
	Title      string  `json:"title" binding:"required"`
	Subject    string  `json:"subject" binding:"required"`
	Deadline   string  `json:"deadline" binding:"required"`
	TotalHours float64 `json:"total_hours" binding:"required,gt=0"`
}

type UpdateTaskRequest struct {
	Title          *string  `json:"title"`
	Subject        *string  `json:"subject"`
	Deadline       *string  `json:"deadline"`
	TotalHours     *float64 `json:"total_hours"`
	CompletedHours *float64 `json:"completed_hours"`
}
