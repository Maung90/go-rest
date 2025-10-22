package service

import (
	"context"
	"go-rest/domain" // Ganti dengan path domain Anda
)

// customerService provides business logic for entity
type customerService interface {
	GetAll(ctx context.Context) ([]domain.CustomerService, error)
	GetByID(ctx context.Context, id int) (domain.CustomerService, error)
	Store(ctx context.Context, entity *domain.CustomerService) error
	Update(ctx context.Context, entity *domain.CustomerService) error
	Delete(ctx context.Context, id int) error
}

type customerService struct {
	customerRepository domain.CustomerRepository
}

// NewcustomerService creates a new service
func Newcustomer(customerRepository domain.CostomerRepository) domain.CustomerService {
	return &customerService{
		customerRepository: customerRepository,
	}
}

func (s *customerService) GetAll(ctx context.Context) ([]domain.CustomerService, error) {
	customers, err := s.customerRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	var customerData []dto.CustomerData
	for _, v := range customers{
		customerData = append(customerData, dto.CustomerData{
			ID :v.ID,
			Code :v.code,
			Name :v.name,
		})
	}
	return customerData, nil
}

func (s *customerService) GetByID(ctx context.Context, id int) (domain.CustomerService, error) {
	return s.customerRepository.FindByID(ctx, id)
}

func (s *customerService) Store(ctx context.Context, entity *domain.CustomerService) error {
	return s.customerRepository.Create(ctx, entity)
}

func (s *customerService) Update(ctx context.Context, entity *domain.CustomerService) error {
	return s.repo.Update(ctx, entity)
}

func (s *customerService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
