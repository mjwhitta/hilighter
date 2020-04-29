package hilighter

import (
	"fmt"
	"io"
)

// Errorf wraps fmt.Errorf().
func Errorf(format string, args ...interface{}) error {
	return fmt.Errorf(format, args...)
}

// Fprint wraps fmt.Fprint().
func Fprint(w io.Writer, args ...interface{}) (int, error) {
	return fmt.Fprint(w, args...)
}

// Fprintf wraps fmt.Fprintf().
func Fprintf(
	w io.Writer,
	format string,
	args ...interface{},
) (int, error) {
	return fmt.Fprintf(w, format, args...)
}

// Fprintln wraps fmt.Fprintln().
func Fprintln(w io.Writer, args ...interface{}) (int, error) {
	return fmt.Fprintln(w, args...)
}

// Print wraps fmt.Print().
func Print(args ...interface{}) {
	fmt.Print(args...)
}

// Printf wraps fmt.Printf().
func Printf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

// PrintHex will take a hex color code and print a string with ANSI
// escape codes.
func PrintHex(hex string, str string) {
	Print(Hex(hex, str))
}

// PrintHilight will print a string with an ANSI escape code.
func PrintHilight(code string, str string) {
	Print(Hilight(code, str))
}

// PrintHilights will print a string with ANSI escape codes.
func PrintHilights(codes []string, str string) {
	// Apply all specified color codes
	for _, code := range codes {
		str = Hilight(code, str)
	}

	// Print the result
	Print(str)
}

// Println wraps fmt.Println().
func Println(args ...interface{}) {
	fmt.Println(args...)
}

// PrintlnHex will take a hex color code and print a line with ANSI
// escape codes.
func PrintlnHex(hex string, str string) {
	Println(Hex(hex, str))
}

// PrintlnHilight will print a line with an ANSI escape code.
func PrintlnHilight(code string, str string) {
	Println(Hilight(code, str))
}

// PrintlnHilights will print a line with ANSI escape codes.
func PrintlnHilights(codes []string, str string) {
	// Apply all specified color codes
	for _, code := range codes {
		str = Hilight(code, str)
	}

	// Print the result
	Println(str)
}

// PrintlnOnHex will take a hex color code and print a line with ANSI
// escape codes.
func PrintlnOnHex(hex string, str string) {
	Println(OnHex(hex, str))
}

// PrintlnOnRainbow will print a line rotating through ANSI color
// codes for a rainbow effect.
func PrintlnOnRainbow(str string) {
	Println(OnRainbow(str))
}

// PrintlnRainbow will print a line rotating through ANSI color codes
// for a rainbow effect.
func PrintlnRainbow(str string) {
	Println(Rainbow(str))
}

// PrintlnWrap will wrap a line to the specified width and print it.
func PrintlnWrap(width int, str string) {
	Println(Wrap(width, str))
}

// PrintOnHex will take a hex color code and print a line with ANSI
// escape codes.
func PrintOnHex(hex string, str string) {
	Print(OnHex(hex, str))
}

// PrintOnRainbow will print a string rotating through ANSI color
// codes for a rainbow effect.
func PrintOnRainbow(str string) {
	Print(OnRainbow(str))
}

// PrintRainbow will print a string rotating through ANSI color codes
// for a rainbow effect.
func PrintRainbow(str string) {
	Print(Rainbow(str))
}

// PrintWrap will wrap a string to the specified width and print it.
func PrintWrap(width int, str string) {
	Print(Wrap(width, str))
}

// Sprint wraps fmt.Sprint().
func Sprint(args ...interface{}) string {
	return fmt.Sprint(args...)
}

// Sprintf wraps fmt.Sprintf().
func Sprintf(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}

// Sprintln wraps fmt.Sprintln().
func Sprintln(args ...interface{}) string {
	return fmt.Sprintln(args...)
}
