package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strings"

	_ "github.com/lib/pq" // Postgres driver
)

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

// The Hardcoded Brain - Zero latency, zero API limits
func fetchGenerativeIdea(clickCount int) string {
	ideas := []string{
		"Build a Go microservice that randomly deletes a repo when your Razorpay payment fails.",
		"Create a Flutter app that shocks you if your Python script throws a syntax error.",
		"Train an AI bot to automatically decline all GitHub pull requests with passive aggressive comments.",
		"Build a distributed system where every node is just a really old Android phone.",
		"Write a dropshipping bot that only sells products nobody wants using Razorpay test mode.",
		"Create a Python script that changes your wallpaper to a random stack trace daily.",
		"Build a Go backend that translates all SQL queries into Shakespearean English before executing.",
		"Make a Flutter widget that slowly shrinks every time you tap it until it vanishes.",
		"Train a neural network to predict when you will abandon your current side project.",
		"Write an operating system in Python that only runs one infinite loop and then crashes.",
		"Build a bot that buys you a coffee via Razorpay every time you fix a bug.",
		"Create a distributed database where data is stored entirely in Discord chat messages.",
		"Write a Go server that deliberately delays responses based on the current weather in Mumbai.",
		"Make an AI that generates worse project ideas than this one and deploys them automatically.",
		"Build a dropshipping store that randomly changes its prices every three seconds.",
		"Create a Flutter app that requires you to scream to increase the phone volume.",
		"Write a Python script that uninstalls a random library every time you hit save.",
		"Build a Go CLI tool that git commits with AI generated lyrics from early 2000s pop.",
		"Train an AI model to detect if you are crying through your webcam while debugging.",
		"Create a database that uses Google Sheets as its primary storage engine.",
		"Build a microservice that sends a Razorpay invoice to anyone who opens a GitHub issue.",
		"Write a Flutter clone of Tinder but for matching orphaned Git branches.",
		"Create a Python bot that replies LGTM to every PR but secretly introduces infinite loops.",
		"Build a Go application that only compiles if the moon is currently full.",
		"Make a distributed system that consensus checks by asking random people on Twitter.",
		"Write a dropshipping script that only sources items that are completely out of stock.",
		"Create an AI bot that explains your code back to you condescendingly.",
		"Build a Go routine that continuously mines crypto but donates it to your enemies.",
		"Write a Python backend that requires users to solve a captcha to log out.",
		"Make a Flutter UI where all the buttons actively run away from your cursor.",
		"Create a Razorpay integration that randomly discounts items by a fraction of a cent.",
		"Build a bot that automatically emails your manager when you browse Reddit during work hours.",
		"Write a Go server that refuses to handle HTTP requests if they are not polite enough.",
		"Train an AI to rewrite your clean Python code into unreadable spaghetti code.",
		"Create a distributed queue where messages are delivered via carrier pigeon APIs.",
		"Build a Flutter app that completely forgets your data every time you close it.",
		"Write a script that automatically buys domain names for every dumb idea you type.",
		"Make a dropshipping AI that writes product descriptions entirely in Morse code.",
		"Create a Go microservice that translates REST calls into SOAP just to inflict pain.",
		"Build a Python tool that changes your system language to Klingon on compile errors.",
		"Write a Razorpay webhook that triggers a physical confetti cannon in your room.",
		"Make a Flutter game where the only objective is to compile the C++ engine.",
		"Create a bot that analyzes your Spotify history and judges your coding music.",
		"Build a distributed cache that purposely forgets things just to keep you guessing.",
		"Write a script that replaces every semicolon in someone else's code with a Greek question mark.",
		"Train an AI to automatically blame the intern in the git commit history.",
		"Create a Python server that only accepts requests written in pure binary.",
		"Build a Flutter calendar app that randomly deletes your meetings for the thrill of it.",
		"Write a Razorpay script that bills you one rupee every time you use console log.",
		"Make a Go CLI that plays sad trombone sounds when your build inevitably fails.",
	}

	// Use modulo to cycle through the array safely without ever going out of bounds
	ideaIndex := clickCount % len(ideas)
	return ideas[ideaIndex]
}

func Handler(w http.ResponseWriter, r *http.Request) {
	dbURL := os.Getenv("DATABASE_URL")
	totalClicks := 0

	// 1. Fetch live clicks from Neon Database
	if dbURL != "" {
		db, err := sql.Open("postgres", dbURL)
		if err == nil {
			defer db.Close()
			db.QueryRow("SELECT clicks FROM melt_stats WHERE id = 1").Scan(&totalClicks)
		}
	}

	// 2. Fetch an idea based on the click count
	idea := fetchGenerativeIdea(totalClicks)
	lines := wrapText(idea, 60)

	textSVG := ""
	yOffset := 65
	if len(lines) == 1 {
		yOffset = 80
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
