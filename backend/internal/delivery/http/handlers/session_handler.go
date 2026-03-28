package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"smart-study-planner/backend/internal/domain"
	"smart-study-planner/backend/internal/service"
)

type SessionHandler struct {
	sessionService *service.SessionService
}

func NewSessionHandler(sessionService *service.SessionService) *SessionHandler {
	return &SessionHandler{sessionService: sessionService}
}

func (h *SessionHandler) CreateSession(c *gin.Context) {
	var request domain.CreateStudySessionRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetInt64("user_id")

	task, err := h.sessionService.CreateSession(userID, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, task)
}

