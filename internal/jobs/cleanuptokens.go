package jobs

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type CleanupJob struct {
	DB *sql.DB
}

func (j CleanupJob) Run() {
	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()

	// jalankan sekali saat aplikasi start
	j.deleteExpiredTokens()

	for range ticker.C {
		j.deleteExpiredTokens()
	}
}

func (j CleanupJob) deleteExpiredTokens() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := j.DB.ExecContext(ctx, `
		DELETE FROM tokens
		WHERE expiry < NOW()
	`)
	if err != nil {
		log.Println(err)
	}
}
