package hilighter

import (
	"errors"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func bgColor(code string, str string, args ...interface{}) string {
	// Strip all other bg color codes and don't extend bg color over
	// newlines
	str = newline.ReplaceAllString(
		plainBg(str, args...),
		"\x1b["+Colors["on_default"]+"m\n\x1b["+Colors[code]+"m",
	)

	// Wrap whole thing with specified color code
	var colorized = "\x1b[" + Colors[code] + "m" + str +
		"\x1b[" + Colors["on_default"] + "m"

	// Remove color codes, if the line only contains color codes
	return onlyCodes.ReplaceAllString(colorized, "$1$4")
}

func colorize(clr string, str string, args ...interface{}) string {
	if Disable {
		// Return the string w/o any color codes
		return Plain(str, args...)
	}

	// Call the appropriate function
	if strings.HasPrefix(clr, "on_") {
		return bgColor(clr, str, args...)
	}
	return fgColor(clr, str, args...)
}

func fgColor(code string, str string, args ...interface{}) string {
	// Strip all other fg color codes and don't extend fg color over
	// newlines
	str = newline.ReplaceAllString(
		plainFg(str, args...),
		"\x1b["+Colors["default"]+"m\n\x1b["+Colors[code]+"m",
	)

	// Wrap whole thing with specified color code
	var colorized = "\x1b[" + Colors[code] + "m" + str +
		"\x1b[" + Colors["default"] + "m"

	// Remove color codes, if the line only contains color codes
	return onlyCodes.ReplaceAllString(colorized, "$1$4")
}

func Hex(hex string, str string, args ...interface{}) string {
	return colorize(hexToXterm256(hex), str, args...)
}

// Convert hex to xterm-256 8-bit value
// https://stackoverflow.com/questions/11765623/convert-hex-hex_to_256bitto-closest-x11-color-number
func hexToXterm256(hex string) string {
	var hasKey bool
	if _, hasKey = cachedCodes[hex]; hasKey {
		return cachedCodes[hex]
	}

	// For simplicity, assume RGB space is perceptually uniform.
	// There are 5 places where one of two outputs needs to be
	// chosen when the input is the exact middle:
	// - The r/g/b channels and the gray value:
	//     - the higher value output is chosen
	// - If the gray and color values have same distance from the
	//   input
	//     - color is chosen

	// Calculate the nearest 0-based color index at 16..231
	var r, g, b = uint64(0), uint64(0), uint64(0)
	var matches = parseHex.FindAllStringSubmatch(hex, -1)
	for _, match := range matches {
		r, _ = strconv.ParseUint(match[1], 16, 8)
		g, _ = strconv.ParseUint(match[2], 16, 8)
		b, _ = strconv.ParseUint(match[3], 16, 8)
	}

	// 0..5 each
	var ir = (r - 35) / 40
	if r < 48 {
		ir = 0
	} else if r < 115 {
		ir = 1
	}
	var ig = (g - 35) / 40
	if g < 48 {
		ig = 0
	} else if g < 115 {
		ig = 1
	}
	var ib = (b - 35) / 40
	if b < 48 {
		ib = 0
	} else if b < 115 {
		ib = 1
	}

	// 0..215 lazy evaluation
	var cidx = (36 * ir) + (6 * ig) + ib + 16

	// Calculate the nearest 0-based gray index at 232..255
	var average = (r + g + b) / 3

	// 0..23
	var gidx = (average - 3) / 10
	if average > 238 {
		gidx = 23
	}

	// Calculate the represented colors back from the index
	var i2cv = []uint64{0, 0x5f, 0x87, 0xaf, 0xd7, 0xff}

	// r/g/b 0..255 each
	var cr = i2cv[ir]
	var cg = i2cv[ig]
	var cb = i2cv[ib]

	// same value for r/g/b 0..255
	var gv = (10 * gidx) + 8

	// Return the one which is nearer to the original rgb values
	var clrErr = math.Pow(float64(cr-r), 2) +
		math.Pow(float64(cg-g), 2) + math.Pow(float64(cb-b), 2)
	var grayErr = math.Pow(float64(gv-r), 2) +
		math.Pow(float64(gv-g), 2) + math.Pow(float64(gv-b), 2)

	if clrErr <= grayErr {
		cachedCodes[hex] = Sprintf("color_%03d", cidx)
	} else {
		cachedCodes[hex] = Sprintf("color_%03d", gidx+232)
	}

	return cachedCodes[hex]
}

func Hilight(code string, str string, args ...interface{}) string {
	// Call the appropriate function
	var hasKey bool
	if _, hasKey = Colors[code]; hasKey {
		return colorize(code, str, args...)
	} else if _, hasKey = Modes[code]; hasKey {
		return modify(code, str, args...)
	} else {
		switch code {
		case "on_rainbow":
			return OnRainbow(str, args...)
		case "plain":
			return Plain(str, args...)
		case "rainbow":
			return Rainbow(str, args...)
		default:
			// Check if hex color code
			var matches = hexCode.FindAllStringSubmatch(code, -1)
			for _, match := range matches {
				var clr = hexToXterm256(match[2])
				if strings.HasPrefix(code, "on_") {
					clr = "on_" + clr
				}
				return colorize(clr, str, args...)
			}

			// Check if wrap
			matches = wrap.FindAllStringSubmatch(code, -1)
			for _, match := range matches {
				// Determine wrap width, default to 80
				var width = 80
				if len(match) == 3 && len(match[2]) > 0 {
					width, _ = strconv.Atoi(match[2])
				}
				return Wrap(width, str, args...)
			}

			// Otherwise panic
			panic(errors.New("Invalid color or mode: " + code))
		}
	}
}

func Hilights(
	codes []string,
	str string,
	args ...interface{},
) string {
	str = Sprintf(str, args...)

	// Apply all specified color codes
	for i := range codes {
		str = Hilight(codes[i], str)
	}

	return str
}

func modify(mode string, str string, args ...interface{}) string {
	if Disable {
		// Return the string w/o any color codes
		return Plain(str, args...)
	}

	// Reverse mode
	var opposite = "no_" + mode
	if strings.HasPrefix(opposite, "no_no_") {
		opposite = doubleno.ReplaceAllString(opposite, "")
	}

	// Store specified mode code for removal
	var rm = Modes[mode]

	// Determine the off color code, if it exists
	var hasKey bool
	var off = ""
	if _, hasKey = Modes[opposite]; hasKey {
		// Store opposite code for removal
		rm += "|" + Modes[opposite]

		// Save opposite mode code sequence if it starts with "no_"
		if strings.HasPrefix(opposite, "no_") {
			off = "\x1b[" + Modes[opposite] + "m"
		}
	}

	// Remove other occurrences of specified mode and opposite
	var tmp = regexp.MustCompile(`\x1b\[(` + rm + ")m")
	str = tmp.ReplaceAllString(Sprintf(str, args...), "")

	// Wrap the whole thing with specified color code
	var modified = "\x1b[" + Modes[mode] + "m" + str + off

	// Remove color codes, if the line only contains color codes
	return onlyCodes.ReplaceAllString(modified, "$1$4")
}

func OnHex(hex string, str string, args ...interface{}) string {
	return colorize("on_"+hexToXterm256(hex), str, args...)
}

func OnRainbow(str string, args ...interface{}) string {
	if Disable {
		// Return the string w/o any color codes
		return Plain(str, args...)
	}

	// Strip all other bg color codes and split on newline
	var lines = strings.Split(plainBg(str, args...), "\n")
	var out []string

	// Loop thru lines and apply bg color codes
	var colors = rainbowColors()
	var end = "\x1b[" + Colors["on_default"] + "m"
	for i := range lines {
		var line []string
		var chars = iterate.FindAllString(lines[i], -1)

		// Loop thru non-color-code bytes and apply on_rainbow
		for idx, char := range chars {
			var code = strconv.Itoa(colors[idx%len(colors)] + 10)
			line = append(line, "\x1b["+code+"m"+char)
		}

		// Put line back together again, ensure on_default code at end
		out = append(out, strings.Join(line, "")+end)
	}

	// Put lines back together, and remove color codes if the line
	// only contains color codes
	return onlyCodes.ReplaceAllString(strings.Join(out, "\n"), "$1$4")
}

func Plain(str string, args ...interface{}) string {
	// Strip all color codes
	return allCodes.ReplaceAllString(Sprintf(str, args...), "")
}

func plainBg(str string, args ...interface{}) string {
	return bgCodes.ReplaceAllString(Sprintf(str, args...), "")
}

func plainFg(str string, args ...interface{}) string {
	return fgCodes.ReplaceAllString(Sprintf(str, args...), "")
}

func Rainbow(str string, args ...interface{}) string {
	if Disable {
		// Return the string w/o any color codes
		return Plain(str, args...)
	}

	// Strip all other fg color codes and split on newline
	var lines = strings.Split(plainFg(str, args...), "\n")
	var out []string

	// Loop thru lines and apply fg color codes
	var colors = rainbowColors()
	for i := range lines {
		var line []string
		var chars = iterate.FindAllString(lines[i], -1)

		// Loop thru non-color-code bytes and apply on_rainbow
		for idx, char := range chars {
			var code = strconv.Itoa(colors[idx%len(colors)])
			line = append(line, "\x1b["+code+"m"+char)
		}

		// Put line back together again
		out = append(out, strings.Join(line, ""))
	}

	// Put lines back together, and remove color codes if the line
	// only contains color codes
	return onlyCodes.ReplaceAllString(strings.Join(out, "\n"), "$1$4")
}

func rainbowColors() []int {
	// Don't include black, white, light_black, and light_white
	return []int{31, 32, 33, 34, 35, 36, 91, 92, 93, 94, 95, 96}
}

func Sample() {
	// Show all bg/fg combos of the first 16 colors
	var fg, bg string
	for f := 0; f < 16; f++ {
		for b := 0; b < 16; b++ {
			fg = Sprintf("color_%03d", f)
			bg = Sprintf("on_color_%03d", b)
			Print(colorize(fg, colorize(bg, " mw ")))
		}
		Print("\n")
	}
}

func Table() {
	// Show a pretty table of all 8-bit colors
	var bg string
	for i := 0; i < 16; i++ {
		bg = Sprintf("on_color_%03d", i)
		Print(
			colorize(
				bg,
				" %s %s ",
				Black("%03d", i),
				White("%03d", i),
			),
		)
		if (i+1)%8 == 0 {
			Print("\n")
		}
	}
	for i := 16; i < 256; i++ {
		bg = Sprintf("on_color_%03d", i)
		Print(
			colorize(
				bg,
				" %s %s ",
				Black("%03d", i),
				White("%03d", i),
			),
		)
		if (i-15)%6 == 0 {
			Print("\n")
		}
	}
}

func Wrap(width int, str string, args ...interface{}) string {
	str = Sprintf(str, args...)

	var line = ""
	var lines []string
	var words = notwhitespace.FindAllString(str, -1)

	// Loop thru words
	for _, word := range words {
		if len(Plain(line))+len(Plain(word)) > width {
			// Wrap if line would be longer than width
			lines = append(lines, line)
			line = word
		} else if len(line) == 0 {
			// Can't wrap less than a single word
			line = word
		} else {
			// Append word to line
			line += " " + word
		}
	}

	// Ensure last line is not forgotten
	if len(line) != 0 {
		lines = append(lines, line)
	}

	// If original string ended with newline, put it back
	if strings.HasSuffix(str, "\n") {
		lines = append(lines, "")
	}

	// Join lines and return
	return strings.Join(lines, "\n")
}
