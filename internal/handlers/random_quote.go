package handlers

import (
	"brand_scout_tt/internal/models"
	"database/sql"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
)

func NewRandomQuoteHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		quote, err := getRandomQuote(db)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				http.Error(w, "No quotes available", http.StatusNotFound)
			} else {
				slog.Error("Failed to get random quote", "error", err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(quote)
	}
}

func getRandomQuote(db *sql.DB) (models.Quote, error) {
	var q models.Quote
	err := db.QueryRow(`
		SELECT id, author, quote FROM quotes
		OFFSET floor(random() * (SELECT count(*) FROM quotes))
		LIMIT 1
	`).Scan(&q.ID, &q.Author, &q.Quote)
	return q, err
}
