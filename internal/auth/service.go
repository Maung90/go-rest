package auth

import (
	"errors"
	"go-rest/internal/user"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(input RegisterInput) (user.User, error)
	FindByEmail(email string) (user.User, error)
	Login(input LoginInput) (user.User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) Register(input RegisterInput) (user.User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return user.User{}, err
	}
	newUser := user.User{
		Name: input.Name,
		Email:    input.Email,
		Password: string(passwordHash),
	}
	return s.repository.Register(newUser)
}

func (s *service) FindByEmail(email string) (user.User, error) {
	return s.repository.FindByEmail(email)
}

func (s *service) Login(input LoginInput) (user.User, error) {
	
	foundUser, err := s.repository.FindByEmail(input.Email)
	if err != nil {
		return user.User{}, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(input.Password))
	if err != nil {
		return user.User{}, errors.New("invalid email or password")
	}
	return foundUser, nil
}