package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"smart-study-planner/backend/internal/delivery/http/handlers"
	"smart-study-planner/backend/internal/delivery/http/middleware"
	"smart-study-planner/backend/internal/service"
)

func SetupRouter(
	authService *service.AuthService,
	taskService *service.TaskService,
	sessionService *service.SessionService,
	jwtSecret string,
) *gin.Engine {
	router := gin.Default()

	// Simple CORS middleware so the React app can call the Go API during development.
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS")
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Smart Study Planner API is healthy"})
	})

	authHandler := handlers.NewAuthHandler(authService)
	taskHandler := handlers.NewTaskHandler(taskService)
	sessionHandler := handlers.NewSessionHandler(sessionService)
	authMiddleware := middleware.NewAuthMiddleware(jwtSecret)

	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/register", authHandler.Register)
		authRoutes.POST("/login", authHandler.Login)
	}

	protectedRoutes := router.Group("/")
	protectedRoutes.Use(authMiddleware.Handle())
	{
		protectedRoutes.POST("/tasks", taskHandler.CreateTask)
		protectedRoutes.GET("/tasks", taskHandler.ListTasks)
		protectedRoutes.GET("/tasks/:id", taskHandler.GetTask)
		protectedRoutes.PUT("/tasks/:id", taskHandler.UpdateTask)
		protectedRoutes.POST("/sessions", sessionHandler.CreateSession)
	}

	return router
}

