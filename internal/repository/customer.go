package repository

import (
	"go-rest/domain"
	"context"
)

type customerRepository struct {

}

func NewCustomer() []domain.CustomerRepository{
	return &customerRepository{}
}

func (c *customerRepository) FindAll(ctx context.Context)([]domain.Customer, error){
	return nil, nil
}
func (c *customerRepository) FindById(ctx context.Context, id int)(domain.Customer, error){
	return nil, nil
}
func (c *customerRepository) Create(ctx context.Context, customer *domain.Customer) error{
	return nil
}
func (c *customerRepository) Update(ctx context.Context, customer *domain.Customer) error{
	return nil
}
func (c *customerRepository) Delete(ctx context.Context, id int) error{
	return nil
}
