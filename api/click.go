package handler

import (
	"database/sql"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	dbURL := os.Getenv("DATABASE_URL")

	if dbURL != "" {
		db, err := sql.Open("postgres", dbURL)
		if err == nil {
			defer db.Close()
			// Fire-and-forget: Increments BOTH ideas and cookies instantly
			db.Exec("UPDATE melt_stats SET clicks = clicks + 1, cookies = cookies + 1 WHERE id = 1")
		}
	}

	// Tell the browser NOT to cache this redirect, then send them back instantly
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	http.Redirect(w, r, "https://github.com/advikdivekar", http.StatusFound)
}
