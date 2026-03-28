package service

import (
	"errors"
	"math"
	"time"

	"smart-study-planner/backend/internal/domain"
	"smart-study-planner/backend/internal/repository"
)

type TaskService struct {
	taskRepository *repository.TaskRepository
}

func NewTaskService(taskRepository *repository.TaskRepository) *TaskService {
	return &TaskService{taskRepository: taskRepository}
}

func (s *TaskService) CreateTask(userID int64, request domain.CreateTaskRequest) (*domain.TaskView, error) {
	deadline, err := parseDateInput(request.Deadline)
	if err != nil {
		return nil, err
	}

	task := &domain.Task{
		UserID:         userID,
		Title:          request.Title,
		Subject:        request.Subject,
		Deadline:       deadline,
		TotalHours:     request.TotalHours,
		CompletedHours: 0,
	}

	s.applyCalculatedFields(task)

	if err := s.taskRepository.Create(task); err != nil {
		return nil, err
	}

	s.applyCalculatedFields(task)
	return s.toTaskView(task), nil
}

func (s *TaskService) GetTasks(userID int64) ([]domain.TaskView, error) {
	tasks, err := s.taskRepository.ListByUser(userID)
	if err != nil {
		return nil, err
	}

	result := make([]domain.TaskView, 0, len(tasks))
	for _, item := range tasks {
		task := item
		s.applyCalculatedFields(&task)
		_ = s.taskRepository.Update(&task)
		result = append(result, *s.toTaskView(&task))
	}

	return result, nil
}

func (s *TaskService) GetTask(userID, taskID int64) (*domain.TaskView, error) {
	task, err := s.taskRepository.FindByID(taskID, userID)
	if err != nil {
		return nil, err
	}

	s.applyCalculatedFields(task)
	_ = s.taskRepository.Update(task)
	return s.toTaskView(task), nil
}

func (s *TaskService) UpdateTask(userID, taskID int64, request domain.UpdateTaskRequest) (*domain.TaskView, error) {
	task, err := s.taskRepository.FindByID(taskID, userID)
	if err != nil {
		return nil, err
	}

	if request.Title != nil {
		task.Title = *request.Title
	}
	if request.Subject != nil {
		task.Subject = *request.Subject
	}
	if request.Deadline != nil {
		deadline, err := parseDateInput(*request.Deadline)
		if err != nil {
			return nil, err
		}
		task.Deadline = deadline
	}
	if request.TotalHours != nil {
		task.TotalHours = math.Max(*request.TotalHours, task.CompletedHours)
	}
	if request.CompletedHours != nil {
		if *request.CompletedHours < 0 {
			return nil, errors.New("completed hours cannot be negative")
		}
		task.CompletedHours = math.Min(*request.CompletedHours, task.TotalHours)
	}

	s.applyCalculatedFields(task)

	if err := s.taskRepository.Update(task); err != nil {
		return nil, err
	}

	return s.toTaskView(task), nil
}

func (s *TaskService) RecalculateAndSave(task *domain.Task) error {
	s.applyCalculatedFields(task)
	return s.taskRepository.Update(task)
}

func (s *TaskService) applyCalculatedFields(task *domain.Task) {
	remainingHours := math.Max(task.TotalHours-task.CompletedHours, 0)
	daysLeft := calculateDaysLeft(task.Deadline)

	if remainingHours == 0 {
		task.DailyTarget = 0
		task.Status = "on-track"
		task.Urgency = calculateUrgency(daysLeft)
		return
	}

	task.DailyTarget = roundToTwoDecimals(remainingHours / float64(daysLeft))
	task.Urgency = calculateUrgency(daysLeft)

	expectedHours := calculateExpectedHours(task.TotalHours, task.CreatedAt, task.Deadline)
	if task.CompletedHours+0.01 < expectedHours {
		task.Status = "behind"
	} else {
		task.Status = "on-track"
	}
}

func (s *TaskService) toTaskView(task *domain.Task) *domain.TaskView {
	daysLeft := calculateDaysLeft(task.Deadline)
	remainingHours := roundToTwoDecimals(math.Max(task.TotalHours-task.CompletedHours, 0))
	progressPct := 0.0
	if task.TotalHours > 0 {
		progressPct = roundToTwoDecimals((task.CompletedHours / task.TotalHours) * 100)
	}

	expectedHours := roundToTwoDecimals(calculateExpectedHours(task.TotalHours, task.CreatedAt, task.Deadline))

	return &domain.TaskView{
		Task:           *task,
		DaysLeft:       daysLeft,
		RemainingHours: remainingHours,
		ProgressPct:    progressPct,
		ExpectedHours:  expectedHours,
		IsBehind:       task.Status == "behind",
	}
}

func calculateDaysLeft(deadline time.Time) int {
	today := truncateToDate(time.Now())
	deadlineDate := truncateToDate(deadline)

	if deadlineDate.Before(today) {
		return 1
	}

	difference := deadlineDate.Sub(today).Hours() / 24
	return int(math.Max(math.Ceil(difference), 1))
}

func calculateUrgency(daysLeft int) string {
	if daysLeft < 3 {
		return "high"
	}
	if daysLeft <= 7 {
		return "medium"
	}
	return "low"
}

func calculateExpectedHours(totalHours float64, createdAt, deadline time.Time) float64 {
	if createdAt.IsZero() {
		return 0
	}

	start := truncateToDate(createdAt)
	today := truncateToDate(time.Now())
	end := truncateToDate(deadline)

	totalDays := int(math.Max(end.Sub(start).Hours()/24, 1))
	elapsedDays := int(math.Max(today.Sub(start).Hours()/24, 0))

	expected := (totalHours / float64(totalDays)) * float64(elapsedDays)
	return math.Min(roundToTwoDecimals(expected), totalHours)
}

func truncateToDate(value time.Time) time.Time {
	return time.Date(value.Year(), value.Month(), value.Day(), 0, 0, 0, 0, value.Location())
}

func parseDateInput(value string) (time.Time, error) {
	// Accept the plain YYYY-MM-DD strings sent by HTML date inputs.
	if parsed, err := time.Parse("2006-01-02", value); err == nil {
		return parsed, nil
	}

	// Accept full timestamps too, so the API stays flexible.
	return time.Parse(time.RFC3339, value)
}

func roundToTwoDecimals(value float64) float64 {
	return math.Round(value*100) / 100
}
