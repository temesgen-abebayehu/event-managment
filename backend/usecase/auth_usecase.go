package usecase

import (
	"errors"
	"event-management-backend/domain"
	"event-management-backend/repository"
	"event-management-backend/utils"
)

type AuthUsecase struct {
	userRepo  *repository.UserRepository
	jwtSecret string
}

func NewAuthUsecase(userRepo *repository.UserRepository, jwtSecret string) *AuthUsecase {
	return &AuthUsecase{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

func (u *AuthUsecase) Signup(req domain.SignupRequest) (*domain.AuthResponse, error) {
	// Business logic: Validate
	if len(req.Password) < 8 {
		return nil, errors.New("password must be at least 8 characters")
	}

	// Business logic: Hash password
	passwordHash, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// Create user
	user, err := u.userRepo.Create(req.Email, passwordHash, req.FullName)
	if err != nil {
		return nil, errors.New("email already exists")
	}

	// Business logic: Generate JWT
	token, err := utils.GenerateJWT(user.ID, user.Email, u.jwtSecret)
	if err != nil {
		return nil, err
	}

	return &domain.AuthResponse{
		AccessToken: token,
		User:        *user,
	}, nil
}

func (u *AuthUsecase) Login(req domain.LoginRequest) (*domain.AuthResponse, error) {
	// Get user with password hash
	user, passwordHash, err := u.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Business logic: Verify password
	if !utils.CheckPassword(req.Password, passwordHash) {
		return nil, errors.New("invalid email or password")
	}

	// Business logic: Generate JWT
	token, err := utils.GenerateJWT(user.ID, user.Email, u.jwtSecret)
	if err != nil {
		return nil, err
	}

	return &domain.AuthResponse{
		AccessToken: token,
		User:        *user,
	}, nil
}
