package sleep


type Service interface {
	GetByID(userID, id int) (Sleep, error)
	Create(sleep Sleep) (Sleep, error)
	Update(sleep Sleep) (Sleep, error)
	Delete(id int) error
	GetByUserID(userID int) ([]Sleep, error)
	GetWeeklyStats(userID int) ([]SleepStat, error)
	GetMonthlyStats(userID int) ([]SleepStat, error)
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

func (s *service) GetWeeklyStats(userID int) ([]SleepStat, error) {
	return s.repository.GetWeeklyStats(userID)
}

func (s *service) GetMonthlyStats(userID int) ([]SleepStat, error) {
	return s.repository.GetMonthlyStats(userID)
}
