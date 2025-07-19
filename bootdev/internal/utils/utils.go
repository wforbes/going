package utils

import "fmt"

func Ternary[T any](condition bool, ifTrue T, ifFalse T) T {
	if condition {
		return ifTrue
	}
	return ifFalse
}

func PrintSectionStart(displayTxt string, thick bool) {
	separator := Ternary(thick, "=", "-")
	fmt.Println()
	printSeparator(separator)
	fmt.Println(displayTxt)
	printSeparator(separator)
	fmt.Println()
}

func PrintSectionEnd(thick bool) {
	separator := Ternary(thick, "=", "-")
	fmt.Println()
	printSeparator(separator)
	fmt.Println()
}

func printSeparator(char string) {
	if len(char) == 0 {
		char = "-"
	}
	if len(char) > 1 {
		// get first character of the string and set that as the 'char' var
		char = string(char[0]) // access first byte, then convert it to string?
	}
	width := 60
	var output string
	// could also use range like:
	// for range width {}
	for i := 0; i < width; i++ {
		output += char
	}
	fmt.Println(output)
}
