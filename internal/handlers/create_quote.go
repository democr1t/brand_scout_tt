package handlers

import (
	"brand_scout_tt/internal/models"
	"database/sql"
	"encoding/json"
	"log/slog"
	"net/http"
)

func NewQuoteCreateHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var quote models.Quote
		if err := json.NewDecoder(r.Body).Decode(&quote); err != nil {
			slog.Warn("Invalid request body", "error", err)
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if quote.Author == "" || quote.Quote == "" {
			http.Error(w, "Author and text are required", http.StatusBadRequest)
			return
		}

		var id int
		err := db.QueryRow(
			"INSERT INTO quotes (author, quote) VALUES ($1, $2) RETURNING id",
			quote.Author, quote.Quote,
		).Scan(&id)

		if err != nil {
			slog.Error("Failed to create quote", "error", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		quote.ID = id
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(quote)
	}
}
