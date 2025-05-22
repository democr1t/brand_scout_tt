package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"
)

func NewQuoteDeleteHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "Invalid quote ID", http.StatusBadRequest)
			return
		}

		if err := deleteQuote(db, id); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				http.Error(w, "Quote not found", http.StatusNotFound)
			} else {
				slog.Error("Failed to delete quote", "error", err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
			return
		}

		err = json.NewEncoder(w).Encode(map[string]string{
			"message": "Quote deleted successfully",
			"id":      strconv.Itoa(id),
		})

		if err != nil {
			http.Error(w, "Cant prepare response for you", http.StatusInternalServerError)
			slog.Error("Failed to write response", "error", err)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func deleteQuote(db *sql.DB, id int) error {
	res, err := db.Exec("DELETE FROM quotes WHERE id = $1", id)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}
