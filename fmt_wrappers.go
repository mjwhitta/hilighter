package hilighter

import "fmt"

func Print(args ...interface{}) {
	fmt.Print(args...)
}

func Printf(str string, args ...interface{}) {
	fmt.Printf(str, args...)
}

func PrintHex(hex string, str string, args ...interface{}) {
	fmt.Print(Hex(hex, str, args...))
}

func PrintHilight(code string, str string, args ...interface{}) {
	fmt.Print(Hilight(code, str, args...))
}

func PrintHilights(codes []string, str string, args ...interface{}) {
	str = fmt.Sprintf(str, args...)

	// Apply all specified color codes
	for i := range codes {
		str = Hilight(codes[i], str)
	}

	// Print the result
	fmt.Print(str)
}

func Println(args ...interface{}) {
	fmt.Println(args...)
}

func PrintlnHex(hex string, str string, args ...interface{}) {
	fmt.Println(Hex(hex, str, args...))
}

func PrintlnHilight(code string, str string, args ...interface{}) {
	fmt.Println(Hilight(code, str, args...))
}

func PrintlnHilights(codes []string, str string, args ...interface{}) {
	str = fmt.Sprintf(str, args...)

	// Apply all specified color codes
	for i := range codes {
		str = Hilight(codes[i], str)
	}

	// Print the result
	fmt.Println(str)
}

func PrintlnOnHex(hex string, str string, args ...interface{}) {
	fmt.Println(OnHex(hex, str, args...))
}

func PrintlnOnRainbow(str string, args ...interface{}) {
	fmt.Println(OnRainbow(str, args...))
}

func PrintlnRainbow(str string, args ...interface{}) {
	fmt.Println(Rainbow(str, args...))
}

func PrintlnWrap(width int, str string, args ...interface{}) {
	fmt.Println(Wrap(width, str, args...))
}

func PrintOnHex(hex string, str string, args ...interface{}) {
	fmt.Print(OnHex(hex, str, args...))
}

func PrintOnRainbow(str string, args ...interface{}) {
	fmt.Print(OnRainbow(str, args...))
}

func PrintRainbow(str string, args ...interface{}) {
	fmt.Print(Rainbow(str, args...))
}

func PrintWrap(width int, str string, args ...interface{}) {
	fmt.Print(Wrap(width, str, args...))
}

func Sprint(args ...interface{}) string {
	return fmt.Sprint(args...)
}

func Sprintf(str string, args ...interface{}) string {
	return fmt.Sprintf(str, args...)
}

func Sprintln(args ...interface{}) string {
	return fmt.Sprintln(args...)
}
