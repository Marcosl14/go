package services

import (
	"example.com/api/internal/core/domain"
	"example.com/api/internal/core/ports"
)

type eventService struct {
	eventRepo ports.EventRepository
}

func NewEventService(repo ports.EventRepository) *eventService {
	return &eventService{
		eventRepo: repo,
	}
}

func (s *eventService) CreateEvent(event *domain.Event) (*domain.Event, error) {
	return s.eventRepo.Save(event)
}

func (s *eventService) GetEvents() ([]domain.Event, error) {
	return s.eventRepo.GetAll()
}

func (s *eventService) GetEvent(id int64) (domain.Event, error) {
	return s.eventRepo.GetById(id)
}
