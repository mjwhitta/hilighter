package hilighter

import (
	"fmt"
	"regexp"
)

// Version is the package version
const Version string = "1.14.1"

var (
	// Colors maps color names to color codes
	Colors = map[string]string{
		"default":   "39",
		"ondefault": "49",

		// Foregrounds
		"black":        "30",
		"red":          "31",
		"green":        "32",
		"yellow":       "33",
		"blue":         "34",
		"magenta":      "35",
		"cyan":         "36",
		"white":        "37",
		"lightblack":   "90",
		"lightred":     "91",
		"lightgreen":   "92",
		"lightyellow":  "93",
		"lightblue":    "94",
		"lightmagenta": "95",
		"lightcyan":    "96",
		"lightwhite":   "97",

		// Backgrounds
		"onblack":        "40",
		"onred":          "41",
		"ongreen":        "42",
		"onyellow":       "43",
		"onblue":         "44",
		"onmagenta":      "45",
		"oncyan":         "46",
		"onwhite":        "47",
		"onlightblack":   "100",
		"onlightred":     "101",
		"onlightgreen":   "102",
		"onlightyellow":  "103",
		"onlightblue":    "104",
		"onlightmagenta": "105",
		"onlightcyan":    "106",
		"onlightwhite":   "107",
	}

	// Modes maps mode names to mode codes
	Modes = map[string]string{
		"normal": "0",
		"reset":  "0",

		// On
		"bold":          "1",
		"dim":           "2",
		"faint":         "2",
		"italic":        "3",
		"underline":     "4",
		"blink":         "5",
		"blinkslow":     "5",
		"blinkrapid":    "6",
		"inverse":       "7",
		"negative":      "7",
		"swap":          "7",
		"conceal":       "8",
		"hide":          "8",
		"crossed_out":   "9",
		"strikethrough": "9",
		"fraktur":       "20",

		// Off
		"nobold":          "21",
		"nodim":           "22",
		"nofaint":         "22",
		"noitalic":        "23",
		"nofraktur":       "23",
		"nounderline":     "24",
		"noblink":         "25",
		"noblinkslow":     "25",
		"noblinkrapid":    "26",
		"noinverse":       "27",
		"nonegative":      "27",
		"noswap":          "27",
		"noconceal":       "28",
		"nohide":          "28",
		"nocrossed_out":   "29",
		"nostrikethrough": "29",
	}

	// Cached hex to xterm-256 8-bit mappings
	cachedXterm = map[string]string{}

	// Used to disable all color codes
	disable = false

	hexByte string = "([0-9a-f]{2})"

	// Various regular expressions
	reAllCodes = regexp.MustCompile(`\x1b\[([0-9;]*m|K)`)
	reBgCodes  = regexp.MustCompile(`\x1b\[(4|10)[0-9;]+m`)
	reFgCodes  = regexp.MustCompile(`\x1b\[[39][0-9;]+m`)
	reHexCodes = regexp.MustCompile(`(?i)(on_)?([0-9a-f]{6})`)
	reIterate  = regexp.MustCompile(
		`(\x1b\[([0-9;]*m|K))*[^\x1b](\x1b\[([0-9;]*m|K))*`,
	)
	reNewline   = regexp.MustCompile(`\n`)
	reOnlyCodes = regexp.MustCompile(
		`(^|\n)(\x1b\[([0-9;]+m|K))+(\n|$)`,
	)
	reParseHex = regexp.MustCompile(
		`(?i)^#?` + hexByte + hexByte + hexByte + hexByte + `?$`,
	)
	reWrap = regexp.MustCompile(`wrap(_(\d+))?`)
)

func init() {
	var key string
	var val string

	// Add all 8-bit colors, fg and bg
	for i := range 256 {
		key = fmt.Sprintf("color%03d", i)
		val = fmt.Sprintf("38;5;%03d", i)
		Colors[key] = val

		key = fmt.Sprintf("oncolor%03d", i)
		val = fmt.Sprintf("48;5;%03d", i)
		Colors[key] = val
	}
}
