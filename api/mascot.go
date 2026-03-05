package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type State struct {
	Cookies        int    `json:"cookies"`
	LastInteractor string `json:"last_interactor"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	stateURL := "https://raw.githubusercontent.com/advikdivekar/advikdivekar/main/bot_state.json"

	var botState State
	resp, err := http.Get(stateURL)
	if err == nil {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		json.Unmarshal(body, &botState)
	}

	if botState.LastInteractor == "" {
		botState.LastInteractor = "Nobody"
	}

	rand.Seed(time.Now().UnixNano())

	// Unhinged Project Ideas (using | for line breaks)
	dialogues := []string{
		"Build a dating app where orphaned|micro-services can find each other.",
		"An AI that just replies 'LGTM' and merges|every PR. Call it Tadashi-Evil.",
		"A database that randomly drops 10% of your|tables to keep you on your toes.",
		"Create an OS that bricks your PC|if you write bad C++ code.",
		"A distributed system where the nodes|just gossip about Advik's search history.",
		"Make a Web3 to-do list where it costs|$50 in gas fees to check off a task.",
		"An e-commerce dropshipping store|that only sells stolen API keys.",
		"Write your own programming language|where every syntax error deletes a file.",
		fmt.Sprintf("Click me for an idea. Or just give me|cookies. I currently have %d.", botState.Cookies),
	}

	selectedText := dialogues[rand.Intn(len(dialogues))]
	lines := strings.Split(selectedText, "|")
	line1 := lines[0]
	line2 := ""
	if len(lines) > 1 {
		line2 = lines[1]
	}

	w.Header().Set("Content-Type", "image/svg+xml")
	w.Header().Set("Cache-Control", "no-cache, max-age=0, no-store, s-maxage=0, must-revalidate")

	svg := fmt.Sprintf(`
	<svg width="850" height="220" viewBox="0 0 850 220" fill="none" xmlns="http://www.w3.org/2000/svg">
		<style>
			.text-main { font: 600 18px 'Segoe UI', Ubuntu, sans-serif; fill: #ffffff; }
			.text-sub { font: 400 13px 'Segoe UI', Ubuntu, sans-serif; fill: #a0a0a0; }
			.bubble { fill: #1a1b26; stroke: #ff708d; stroke-width: 2px; rx: 15px; }
		</style>
		
		<circle cx="100" cy="120" r="60" fill="#E4405F" />
		<circle cx="100" cy="132" r="54" fill="#D14836" />
		<circle cx="80" cy="115" r="7" fill="#121011" /> 
		<circle cx="120" cy="115" r="7" fill="#121011" /> 
		<line x1="80" y1="115" x2="120" y2="115" stroke="#121011" stroke-width="2.5" /> 
		
		<circle cx="65" cy="128" r="5" fill="#ff708d" opacity="0.6"/>
		<circle cx="135" cy="128" r="5" fill="#ff708d" opacity="0.6"/>
		
		<text x="65" y="200" class="text-sub">Melt-Chan</text>

		<rect x="200" y="30" width="620" height="100" class="bubble" />
		<path d="M 200 80 L 175 95 L 200 110 Z" fill="#1a1b26" stroke="#ff708d" stroke-width="2" stroke-linejoin="round" />
		<path d="M 201 82 L 178 95 L 201 108 Z" fill="#1a1b26" /> 
		
		<text x="225" y="75" class="text-main">%s</text>
		<text x="225" y="105" class="text-main">%s</text>
		
		<text x="210" y="160" class="text-sub">🍪 Cookies Eaten: %d   |   👤 Last Fed By: @%s</text>
		<text x="210" y="180" class="text-sub">👆 Click me for a wild project idea</text>
	</svg>`, line1, line2, botState.Cookies, botState.LastInteractor)

	fmt.Fprint(w, svg)
}
