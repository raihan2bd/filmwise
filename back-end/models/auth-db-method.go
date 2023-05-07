package models

import (
	"errors"
	"time"
)

func (m *DBModel) InsertUser(username, email, password string) error {
	stmt := `INSERT INTO users (username, email, password, created_at, updated_at)
	VALUES($1, $2, $3, $4, $5)`

	_, err := m.DB.Exec(stmt, username, email, password, time.Now(), time.Now())
	if err != nil {
		return errors.New("failed to save the credientials")
	}

	return nil
}

// GetUserByEmail gets user by email
func (m *DBModel) GetUserByEmail(email string) (*User, error) {
	stmt := `SELECT id, name, email, password, user_type FROM users
	WHERE email = $1`

	row := m.DB.QueryRow(stmt, email)

	u := &User{}

	err := row.Scan(&u.ID, &u.FullName, &u.Email, &u.Password, &u.UserType)
	if err != nil {
		return nil, err
	}

	return u, nil
}
