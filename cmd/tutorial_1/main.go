package main

import "fmt"

func main() {
	fmt.Println("Going: the act of learning to code with Go")
	bootDevLessons()
}

func bootDevLessons() {

	printSectionStart("Chapter 1: Variables", true)
	bootDevC1L2()
	bootDevC1L3()
	bootDevC1L4()
	/*
		boot.dev Chapter 1, Lesson 5
		This is a multi-line comment,
		same as it's always been
		(https://www.boot.dev/lessons/a581676a-054c-4c8f-a95d-396d626b0803)
	*/

	/*
		boot.dev Chapter 1, Lesson 6
		Talked about Compilation, the parts of the basic hello world file
		- package main lets the Go compiler know that we want this code to compile
			and run as a standalone program, as opposed to being a library that's
			imported by other programs
		- import "fmt" imports the formatting package from the standard library
		- func main() defines the main function, the entry point for the program
	*/

	// boot.dev CH1L7 - Go compiles faster than other compiled languages, but runs slower
	//	than others like Rust, C, Zig. Though since its a compiled language it runs faster
	//	than interpreted languages like python, js, and ruby.
	bootDevC1L8()

	printSectionEnd(true)
}

func bootDevC1L2() {
	printSectionStart("Chapter 1: Lesson 2", false)
	// initializing variables with zero values
	// (https://www.boot.dev/lessons/a1eae01c-0a40-47d5-9b98-94fe48199073)
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string
	fmt.Println("Initializing with zero values:")
	fmt.Printf("%v %.2f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}

func bootDevC1L3() {
	printSectionStart("Chapter 1: Lesson 3", false)
	// using the walrus operator to declare new variables and assign values
	// (can't be used outside of function scope)
	// (https://www.boot.dev/lessons/62881ae0-89f4-44e6-b69e-50a7652a7da3)
	messageStart := "Happy birthday! You are now"
	age := 21
	messageEnd := "years old!"
	fmt.Println(messageStart, age, messageEnd)
	printSectionEnd(false)
}

func bootDevC1L4() {
	printSectionStart("Chapter 1: Lesson 4", false)
	// casting a variable to a different type and using it in an expression
	// (https://www.boot.dev/lessons/73145333-7245-4643-ae6b-e65a5f719906)
	numMessagesFromDoris := 72
	costPerMessage := .02
	totalCost := costPerMessage * float64(numMessagesFromDoris)
	fmt.Printf("Doris spent %.2f on text messages today\n", totalCost)
	printSectionEnd(false)
}

func bootDevC1L8() {
	printSectionStart("Chapter 1: Lesson 8", false)
	// whole numbers
	//	int int8 int16 int32 int64

	// positive whole numbers (unsigned integers)
	//	uint uint8 uint16 uint32 uint64 uintptr

	// signed decimal numbers
	//	float32 float64

	// complex numbers (with real and imaginary part)
	//	complex64 complex128

	// the sizes (8, 16, 32, 64, 128) are how many bits in memory will be used to store the variable
	// the default int and uint types refer to 32 or 64 depending on the users environment
	// Use the standard sizes unless you have a performance requirement
	//	int, uint, float64, complex128

	// Some types can be easily converted:
	temperatureFloat := 88.26
	temperatureInt := int64(temperatureFloat)
	fmt.Println("Converted from float:(", temperatureFloat, ") to int:(", temperatureInt, ")")
	// Casting a float to an int just truncates off the floating point part

	accountAgeFloat := 2.6
	accountAgeInt := int64(accountAgeFloat)
	fmt.Println("Your account has existed for", accountAgeInt, "years")
	printSectionEnd(false)
}

func Ternary(condition bool, ifTrue any, ifFalse any) any {
	if condition {
		return ifTrue
	}
	return ifFalse
}

func printSectionStart(displayTxt string, thick bool) {
	seperator, ok := Ternary(thick, "=", "-").(string) // assert that string is returned from Ternary
	if !ok {                                           // ok is a boolean that will be false if assertion failed
		fmt.Println(displayTxt)
		return
	}
	fmt.Println()
	printSeperator(seperator)
	fmt.Println(displayTxt)
	printSeperator(seperator)
	fmt.Println()
}

func printSectionEnd(thick bool) {
	seperator, ok := Ternary(thick, "=", "-").(string) // assert that string is returned from Ternary
	if !ok {                                           // ok is a boolean that will be false if assertion failed
		return
	}
	fmt.Println()
	printSeperator(seperator)
	fmt.Println()
}

func printSeperator(char string) {
	if len(char) == 0 {
		char = "-"
	}
	if len(char) > 1 {
		// get first character of the string and set that as the 'char' var
		char = string(char[0]) // access first byte, then convert it to string?
	}
	width := 36
	var output string
	// could also use range like:
	// for range width {}
	for i := 0; i < width; i++ {
		output += char
	}
	fmt.Println(output)
}
