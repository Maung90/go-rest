package activity

import (
	// "time"
	// "errors"
)

type Service interface {
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}