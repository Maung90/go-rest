package user

import "golang.org/x/crypto/bcrypt"

type Service interface {
	GetAllUsers() ([]User, error)
	GetUserByID(id int) (User, error)
	CreateUser(input CreateUserInput) (User, error)
	UpdateUser(id int, input UpdateUserInput) (User, error)
	DeleteUser(id int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) GetAllUsers() ([]User, error) {
	return s.repository.FindAll()
}

func (s *service) GetUserByID(id int) (User, error) {
	return s.repository.FindByID(id)
}

func (s *service) CreateUser(input CreateUserInput) (User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}

	user := User{
		Username:     input.Username,
		Email:    input.Email,
		Password: string(passwordHash),
	}
	return s.repository.Save(user)
}

func (s *service) UpdateUser(id int, input UpdateUserInput) (User, error) {
	user, err := s.repository.FindByID(id)
	if err != nil {
		return user, err
	}

	user.Username = input.Username
	user.Email = input.Email

	return s.repository.Update(user)
}

func (s *service) DeleteUser(id int) error {
	_, err := s.repository.FindByID(id)
	if err != nil {
		return err
	}
	return s.repository.Delete(id)
}