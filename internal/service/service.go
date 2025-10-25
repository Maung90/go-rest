package service

import "go-rest/internal/repository"

// Service adalah implementasi generic untuk logika bisnis CRUD.
// Ia membutuhkan sebuah repository yang sesuai dengan tipe T.
type Service[T any] struct {
	repo repository.Repository[T]
}

// NewService membuat instance Service generic baru.
func NewService[T any](repo repository.Repository[T]) *Service[T] {
	return &Service[T]{repo: repo}
}

func (s *Service[T]) GetAll() ([]T, error) {
	return s.repo.FindAll()
}

func (s *Service[T]) GetByID(id int) (T, error) {
	return s.repo.FindByID(id)
}

// Catatan: Untuk Create dan Update, kita akan sederhanakan
// dengan menerima model utuh. Penanganan DTO (Data Transfer Object)
// bisa ditambahkan jika perlu, tapi ini intinya.
func (s *Service[T]) Create(model T) (T, error) {
	// Di sini bisa ditambahkan logika bisnis umum jika ada
	return s.repo.Save(model)
}

func (s *Service[T]) Update(model T) (T, error) {
	// Di sini bisa ditambahkan logika bisnis umum jika ada
	return s.repo.Update(model)
}

func (s *Service[T]) Delete(id int) error {
	// Lakukan pengecekan apakah data ada sebelum dihapus
	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}