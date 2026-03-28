package service

import (
	"errors"
	"math"

	"smart-study-planner/backend/internal/domain"
	"smart-study-planner/backend/internal/repository"
)

type SessionService struct {
	sessionRepository *repository.SessionRepository
	taskRepository    *repository.TaskRepository
	taskService       *TaskService
}

func NewSessionService(
	sessionRepository *repository.SessionRepository,
	taskRepository *repository.TaskRepository,
	taskService *TaskService,
) *SessionService {
	return &SessionService{
		sessionRepository: sessionRepository,
		taskRepository:    taskRepository,
		taskService:       taskService,
	}
}

func (s *SessionService) CreateSession(userID int64, request domain.CreateStudySessionRequest) (*domain.TaskView, error) {
	task, err := s.taskRepository.FindByID(request.TaskID, userID)
	if err != nil {
		return nil, err
	}

	sessionDate, err := parseDateInput(request.SessionDate)
	if err != nil {
		return nil, err
	}

	session := &domain.StudySession{
		TaskID:       request.TaskID,
		UserID:       userID,
		SessionDate:  sessionDate,
		HoursStudied: request.HoursStudied,
	}

	if err := s.sessionRepository.Create(session); err != nil {
		return nil, err
	}

	task.CompletedHours = math.Min(task.CompletedHours+request.HoursStudied, task.TotalHours)

	if err := s.taskService.RecalculateAndSave(task); err != nil {
		return nil, err
	}

	view, err := s.taskService.GetTask(userID, task.ID)
	if err != nil {
		return nil, err
	}

	return view, nil
}

func (s *SessionService) ValidateOwnership(userID, taskID int64) error {
	_, err := s.taskRepository.FindByID(taskID, userID)
	if err != nil {
		return errors.New("task not found for this user")
	}
	return nil
}
