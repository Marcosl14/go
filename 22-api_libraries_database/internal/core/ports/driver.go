package ports

import "example.com/api/internal/core/domain"

type EventService interface {
	CreateEvent(event *domain.Event) (*domain.Event, error)
	GetEvents() ([]domain.Event, error)
	GetEvent(id int64) (domain.Event, error)
}
