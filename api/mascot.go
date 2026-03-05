package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/lib/pq"
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

	ideaIndex := clickCount % len(ideas)
	return ideas[ideaIndex]
}

type BotState struct {
	Cookies int `json:"cookies"`
}

func getCookieCount() int {
	cwd, _ := os.Getwd()
	path := filepath.Join(cwd, "bot_state.json")
	data, err := os.ReadFile(path)

	var state BotState
	if err == nil {
		json.Unmarshal(data, &state)
		return state.Cookies
	}
	return 0
}

func Handler(w http.ResponseWriter, r *http.Request) {
	dbURL := os.Getenv("DATABASE_URL")
	totalClicks := 0

	if dbURL != "" {
		db, err := sql.Open("postgres", dbURL)
		if err == nil {
			defer db.Close()
			db.QueryRow("SELECT clicks FROM melt_stats WHERE id = 1").Scan(&totalClicks)
		}
	}

	totalCookies := getCookieCount()

	idea := fetchGenerativeIdea(totalClicks)
	// Reduced max chars from 60 to 45 so text fits in the new right-aligned bubble
	lines := wrapText(idea, 45)

	textSVG := ""
	yOffset := 65
	if len(lines) == 1 {
		yOffset = 80
	}

	for _, line := range lines {
		// Moved text start X coordinate to 400
		textSVG += fmt.Sprintf(`<tspan x="400" y="%d">%s</tspan>`, yOffset, line)
		yOffset += 22
	}

	w.Header().Set("Content-Type", "image/svg+xml")
	w.Header().Set("Cache-Control", "no-cache, max-age=0, no-store, must-revalidate")

	svg := fmt.Sprintf(`
	<svg width="850" height="220" viewBox="0 0 850 220" fill="none" xmlns="http://www.w3.org/2000/svg">
		<defs>
			<linearGradient id="armorGradient" x1="0%%" y1="0%%" x2="0%%" y2="100%%">
				<stop offset="0%%" stop-color="#E4405F" />
				<stop offset="100%%" stop-color="#A82A43" />
			</linearGradient>
			<filter id="glow">
				<feGaussianBlur stdDeviation="1.5" result="coloredBlur"/>
				<feMerge>
					<feMergeNode in="coloredBlur"/>
					<feMergeNode in="SourceGraphic"/>
				</feMerge>
			</filter>
		</defs>
		<style>
			.text-main { font: 600 16px 'Segoe UI', Ubuntu, sans-serif; fill: #ffffff; }
			.text-sub { font: 500 14px 'Segoe UI', Ubuntu, sans-serif; fill: #a0a0a0; }
			.bubble { fill: #1a1b26; stroke: #8080ff; stroke-width: 2.5px; filter: drop-shadow(0px 0px 3px #8080ff); }
			.text-intro { font: italic 600 13px 'Segoe UI', Ubuntu, sans-serif; fill: #8080ff; }
			
			@keyframes wave {
				0%% { transform: rotate(0deg); }
				25%% { transform: rotate(30deg); }
				50%% { transform: rotate(0deg); }
				75%% { transform: rotate(30deg); }
				100%% { transform: rotate(0deg); }
			}
			.arm-wave {
				transform-origin: 155px 105px;
				animation: wave 3s infinite ease-in-out;
			}
		</style>
		
		<circle cx="150" cy="70" r="4" fill="#1a1b26" stroke="#8080ff" stroke-width="1.5"/>
		<circle cx="170" cy="55" r="7" fill="#1a1b26" stroke="#8080ff" stroke-width="1.5"/>

		<path d="M 190,50 A 25,25 0 0,1 215,20 A 35,35 0 0,1 275,10 A 35,35 0 0,1 335,20 A 25,25 0 0,1 360,50 A 25,25 0 0,1 335,80 L 215,80 A 25,25 0 0,1 190,50 Z" fill="#1a1b26" stroke="#8080ff" stroke-width="2.5" />
		<text x="210" y="45" class="text-intro">Hi, I'm Melt-Chan!</text>
		<text x="200" y="65" class="text-intro">Advik's Idea Generator</text>

		<ellipse cx="65" cy="115" rx="18" ry="40" fill="url(#armorGradient)" transform="rotate(35 65 115)" stroke="#121011" stroke-width="0.5"/>
		<circle cx="73" cy="95" r="7" fill="#8080ff" opacity="0.8" filter="url(#glow)"/> 
		<rect x="55" y="125" width="20" height="25" fill="#121011" rx="2" transform="rotate(35 65 115)" /> 

		<path d="M 85 95 C 60 140, 50 210, 115 210 C 180 210, 170 140, 145 95 Z" fill="url(#armorGradient)" stroke="#121011" stroke-width="0.5" />
		<path d="M 90 100 C 70 140, 62 200, 115 200 C 168 200, 160 140, 140 100 Z" fill="#D14836" stroke="#121011" stroke-width="0.5" opacity="0.6"/>
		<ellipse cx="115" cy="130" rx="22" ry="18" fill="#121011" opacity="0.4"/> 
		<circle cx="115" cy="130" r="10" fill="#8080ff" opacity="0.8" filter="url(#glow)"/> 

		<g class="arm-wave">
			<ellipse cx="165" cy="115" rx="18" ry="40" fill="url(#armorGradient)" transform="rotate(-35 165 115)" stroke="#121011" stroke-width="0.5"/>
			<circle cx="157" cy="95" r="7" fill="#8080ff" opacity="0.8" filter="url(#glow)"/> 
			<rect x="155" y="125" width="20" height="25" fill="#121011" rx="2" transform="rotate(-35 165 115)" /> 
		</g>

		<ellipse cx="115" cy="70" rx="42" ry="30" fill="url(#armorGradient)" stroke="#121011" stroke-width="0.5"/>
		<ellipse cx="115" cy="73" rx="34" ry="24" fill="#D14836" stroke="#121011" stroke-width="0.5" />
		<ellipse cx="78" cy="70" rx="8" ry="12" fill="#E4405F" transform="rotate(-10 78 70)"/>
		<ellipse cx="152" cy="70" rx="8" ry="12" fill="#E4405F" transform="rotate(10 152 70)"/>
		
		<circle cx="97" cy="73" r="5.5" fill="#121011" /> 
		<circle cx="133" cy="73" r="5.5" fill="#121011" /> 
		<line x1="97" y1="73" x2="133" y2="73" stroke="#121011" stroke-width="3.5" /> 
		
		<rect x="380" y="30" width="450" height="110" class="bubble" />
		<path d="M 380 80 L 355 95 L 380 110 Z" fill="#1a1b26" stroke="#8080ff" stroke-width="2.5" stroke-linejoin="round" />
		<path d="M 381 82 L 358 95 L 381 108 Z" fill="#1a1b26" /> 
		
		<text class="text-main">%s</text>
		
		<text x="400" y="175" class="text-sub">✨ Ideas: %d   |   🍪 Cookies: %d   |   👆 Click me!</text>
	</svg>`, textSVG, totalClicks, totalCookies)

	fmt.Fprint(w, svg)
}
