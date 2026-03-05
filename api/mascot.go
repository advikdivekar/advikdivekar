package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

// Wraps text so it never breaks out of the bubble
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
	return ideas[clickCount%len(ideas)]
}

type BotState struct {
	Cookies int `json:"cookies"`
}

// Bulletproof cookie file reader for Vercel
func getCookieCount() int {
	data, err := os.ReadFile("bot_state.json")
	if err != nil {
		data, err = os.ReadFile("../bot_state.json")
	}
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
	lines := wrapText(idea, 45)

	// Dynamic centering math so text is always perfectly in the middle of the box
	startY := 105
	if len(lines) == 2 {
		startY = 92
	} else if len(lines) == 3 {
		startY = 80
	} else if len(lines) == 4 {
		startY = 68
	}

	textSVG := ""
	for _, line := range lines {
		textSVG += fmt.Sprintf(`<tspan x="605" y="%d" text-anchor="middle">%s</tspan>`, startY, line)
		startY += 26
	}

	w.Header().Set("Content-Type", "image/svg+xml")
	w.Header().Set("Cache-Control", "no-cache, max-age=0, no-store, must-revalidate")

	svg := fmt.Sprintf(`
	<svg width="850" height="240" viewBox="0 0 850 240" fill="none" xmlns="http://www.w3.org/2000/svg">
		<defs>
			<linearGradient id="armorGradient" x1="0%%" y1="0%%" x2="0%%" y2="100%%">
				<stop offset="0%%" stop-color="#ff4b6e" />
				<stop offset="100%%" stop-color="#a81533" />
			</linearGradient>
			<filter id="glow">
				<feGaussianBlur stdDeviation="2.5" result="coloredBlur"/>
				<feMerge>
					<feMergeNode in="coloredBlur"/>
					<feMergeNode in="SourceGraphic"/>
				</feMerge>
			</filter>
			<filter id="shadow" x="-10%%" y="-10%%" width="120%%" height="120%%">
				<feDropShadow dx="0" dy="5" stdDeviation="5" flood-color="#8b5cf6" flood-opacity="0.3"/>
			</filter>
		</defs>
		<style>
			.text-main { font: 600 17px 'Segoe UI', Ubuntu, sans-serif; fill: #ffffff; }
			.text-sub { font: 500 14px 'Segoe UI', Ubuntu, sans-serif; fill: #a0a0a0; }
			.text-intro { font: italic 600 13px 'Segoe UI', Ubuntu, sans-serif; fill: #a78bfa; }
			
			@keyframes wave {
				0%% { transform: rotate(0deg); }
				25%% { transform: rotate(40deg); }
				50%% { transform: rotate(0deg); }
				75%% { transform: rotate(40deg); }
				100%% { transform: rotate(0deg); }
			}
			.arm-wave {
				transform-origin: 160px 105px;
				animation: wave 2.5s infinite ease-in-out;
			}
		</style>
		
		<circle cx="150" cy="55" r="4" fill="#1a1b26" stroke="#8b5cf6" stroke-width="2"/>
		<circle cx="165" cy="40" r="7" fill="#1a1b26" stroke="#8b5cf6" stroke-width="2"/>
		
		<rect x="185" y="10" width="160" height="60" rx="30" fill="#1a1b26" stroke="#8b5cf6" stroke-width="2.5" filter="url(#shadow)"/>
		<text x="265" y="34" class="text-intro" text-anchor="middle">Hi, I'm Melt-Chan!</text>
		<text x="265" y="54" class="text-intro" text-anchor="middle">Advik's Idea Generator</text>

		<ellipse cx="60" cy="135" rx="22" ry="45" fill="url(#armorGradient)" transform="rotate(25 60 135)" />
		<circle cx="68" cy="105" r="7" fill="#8b5cf6" opacity="0.8" filter="url(#glow)"/> 

		<path d="M 70 100 C 40 160, 45 230, 110 230 C 175 230, 180 160, 150 100 Z" fill="url(#armorGradient)" />
		<path d="M 75 110 C 55 160, 60 220, 110 220 C 160 220, 165 160, 145 110 Z" fill="#7a0f25" opacity="0.4"/>
		
		<circle cx="110" cy="140" r="16" fill="#121011" opacity="0.6"/>
		<circle cx="110" cy="140" r="9" fill="#8b5cf6" filter="url(#glow)"/> 

		<g class="arm-wave">
			<ellipse cx="160" cy="135" rx="22" ry="45" fill="url(#armorGradient)" />
			<rect x="150" y="145" width="20" height="25" fill="#121011" rx="4" opacity="0.5" /> 
		</g>

		<ellipse cx="110" cy="75" rx="45" ry="32" fill="url(#armorGradient)" />
		<ellipse cx="110" cy="78" rx="36" ry="24" fill="#D14836" />
		<ellipse cx="70" cy="78" rx="6" ry="12" fill="#E4405F" transform="rotate(-10 70 78)"/>
		<ellipse cx="150" cy="78" rx="6" ry="12" fill="#E4405F" transform="rotate(10 150 78)"/>
		
		<circle cx="92" cy="78" r="6" fill="#121011" /> 
		<circle cx="128" cy="78" r="6" fill="#121011" /> 
		<line x1="92" y1="78" x2="128" y2="78" stroke="#121011" stroke-width="4" /> 
		
		<circle cx="75" cy="86" r="5" fill="#ff708d" opacity="0.8"/>
		<circle cx="145" cy="86" r="5" fill="#ff708d" opacity="0.8"/>

		<rect x="375" y="15" width="460" height="170" rx="16" fill="#1a1b26" stroke="#8b5cf6" stroke-width="2.5" filter="url(#shadow)" />
		
		<path d="M 375 75 L 350 85 L 375 95 Z" fill="#1a1b26" stroke="#8b5cf6" stroke-width="2.5" stroke-linejoin="round"/>
		<path d="M 377 77 L 354 85 L 377 93 Z" fill="#1a1b26" /> 
		
		<text class="text-main">%s</text>
		
		<text x="605" y="220" class="text-sub" text-anchor="middle">✨ Ideas generated: %d   |   🍪 Cookies eaten: %d   |   👆 Click me!</text>
	</svg>`, textSVG, totalClicks, totalCookies)

	fmt.Fprint(w, svg)
}
