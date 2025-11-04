package dailyStory

import (
	"time"
	"errors"
)

type Service interface {
	FindAll(userID int) ([]DailyStory, error)
	FindByID(userID, id int) (DailyStory, error)
	FindByDate(userID int, date string) ([]DailyStory, error)
	Save(userID int, input StoryInput) (DailyStory, error)
	Update(id int, userID int, input StoryInput) (DailyStory, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) FindAll(userID int) ([]DailyStory, error) {
	return s.repository.FindAll(userID)
}

// FindByDate implements Service.
func (s *service) FindByDate(userID int, date string) ([]DailyStory, error) {
	return s.repository.FindByDate(userID, date)
}

// FindByID implements Service.
func (s *service) FindByID(userID int, id int) (DailyStory, error) {
	return s.repository.FindByID(userID, id)
}

// Save implements Service.
func (s *service) Save(userID int, input StoryInput) (DailyStory, error) {
	parsedStoryDate, err := time.Parse("2006-01-02", input.StoryDate)
	if err != nil {
		return DailyStory{}, errors.New("invalid date format, please use YYYY-MM-DD")
	}

	dailyStory := DailyStory{
		User_id:   userID, 
		StoryDate: parsedStoryDate,
		StoryText: input.StoryText,
		Mood:      input.Mood,
	}
	return s.repository.Save(dailyStory)
}

// Implementasi fungsi Update
func (s *service) Update(id int, userID int, input StoryInput) (DailyStory, error) {

	parsedStoryDate, err := time.Parse("2006-01-02", input.StoryDate)
	if err != nil {
		return DailyStory{}, errors.New("invalid date format, please use YYYY-MM-DD")
	}
	
	dailyStory := DailyStory{
		ID:        id,     // ID dari URL
		User_id:   userID, // UserID dari middleware
		StoryDate: parsedStoryDate,
		StoryText: input.StoryText,
		Mood:      input.Mood,
	}

	return s.repository.Update(dailyStory)
}
func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}