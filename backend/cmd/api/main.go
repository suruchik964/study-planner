package main

import (
	"log"

	"github.com/joho/godotenv"

	"smart-study-planner/backend/internal/config"
	httpdelivery "smart-study-planner/backend/internal/delivery/http"
	"smart-study-planner/backend/internal/repository"
	"smart-study-planner/backend/internal/service"
)

func main() {
	// Load .env file when it exists. If it does not exist, normal environment variables still work.
	_ = godotenv.Load()

	cfg := config.Load()

	db, err := repository.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)
	sessionRepository := repository.NewSessionRepository(db)

	authService := service.NewAuthService(userRepository, cfg.JWTSecret)
	taskService := service.NewTaskService(taskRepository)
	sessionService := service.NewSessionService(sessionRepository, taskRepository, taskService)

	router := httpdelivery.SetupRouter(authService, taskService, sessionService, cfg.JWTSecret)

	log.Printf("backend is running on http://localhost:%s", cfg.AppPort)
	if err := router.Run(":" + cfg.AppPort); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

