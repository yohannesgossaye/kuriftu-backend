package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/rs/zerolog"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/yohannesgossaye/kuriftu-backend/internal/application/auth"
	"github.com/yohannesgossaye/kuriftu-backend/internal/infrastructure/api/handlers"
	"github.com/yohannesgossaye/kuriftu-backend/internal/infrastructure/api/middleware"
)

func SetupRoutes(svc *auth.Service, log *zerolog.Logger) *chi.Mux {
	r := chi.NewRouter()

	// CORS middleware
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8081", "http://localhost:8080"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Use(middleware.Logger(log))

	// Serve Swagger JSON and UI
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/swagger.yaml"), // Relative path
	))
	r.Get("/swagger/swagger.yaml", func(w http.ResponseWriter, r *http.Request) {
		log.Info().Msg("Serving swagger.yaml from ./docs/swagger.yaml")
		http.ServeFile(w, r, "./docs/swagger.yaml")
	})

	r.Post("/auth/register", handlers.RegisterHandler(svc, log))

	return r
}
