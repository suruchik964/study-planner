# Project Guide

This file explains every file created in this project and why it exists.

## Root files

- `.gitignore`: prevents generated files like `node_modules` and `.env` files from being committed.
- `docker-compose.yml`: starts PostgreSQL in a container so you do not need to install PostgreSQL manually.
- `README.md`: beginner-friendly setup instructions.
- `PROJECT_GUIDE.md`: this explanation file.

## Backend files

- `backend/.env.example`: sample environment variables for the backend.
- `backend/go.mod`: tells Go which project this is and which libraries it uses.
- `backend/cmd/api/main.go`: backend entry point; it loads config, connects to the database, and starts Gin.
- `backend/internal/config/config.go`: reads environment variables.
- `backend/internal/domain/user.go`: user model used by the application.
- `backend/internal/domain/task.go`: task model and helper request/response shapes.
- `backend/internal/domain/study_session.go`: study session model.
- `backend/internal/repository/postgres.go`: opens the PostgreSQL database connection.
- `backend/internal/repository/user_repository.go`: database queries for users.
- `backend/internal/repository/task_repository.go`: database queries for tasks.
- `backend/internal/repository/session_repository.go`: database queries for study sessions.
- `backend/internal/service/auth_service.go`: signup/login business logic, password hashing, token creation.
- `backend/internal/service/task_service.go`: task rules like days left, urgency, daily target, and behind detection.
- `backend/internal/service/session_service.go`: stores study sessions and updates task progress.
- `backend/internal/utils/password.go`: password hashing and password comparison helpers.
- `backend/internal/utils/jwt.go`: creates and reads JWT tokens.
- `backend/internal/delivery/http/router.go`: all HTTP routes are registered here.
- `backend/internal/delivery/http/handlers/auth_handler.go`: handles register and login requests.
- `backend/internal/delivery/http/handlers/task_handler.go`: handles task routes.
- `backend/internal/delivery/http/handlers/session_handler.go`: handles session routes.
- `backend/internal/delivery/http/middleware/auth_middleware.go`: checks JWT tokens before protected routes.
- `backend/migrations/001_init.sql`: creates the PostgreSQL tables.

## Frontend files

- `frontend/.env.example`: sample environment variables for the frontend.
- `frontend/package.json`: frontend dependencies and scripts.
- `frontend/tsconfig.json`: TypeScript compiler settings.
- `frontend/tsconfig.node.json`: TypeScript settings for build tools like Vite.
- `frontend/vite.config.ts`: Vite development server config.
- `frontend/index.html`: HTML shell where React mounts.
- `frontend/src/main.tsx`: starts the React app.
- `frontend/src/App.tsx`: defines the app routes and page layout.
- `frontend/src/styles.css`: all app styling.
- `frontend/src/types/index.ts`: shared TypeScript types for users, tasks, and sessions.
- `frontend/src/api/client.ts`: low-level function for talking to the backend API.
- `frontend/src/api/auth.ts`: helper functions for register/login.
- `frontend/src/api/tasks.ts`: helper functions for task and session requests.
- `frontend/src/context/AuthContext.tsx`: stores login state for the whole app.
- `frontend/src/components/Layout.tsx`: main navigation layout.
- `frontend/src/components/ProtectedRoute.tsx`: stops logged-out users from opening protected pages.
- `frontend/src/components/TaskCard.tsx`: reusable task card shown on the dashboard.
- `frontend/src/pages/AuthPage.tsx`: login and register screen.
- `frontend/src/pages/DashboardPage.tsx`: task list, daily targets, and warnings.
- `frontend/src/pages/AddTaskPage.tsx`: form to create a new study task.
- `frontend/src/pages/ProgressPage.tsx`: stats and charts.
- `frontend/src/pages/TimetablePage.tsx`: daily study plan view.

## Simple concept explanations

### Clean architecture

This means code is split by responsibility:

- `domain`: what the data looks like
- `repository`: how data is saved/read from the database
- `service`: business rules
- `delivery`: HTTP request/response layer

This structure keeps the project easier to grow later.

### JWT authentication

JWT is a signed token. After login, the backend gives the frontend a token. The frontend sends it with future requests, and the backend checks it to know which user is making the request.

### Password hashing

Passwords should never be saved as plain text. Hashing turns the password into a secure form before storing it.

### Daily target and pressure system

If a task has not received enough study hours yet, the remaining hours are divided by the remaining days. That makes the daily target increase automatically when the student falls behind.

## Short summaries

- Backend summary: secure API with task planning logic.
- Frontend summary: visual app for managing tasks and progress.
- Database summary: stores users, tasks, and study sessions with relationships.
