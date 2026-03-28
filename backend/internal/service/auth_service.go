package service

import (
	"database/sql"
	"errors"
	"strings"

	"smart-study-planner/backend/internal/domain"
	"smart-study-planner/backend/internal/repository"
	"smart-study-planner/backend/internal/utils"
)

type AuthService struct {
	userRepository *repository.UserRepository
	jwtSecret      string
}

func NewAuthService(userRepository *repository.UserRepository, jwtSecret string) *AuthService {
	return &AuthService{userRepository: userRepository, jwtSecret: jwtSecret}
}

func (s *AuthService) Register(request domain.RegisterRequest) (*domain.AuthResponse, error) {
	existingUser, err := s.userRepository.FindByEmail(strings.ToLower(request.Email))
	if err == nil && existingUser != nil {
		return nil, errors.New("email already exists")
	}
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	passwordHash, err := utils.HashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Name:         request.Name,
		Email:        strings.ToLower(request.Email),
		PasswordHash: passwordHash,
	}

	if err := s.userRepository.Create(user); err != nil {
		return nil, err
	}

	token, err := utils.GenerateToken(user.ID, user.Email, s.jwtSecret)
	if err != nil {
		return nil, err
	}

	user.PasswordHash = ""

	return &domain.AuthResponse{
		Token: token,
		User:  *user,
	}, nil
}

func (s *AuthService) Login(request domain.LoginRequest) (*domain.AuthResponse, error) {
	user, err := s.userRepository.FindByEmail(strings.ToLower(request.Email))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("invalid email or password")
		}
		return nil, err
	}

	if !utils.CheckPassword(request.Password, user.PasswordHash) {
		return nil, errors.New("invalid email or password")
	}

	token, err := utils.GenerateToken(user.ID, user.Email, s.jwtSecret)
	if err != nil {
		return nil, err
	}

	user.PasswordHash = ""

	return &domain.AuthResponse{
		Token: token,
		User:  *user,
	}, nil
}

