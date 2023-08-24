package models

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func (m *DBModel) InsertUser(name, email, password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `INSERT INTO users (name, email, password, created_at, updated_at)
	VALUES($1, $2, $3, $4, $5)`

	_, err := m.DB.ExecContext(ctx, stmt, name, email, password, time.Now(), time.Now())
	if err != nil {
		fmt.Println(err)
		return errors.New("failed to save the credentials")
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
