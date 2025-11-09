package dashboard

import (
	"go-rest/internal/activity"
	"go-rest/internal/dailyStory"
	"go-rest/internal/habitLog"
	"go-rest/internal/sleep"
)

// Service interface
type Service interface {
	GetDashboard(userID int, date string) (*DashboardData, error)
}

// Implementation struct
type service struct {
	activityRepo activity.Repository
	storyRepo    dailyStory.Repository
	sleepRepo    sleep.Repository
	habitRepo    habitLog.Repository
}

// Constructor
func NewService(
	h habitLog.Repository,
	a activity.Repository,
	s dailyStory.Repository,
	sl sleep.Repository,
) Service {
	return &service{
		habitRepo:    h,
		activityRepo: a,
		storyRepo:    s,
		sleepRepo:    sl,
	}
}

// Response DTO
type DashboardData struct {
	HabitsDoneCount int                    `json:"habits_done"`
	HabitsDone      []habitLog.HabitLog    `json:"habits"`
	Activities      []activity.Activity    `json:"activities"`
	Story           []dailyStory.DailyStory `json:"story,omitempty"`
	Sleep           []sleep.Sleep           `json:"sleep,omitempty"`
}

// Service Method
func (svc *service) GetDashboard(userID int, date string) (*DashboardData, error) {
	habitsDone, err := svc.habitRepo.FindDoneByDate(userID, date)
	if err != nil {
		return nil, err
	}
	habitCount := len(habitsDone)

	activities, err := svc.activityRepo.FindByDate(date, userID)
	if err != nil {
		return nil, err
	}

	story, err := svc.storyRepo.FindByDate(userID, date)
	if err != nil {
		return nil, err
	}

	sleep, err := svc.sleepRepo.FindByDate(userID, date)
	if err != nil {
		return nil, err
	}

	return &DashboardData{
		HabitsDoneCount: habitCount,
		HabitsDone:      habitsDone,
		Activities:      activities,
		Story:           story,
		Sleep:           sleep,
	}, nil
}