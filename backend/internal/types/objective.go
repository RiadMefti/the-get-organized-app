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
	Abandoned bool      `json:"abandoned"`
	Goals      []Goal    `json:"goals"`
}

type Goal struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	Abandoned  bool      `json:"abandoned"`
}

type ObjectiveDB struct {
	ID         uuid.UUID `json:"id"`
	UserID     uuid.UUID `json:"user_id"`
	StartDate  time.Time `json:"start_date"`
	Type       string    `json:"type"`
	Abandoned bool      `json:"abandoned"`
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

// CreateObjectiveRequest represents the request structure for creating an objective.
type CreateObjectiveRequest struct {
	ObjType string    `json:"objType"` 
	UserID  uuid.UUID `json:"userID"`
    StartDate time.Time `json:"start_date"`
}

// AbandonObjectiveRequest represents the request structure for abandoning an objective.
type AbandonObjectiveRequest struct {
	ObjectiveID uuid.UUID `json:"objectiveID"`
}

// GetObjectivesRequest represents the request structure for getting objectives by user ID.
type GetObjectivesRequest struct {
	UserID uuid.UUID `json:"userID"`
}

// CreateGoalRequest represents the request structure for creating a goal.
type CreateGoalRequest struct {
	Name         string    `json:"name"`        
	Description  string    `json:"description"`  
	ObjectiveID  uuid.UUID `json:"objectiveID"`  
}

// UpdateGoalRequest represents the request structure for updating a goal.
type UpdateGoalRequest struct {
	ID          uuid.UUID `json:"id"`          
	Name        string    `json:"name"`        
	Description string    `json:"description"` 
	Done        bool      `json:"done"`        
}

// AbandonGoalRequest represents the request structure for abandoning a goal.
type AbandonGoalRequest struct {
	ID uuid.UUID `json:"id"` 
}