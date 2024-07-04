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
	"github.com/go-chi/cors"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Get("/", s.HelloWorldHandler)

	r.Get("/health", s.healthHandler)

	r.Post("/register", s.registerHandler)

	r.Post("/login", s.loginHandler)

	r.Post("/isAuthenticated", s.isAuthenticatedHandler)

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
