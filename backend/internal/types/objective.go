package types

import (
	"time"

	"github.com/google/uuid"
)

type Objective struct {
	ID         uuid.UUID `json:"id"`
	UserID     uuid.UUID `json:"user_id"`
	Type       string    `json:"type"`
	StartDate  time.Time `json:"start_date"`
	Abandonned bool      `json:"abandonned"`
	Goals      []Goal    `json:"goals"`
}

type Goal struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	Abandonned  bool      `json:"abandonned"`
}


type ObjectiveDB struct {
    ID         uuid.UUID `json:"id"`
    UserID     uuid.UUID `json:"user_id"`
    StartDate  time.Time `json:"start_date"`
    Type       string    `json:"type"`
    Abandoned  bool      `json:"abandoned"`
    CreatedAt  time.Time `json:"created_at"`
    UpdatedAt  time.Time `json:"updated_at"`
}

type GoalDB struct {
    ID          uuid.UUID `json:"id"`
    Name        string    `json:"name"`
    Description string    `json:"description"`
    Done        bool      `json:"done"`
    Abandoned   bool      `json:"abandoned"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}