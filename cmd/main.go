package main

import (
	"brand_scout_tt/internal/handlers"
	"database/sql"
	_ "github.com/lib/pq"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}
func main() {
	db := setupDatabase()
	defer db.Close()

	router := http.NewServeMux()

	router.HandleFunc("POST /quotes", handlers.NewQuoteCreateHandler(db))
	router.HandleFunc("GET /quotes", handlers.NewQuoteHandler(db))
	router.HandleFunc("GET /quotes/random", handlers.NewRandomQuoteHandler(db))
	router.HandleFunc("DELETE /quotes/{id}", handlers.NewQuoteDeleteHandler(db))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	slog.Info("Starting server", "port", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		slog.Error("Server failed", "error", err)
		os.Exit(1)
	}
}

func setupDatabase() *sql.DB {
	dsn := os.Getenv("DSN")
	if dsn == "" {
		slog.Error("DSN environment variable is not set")
		os.Exit(1)
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		slog.Error("Failed to connect to database", "error", err)
		os.Exit(1)
	}

	if err := db.Ping(); err != nil {
		slog.Error("Failed to ping database", "error", err)
		os.Exit(1)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS quotes (
			id SERIAL PRIMARY KEY,
			author TEXT NOT NULL,
			quote TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT NOW()
		)
	`)
	if err != nil {
		slog.Error("Failed to create table", "error", err)
		os.Exit(1)
	}

	_, err = db.Exec(`
    INSERT INTO quotes (author, quote) VALUES
    ('Альберт Эйнштейн', 'Воображение важнее знания.'),
    ('Марк Твен', 'Никогда не откладывай на завтра то, что можно сделать послезавтра.'),
    ('Фрэнк Заппа', 'Так много книг — так мало времени.'),
    ('Оскар Уайльд', 'Будь собой, все остальные роли уже заняты.'),
    ('Конфуций', 'Выбери работу по душе, и тебе не придется работать ни дня в жизни.'),
    ('Теодор Рузвельт', 'Верь, что можешь — и ты уже на полпути.'),
    ('Махатма Ганди', 'Будь тем изменением, которое хочешь видеть в мире.'),
    ('Стив Джобс', 'Оставайтесь голодными. Оставайтесь безрассудными.'),
    ('Лев Толстой', 'Все счастливые семьи похожи друг на друга, каждая несчастливая семья несчастлива по-своему.'),
    ('Федор Достоевский', 'Красота спасет мир.')
`)
	if err != nil {
		slog.Error("Failed to insert test quotes", "error", err)
		os.Exit(1)
	}

	return db
}
