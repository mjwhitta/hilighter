package hilighter

import "fmt"

// Print wraps fmt.Print().
func Print(args ...interface{}) {
	fmt.Print(args...)
}

// Printf wraps fmt.Printf(str string, args ...interface{}).
func Printf(str string, args ...interface{}) {
	fmt.Printf(str, args...)
}

// PrintHex will take a hex color code and print a string with ANSI
// escape codes.
func PrintHex(hex string, str string, args ...interface{}) {
	fmt.Print(Hex(hex, str, args...))
}

// PrintHilight will print a string with an ANSI escape code.
func PrintHilight(code string, str string, args ...interface{}) {
	fmt.Print(Hilight(code, str, args...))
}

// PrintHilights will print a string with ANSI escape codes.
func PrintHilights(codes []string, str string, args ...interface{}) {
	str = fmt.Sprintf(str, args...)

	// Apply all specified color codes
	for i := range codes {
		str = Hilight(codes[i], str)
	}

	// Print the result
	fmt.Print(str)
}

// Println wraps fmt.Println(args ...interface{}).
func Println(args ...interface{}) {
	fmt.Println(args...)
}

// PrintlnHex will take a hex color code and print a line with ANSI
// escape codes.
func PrintlnHex(hex string, str string, args ...interface{}) {
	fmt.Println(Hex(hex, str, args...))
}

// PrintlnHilight will print a line with an ANSI escape code.
func PrintlnHilight(code string, str string, args ...interface{}) {
	fmt.Println(Hilight(code, str, args...))
}

// PrintlnHilights will print a line with ANSI escape codes.
func PrintlnHilights(
	codes []string,
	str string,
	args ...interface{},
) {
	str = fmt.Sprintf(str, args...)

	// Apply all specified color codes
	for i := range codes {
		str = Hilight(codes[i], str)
	}

	// Print the result
	fmt.Println(str)
}

// PrintlnOnHex will take a hex color code and print a line with ANSI
// escape codes.
func PrintlnOnHex(hex string, str string, args ...interface{}) {
	fmt.Println(OnHex(hex, str, args...))
}

// PrintlnOnRainbow will print a line rotating through ANSI color
// codes for a rainbow effect.
func PrintlnOnRainbow(str string, args ...interface{}) {
	fmt.Println(OnRainbow(str, args...))
}

// PrintlnRainbow will print a line rotating through ANSI color codes
// for a rainbow effect.
func PrintlnRainbow(str string, args ...interface{}) {
	fmt.Println(Rainbow(str, args...))
}

// PrintlnWrap will wrap a line to the specified width and print it.
func PrintlnWrap(width int, str string, args ...interface{}) {
	fmt.Println(Wrap(width, str, args...))
}

// PrintOnHex will take a hex color code and print a line with ANSI
// escape codes.
func PrintOnHex(hex string, str string, args ...interface{}) {
	fmt.Print(OnHex(hex, str, args...))
}

// PrintOnRainbow will print a string rotating through ANSI color
// codes for a rainbow effect.
func PrintOnRainbow(str string, args ...interface{}) {
	fmt.Print(OnRainbow(str, args...))
}

// PrintRainbow will print a string rotating through ANSI color codes
// for a rainbow effect.
func PrintRainbow(str string, args ...interface{}) {
	fmt.Print(Rainbow(str, args...))
}

// PrintWrap will wrap a string to the specified width and print it.
func PrintWrap(width int, str string, args ...interface{}) {
	fmt.Print(Wrap(width, str, args...))
}

// Sprint wraps fmt.Sprint(args ...interface{}).
func Sprint(args ...interface{}) string {
	return fmt.Sprint(args...)
}

// Sprintf wraps fmt.Sprintf(str string, args ...interface{}).
func Sprintf(str string, args ...interface{}) string {
	return fmt.Sprintf(str, args...)
}

// Sprintln wraps fmt.Sprintln(args ...interface{}).
func Sprintln(args ...interface{}) string {
	return fmt.Sprintln(args...)
}
