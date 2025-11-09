package activity

import (
	"errors"
	"go-rest/pkg/parser"
)

type Service interface {
	FindAll(userID int) ([]Activity, error)
	FindById(userID, id int) (Activity, error)
	Save(userID int, input ActivityInput) (Activity, error)
	Update(id, userID int, input ActivityInput) (Activity, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) FindAll(userID int) ([]Activity, error) {
	return s.repository.FindAll(userID)
}

func (s *service) FindById(userID, id int) (Activity, error) {
	return s.repository.FindById(id, userID)
}

func (s *service) Save(userID int, input ActivityInput) (Activity, error) {
	parseDate, err := parser.ParseDateString(input.ActivityDate)
	if err != nil {
		return Activity{}, errors.New("Format tanggal salah, gunakan format 	YYYY-MM-DD")
	}
	activity := Activity{
		User_id:         userID,
		ActivityDate:    parseDate,
		Title:           input.Title,
		DurationMinutes: input.DurationMinutes,
		Notes:           input.Notes,
	}

	return s.repository.Save(activity)
}

func (s *service) Update(id, userID int, input ActivityInput) (Activity, error) {
	parseDate, err := parser.ParseDateString(input.ActivityDate)
	if err != nil {
		return Activity{}, errors.New("Format tanggal salah, gunakan format 	YYYY-MM-DD")
	}
	activity := Activity{
		ID:              id,
		User_id:         userID,
		ActivityDate:    parseDate,
		Title:           input.Title,
		DurationMinutes: input.DurationMinutes,
		Notes:           input.Notes,
	}

	return s.repository.Save(activity)
}

func (s *service) Delete(id int) error {
    return s.repository.Delete(id)
}