package activity

type Service interface {
	GetAll() ([]Activity, error)
	GetByID(id int) (Activity, error)
	Create(input CreateActivityInput) (Activity, error)
	Update(id int, input UpdateActivityInput) (Activity, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) GetAll() ([]Activity, error) {
	return s.repository.FindAll()
}

func (s *service) GetByID(id int) (Activity, error) {
	return s.repository.FindByID(id)
}

func (s *service) Create(input CreateActivityInput) (Activity, error) {
	activity := Activity{
		User_id:  input.User_id,
		Activity: input.Activity,
	}
	return s.repository.Save(activity)
}

func (s *service) Update(id int, input UpdateActivityInput) (Activity, error) {
	Activity, err := s.repository.FindByID(id)
	if err != nil {
		return Activity, err
	}

	Activity.User_id = input.User_id
	Activity.Activity = input.Activity

	return s.repository.Update(Activity)
}

func (s *service) Delete(id int) error {
	_, err := s.repository.FindByID(id)
	if err != nil {
		return err
	}
	return s.repository.Delete(id)
}