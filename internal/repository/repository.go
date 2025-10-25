package repository

// Repository adalah interface generic untuk operasi CRUD.
// [T any] berarti T bisa berupa tipe data apa pun (User, Activity, dll).
type Repository[T any] interface {
	FindAll() ([]T, error)
	FindByID(id int) (T, error)
	Save(model T) (T, error)
	Update(model T) (T, error)
	Delete(id int) error
}