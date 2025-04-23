package ports

import "example.com/api/internal/core/domain"

type EventRepository interface {
	Save(event *domain.Event) (*domain.Event, error)
	GetAll() ([]domain.Event, error)
	GetById(in int64) (domain.Event, error)
	Update(event *domain.Event) (domain.Event, error)
}
