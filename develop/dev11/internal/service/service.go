package service

import (
	"time"

	"github.com/mirhijinam/wb-l2/develop/dev11/internal/models"
)

type calendarRepo interface {
	CreateEvent(models.Event) models.Event
	UpdateEvent(int, string, string, time.Time) (models.Event, error)
	DeleteEvent(int) (models.Event, error)
	DailyEvents() []models.Event
	WeeklyEvents() []models.Event
	MonthlyEvents() []models.Event
}

type Service struct {
	calendarRepo calendarRepo
}

func New(repo calendarRepo) *Service {
	return &Service{
		calendarRepo: repo,
	}
}

func (s *Service) CreateEvent(e models.Event) (models.Event, error) {
	err := ValidateEvent(e)
	if err != nil {
		return models.Event{}, err
	}

	createdEvent := s.calendarRepo.CreateEvent(e)
	return createdEvent, nil
}

func (s *Service) UpdateEvent(id int, e models.Event) error {
	err := ValidateEvent(e)
	if err != nil {
		return err
	}

	_, err = s.calendarRepo.UpdateEvent(id, e.Name, e.Data, e.Date)
	return err
}

func (s *Service) DeleteEvent(id int) error {
	_, err := s.calendarRepo.DeleteEvent(id)
	return err
}

func (s *Service) GetDailyEvents() ([]models.Event, error) {
	events := s.calendarRepo.DailyEvents()
	return events, nil
}

func (s *Service) GetWeeklyEvents() ([]models.Event, error) {
	events := s.calendarRepo.WeeklyEvents()
	return events, nil
}

func (s *Service) GetMonthlyEvents() ([]models.Event, error) {
	events := s.calendarRepo.MonthlyEvents()
	return events, nil
}
