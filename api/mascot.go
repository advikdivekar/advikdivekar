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
			// Increment click count on each request, return updated value
			db.QueryRow(`
				INSERT INTO melt_stats (id, clicks) VALUES (1, 1)
				ON CONFLICT (id) DO UPDATE SET clicks = melt_stats.clicks + 1
				RETURNING clicks
			`).Scan(&totalClicks)
		}
	}

	totalCookies := getCookieCount()
	idea := fetchGenerativeIdea(totalClicks)

	// Wrap idea text at 32 chars — tuned for 19px font in 596px bubble
	lines := wrapText(idea, 32)

	// Bubble inner area: y=14..172, centre y=93
	// Text block centred vertically
	lineHeight := 28
	totalTextH := len(lines) * lineHeight
	startY := 93 - totalTextH/2 + lineHeight/2 + 4

	textSVG := ""
	for _, line := range lines {
		textSVG += fmt.Sprintf(
			`<tspan x="492" y="%d" text-anchor="middle">%s</tspan>`,
			startY, line,
		)
		startY += lineHeight
	}

	w.Header().Set("Content-Type", "image/svg+xml")
	w.Header().Set("Cache-Control", "no-cache, max-age=0, no-store, must-revalidate")

	svg := fmt.Sprintf(`<svg width="800" height="255" viewBox="0 0 800 255" xmlns="http://www.w3.org/2000/svg">
<defs>
  <linearGradient id="bodyGrad" x1="0.3" y1="0" x2="0.7" y2="1">
    <stop offset="0%%" stop-color="#ff4d4d"/>
    <stop offset="50%%" stop-color="#e60000"/>
    <stop offset="100%%" stop-color="#8a0000"/>
  </linearGradient>
  <linearGradient id="armLeft" x1="0" y1="0" x2="1" y2="1">
    <stop offset="0%%" stop-color="#ff4d4d"/>
    <stop offset="100%%" stop-color="#8a0000"/>
  </linearGradient>
  <linearGradient id="armRight" x1="1" y1="0" x2="0" y2="1">
    <stop offset="0%%" stop-color="#ff4d4d"/>
    <stop offset="100%%" stop-color="#8a0000"/>
  </linearGradient>
  <linearGradient id="legGrad" x1="0" y1="0" x2="0" y2="1">
    <stop offset="0%%" stop-color="#e60000"/>
    <stop offset="100%%" stop-color="#5c0000"/>
  </linearGradient>
</defs>

<style>
  .idea { font: 700 19px -apple-system, BlinkMacSystemFont, 'Segoe UI', Helvetica, Arial, sans-serif; fill: #e6edf3; }
  .sub { font: 400 11px -apple-system, BlinkMacSystemFont, 'Segoe UI', Helvetica, Arial, sans-serif; fill: #8b949e; }
  .nametag { font: 600 10px -apple-system, BlinkMacSystemFont, 'Segoe UI', Helvetica, Arial, sans-serif; fill: #8b949e; }
  .btn { font: 600 10px -apple-system, BlinkMacSystemFont, 'Segoe UI', Helvetica, Arial, sans-serif; fill: #c9d1d9; letter-spacing: 0.06em; }
  
  .bubble-bg { fill: rgba(30, 33, 39, 0.85); stroke: rgba(60, 68, 77, 0.85); stroke-width: 1.5; }
  .bubble-arrow { fill: rgba(30, 33, 39, 0.85); }
  .bubble-arrow-outline { fill: rgba(30, 33, 39, 0.85); stroke: rgba(60, 68, 77, 0.85); stroke-width: 1.5; stroke-linejoin: round; }
  .btn-bg { fill: rgba(40, 44, 52, 0.85); stroke: rgba(60, 68, 77, 0.85); stroke-width: 1; }
</style>

<!-- BACKGROUND -->
<!-- Fully transparent background to seamlessly blend into README! -->

<!-- ═══ MELT — RED BAYMAX ═══ -->
<!-- SHADOW -->
<ellipse cx="90" cy="208" rx="60" ry="6" fill="#000" opacity="0.15"/>

<!-- LEGS -->
<ellipse cx="65" cy="186" rx="16" ry="24" fill="url(#legGrad)"/>
<ellipse cx="115" cy="186" rx="16" ry="24" fill="url(#legGrad)"/>

<!-- LEFT ARM -->
<ellipse cx="38" cy="125" rx="22" ry="60" fill="url(#armLeft)" transform="rotate(15 38 125)"/>

<!-- RIGHT ARM -->
<ellipse cx="142" cy="125" rx="22" ry="60" fill="url(#armRight)" transform="rotate(-15 142 125)"/>

<!-- BODY -->
<ellipse cx="90" cy="130" rx="58" ry="70" fill="url(#bodyGrad)"/>

<!-- HEAD -->
<ellipse cx="90" cy="58" rx="42" ry="28" fill="url(#bodyGrad)"/>

<!-- EYES (Black, connected by line) -->
<line x1="72" y1="58" x2="108" y2="58" stroke="#111" stroke-width="2.5"/>
<circle cx="72" cy="58" r="5.5" fill="#111"/>
<circle cx="108" cy="58" r="5.5" fill="#111"/>

<!-- CHEST BADGE -->
<circle cx="120" cy="100" r="7" fill="#cc0000" stroke="#8a0000" stroke-width="1.5"/>
<line x1="116" y1="100" x2="124" y2="100" stroke="#8a0000" stroke-width="1.5"/>

<!-- NAME BUBBLE -->
<rect x="25" y="4" width="130" height="32" rx="8" class="bubble-bg"/>
<polygon points="76,36 84,46 92,36" class="bubble-arrow-outline"/>
<polygon points="77,36 84,45 91,36" class="bubble-arrow"/>
<text x="90" y="18" class="nametag" text-anchor="middle">Hello! I'm Melt 👋</text>
<text x="90" y="30" class="nametag" text-anchor="middle">by Advik ✨</text>

<!-- ═══ IDEA BUBBLE ═══ -->
<rect x="194" y="14" width="596" height="158" rx="12" class="bubble-bg"/>
<!-- Arrow pointing from bubble to Mascot -->
<polygon points="194,93 179,85 194,101" class="bubble-arrow-outline"/>
<polygon points="195,87 181,93 195,99" class="bubble-arrow"/>

<!-- IDEA TEXT -->
<text class="idea">%s</text>

<!-- ═══ COUNTERS BADGE (Centered to Idea Bubble) ═══ -->
<rect x="392" y="185" width="200" height="26" rx="13" class="btn-bg"/>
<text x="492" y="202" class="sub" text-anchor="middle">✨ %d ideas  •  🍪 %d cookies</text>

<!-- ═══ FEED COOKIE BUTTON (Below Melt) ═══ -->
<rect x="5" y="222" width="180" height="26" rx="4" fill="#1a1b26" stroke="rgba(60, 68, 77, 0.85)" stroke-width="1"/>
<text x="90" y="239" class="btn" text-anchor="middle" fill="#fff">FEED MELT A COOKIE🍪</text>

</svg>`, textSVG, totalClicks, totalCookies)

	fmt.Fprint(w, svg)
}
