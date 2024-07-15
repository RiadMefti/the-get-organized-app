package server

import (
	"backend/internal/types"
	"backend/internal/utils"
	"net/http"
)

func (s *Server) createObjectiveHandler(w http.ResponseWriter, r *http.Request) {
	var createObjectiveRequest types.CreateObjectiveRequest
	err := utils.ParseJSON(r, &createObjectiveRequest)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = s.db.CreateObjective(createObjectiveRequest.StartDate, createObjectiveRequest.ObjType, createObjectiveRequest.UserID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, struct {
		Message string `json:"message"`
	}{Message: "Objective created successfully"})
}

func (s *Server) abandonObjectiveHandler(w http.ResponseWriter, r *http.Request) {
	var abandonObjectiveRequest types.AbandonObjectiveRequest
	err := utils.ParseJSON(r, &abandonObjectiveRequest)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = s.db.AbandonObjective(abandonObjectiveRequest.ObjectiveID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, struct {
		Message string `json:"message"`
	}{Message: "Objective abandoned successfully"})
}

func (s *Server) getObjectivesHandler(w http.ResponseWriter, r *http.Request) {
	var getObjectivesRequest types.GetObjectivesRequest
	err := utils.ParseJSON(r, &getObjectivesRequest)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	objectives, err := s.db.GetObjective(getObjectivesRequest.UserID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, objectives)
}

func (s *Server) createGoalHandler(w http.ResponseWriter, r *http.Request) {
	var createGoalRequest types.CreateGoalRequest
	err := utils.ParseJSON(r, &createGoalRequest)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = s.db.CreateGoal(createGoalRequest.Name, createGoalRequest.Description, createGoalRequest.ObjectiveID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, struct {
		Message string `json:"message"`
	}{Message: "Goal created successfully"})
}

func (s *Server) updateGoalHandler(w http.ResponseWriter, r *http.Request) {
	var updateGoalRequest types.UpdateGoalRequest
	err := utils.ParseJSON(r, &updateGoalRequest)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = s.db.UpdateGoal(updateGoalRequest.ID, updateGoalRequest.Name, updateGoalRequest.Description, updateGoalRequest.Done)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, struct {
		Message string `json:"message"`
	}{Message: "Goal updated successfully"})
}

func (s *Server) abandonGoalHandler(w http.ResponseWriter, r *http.Request) {
	var abandonGoalRequest types.AbandonGoalRequest
	err := utils.ParseJSON(r, &abandonGoalRequest)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = s.db.AbandonGoal(abandonGoalRequest.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, struct {
		Message string `json:"message"`
	}{Message: "Goal abandoned successfully"})
}
