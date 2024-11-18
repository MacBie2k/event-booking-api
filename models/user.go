package models

import (
	"errors"

	"github.com/MacBie2k/event-booking-api/db"
	"github.com/MacBie2k/event-booking-api/utils"
)

type User struct {
	Id       int64
	Email    string `binding:"required" validate:"email"`
	Password string `binding:"required" validate:"password"`
}

func (u *User) Save() error {
	query := "INSERT INTO users(email, password) VALUES(?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	u.Id, err = result.LastInsertId()
	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)
	var retrievedPassword string
	err := row.Scan(&u.Id, &retrievedPassword)
	if err != nil {
		return errors.New("User not verified")
	}

	if !utils.CheckPasswordHash(u.Password, retrievedPassword) {
		return errors.New("User not verified")
	}

	return nil
}
