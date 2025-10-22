package repository

import (
	"context"
	"database/sql"
	"go-rest/domain"
	"log"
)

type customerRepository struct {
	db *sql.DB // Tambahkan properti untuk menampung koneksi DB
}

// Ubah constructor untuk menerima koneksi DB dan mengembalikan interface yang benar
func NewCustomer(db *sql.DB) domain.CustomerRepository {
	return &customerRepository{
		db: db,
	}
}

// Implementasikan fungsi FindAll untuk mengambil data
func (c *customerRepository) FindAll(ctx context.Context) ([]domain.Customer, error) {
	query := "SELECT id, code, name, created_at, updated_at FROM customers WHERE deleted = 0"
	rows, err := c.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	customers := make([]domain.Customer, 0)
	for rows.Next() {
		var customer domain.Customer
		err := rows.Scan(
			&customer.ID,
			&customer.Code,
			&customer.Name,
			&customer.Created_at,
			&customer.Updated_at,
		)
		if err != nil {
			log.Printf("Error scanning customer row: %v", err)
			continue
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

// Fungsi lain (biarkan dulu, fokus pada FindAll)
func (c *customerRepository) FindById(ctx context.Context, id int) (domain.Customer, error) {
	// Implementasi nanti
	return domain.Customer{}, nil
}
func (c *customerRepository) Create(ctx context.Context, customer *domain.Customer) error {
	return nil
}
func (c *customerRepository) Update(ctx context.Context, customer *domain.Customer) error {
	return nil
}

// Perbaiki tipe parameter id di interface dan implementasi
func (c *customerRepository) Delete(ctx context.Context, id int) error {
	return nil
}
