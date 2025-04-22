package ports

import "example.com/api/internal/core/domain"

type EventRepository interface {
	Save(event *domain.Event) (*domain.Event, error)
	GetAll() ([]domain.Event, error)
	// Puedes agregar más métodos según tus necesidades (GetByID, Update, Delete, etc.)
}
