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

	// ─────────────────────────────────────────────────────────────────────
	// Canvas: 800 × 220
	// Background: #0d1117 (GitHub dark exact)
	// Left zone (Melt): x=0..186, centred x=90
	// Right zone (bubble): x=194..790
	// Bottom bar: y=178..210 (stats + cookie button)
	//
	// GitHub-SAFE SVG:
	//   ✓ linearGradient
	//   ✓ basic shapes
	//   ✓ text / tspan
	//   ✗ NO filters, feDropShadow, feGaussianBlur
	//   ✗ NO radialGradient (stripped by GitHub)
	//   ✗ NO <style> external fonts (use system stack)
	// ─────────────────────────────────────────────────────────────────────
	svg := fmt.Sprintf(`<svg width="800" height="220" viewBox="0 0 800 220" xmlns="http://www.w3.org/2000/svg">
<defs>
  <linearGradient id="redTop" x1="0" y1="0" x2="0" y2="1">
    <stop offset="0%%" stop-color="#ff5c6e"/>
    <stop offset="100%%" stop-color="#c0192e"/>
  </linearGradient>
  <linearGradient id="redSide" x1="0" y1="0" x2="1" y2="0">
    <stop offset="0%%" stop-color="#ff5c6e"/>
    <stop offset="100%%" stop-color="#9e0f22"/>
  </linearGradient>
  <linearGradient id="redDark" x1="0" y1="0" x2="0" y2="1">
    <stop offset="0%%" stop-color="#d41e35"/>
    <stop offset="100%%" stop-color="#8a0c1d"/>
  </linearGradient>
  <linearGradient id="bubbleGrad" x1="0" y1="0" x2="0" y2="1">
    <stop offset="0%%" stop-color="#161b22"/>
    <stop offset="100%%" stop-color="#0d1117"/>
  </linearGradient>
  <linearGradient id="btnGrad" x1="0" y1="0" x2="0" y2="1">
    <stop offset="0%%" stop-color="#21262d"/>
    <stop offset="100%%" stop-color="#161b22"/>
  </linearGradient>
  <linearGradient id="coreGrad" x1="0" y1="0" x2="0" y2="1">
    <stop offset="0%%" stop-color="#a78bfa"/>
    <stop offset="100%%" stop-color="#7c3aed"/>
  </linearGradient>
</defs>

<style>
  .idea    { font: 700 19px -apple-system, BlinkMacSystemFont, 'Segoe UI', Helvetica, Arial, sans-serif; fill: #e6edf3; }
  .sub     { font: 400 11px -apple-system, BlinkMacSystemFont, 'Segoe UI', Helvetica, Arial, sans-serif; fill: #484f58; }
  .nametag { font: 600 10px -apple-system, BlinkMacSystemFont, 'Segoe UI', Helvetica, Arial, sans-serif; fill: #8b949e; }
  .btn     { font: 600 10px -apple-system, BlinkMacSystemFont, 'Segoe UI', Helvetica, Arial, sans-serif; fill: #c9d1d9; letter-spacing: 0.06em; }
</style>

<!-- CARD BACKGROUND -->
<rect x="1" y="1" width="798" height="218" rx="12" fill="#0d1117" stroke="#30363d" stroke-width="1.5"/>

<!-- red accent top -->
<rect x="1" y="1" width="798" height="3" rx="2" fill="#c0192e" opacity="0.7"/>

<!-- ═══ MELT — SITTING RED BAYMAX ═══ -->

<!-- LEFT FOOT -->
<ellipse cx="56" cy="202" rx="26" ry="13" fill="url(#redDark)"/>
<ellipse cx="50" cy="197" rx="10" ry="5" fill="#ff8090" opacity="0.28"/>

<!-- RIGHT FOOT -->
<ellipse cx="124" cy="202" rx="26" ry="13" fill="url(#redDark)"/>
<ellipse cx="118" cy="197" rx="10" ry="5" fill="#ff8090" opacity="0.28"/>

<!-- LEFT THIGH -->
<ellipse cx="62" cy="183" rx="30" ry="22" fill="url(#redTop)"/>
<ellipse cx="56" cy="174" rx="12" ry="8" fill="#ff8090" opacity="0.22"/>

<!-- RIGHT THIGH -->
<ellipse cx="118" cy="183" rx="30" ry="22" fill="url(#redTop)"/>
<ellipse cx="112" cy="174" rx="12" ry="8" fill="#ff8090" opacity="0.22"/>

<!-- BODY -->
<ellipse cx="90" cy="158" rx="52" ry="56" fill="url(#redTop)"/>
<ellipse cx="74" cy="136" rx="18" ry="14" fill="#ff8090" opacity="0.2"/>
<ellipse cx="116" cy="168" rx="16" ry="22" fill="#7a0b1c" opacity="0.3"/>

<!-- LEFT ARM -->
<ellipse cx="48" cy="154" rx="20" ry="30" fill="url(#redSide)" transform="rotate(-25 48 154)"/>
<ellipse cx="42" cy="144" rx="8" ry="11" fill="#ff8090" opacity="0.18" transform="rotate(-25 42 144)"/>
<ellipse cx="74" cy="177" rx="24" ry="16" fill="url(#redDark)" transform="rotate(-10 74 177)"/>

<!-- RIGHT ARM (crosses over left) -->
<ellipse cx="132" cy="154" rx="20" ry="30" fill="url(#redSide)" transform="rotate(25 132 154)"/>
<ellipse cx="126" cy="144" rx="8" ry="11" fill="#ff8090" opacity="0.18" transform="rotate(25 126 144)"/>
<ellipse cx="108" cy="174" rx="26" ry="17" fill="url(#redTop)" transform="rotate(8 108 174)"/>

<!-- CLASPED HANDS -->
<ellipse cx="90" cy="188" rx="16" ry="11" fill="url(#redDark)"/>
<ellipse cx="86" cy="184" rx="6" ry="4" fill="#ff8090" opacity="0.18"/>

<!-- CHEST CORE -->
<circle cx="90" cy="157" r="11" fill="#12101a"/>
<circle cx="90" cy="157" r="7" fill="url(#coreGrad)"/>
<circle cx="90" cy="157" r="3" fill="#ddd6fe"/>
<circle cx="88" cy="155" r="1.5" fill="white" opacity="0.45"/>

<!-- NECK -->
<rect x="72" y="107" width="36" height="18" rx="9" fill="url(#redTop)"/>

<!-- HEAD -->
<ellipse cx="90" cy="88" rx="52" ry="50" fill="url(#redTop)"/>
<ellipse cx="70" cy="66" rx="20" ry="16" fill="#ff8090" opacity="0.2"/>
<ellipse cx="108" cy="108" rx="18" ry="12" fill="#7a0b1c" opacity="0.28"/>

<!-- FACE PLATE -->
<ellipse cx="90" cy="92" rx="36" ry="28" fill="url(#redDark)"/>
<ellipse cx="90" cy="96" rx="30" ry="22" fill="#6a0918" opacity="0.38"/>

<!-- EYES -->
<ellipse cx="77"  cy="89" rx="8.5" ry="9.5" fill="#0d1117"/>
<ellipse cx="103" cy="89" rx="8.5" ry="9.5" fill="#0d1117"/>
<ellipse cx="74"  cy="85" rx="3"   ry="2.5" fill="white" opacity="0.55"/>
<ellipse cx="100" cy="85" rx="3"   ry="2.5" fill="white" opacity="0.55"/>

<!-- CHEEKS -->
<ellipse cx="62"  cy="103" rx="7" ry="4" fill="#ff8fa3" opacity="0.4"/>
<ellipse cx="118" cy="103" rx="7" ry="4" fill="#ff8fa3" opacity="0.4"/>

<!-- NAME BUBBLE -->
<rect x="20" y="6" width="140" height="34" rx="8" fill="#161b22" stroke="#30363d" stroke-width="1"/>
<polygon points="76,40 84,52 92,40" fill="#161b22" stroke="#30363d" stroke-width="1" stroke-linejoin="round"/>
<polygon points="77,41 84,50 91,41" fill="#161b22"/>
<text x="90" y="21" class="nametag" text-anchor="middle">Hello! I'm Melt 👋</text>
<text x="90" y="34" class="nametag" text-anchor="middle">by Advik-chan ✨</text>

<!-- DIVIDER -->
<line x1="186" y1="14" x2="186" y2="206" stroke="#21262d" stroke-width="1"/>

<!-- ═══ IDEA BUBBLE ═══ -->
<rect x="194" y="14" width="596" height="158" rx="10" fill="url(#bubbleGrad)" stroke="#21262d" stroke-width="1.5"/>
<!-- arrow -->
<polygon points="194,93 179,85 179,101" fill="#161b22" stroke="#21262d" stroke-width="1.5" stroke-linejoin="round"/>
<polygon points="195,87 183,85 195,93" fill="#161b22"/>
<polygon points="195,99 183,101 195,93" fill="#161b22"/>

<!-- IDEA TEXT — big, centred -->
<text class="idea">%s</text>

<!-- ═══ BOTTOM BAR ═══ -->
<text x="200" y="198" class="sub">✨ %d ideas   🍪 %d cookies   👆 click the image for a new idea</text>

<rect x="574" y="183" width="210" height="28" rx="6" fill="url(#btnGrad)" stroke="#30363d" stroke-width="1"/>
<text x="679" y="201" class="btn" text-anchor="middle">🍪 Feed Melt a Cookie</text>

</svg>`, textSVG, totalClicks, totalCookies)

	fmt.Fprint(w, svg)
}
