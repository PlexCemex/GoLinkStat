package auth

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"projects/GoLinkStat/internal/user"
	"projects/GoLinkStat/pkg/di"
)

type AuthService struct {
	UserRepository di.IUserRepository
}

func NewAuthService(userRepository di.IUserRepository) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
	}
}

func (service *AuthService) Register(email, password, name string) (string, error) {
	existedUser, _ := service.UserRepository.GetByEmail(email)
	if existedUser != nil {
		return "", errors.New(ErrUserExists)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user := user.User{
		Email:    email,
		Password: string(hashedPassword),
		Name:     name,
	}
	_, err = service.UserRepository.Create(&user)
	if err != nil {
		return "", err
	}
	return user.Email, nil
}

func (service *AuthService) Login(email, password string) (string, error) {
	existedUser, err := service.UserRepository.GetByEmail(email)
	if err != nil || existedUser == nil {
		return "", errors.New(ErrWrongCredentials)
	}
	err = bcrypt.CompareHashAndPassword([]byte(existedUser.Password), []byte(password))
	if err != nil {
		return "", errors.New(ErrWrongCredentials)
	}
	return existedUser.Email, nil
}
