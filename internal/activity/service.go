package activity

// import (
	// "errors"
	// "go-rest/pkg/parser"
// )

type Service interface {
	FindAll(userID int) ([]Activity, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) FindAll(userID int) ([]Activity, error) {
	return s.repository.FindAll(userID)
}