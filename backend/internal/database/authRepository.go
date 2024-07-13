package database

import (
	"backend/internal/types"
	"database/sql"

	"github.com/google/uuid"
)

func (s *service) GetUserByEmail(email string) (types.UserDB, error) {
	var user types.UserDB

	err := s.db.QueryRow("SELECT id, email, hashed_password, created_at, updated_at FROM users WHERE email = $1", email).Scan(&user.ID, &user.Email, &user.HashedPassword, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {

		if err == sql.ErrNoRows {

			return types.UserDB{}, nil
		}

		return types.UserDB{}, err
	}

	return user, nil
}
func (s *service) CreateUser(email string, hashedPassword string) (string, error) {
	newUUID, err := uuid.NewRandom() // Generate a new random UUID.
	if err != nil {
		return "", err
	}

	_, err = s.db.Exec("INSERT INTO users (id, email, hashed_password) VALUES ($1, $2, $3)", newUUID, email, hashedPassword)
	if err != nil {
		return "", err
	}
	return newUUID.String(), nil
}
