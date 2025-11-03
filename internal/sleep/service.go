package sleep


type Service interface {
	GetByID(userID, id int) (Sleep, error)
	Create(sleep Sleep) (Sleep, error)
	Update(sleep Sleep) (Sleep, error)
	Delete(id int) error
	GetByUserID(userID int) ([]Sleep, error)
}

type service struct {
	repository Repository 
}

func NewService(repository Repository) Service {
	return &service{ repository: repository}
}

func (s *service) GetByID(userID, id int) (Sleep, error) {
	return s.repository.FindByID(userID, id)
}

func (s *service) Create(sleep Sleep) (Sleep, error) {
	return s.repository.Save(sleep)
}

func (s *service) Update(sleep Sleep) (Sleep, error) {
	return s.repository.Update(sleep)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

func (s *service) GetByUserID(userID int) ([]Sleep, error) {
	return s.repository.FindAll(userID)
}