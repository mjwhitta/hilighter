package hilighter

import "fmt"

var Colors = map[string]string{
	"black":         "30",
	"red":           "31",
	"green":         "32",
	"yellow":        "33",
	"blue":          "34",
	"magenta":       "35",
	"cyan":          "36",
	"white":         "37",
	"light_black":   "90",
	"light_red":     "91",
	"light_green":   "92",
	"light_yellow":  "93",
	"light_blue":    "94",
	"light_magenta": "95",
	"light_cyan":    "96",
	"light_white":   "97",

	"on_black":         "40",
	"on_red":           "41",
	"on_green":         "42",
	"on_yellow":        "43",
	"on_blue":          "44",
	"on_magenta":       "45",
	"on_cyan":          "46",
	"on_white":         "47",
	"on_light_black":   "100",
	"on_light_red":     "101",
	"on_light_green":   "102",
	"on_light_yellow":  "103",
	"on_light_blue":    "104",
	"on_light_magenta": "105",
	"on_light_cyan":    "106",
	"on_light_white":   "107",

	"default":    "39",
	"on_default": "49",
}

var Modes = map[string]string{
	"reset":         "0",
	"normal":        "0",
	"bold":          "1",
	"dim":           "2",
	"faint":         "2",
	"italic":        "3",
	"underline":     "4",
	"blink":         "5",
	"blink_slow":    "5",
	"blink_rapid":   "6",
	"inverse":       "7",
	"negative":      "7",
	"swap":          "7",
	"hide":          "8",
	"conceal":       "8",
	"crossed_out":   "9",
	"strikethrough": "9",
	"fraktur":       "20",

	"no_bold":          "21",
	"no_dim":           "22",
	"no_faint":         "22",
	"no_italic":        "23",
	"no_fraktur":       "23",
	"no_underline":     "24",
	"no_blink":         "25",
	"no_blink_slow":    "25",
	"no_blink_rapid":   "26",
	"no_inverse":       "27",
	"no_negative":      "27",
	"no_swap":          "27",
	"no_hide":          "28",
	"no_conceal":       "28",
	"no_crossed_out":   "29",
	"no_strikethrough": "29",
}

func init() {
	// Add all 8-bit colors, fg and bg
	for i := 0; i < 256; i++ {
		var key = fmt.Sprintf("color_%03d", i)
		var val = fmt.Sprintf("38;5;%03d", i)
		Colors[key] = val

		key = fmt.Sprintf("on_color_%03d", i)
		val = fmt.Sprintf("48;5;%03d", i)
		Colors[key] = val
	}
}
