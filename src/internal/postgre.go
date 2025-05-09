package postgres

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

)

const (
	maxRetries = 5
	retryDelay = 5 * time.Second
)

func ConnectWithRetries(cfg config.DB) (*sqlx.DB, error) {
	var db *sqlx.DB
	var err error

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Username, cfg.Name, cfg.Password)
	logger.Debug("preparing connect to PostgreSQL", slog.String("dsn", dsn))
	for i := 0; i < maxRetries; i++ {
		db, err = sqlx.Open("postgres", dsn)
		if err == nil {
			if err = db.Ping(); err == nil {
				logger.Info("Successfully connected to PostgreSQL",
					slog.String("host", cfg.Host),
					slog.String("db", cfg.Name))
				return db, nil
			}
		}

		logger.Warn("Failed to connect to PostgreSQL, retrying...",
			slog.Int("attempt", i+1),
			slog.String("error", err.Error()))

		time.Sleep(retryDelay)
	}

	return nil, fmt.Errorf("failed to connect after %d attempts: %v", maxRetries, err)
}
