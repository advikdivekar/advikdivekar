#!/usr/bin/env python3
"""
make_ascii_svg.py
High-detail colored ASCII SVG generator.
Requires: pip install Pillow numpy
"""
import sys
import numpy as np
from PIL import Image, ImageOps

# --- CONFIGURATION ---
RAMP = " .`:-=+*cs#%@"   # Detailed character ramp
COLS = 160               # HIGHER resolution for maximum face detail
CHAR_W = 5.0             # Character width in SVG units
CHAR_H = 9               # Character height in SVG units
FONT_SIZE = 9            # Font size
BG_COLOR = "#0d1117"     # GitHub Dark background
ROW_DURATION = 0.55      # seconds per row wipe
ROW_STAGGER = 0.045      # seconds between successive rows starting
COLOR_QUANTIZATION = 64  # INCREASED to 64 for finer skin tones and shadows
# ---------------------

def escape(s: str) -> str:
    return s.replace("&", "&amp;").replace("<", "&lt;").replace(">", "&gt;")

def image_to_colored_rows(path: str):
    try:
        img = Image.open(path).convert("RGB")
    except Exception as e:
        print(f"Error opening image: {e}")
        return []
    
    # Auto-contrast makes the face detailing pop in ASCII
    img = ImageOps.autocontrast(img, cutoff=1)
    
    w, h = img.size
    rows = int(COLS * (h / w) * 0.5)
    
    img = img.resize((COLS, rows))
    
    # Quantize colors to 64 for detailed face shading
    q_img = img.quantize(colors=COLOR_QUANTIZATION, method=Image.Quantize.MEDIANCUT)
    palette = q_img.getpalette()
    
    # Grayscale mapping for character density
    gray_img = img.convert("L")
    gray_arr = np.array(gray_img, dtype=np.float32) / 255.0
    
    colored_rows = []
    for r in range(rows):
        line = []
        for c in range(COLS):
            idx = int((1.0 - gray_arr[r, c]) * (len(RAMP) - 1))
            char = RAMP[idx]
            
            color_idx = q_img.getpixel((c, r))
            r_col = palette[color_idx * 3]
            g_col = palette[color_idx * 3 + 1]
            b_col = palette[color_idx * 3 + 2]
            hex_color = f"#{r_col:02x}{g_col:02x}{b_col:02x}"
            
            line.append((char, hex_color))
        colored_rows.append(line)
    return colored_rows

def build_svg(colored_rows) -> str:
    if not colored_rows:
        return ""
        
    num_rows = len(colored_rows)
    width = COLS * CHAR_W + 20
    height = num_rows * CHAR_H + 20

    parts = [
        f'<svg viewBox="0 0 {width:.0f} {height:.0f}" xmlns="http://www.w3.org/2000/svg" '
        f'font-family="Menlo, Consolas, monospace" font-size="{FONT_SIZE}">',
        f'<rect width="100%" height="100%" fill="{BG_COLOR}"/>',
        "<style>",
        f"text {{ white-space: pre; }}",
        "</style>",
    ]

    for r, row_data in enumerate(colored_rows):
        y = 12 + r * CHAR_H
        start = ROW_STAGGER * r
        clip_id = f"clip{r}"
        
        text_content = ""
        current_color = None
        current_text = ""
        
        # Grouping same-colored characters keeps file size small
        for char, color in row_data:
            if color == current_color:
                current_text += char
            else:
                if current_text:
                    text_content += f'<tspan fill="{current_color}">{escape(current_text)}</tspan>'
                current_color = color
                current_text = char
        if current_text:
            text_content += f'<tspan fill="{current_color}">{escape(current_text)}</tspan>'
        
        row_width = len(row_data) * CHAR_W
        
        parts.append(
            f'<clipPath id="{clip_id}">'
            f'<rect x="10" y="{y - FONT_SIZE:.1f}" width="0" height="{CHAR_H:.1f}">'
            f'<animate attributeName="width" from="0" to="{row_width:.1f}" '
            f'begin="{start:.3f}s" dur="{ROW_DURATION}s" fill="freeze" '
            f'calcMode="spline" keySplines="0.25 0.1 0.25 1"/>'
            f"</rect>"
            f"</clipPath>"
        )
        
        parts.append(
            f'<text x="10" y="{y:.1f}" clip-path="url(#{clip_id})">{text_content}</text>'
        )
        
        parts.append(
            f'<rect x="10" y="{y - FONT_SIZE:.1f}" width="{CHAR_W}" height="{CHAR_H:.1f}" fill="#ffffff" opacity="0.8">'
            f'<animate attributeName="x" from="10" to="{10 + row_width:.1f}" '
            f'begin="{start:.3f}s" dur="{ROW_DURATION}s" fill="freeze" '
            f'calcMode="spline" keySplines="0.25 0.1 0.25 1"/>'
            f'<animate attributeName="opacity" from="0.8" to="0" '
            f'begin="{start + ROW_DURATION:.3f}s" dur="0.15s" fill="freeze"/>'
            f"</rect>"
        )

    parts.append("</svg>")
    return "\n".join(parts)

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: python make_ascii_svg.py <path_to_image>")
        sys.exit(1)
        
    src = sys.argv[1]
    print(f"Processing {src}...")
    colored_rows = image_to_colored_rows(src)
    
    if colored_rows:
        svg = build_svg(colored_rows)
        output_name = "avi-ascii.svg"
        with open(output_name, "w") as f:
            f.write(svg)
        print(f"Success! Wrote {output_name}")