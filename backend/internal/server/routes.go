package server

import (
	"backend/internal/utils"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/google/uuid"
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

	r.Use(AuthorizationMiddlware)

	r.Get("/", s.HelloWorldHandler)

	r.Get("/health", s.healthHandler)

	r.Post("/register", s.registerHandler)

	r.Post("/login", s.loginHandler)

	r.Post("/isAuthenticated", s.isAuthenticatedHandler)

	r.Post("/createObjective", s.createObjectiveHandler)

	r.Put("/abandonObjective", s.abandonObjectiveHandler)

	r.Get("/getObjectives", s.getObjectivesHandler)

	r.Post("/createGoalHandler", s.createGoalHandler)

	r.Put("/updateGoalHandler", s.updateGoalHandler)

	r.Put("/abandonGoalHandler", s.abandonGoalHandler)

	return r
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	uuidStr := "29e7c3f5-4e72-4411-817f-15b2e11ead74"
	parsedUUID, err := uuid.Parse(uuidStr)

	if err != nil {
		log.Fatalf("error parsing UUID. Err: %v", err)
	}

	err = s.db.AbandonGoal(parsedUUID)

	if err != nil {

		log.Fatalf("error creating objective. Err: %v", err)
	}

	utils.WriteJSON(w, http.StatusOK, resp)

}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}
