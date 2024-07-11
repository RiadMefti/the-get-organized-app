package server

import (
	"backend/internal/auth-service"
	"backend/internal/types"
	"backend/internal/utils"
	"fmt"
	"net/http"
)

func (s *Server) registerHandler(w http.ResponseWriter, r *http.Request) {

	var userRegistration types.UserRegistration
	err := utils.ParseJSON(r, &userRegistration)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	id, err := auth.CreateUser(s.db, userRegistration)
	if err != nil {
		utils.WriteError(w, http.StatusConflict, err)
		return
	}

	jwt, jwtErr := auth.GenerateToken(fmt.Sprint(id), userRegistration.Email)
	if jwtErr != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, struct {
		Jwt string `json:"jwt"`
	}{Jwt: jwt})
}

func (s *Server) loginHandler(w http.ResponseWriter, r *http.Request) {
	var userLogin types.UserLogin
	err := utils.ParseJSON(r, &userLogin)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	id, err := auth.AuthentificateUser(s.db, userLogin)

	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, err)
		return
	}

	jwt, jwtErr := auth.GenerateToken(fmt.Sprint(id), userLogin.Email)
	if jwtErr != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, struct {
		Jwt string `json:"jwt"`
	}{Jwt: jwt})
}

func (s *Server) isAuthenticatedHandler(w http.ResponseWriter, r *http.Request) {

	var jwtToken types.JwtToken

	err := utils.ParseJSON(r, &jwtToken)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	_, err = auth.ValidateToken(jwtToken.Jwt)

	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, struct {
		Message string `json:"message"`
	}{Message: "Authenticated"})

}
