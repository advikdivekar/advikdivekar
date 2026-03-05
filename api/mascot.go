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
	lines := wrapText(idea, 34)

	// Bubble: y=52..200, centre=126
	lineHeight := 23
	totalTextHeight := len(lines) * lineHeight
	startY := 126 - totalTextHeight/2 + lineHeight/2

	textSVG := ""
	for _, line := range lines {
		textSVG += fmt.Sprintf(`<tspan x="535" y="%d" text-anchor="middle">%s</tspan>`, startY, line)
		startY += lineHeight
	}

	w.Header().Set("Content-Type", "image/svg+xml")
	w.Header().Set("Cache-Control", "no-cache, max-age=0, no-store, must-revalidate")

	svg := fmt.Sprintf(`<svg width="860" height="300" viewBox="0 0 860 300" fill="none" xmlns="http://www.w3.org/2000/svg">
<defs>
  <radialGradient id="bodyRad" cx="35%%" cy="30%%" r="65%%">
    <stop offset="0%%"   stop-color="#ff8080"/>
    <stop offset="55%%"  stop-color="#e02040"/>
    <stop offset="100%%" stop-color="#8b0f22"/>
  </radialGradient>
  <radialGradient id="headRad" cx="38%%" cy="28%%" r="60%%">
    <stop offset="0%%"   stop-color="#ff9090"/>
    <stop offset="50%%"  stop-color="#e02040"/>
    <stop offset="100%%" stop-color="#8b0f22"/>
  </radialGradient>
  <radialGradient id="limbRad" cx="30%%" cy="25%%" r="65%%">
    <stop offset="0%%"   stop-color="#ff7070"/>
    <stop offset="60%%"  stop-color="#cc1a35"/>
    <stop offset="100%%" stop-color="#7a0b1c"/>
  </radialGradient>
  <radialGradient id="faceRad" cx="40%%" cy="35%%" r="60%%">
    <stop offset="0%%"   stop-color="#d43050"/>
    <stop offset="100%%" stop-color="#7a0b1c"/>
  </radialGradient>
  <linearGradient id="borderGrad" x1="0%%" y1="0%%" x2="100%%" y2="0%%">
    <stop offset="0%%"   stop-color="#4c1d95" stop-opacity="0"/>
    <stop offset="20%%"  stop-color="#7c3aed"/>
    <stop offset="50%%"  stop-color="#a855f7"/>
    <stop offset="80%%"  stop-color="#7c3aed"/>
    <stop offset="100%%" stop-color="#4c1d95" stop-opacity="0"/>
  </linearGradient>
  <radialGradient id="shineHead" cx="38%%" cy="22%%" r="40%%">
    <stop offset="0%%"   stop-color="#ffffff" stop-opacity="0.22"/>
    <stop offset="100%%" stop-color="#ffffff" stop-opacity="0"/>
  </radialGradient>
  <radialGradient id="shineBody" cx="32%%" cy="20%%" r="45%%">
    <stop offset="0%%"   stop-color="#ffffff" stop-opacity="0.18"/>
    <stop offset="100%%" stop-color="#ffffff" stop-opacity="0"/>
  </radialGradient>
  <radialGradient id="shineArm" cx="28%%" cy="22%%" r="50%%">
    <stop offset="0%%"   stop-color="#ffffff" stop-opacity="0.2"/>
    <stop offset="100%%" stop-color="#ffffff" stop-opacity="0"/>
  </radialGradient>
  <filter id="coreGlow" x="-40%%" y="-40%%" width="180%%" height="180%%">
    <feGaussianBlur stdDeviation="4" result="blur"/>
    <feMerge><feMergeNode in="blur"/><feMergeNode in="SourceGraphic"/></feMerge>
  </filter>
  <filter id="cardShadow">
    <feDropShadow dx="0" dy="8" stdDeviation="14" flood-color="#6d28d9" flood-opacity="0.3"/>
  </filter>
  <filter id="bubbleShadow">
    <feDropShadow dx="0" dy="4" stdDeviation="8" flood-color="#6d28d9" flood-opacity="0.4"/>
  </filter>
  <filter id="btnGlow">
    <feDropShadow dx="0" dy="2" stdDeviation="7" flood-color="#a855f7" flood-opacity="0.55"/>
  </filter>
</defs>

<style>
  .idea-text { font: 600 14.5px 'Segoe UI', Ubuntu, sans-serif; fill: #ede9fe; }
  .sub-text  { font: 500 11px   'Segoe UI', Ubuntu, sans-serif; fill: #9d84c7; }
  .tag-text  { font: italic 600 10.5px 'Segoe UI', Ubuntu, sans-serif; fill: #c4b5fd; }
  .btn-text  { font: 700 10.5px 'Segoe UI', Ubuntu, sans-serif; fill: #ede9fe; letter-spacing: 0.09em; }
  .lbl-text  { font: 700 9px   'Segoe UI', Ubuntu, sans-serif; fill: #6d28d9; letter-spacing: 0.18em; }
</style>

<!-- TOP BORDER -->
<rect x="0" y="0" width="860" height="6" fill="url(#borderGrad)"/>

<!-- CARD -->
<rect x="0" y="6" width="860" height="288" fill="#0d0c18" filter="url(#cardShadow)"/>
<rect x="0"   y="6" width="3" height="288" fill="#6d28d9" opacity="0.4"/>
<rect x="857" y="6" width="3" height="288" fill="#6d28d9" opacity="0.4"/>
<text x="14" y="23" class="lbl-text">✦ MELT · IDEA BOT</text>
<line x1="215" y1="16" x2="215" y2="284" stroke="#3b1f6e" stroke-width="1" stroke-dasharray="4 4" opacity="0.45"/>

<!-- ── LEFT THIGH / LEG ── -->
<ellipse cx="72" cy="258" rx="38" ry="26" fill="url(#limbRad)"/>
<ellipse cx="72" cy="258" rx="38" ry="26" fill="url(#shineArm)"/>
<ellipse cx="52" cy="278" rx="28" ry="16" fill="url(#limbRad)"/>
<ellipse cx="52" cy="278" rx="28" ry="16" fill="url(#shineArm)"/>

<!-- ── RIGHT THIGH / LEG ── -->
<ellipse cx="148" cy="258" rx="38" ry="26" fill="url(#limbRad)"/>
<ellipse cx="148" cy="258" rx="38" ry="26" fill="url(#shineArm)"/>
<ellipse cx="164" cy="278" rx="28" ry="16" fill="url(#limbRad)"/>
<ellipse cx="164" cy="278" rx="28" ry="16" fill="url(#shineArm)"/>

<!-- ── BODY ── -->
<ellipse cx="108" cy="200" rx="62" ry="68" fill="url(#bodyRad)"/>
<ellipse cx="108" cy="200" rx="62" ry="68" fill="url(#shineBody)"/>

<!-- ── LEFT ARM ── -->
<ellipse cx="60" cy="195" rx="26" ry="36" fill="url(#limbRad)" transform="rotate(-30 60 195)"/>
<ellipse cx="60" cy="195" rx="26" ry="36" fill="url(#shineArm)" transform="rotate(-30 60 195)"/>
<ellipse cx="88" cy="222" rx="28" ry="20" fill="url(#limbRad)" transform="rotate(-15 88 222)"/>
<ellipse cx="88" cy="222" rx="28" ry="20" fill="url(#shineArm)" transform="rotate(-15 88 222)"/>

<!-- ── RIGHT ARM ── -->
<ellipse cx="156" cy="195" rx="26" ry="36" fill="url(#limbRad)" transform="rotate(30 156 195)"/>
<ellipse cx="156" cy="195" rx="26" ry="36" fill="url(#shineArm)" transform="rotate(30 156 195)"/>
<ellipse cx="122" cy="218" rx="30" ry="21" fill="url(#limbRad)" transform="rotate(10 122 218)"/>
<ellipse cx="122" cy="218" rx="30" ry="21" fill="url(#shineArm)" transform="rotate(10 122 218)"/>

<!-- ── HANDS CLASPED ── -->
<ellipse cx="105" cy="232" rx="18" ry="14" fill="url(#limbRad)"/>
<ellipse cx="105" cy="232" rx="18" ry="14" fill="url(#shineArm)"/>

<!-- ── CHEST CORE ── -->
<circle cx="108" cy="196" r="13"  fill="#12101a"/>
<circle cx="108" cy="196" r="8.5" fill="#8b5cf6" filter="url(#coreGlow)"/>
<circle cx="108" cy="196" r="3.5" fill="#c4b5fd"/>

<!-- ── NECK ── -->
<ellipse cx="108" cy="145" rx="22" ry="14" fill="url(#bodyRad)"/>

<!-- ── HEAD ── -->
<ellipse cx="108" cy="96" rx="58" ry="54" fill="url(#headRad)"/>
<ellipse cx="108" cy="96" rx="58" ry="54" fill="url(#shineHead)"/>
<path d="M 86 138 Q 108 148 130 138" stroke="#8b0f22" stroke-width="2.5" fill="none" opacity="0.6"/>

<!-- ── FACE PLATE ── -->
<ellipse cx="108" cy="100" rx="40" ry="32" fill="url(#faceRad)"/>
<ellipse cx="108" cy="104" rx="36" ry="27" fill="#6a0918" opacity="0.3"/>

<!-- ── EYES ── -->
<ellipse cx="93"  cy="96" rx="9"   ry="10" fill="#1a1020"/>
<ellipse cx="123" cy="96" rx="9"   ry="10" fill="#1a1020"/>
<ellipse cx="90"  cy="92" rx="3.5" ry="3"  fill="white" opacity="0.65"/>
<ellipse cx="120" cy="92" rx="3.5" ry="3"  fill="white" opacity="0.65"/>

<!-- ── CHEEKS ── -->
<ellipse cx="76"  cy="112" rx="8" ry="4.5" fill="#ff8fa3" opacity="0.5"/>
<ellipse cx="140" cy="112" rx="8" ry="4.5" fill="#ff8fa3" opacity="0.5"/>

<!-- ── NAME BUBBLE ── -->
<rect x="22" y="8" width="152" height="38" rx="14" fill="#150f28" stroke="#6d28d9" stroke-width="1.5"/>
<path d="M 90 46 L 83 56 L 102 46 Z" fill="#150f28" stroke="#6d28d9" stroke-width="1.5" stroke-linejoin="round"/>
<path d="M 91 47 L 85 55 L 100 47 Z" fill="#150f28"/>
<text x="98" y="27" class="tag-text" text-anchor="middle">Hello! I'm Melt 👋</text>
<text x="98" y="42" class="tag-text" text-anchor="middle">Built by Advik-chan ✨</text>

<!-- ── IDEA BUBBLE ── -->
<rect x="222" y="52" width="626" height="148" rx="16"
      fill="#130d24" stroke="#6d28d9" stroke-width="1.8" filter="url(#bubbleShadow)"/>
<path d="M 222 124 L 206 114 L 222 104 Z"
      fill="#130d24" stroke="#6d28d9" stroke-width="1.8" stroke-linejoin="round"/>
<path d="M 224 106 L 209 114 L 224 122 Z" fill="#130d24"/>

<text class="idea-text">%s</text>

<!-- STATS -->
<text x="390" y="222" class="sub-text" text-anchor="middle">
  ✨ Ideas: %d   |   🍪 Cookies: %d   |   👆 Click the widget!
</text>

<!-- COOKIE BUTTON -->
<rect x="572" y="234" width="260" height="34" rx="10"
      fill="#1a1035" stroke="#7c3aed" stroke-width="1.5" filter="url(#btnGlow)"/>
<rect x="573" y="235" width="258" height="14" rx="9" fill="#ffffff" opacity="0.03"/>
<text x="702" y="256" class="btn-text" text-anchor="middle">🍪  FEED MELT-CHAN A COOKIE</text>

<!-- BOTTOM BORDER -->
<rect x="0" y="294" width="860" height="6" fill="url(#borderGrad)"/>
</svg>`, textSVG, totalClicks, totalCookies)

	fmt.Fprint(w, svg)
}
