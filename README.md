# Smart Study Planner

Smart Study Planner is a full-stack learning project with:

- A Go + Gin backend
- A PostgreSQL database running in Docker
- A React + TypeScript frontend

This repository is structured to be easy to read if you are a beginner.

## What this project does

The app helps a student:

- create an account
- log in securely
- add study tasks
- track hours studied each day
- automatically calculate how much to study per day
- see when a task is getting risky or falling behind

## Project folders

- `backend/`: Go API server
- `frontend/`: React user interface
- `docker-compose.yml`: starts PostgreSQL in Docker
- `PROJECT_GUIDE.md`: explains each file and major concepts in simple words

## Beginner setup

### 1. Install tools

You need these tools on your computer:

- Go 1.22 or newer
- Node.js 20 or newer
- Docker Desktop

### 2. Start PostgreSQL

From the project root:

```powershell
docker-compose up -d
```

This starts PostgreSQL on port `5432`.

### 3. Start the backend

```powershell
cd backend
copy .env.example .env
go mod tidy
go run ./cmd/api
```

The backend will run on `http://localhost:8080`.

### 4. Start the frontend

Open a second terminal:

```powershell
cd frontend
copy .env.example .env
npm install
npm run dev
```

The frontend will run on `http://localhost:5173`.

## API routes

- `POST /auth/register`
- `POST /auth/login`
- `POST /tasks`
- `GET /tasks`
- `GET /tasks/:id`
- `PUT /tasks/:id`
- `POST /sessions`

## Short summary

- Backend: handles security, database access, and study logic
- Frontend: shows forms, dashboard cards, charts, and timetable
- Database: stores users, tasks, and study sessions

