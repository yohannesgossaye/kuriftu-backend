package main

import (
	"net/http"

	"github.com/rs/zerolog/log"
	app_auth "github.com/yohannesgossaye/kuriftu-backend/internal/application/auth"
	"github.com/yohannesgossaye/kuriftu-backend/internal/config"
	"github.com/yohannesgossaye/kuriftu-backend/internal/infrastructure/api"
	"github.com/yohannesgossaye/kuriftu-backend/internal/infrastructure/database"
	db_auth "github.com/yohannesgossaye/kuriftu-backend/internal/infrastructure/database/auth"
	"github.com/yohannesgossaye/kuriftu-backend/internal/infrastructure/database/sqlc"
	"github.com/yohannesgossaye/kuriftu-backend/internal/logger"
)

// @title Kuriftu Loyalty Program API
// @version 1.0
// @description API for Kuriftu Loyalty Program
// @host localhost:8080
// @BasePath /
func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}

	logger := logger.NewLogger()

	dbPool, err := database.NewDBPool(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}
	defer dbPool.Close()

	dbQueries := sqlc.New(dbPool)
	authRepo := db_auth.NewRepository(dbQueries)
	authSvc := app_auth.NewService(authRepo)

	r := api.SetupRoutes(authSvc, logger)

	log.Info().Msgf("Starting server on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		log.Fatal().Err(err).Msg("Server failed to start")
	}
}
