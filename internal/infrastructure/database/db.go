package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
	"github.com/yohannesgossaye/kuriftu-backend/internal/config"
)

func NewDBPool(cfg *config.Config) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), cfg.DBURL)
	if err != nil {
		log.Error().Err(err).Msg("Failed to connect to database")
		return nil, err
	}
	return pool, nil
}
