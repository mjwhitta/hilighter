package hilighter

import (
	"regexp"
	"strings"
)

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
	modified = onlyCodes.ReplaceAllString(modified, "")

	return modified
}
