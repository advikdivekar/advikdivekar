package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strings"

	_ "github.com/lib/pq" // Postgres driver
)

// Helper to wrap text so it fits perfectly in the bubble
func wrapText(text string, maxChars int) []string {
	words := strings.Fields(text)
	var lines []string
	var currentLine string

	for _, word := range words {
		if len(currentLine)+len(word)+1 > maxChars {
			lines = append(lines, currentLine)
			currentLine = word
		} else {
			if currentLine == "" {
				currentLine = word
			} else {
				currentLine += " " + word
			}
		}
	}
	if currentLine != "" {
		lines = append(lines, currentLine)
	}
	return lines
}

// TODO: We will replace this with the Gemini API call next!
func fetchGenerativeIdea() string {
	return "Build a micro-service that randomly deletes one line of code from your repo every time you drink coffee. Call it Chaos Driven Development."
}

func Handler(w http.ResponseWriter, r *http.Request) {
	dbURL := os.Getenv("DATABASE_URL")
	totalClicks := 0

	// Fetch actual clicks from Neon DB
	if dbURL != "" {
		db, err := sql.Open("postgres", dbURL)
		if err == nil {
			defer db.Close()
			db.QueryRow("SELECT clicks FROM melt_stats WHERE id = 1").Scan(&totalClicks)
		}
	}

	idea := fetchGenerativeIdea()
	lines := wrapText(idea, 60) // Wraps text at 60 characters

	// Build the dynamic <tspan> lines for the SVG
	textSVG := ""
	yOffset := 65
	if len(lines) == 1 {
		yOffset = 80 // Center if it's only one line
	}

	for _, line := range lines {
		textSVG += fmt.Sprintf(`<tspan x="220" y="%d">%s</tspan>`, yOffset, line)
		yOffset += 22
	}

	w.Header().Set("Content-Type", "image/svg+xml")
	w.Header().Set("Cache-Control", "no-cache, max-age=0, no-store, must-revalidate")

	svg := fmt.Sprintf(`
	<svg width="850" height="220" viewBox="0 0 850 220" fill="none" xmlns="http://www.w3.org/2000/svg">
		<style>
			.text-main { font: 600 16px 'Segoe UI', Ubuntu, sans-serif; fill: #ffffff; }
			.text-sub { font: 500 14px 'Segoe UI', Ubuntu, sans-serif; fill: #a0a0a0; }
			.bubble { fill: #1a1b26; stroke: #ff708d; stroke-width: 2px; rx: 12px; }
		</style>
		
		<circle cx="100" cy="110" r="55" fill="#E4405F" />
		<circle cx="100" cy="120" r="48" fill="#D14836" />
		<circle cx="80" cy="105" r="7" fill="#121011" /> 
		<circle cx="120" cy="105" r="7" fill="#121011" /> 
		<line x1="80" y1="105" x2="120" y2="105" stroke="#121011" stroke-width="2.5" /> 
		
		<circle cx="68" cy="115" r="5" fill="#ff708d" opacity="0.6"/>
		<circle cx="132" cy="115" r="5" fill="#ff708d" opacity="0.6"/>
		<text x="65" y="185" class="text-sub" fill="#ff708d">Melt-Chan</text>

		<rect x="195" y="30" width="600" height="110" class="bubble" />
		<path d="M 195 80 L 170 95 L 195 110 Z" fill="#1a1b26" stroke="#ff708d" stroke-width="2" stroke-linejoin="round" />
		<path d="M 196 82 L 173 95 L 196 108 Z" fill="#1a1b26" /> 
		
		<text class="text-main">%s</text>
		
		<text x="210" y="175" class="text-sub">✨ Ideas generated: %d   |   👆 Click my face to generate another!</text>
	</svg>`, textSVG, totalClicks)

	fmt.Fprint(w, svg)
}
