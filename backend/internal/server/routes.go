package server

import (
	"backend/internal/auth-service"
	"backend/internal/types"
	"backend/internal/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", s.HelloWorldHandler)

	r.Get("/health", s.healthHandler)

	r.Post("/register", s.registerHandler)

	r.Post("/login", s.loginHandler)

	return r
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}

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
