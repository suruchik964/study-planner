package repository

import (
	"database/sql"

	"smart-study-planner/backend/internal/domain"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Create(task *domain.Task) error {
	query := `
		INSERT INTO tasks (
			user_id, title, subject, deadline, total_hours, completed_hours, daily_target, urgency, status
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, created_at, updated_at
	`

	return r.db.QueryRow(
		query,
		task.UserID,
		task.Title,
		task.Subject,
		task.Deadline,
		task.TotalHours,
		task.CompletedHours,
		task.DailyTarget,
		task.Urgency,
		task.Status,
	).Scan(&task.ID, &task.CreatedAt, &task.UpdatedAt)
}

func (r *TaskRepository) Update(task *domain.Task) error {
	query := `
		UPDATE tasks
		SET title = $1,
			subject = $2,
			deadline = $3,
			total_hours = $4,
			completed_hours = $5,
			daily_target = $6,
			urgency = $7,
			status = $8,
			updated_at = NOW()
		WHERE id = $9 AND user_id = $10
		RETURNING updated_at
	`

	return r.db.QueryRow(
		query,
		task.Title,
		task.Subject,
		task.Deadline,
		task.TotalHours,
		task.CompletedHours,
		task.DailyTarget,
		task.Urgency,
		task.Status,
		task.ID,
		task.UserID,
	).Scan(&task.UpdatedAt)
}

func (r *TaskRepository) FindByID(id, userID int64) (*domain.Task, error) {
	query := `
		SELECT id, user_id, title, subject, deadline, total_hours, completed_hours, daily_target, urgency, status, created_at, updated_at
		FROM tasks
		WHERE id = $1 AND user_id = $2
	`

	task := &domain.Task{}
	err := r.db.QueryRow(query, id, userID).Scan(
		&task.ID,
		&task.UserID,
		&task.Title,
		&task.Subject,
		&task.Deadline,
		&task.TotalHours,
		&task.CompletedHours,
		&task.DailyTarget,
		&task.Urgency,
		&task.Status,
		&task.CreatedAt,
		&task.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (r *TaskRepository) ListByUser(userID int64) ([]domain.Task, error) {
	query := `
		SELECT id, user_id, title, subject, deadline, total_hours, completed_hours, daily_target, urgency, status, created_at, updated_at
		FROM tasks
		WHERE user_id = $1
		ORDER BY deadline ASC, created_at ASC
	`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []domain.Task{}
	for rows.Next() {
		var task domain.Task
		if err := rows.Scan(
			&task.ID,
			&task.UserID,
			&task.Title,
			&task.Subject,
			&task.Deadline,
			&task.TotalHours,
			&task.CompletedHours,
			&task.DailyTarget,
			&task.Urgency,
			&task.Status,
			&task.CreatedAt,
			&task.UpdatedAt,
		); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, rows.Err()
}

