package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

// State represents the data in bot_state.json
type State struct {
	Cookies        int    `json:"cookies"`
	Roasts         int    `json:"roasts"`
	LastInteractor string `json:"last_interactor"`
	CurrentMood    string `json:"current_mood"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// 1. Fetch the latest state directly from your GitHub repo
	// Note: Once you push your repo, this URL will work natively.
	stateURL := "https://raw.githubusercontent.com/advikdivekar/advikdivekar/main/bot_state.json"

	var botState State
	resp, err := http.Get(stateURL)
	if err == nil {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		json.Unmarshal(body, &botState)
	}

	// Fallback data if repo isn't pushed yet
	if botState.LastInteractor == "" {
		botState.LastInteractor = "Nobody"
		botState.Cookies = 0
		botState.Roasts = 0
	}

	// 2. Random Dialogue Generator
	rand.Seed(time.Now().UnixNano())
	dialogues := []string{
		"I'm made of clay, but Advik's code is spaghetti.",
		fmt.Sprintf("I demand more cookies. I only have %d.", botState.Cookies),
		"Distributed systems? I can't even distribute my own weight.",
		"Follow him for a cookie, or I'll melt onto your keyboard.",
		fmt.Sprintf("Shoutout to %s for interacting with me!", botState.LastInteractor),
		"Give me a project idea, I'm bored.",
	}
	selectedText := dialogues[rand.Intn(len(dialogues))]

	// 3. Set Headers so GitHub doesn't cache the image aggressively
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Header().Set("Cache-Control", "no-cache, max-age=0, no-store, s-maxage=0, must-revalidate")

	// 4. Render the SVG
	svg := fmt.Sprintf(`
	<svg width="600" height="200" viewBox="0 0 600 200" fill="none" xmlns="http://www.w3.org/2000/svg">
		<style>
			.text-main { font: 600 16px 'Segoe UI', Ubuntu, sans-serif; fill: #ffffff; }
			.text-sub { font: 400 12px 'Segoe UI', Ubuntu, sans-serif; fill: #a0a0a0; }
			.bubble { fill: #1a1b26; stroke: #ff708d; stroke-width: 2px; rx: 15px; }
		</style>
		
		<circle cx="80" cy="120" r="50" fill="#E4405F" />
		<circle cx="80" cy="130" r="45" fill="#D14836" /> <circle cx="65" cy="115" r="6" fill="#121011" /> <circle cx="95" cy="115" r="6" fill="#121011" /> <line x1="65" y1="115" x2="95" y2="115" stroke="#121011" stroke-width="2" /> <circle cx="55" cy="125" r="4" fill="#ff708d" opacity="0.6"/>
		<circle cx="105" cy="125" r="4" fill="#ff708d" opacity="0.6"/>
		
		<text x="50" y="190" class="text-sub">Melt-Chan</text>

		<rect x="150" y="40" width="420" height="80" class="bubble" />
		<path d="M 150 80 L 130 90 L 150 100 Z" fill="#1a1b26" stroke="#ff708d" stroke-width="2" stroke-linejoin="round" />
		<path d="M 151 82 L 133 90 L 151 98 Z" fill="#1a1b26" /> <text x="170" y="85" class="text-main">%s</text>
		
		<text x="170" y="160" class="text-sub">🍪 Cookies: %d   |   🔥 Roasts: %d   |   👤 Last interacted: @%s</text>
	</svg>`, selectedText, botState.Cookies, botState.Roasts, botState.LastInteractor)

	fmt.Fprint(w, svg)
}
