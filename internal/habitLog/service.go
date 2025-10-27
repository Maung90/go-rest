package habitLog
import "time"

type HabitLogService interface {
	CreateLogs(input CreateHabitLogInput) (HabitLog, error)
	FindHabitLogs(input GetHabitLogInput) ([]HabitLog, error)
}

type habitLogService struct {
	repository Repository
}

func NewHabitLogService(repository Repository) HabitLogService {
	return &habitLogService{repository: repository}
}

func (s *habitLogService) CreateLogs(input CreateHabitLogInput) (HabitLog, error) {
	parsedDate, err := time.Parse("2006-01-02", input.LogDate)
	if err != nil {
		return HabitLog{}, err
	}
	newLogs := HabitLog{
		Habit_id: input.Habit_id,
		User_id:    input.User_id,
		Log_date: parsedDate,
	}
	return s.repository.CreateLogs(newLogs)
}


func (s *habitLogService) FindHabitLogs(input GetHabitLogInput) ([]HabitLog, error) {
	parsedDate, err := time.Parse("2006-01-02", input.LogDate)
	if err != nil {
		return []HabitLog{}, err
	}
		return s.repository.FindHabitLogs(parsedDate)
}

