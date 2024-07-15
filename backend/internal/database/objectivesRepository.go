package database

import (
	"backend/internal/types"
	"time"

	"github.com/google/uuid"
)

func (s *service) CreateObjective(start_date time.Time, objType string, userID uuid.UUID) error {

	newUUID, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	_, err = s.db.Exec("INSERT INTO objectives (id,start_date,type,user_id) VALUES ($1, $2, $3, $4)", newUUID, start_date, objType, userID)

	if err != nil {
		return err
	}

	return nil

}

func (s *service) AbandonObjective(objectiveId uuid.UUID) error {

	_, err := s.db.Exec("UPDATE objectives SET abandoned= $1 WHERE id = $2 ", true, objectiveId)
	if err != nil {
		return err
	}

	return nil

}

func (s *service) GetObjective(userId uuid.UUID) ([]types.Objective, error) {

	rows, err := s.db.Query("SELECT * FROM objectives WHERE user_id = $1", userId)
	if err != nil {
		return nil, err
	}

	objectivesDb := []types.ObjectiveDB{}

	for rows.Next() {
		var obj types.ObjectiveDB
		err = rows.Scan(&obj.ID, &obj.UserID, &obj.StartDate, &obj.Type, &obj.Abandoned, &obj.CreatedAt, &obj.UpdatedAt)
		if err != nil {
			return nil, err
		}
		objectivesDb = append(objectivesDb, obj)
	}

	objectives := []types.Objective{}
	for _, obj := range objectivesDb {
		goals, err := s.getGoalsByObjective(obj.ID)
		if err != nil {
			return nil, err
		}

		objective := types.ObjectiveDbToObjective(obj, goals)

		objectives = append(objectives, objective)

	}

	return objectives, nil

}

func (s *service) CreateGoal(name string, description string, objectiveID uuid.UUID) error {
	// Start a transaction
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	// Use defer to ensure the transaction is closed properly in case of an error
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			tx.Rollback() // err is non-nil; don't change it
		} else {
			err = tx.Commit() // if Commit returns error update err with commit err
		}
	}()

	newUUID, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	// Execute the first INSERT operation within the transaction
	_, err = tx.Exec("INSERT INTO goals (id, name, description) VALUES ($1, $2, $3)", newUUID, name, description)
	if err != nil {
		return err
	}

	// Execute the second INSERT operation within the transaction
	_, err = tx.Exec("INSERT INTO objectives_goals (objective_id, goal_id) VALUES ($1, $2)", objectiveID, newUUID)
	if err != nil {
		return err
	}

	// If no errors occurred, the transaction will be committed at the end of the function
	return nil
}

func (s *service) UpdateGoal(id uuid.UUID, name string, description string, done bool) error {
	_, err := s.db.Exec("UPDATE goals SET name=$1, description=$2, done=$3 WHERE id = $4", name, description, done, id)
	if err != nil {
		return err
	}

	return nil

}

func (s *service) AbandonGoal(id uuid.UUID) error{

	_, err := s.db.Exec("UPDATE goals SET abandoned= $1 WHERE id = $2 ", true, id)
	if err != nil {
		return err
	}
	return nil

}

func (s *service) getGoalsByObjective(objectiveID uuid.UUID) ([]types.GoalDB, error) {
	query := `
        SELECT g.id, g.name, g.description, g.done, g.abandoned, g.created_at, g.updated_at
        FROM goals g
        JOIN objectives_goals og ON g.id = og.goal_id
        WHERE og.objective_id = $1
    `

	rows, err := s.db.Query(query, objectiveID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	goals := []types.GoalDB{}

	for rows.Next() {
		var goal types.GoalDB
		err := rows.Scan(&goal.ID, &goal.Name, &goal.Description, &goal.Done, &goal.Abandoned, &goal.CreatedAt, &goal.UpdatedAt)
		if err != nil {
			return nil, err
		}
		goals = append(goals, goal)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return goals, nil
}
