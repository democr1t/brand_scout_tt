package handlers

import (
	"brand_scout_tt/internal/models"
	"database/sql"
	"encoding/json"
	"log/slog"
	"net/http"
)

func NewQuoteHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		author := r.URL.Query().Get("author")
		var quotes []models.Quote
		var err error

		if author != "" {
			quotes, err = getQuotesByAuthor(db, author)
		} else {
			quotes, err = getAllQuotes(db)
		}

		if err != nil {
			slog.Error("Failed to get quotes", "error", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(quotes)
	}
}

func getAllQuotes(db *sql.DB) ([]models.Quote, error) {
	rows, err := db.Query("SELECT id, author, quote FROM quotes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quotes []models.Quote
	for rows.Next() {
		var q models.Quote
		if err := rows.Scan(&q.ID, &q.Author, &q.Quote); err != nil {
			return nil, err
		}
		quotes = append(quotes, q)
	}

	return quotes, nil
}

func getQuotesByAuthor(db *sql.DB, author string) ([]models.Quote, error) {
	rows, err := db.Query("SELECT id, author, quote FROM quotes WHERE author = $1", author)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quotes []models.Quote
	for rows.Next() {
		var q models.Quote
		if err := rows.Scan(&q.ID, &q.Author, &q.Quote); err != nil {
			return nil, err
		}
		quotes = append(quotes, q)
	}

	return quotes, nil
}
