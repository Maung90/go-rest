package domain

import (
	"context"
	"database/sql"
)

type Customer struct{
	ID 			int  		 `db:"id"`
	Code 		string  	 `db:"code"`
	Name		string  	 `db:"name"`
	Created_at sql.NullTime  `db:"created_at"`
	Updated_at sql.NullTime  `db:"updated_at"`
	Deleted 	bool		 `db:"deleted"`
}
type CustomerRepository interface {
	FindAll(ctx context.Context) ([]Customer, error)
	FindById(ctx context.Context, id int) (Customer, error)
	Create(ctx context.Context, c *Customer) error
	Update(ctx context.Context, c *Customer) error
	Delete(ctx context.Context, id int) error
}


type CustomerService interface{
	Index(ctx context.Context)([] dto.CustomerData, error)
}
