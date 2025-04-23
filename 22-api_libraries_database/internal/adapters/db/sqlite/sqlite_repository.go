package sqlite

import (
	"database/sql"
	"fmt"

	"example.com/api/internal/core/domain"
)

type sqliteEventRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *sqliteEventRepository {
	return &sqliteEventRepository{db: db}
}

func (r *sqliteEventRepository) Save(event *domain.Event) (*domain.Event, error) {
	query := `INSERT INTO events (name, description, location, date_time, user_id) VALUES (?, ?, ?, ?, ?)`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserID)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	event.ID = id
	return event, nil
}

func (r *sqliteEventRepository) GetAll() ([]domain.Event, error) {
	query := `SELECT id, name, description, location, date_time, user_id FROM events`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []domain.Event
	for rows.Next() {
		var event domain.Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return events, nil
}

func (r *sqliteEventRepository) GetById(id int64) (domain.Event, error) {
	query := `SELECT id, name, description, location, date_time, user_id FROM events WHERE id = ?`
	row := r.db.QueryRow(query, id)

	var event domain.Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		if err == sql.ErrNoRows {
			return event, nil
		}

		return event, err
	}

	return event, nil
}

func (r *sqliteEventRepository) Update(event *domain.Event) (domain.Event, error) {
	fmt.Println("Updating event:", event)

	query := `UPDATE events SET name = ?, description = ?, location = ?, date_time = ?, user_id = ? WHERE id = ?`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return domain.Event{}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserID, event.ID)
	if err != nil {
		return domain.Event{}, err
	}

	return *event, nil
}
