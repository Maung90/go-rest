package sleep

import "go-rest/internal/service"

type SleepService interface {
	GetByID(id int) (Sleep, error)
	Create(sleep Sleep) (Sleep, error)
	Update(sleep Sleep) (Sleep, error)
	Delete(id int) error

	GetSleepsByUserId(userID int) ([]Sleep, error)
}

type sleepService struct {
	genericService *service.Service[Sleep]
	repository Repository
}

func NewSleepService(repo Repository) SleepService {
	return &sleepService{
		genericService: service.NewService[Sleep](repo),
		repository: repo,
	}
}

func (s *sleepService) GetByID(id int) (Sleep, error) {
	return s.genericService.GetByID(id)
}

func (s *sleepService) Create(sleep Sleep) (Sleep, error) {
	return s.genericService.Create(sleep)
}

func (s *sleepService) Update(sleep Sleep) (Sleep, error) {
	return s.genericService.Update(sleep)
}

func (s *sleepService) Delete(id int) error {
	return s.genericService.Delete(id)
}

func (s *sleepService) GetSleepsByUserId(userID int) ([]Sleep, error) {
	return s.repository.FindAll(userID)
}