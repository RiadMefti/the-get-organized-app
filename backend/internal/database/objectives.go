package database

import (
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

func (s *service) AbandonObjective(objectiveId string) error {

	_, err := s.db.Exec("UPDATE objectives SET abandonned= $1 WHERE id = $2 ", false, objectiveId)
	if err != nil {
		return err
	}

	return nil

}

func (s *service) GetObjective(userId uuid.UUID) {

}

func (s *service) CreateGoal() {

}

func (s *service) UpdateGoal() {

}

func (s *service) AbandonGoal() {

}
