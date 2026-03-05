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
			db.QueryRow("SELECT clicks FROM melt_stats WHERE id = 1").Scan(&totalClicks)
		}
	}

	totalCookies := getCookieCount()
	idea := fetchGenerativeIdea(totalClicks)
	lines := wrapText(idea, 36)

	// Vertically center text in the bubble (bubble center Y ~115, height 130)
	lineHeight := 24
	totalTextHeight := len(lines) * lineHeight
	startY := 115 - totalTextHeight/2 + lineHeight/2

	textSVG := ""
	for _, line := range lines {
		textSVG += fmt.Sprintf(`<tspan x="570" dy="0" y="%d" text-anchor="middle">%s</tspan>`, startY, line)
		startY += lineHeight
	}

	w.Header().Set("Content-Type", "image/svg+xml")
	w.Header().Set("Cache-Control", "no-cache, max-age=0, no-store, must-revalidate")

	// ─────────────────────────────────────────────────────────────────────────
	// SVG: 860 × 230  (wider to give the bubble more room)
	// Melt (red Baymax) occupies x=10..185, centred at x=100
	// Speech bubble: x=210..860, centred at x=535
	// ─────────────────────────────────────────────────────────────────────────
	svg := fmt.Sprintf(`<svg width="860" height="230" viewBox="0 0 860 230" fill="none"
	xmlns="http://www.w3.org/2000/svg">
<defs>
  <!-- Red gradient for Melt's body -->
  <linearGradient id="bodyGrad" x1="0%%" y1="0%%" x2="30%%" y2="100%%">
    <stop offset="0%%"   stop-color="#ff6b6b"/>
    <stop offset="100%%" stop-color="#b91c3c"/>
  </linearGradient>
  <!-- Slightly lighter red for chest / limbs shading -->
  <linearGradient id="limbGrad" x1="0%%" y1="0%%" x2="0%%" y2="100%%">
    <stop offset="0%%"   stop-color="#e83a5a"/>
    <stop offset="100%%" stop-color="#991030"/>
  </linearGradient>
  <!-- Inner face plate -->
  <linearGradient id="faceGrad" x1="0%%" y1="0%%" x2="0%%" y2="100%%">
    <stop offset="0%%"   stop-color="#c0283e"/>
    <stop offset="100%%" stop-color="#8b1228"/>
  </linearGradient>
  <!-- Glow -->
  <filter id="glow" x="-20%%" y="-20%%" width="140%%" height="140%%">
    <feGaussianBlur stdDeviation="3" result="blur"/>
    <feMerge><feMergeNode in="blur"/><feMergeNode in="SourceGraphic"/></feMerge>
  </filter>
  <!-- Soft shadow for the whole card -->
  <filter id="cardShadow" x="-5%%" y="-5%%" width="110%%" height="120%%">
    <feDropShadow dx="0" dy="6" stdDeviation="8" flood-color="#7c3aed" flood-opacity="0.25"/>
  </filter>
  <!-- Bubble shadow -->
  <filter id="bubbleShadow" x="-5%%" y="-10%%" width="110%%" height="130%%">
    <feDropShadow dx="0" dy="4" stdDeviation="6" flood-color="#7c3aed" flood-opacity="0.35"/>
  </filter>
</defs>

<style>
  .card-bg   { font-family: 'Segoe UI', Ubuntu, sans-serif; }
  .idea-text { font: 600 15px 'Segoe UI', Ubuntu, sans-serif; fill: #f0e6ff; }
  .sub-text  { font: 500 12px 'Segoe UI', Ubuntu, sans-serif; fill: #9d84c7; }
  .tag-text  { font: italic 600 11.5px 'Segoe UI', Ubuntu, sans-serif; fill: #c4b5fd; }
</style>

<!-- ░░ BACKGROUND CARD ░░ -->
<rect x="4" y="4" width="852" height="222" rx="20"
      fill="#0f0e17" stroke="#3b1f6e" stroke-width="1.6" filter="url(#cardShadow)"/>

<!-- subtle purple shimmer strip at top -->
<rect x="4" y="4" width="852" height="4" rx="2" fill="#7c3aed" opacity="0.6"/>

<!-- ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░
     RED BAYMAX  —  "MELT"
     Body centre: (108, 148)  Head centre: (108, 72)
     ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░ -->

<!-- LEFT ARM -->
<ellipse cx="62" cy="148" rx="18" ry="38"
         fill="url(#limbGrad)" transform="rotate(-18 62 148)"/>
<!-- shading crease -->
<ellipse cx="62" cy="148" rx="8" ry="22"
         fill="#7a0f25" opacity="0.35" transform="rotate(-18 62 148)"/>

<!-- RIGHT ARM -->
<ellipse cx="154" cy="148" rx="18" ry="38"
         fill="url(#limbGrad)" transform="rotate(18 154 148)"/>
<ellipse cx="154" cy="148" rx="8" ry="22"
         fill="#7a0f25" opacity="0.35" transform="rotate(18 154 148)"/>

<!-- BODY (large rounded rect / pill) -->
<rect x="58" y="110" width="100" height="108" rx="50"
      fill="url(#bodyGrad)"/>
<!-- body shading overlay -->
<rect x="68" y="118" width="80" height="92" rx="40"
      fill="#7a0f25" opacity="0.18"/>

<!-- CHEST CIRCLE -->
<circle cx="108" cy="158" r="18" fill="#12101a"/>
<circle cx="108" cy="158" r="11" fill="#8b5cf6" filter="url(#glow)"/>
<circle cx="108" cy="158" r="5"  fill="#c4b5fd"/>

<!-- LEFT LEG -->
<rect x="74"  y="208" width="24" height="16" rx="10" fill="url(#limbGrad)"/>
<!-- RIGHT LEG -->
<rect x="118" y="208" width="24" height="16" rx="10" fill="url(#limbGrad)"/>

<!-- HEAD (big rounded circle – Baymax proportions) -->
<ellipse cx="108" cy="72" rx="52" ry="46" fill="url(#bodyGrad)"/>
<!-- head shading -->
<ellipse cx="108" cy="72" rx="42" ry="36" fill="#c0283e" opacity="0.25"/>

<!-- FACE PLATE -->
<ellipse cx="108" cy="76" rx="38" ry="28" fill="url(#faceGrad)"/>

<!-- EYES -->
<circle cx="94"  cy="72" r="7" fill="#12101a"/>
<circle cx="122" cy="72" r="7" fill="#12101a"/>
<!-- eye shine -->
<circle cx="96"  cy="70" r="2.5" fill="white" opacity="0.7"/>
<circle cx="124" cy="70" r="2.5" fill="white" opacity="0.7"/>

<!-- MOUTH LINE -->
<path d="M 94 83 Q 108 90 122 83" stroke="#12101a" stroke-width="3.5"
      stroke-linecap="round" fill="none"/>

<!-- CHEEK BLUSH -->
<ellipse cx="82"  cy="85" rx="7" ry="4" fill="#ff8fa3" opacity="0.55"/>
<ellipse cx="134" cy="85" rx="7" ry="4" fill="#ff8fa3" opacity="0.55"/>

<!-- "Hello! I'm Melt" speech bubble (small, top-left of head) -->
<rect x="50" y="8" width="164" height="42" rx="18"
      fill="#1a1330" stroke="#6d28d9" stroke-width="1.8"/>
<!-- bubble tail -->
<path d="M 106 50 L 98 58 L 116 50 Z" fill="#1a1330" stroke="#6d28d9"
      stroke-width="1.8" stroke-linejoin="round"/>
<path d="M 107 51 L 99 57 L 115 51 Z" fill="#1a1330"/>

<text x="132" y="27" class="tag-text" text-anchor="middle">Hello! I'm Melt</text>
<text x="132" y="43" class="tag-text" text-anchor="middle">Built by Advik-chan ✨</text>

<!-- ░░ IDEA SPEECH BUBBLE ░░ -->
<rect x="212" y="30" width="630" height="148" rx="18"
      fill="#160f2a" stroke="#6d28d9" stroke-width="2" filter="url(#bubbleShadow)"/>
<!-- Pointer arrow pointing left toward Melt -->
<path d="M 212 115 L 190 105 L 212 95 Z"
      fill="#160f2a" stroke="#6d28d9" stroke-width="2" stroke-linejoin="round"/>
<!-- fill the arrow seam -->
<path d="M 214 97 L 193 105 L 214 113 Z" fill="#160f2a"/>

<!-- Idea text (dynamically wrapped) -->
<text class="idea-text">%s</text>

<!-- Stats bar -->
<text x="527" y="196" class="sub-text" text-anchor="middle">
  ✨ Ideas: %d   |   🍪 Cookies: %d   |   👆 Click me!
</text>
</svg>`, textSVG, totalClicks, totalCookies)

	fmt.Fprint(w, svg)
}
