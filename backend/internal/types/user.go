package types

import (
	"time"

	"github.com/google/uuid"
)

type UserRegistration struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	CopyPassword string `json:"copy_password"`
}

type UserDB struct {
	ID             uuid.UUID `json:"id"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"hashed_password"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JwtToken struct {
	Jwt string `json:"jwt"`
}

type Objective struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Type      string    `json:"type"`
	StartDate time.Time `json:"start_date"`
	Goals     []Goal    `json:"goals"`
}

type Goal struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
}
