package models

import (
	"github.com/MacBie2k/event-booking-api/db"
)

type Registration struct {
	Id      int64
	UserId  int64 `binding:"required"`
	EventId int64 `binding:"required"`
}

func (r *Registration) Save() error {
	query := "INSERT INTO registrations (user_id, event_id) VALUES(?, ?)"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(r.UserId, r.EventId)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return err
	}

	r.Id = id
	return err
}

func GetRegistrationByUserAndEventId(userId, eventId int64) (*Registration, error) {
	query := "SELECT * FROM registrations WHERE user_id = ? AND event_id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(userId, eventId)
	var registration Registration

	err = row.Scan(&registration.Id, &registration.UserId, &registration.EventId)

	if err != nil {
		return nil, nil
	}

	return &registration, nil
}

func (r *Registration) Delete() error {
	query := "DELETE FROM registrations WHERE id = ?"
	_, err := db.DB.Exec(query, r.Id)
	if err != nil {
		return err
	}
	return nil
}
