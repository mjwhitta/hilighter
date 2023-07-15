package hilighter

import (
	"fmt"
	"io"
)

// Errorf wraps fmt.Errorf().
func Errorf(format string, args ...any) error {
	return fmt.Errorf(format, args...)
}

// Fprint wraps fmt.Fprint().
func Fprint(w io.Writer, args ...any) (int, error) {
	return fmt.Fprint(w, args...)
}

// Fprintf wraps fmt.Fprintf().
func Fprintf(
	w io.Writer,
	format string,
	args ...any,
) (int, error) {
	return fmt.Fprintf(w, format, args...)
}

// Fprintln wraps fmt.Fprintln().
func Fprintln(w io.Writer, args ...any) (int, error) {
	return fmt.Fprintln(w, args...)
}

// Print wraps fmt.Print().
func Print(args ...any) {
	fmt.Print(args...)
}

// Printf wraps fmt.Printf().
func Printf(format string, args ...any) {
	fmt.Printf(format, args...)
}

// Println wraps fmt.Println().
func Println(args ...any) {
	fmt.Println(args...)
}

// Sprint wraps fmt.Sprint().
func Sprint(args ...any) string {
	return fmt.Sprint(args...)
}

// Sprintf wraps fmt.Sprintf().
func Sprintf(format string, args ...any) string {
	return fmt.Sprintf(format, args...)
}

// Sprintln wraps fmt.Sprintln().
func Sprintln(args ...any) string {
	return fmt.Sprintln(args...)
}
