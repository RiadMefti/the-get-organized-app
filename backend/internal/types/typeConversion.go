package types

func GoalDbListToGoalList(
	goalDbList []GoalDB,

) []Goal {

	goalList := []Goal{}

	for _, goalDb := range goalDbList {
		goal := Goal{
			ID:          goalDb.ID,
			Name:        goalDb.Name,
			Description: goalDb.Description,
			Done:        goalDb.Done,
			Abandoned:   goalDb.Abandoned,
		}

		goalList = append(goalList, goal)
	}

	return goalList
}

func ObjectiveDbToObjective(
	objectiveDb ObjectiveDB,
	GoalDbList []GoalDB,

) Objective {

	objective := Objective{
		ID:        objectiveDb.ID,
		UserID:    objectiveDb.UserID,
		Type:      objectiveDb.Type,
		StartDate: objectiveDb.StartDate,
		Abandoned: objectiveDb.Abandoned,
		Goals:     GoalDbListToGoalList(GoalDbList),
	}

	return objective

}
