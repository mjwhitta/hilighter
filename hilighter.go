package hilighter

import (
	"errors"
	"runtime"
	"strconv"
	"strings"
)

// Disable will prevent color codes from being used.
func Disable(b bool) {
	disable = b
}

// Hex will take a hex color code and add the appropriate ANSI color
// codes to the provided string.
func Hex(hex string, str string) string {
	return colorize(hexToXterm256(hex), str)
}

// Hilight will add the appropriate ANSI code to the specified
// string.
func Hilight(code string, str string) string {
	var clr string
	var hasKey bool
	var matches [][]string
	var width int

	// Call the appropriate function
	if _, hasKey = Colors[code]; hasKey {
		return colorize(code, str)
	} else if _, hasKey = Modes[code]; hasKey {
		return modify(code, str)
	} else {
		switch code {
		case "on_rainbow":
			return OnRainbow(str)
		case "plain":
			return Plain(str)
		case "rainbow":
			return Rainbow(str)
		default:
			// Check if hex color code
			matches = hexCode.FindAllStringSubmatch(code, -1)
			for _, match := range matches {
				clr = hexToXterm256(match[2])
				if strings.HasPrefix(code, "on_") {
					clr = "on_" + clr
				}
				return colorize(clr, str)
			}

			// Check if wrap
			matches = wrap.FindAllStringSubmatch(code, -1)
			for _, match := range matches {
				// Determine wrap width, default to 80
				width = 80
				if len(match) == 3 && len(match[2]) > 0 {
					width, _ = strconv.Atoi(match[2])
				}
				return Wrap(width, str)
			}

			// Otherwise panic
			panic(errors.New("Invalid color or mode: " + code))
		}
	}
}

// Hilights will add the appropriate ANSI codes to the specified
// string.
func Hilights(codes []string, str string) string {
	// Apply all specified color codes
	for _, code := range codes {
		str = Hilight(code, str)
	}

	return str
}

// IsDisabled will return whether or not color codes are disabled.
// Will always return true if GOOS is windows.
func IsDisabled() bool {
	if runtime.GOOS == "windows" {
		return true
	}

	return disable
}

// OnHex will take a hex color code and add the appropriate ANSI color
// codes to the provided string.
func OnHex(hex string, str string) string {
	return colorize("on_"+hexToXterm256(hex), str)
}

// OnRainbow will add multiple ANSI color codes to a string for a
// rainbow effect.
func OnRainbow(str string) string {
	if IsDisabled() {
		// Return the string w/o any color codes
		return Plain(str)
	}

	var chars []string
	var code string
	var colors []int
	var end string
	var line []string
	var lines []string
	var out []string

	// Strip all other bg color codes and split on newline
	lines = strings.Split(plainBg(str), "\n")

	// Loop thru lines and apply bg color codes
	colors = rainbowColors()
	end = "\x1b[" + Colors["on_default"] + "m"
	for i := range lines {
		chars = iterate.FindAllString(lines[i], -1)
		line = []string{}

		// Loop thru non-color-code bytes and apply on_rainbow
		for idx, char := range chars {
			code = strconv.Itoa(colors[idx%len(colors)] + 10)
			line = append(line, "\x1b["+code+"m"+char)
		}

		// Put line back together again, ensure on_default code at end
		out = append(out, strings.Join(line, "")+end)
	}

	// Put lines back together, and remove color codes if the line
	// only contains color codes
	return onlyCodes.ReplaceAllString(strings.Join(out, "\n"), "$1$4")
}

// Plain will strip all ANSI color codes from a string.
func Plain(str string) string {
	// Strip all color codes
	return allCodes.ReplaceAllString(str, "")
}

// Rainbow will add multiple ANSI color codes to a string for a
// rainbow effect.
func Rainbow(str string) string {
	if IsDisabled() {
		// Return the string w/o any color codes
		return Plain(str)
	}

	var chars []string
	var code string
	var colors []int
	var line []string
	var lines []string
	var out []string

	// Strip all other fg color codes and split on newline
	lines = strings.Split(plainFg(str), "\n")

	// Loop thru lines and apply fg color codes
	colors = rainbowColors()
	for i := range lines {
		chars = iterate.FindAllString(lines[i], -1)
		line = []string{}

		// Loop thru non-color-code bytes and apply on_rainbow
		for idx, char := range chars {
			code = strconv.Itoa(colors[idx%len(colors)])
			line = append(line, "\x1b["+code+"m"+char)
		}

		// Put line back together again
		out = append(out, strings.Join(line, ""))
	}

	// Put lines back together, and remove color codes if the line
	// only contains color codes
	return onlyCodes.ReplaceAllString(strings.Join(out, "\n"), "$1$4")
}

// Sample will show all bg/fg combos of the first 16 8-bit colors.
func Sample() {
	var bg string
	var fg string

	for f := 0; f < 16; f++ {
		for b := 0; b < 16; b++ {
			fg = Sprintf("color_%03d", f)
			bg = Sprintf("on_color_%03d", b)
			Print(colorize(fg, colorize(bg, " mw ")))
		}
		Print("\n")
	}
}

// Table will display a pretty table of all 8-bit colors.
func Table() {
	var bg string

	for i := 0; i < 16; i++ {
		bg = Sprintf("on_color_%03d", i)
		PrintHilightf(
			bg,
			" %s %s ",
			Blackf("%03d", i),
			Whitef("%03d", i),
		)
		if (i+1)%8 == 0 {
			Print("\n")
		}
	}

	for i := 16; i < 256; i++ {
		bg = Sprintf("on_color_%03d", i)
		PrintHilightf(
			bg,
			" %s %s ",
			Blackf("%03d", i),
			Whitef("%03d", i),
		)
		if (i-15)%6 == 0 {
			Print("\n")
		}
	}
}

// Wrap will wrap a string to the specified width.
func Wrap(width int, str string) string {
	var lc int
	var line = ""
	var lines []string
	var wc int
	var words = strings.Fields(str)

	str = Sprintf(str)

	// Loop thru words
	for _, word := range words {
		lc = len([]rune(Plain(line)))
		wc = len([]rune(Plain(word)))

		if lc == 0 {
			// Can't wrap less than a single word
			line = word
		} else if lc+wc+1 > width {
			// Wrap if line would be longer than width
			lines = append(lines, line)
			line = word
		} else {
			// Append word to line
			line += " " + word
		}
	}

	// Ensure last line is not forgotten
	if len([]rune(line)) != 0 {
		lines = append(lines, line)
	}

	// If original string ended with newline, put it back
	if strings.HasSuffix(str, "\n") {
		lines = append(lines, "")
	}

	// Join lines and return
	return strings.Join(lines, "\n")
}
