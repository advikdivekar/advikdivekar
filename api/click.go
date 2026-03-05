package handler

import (
	"database/sql"
	"net/http"
	"os"

	_ "github.com/lib/pq" // Postgres driver
)

func Handler(w http.ResponseWriter, r *http.Request) {
	dbURL := os.Getenv("DATABASE_URL")

	// Connect to Neon DB
	db, err := sql.Open("postgres", dbURL)
	if err == nil {
		defer db.Close()
		// Increment the click count silently
		db.Exec("UPDATE melt_stats SET clicks = clicks + 1 WHERE id = 1")
	}

	// Instantly redirect back to your GitHub profile
	http.Redirect(w, r, "https://github.com/advikdivekar", http.StatusFound)
}
