#!/usr/bin/env python3
"""
gif_to_ascii.py
Converts an animated GIF into an animated ASCII GIF with a cyber style.
Usage: python gif_to_ascii.py input.gif output.gif
Requires: pip install Pillow
"""

import sys
from PIL import Image, ImageDraw, ImageFont, ImageSequence


RAMP = " .`:-=+*cs#%@"   # Bright (sparse) -> Dark (dense)
COLS = 80               # Width in characters (lower = smaller file size)
FONT_SIZE = 10          # Size of the text
FILL = (143, 211, 255)  # Cyber blue (#8fd3ff)
BG = (13, 17, 23)       # GitHub Dark (#0d1117)


def gif_to_ascii(input_path, output_path):
    try:
        gif = Image.open(input_path)
    except Exception as e:
        print(f"Error opening {input_path}: {e}")
        return

    frames = []
    w, h = gif.size
    
    # Calculate rows based on aspect ratio (0.5 accounts for character height)
    ROWS = int(COLS * (h / w) * 0.5)
    
    # Try to load a monospace font (Windows/Mac/Linux common paths)
    font = None
    for font_name in ["cour.ttf", "consola.ttf", "DejaVuSansMono.ttf", "Courier.dfont"]:
        try:
            font = ImageFont.truetype(font_name, FONT_SIZE)
            break
        except IOError:
            continue
    
    if font is None:
        print("Warning: Monospace font not found. Using default (will be small).")
        font = ImageFont.load_default()

    print(f"Processing {gif.n_frames} frames...")

    for i, frame in enumerate(ImageSequence.Iterator(gif)):
        # Convert frame to grayscale and resize
        frame = frame.convert("L").resize((COLS, ROWS))
        pixels = frame.load()
        
        # Create a new blank image for the ASCII frame
        img_w = COLS * (FONT_SIZE // 2) + 10
        img_h = ROWS * FONT_SIZE + 10
        ascii_img = Image.new("RGB", (img_w, img_h), BG)
        draw = ImageDraw.Draw(ascii_img)
        
        # Build the ASCII text block
        for y in range(ROWS):
            line = ""
            for x in range(COLS):
                brightness = pixels[x, y] / 255.0
                idx = int((1.0 - brightness) * (len(RAMP) - 1))
                line += RAMP[idx]
            
            # Draw the row of text
            draw.text((5, y * FONT_SIZE + 5), line, font=font, fill=FILL)
        
        frames.append(ascii_img)

    if frames:
        # Save all frames as a new animated GIF
        frames[0].save(
            output_path,
            save_all=True,
            append_images=frames[1:],
            duration=gif.info.get('duration', 100), # Match original speed
            loop=0 # 0 = infinite loop
        )
        print(f"Success! Saved as {output_path}")
    else:
        print("No frames were processed.")

if __name__ == "__main__":
    if len(sys.argv) != 3:
        print("Usage: python gif_to_ascii.py input.gif output.gif")
    else:
        gif_to_ascii(sys.argv[1], sys.argv[2])