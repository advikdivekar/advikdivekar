func Handler(w http.ResponseWriter, r *http.Request) {
	dbURL := os.Getenv("DATABASE_URL")

	// We update the DB, but we don't want to make the user wait forever
	if dbURL != "" {
		db, err := sql.Open("postgres", dbURL)
		if err == nil {
			defer db.Close()
			// Optimization: Execute and don't spend time scanning results
			_, _ = db.Exec("UPDATE melt_stats SET clicks = clicks + 1 WHERE id = 1")
		}
	}

	// Redirect to your profile immediately
	// Using 302 (Found) instead of 301 ensures browsers don't cache the redirect itself
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	http.Redirect(w, r, "https://github.com/advikdivekar", http.StatusFound)
}