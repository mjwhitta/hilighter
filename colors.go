package hilighter

import (
	"fmt"
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
	colorized = onlyCodes.ReplaceAllString(colorized, "")

	return colorized
}

func colorize(clr string, str string, args ...interface{}) string {
	if Disable {
		// Return the string w/o any color codes
		return Plain(str, args...)
	}

	// Call the appropriate function
	if strings.HasPrefix(clr, "on_") {
		return bgColor(clr, str, args...)
	} else {
		return fgColor(clr, str, args...)
	}
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
	colorized = onlyCodes.ReplaceAllString(colorized, "")

	return colorized
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
	return onlyCodes.ReplaceAllString(strings.Join(out, "\n"), "")
}

func plainBg(str string, args ...interface{}) string {
	return bgCodes.ReplaceAllString(fmt.Sprintf(str, args...), "")
}

func plainFg(str string, args ...interface{}) string {
	return fgCodes.ReplaceAllString(fmt.Sprintf(str, args...), "")
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
	return onlyCodes.ReplaceAllString(strings.Join(out, "\n"), "")
}

func rainbowColors() []int {
	// Don't include black, white, light_black, and light_white
	return []int{31, 32, 33, 34, 35, 36, 91, 92, 93, 94, 95, 96}
}
