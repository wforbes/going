package learnGo

import (
	"fmt"
	"going/internal/utils"
	"unicode/utf8"
)

func runChapter1Lessons() {
	c1_l2()
	c1_l3()
	c1_l4()
	/* boot.dev Chapter 1, Lesson 5
	This is a multi-line comment,
	same as it's always been
	(https://www.boot.dev/lessons/a581676a-054c-4c8f-a95d-396d626b0803)
	*/

	/* boot.dev Chapter 1, Lesson 6
	Talked about Compilation, the parts of the basic hello world file
	- package main lets the Go compiler know that we want this code to compile
		and run as a standalone program, as opposed to being a library that's
		imported by other programs
	- import "fmt" imports the formatting package from the standard library
	- func main() defines the main function, the entry point for the program
	*/

	/* boot.dev CH1L7 - Go compiles faster than other compiled languages, but runs slower
	than others like Rust, C, Zig. Though since its a compiled language it runs faster
	than interpreted languages like python, js, and ruby. */
	c1_l8()
	// lesson 9 and 10 was info about Type selection and whatnot (in the readme)
	c1_l11()
	// lesson 12 and 13 was info about compiled vs interpreted languages
	c1_l14()
	// lesson 15 and 16 was about the Go runtime and its memory usage compared to other languages
	c1_l17()
	c1_l18()
	// lesson 19 compared Go's speed with other languages
	c1_l20()
	c1_l21()
	// lesson 22, 23, and 24 are just practice fixing bugs
	c1_l25()
}

func c1_l2() {
	utils.PrintSectionStart("Chapter 1: Lesson 2 - Basic Variables", false)
	// initializing variables with zero values
	// (https://www.boot.dev/lessons/a1eae01c-0a40-47d5-9b98-94fe48199073)
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string
	fmt.Println("Initializing with zero values:")
	fmt.Printf("%v %.2f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
	utils.PrintSectionEnd(false)
}

func c1_l3() {
	utils.PrintSectionStart("Chapter 1: Lesson 3 - Short Variable Declarations", false)
	// using the walrus operator to declare new variables and assign values
	// (can't be used outside of function scope)
	// (https://www.boot.dev/lessons/62881ae0-89f4-44e6-b69e-50a7652a7da3)
	messageStart := "Happy birthday! You are now"
	age := 21
	messageEnd := "years old!"
	fmt.Println(messageStart, age, messageEnd)
	utils.PrintSectionEnd(false)
}

func c1_l4() {
	utils.PrintSectionStart("Chapter 1: Lesson 4 - Why Go?", false)
	// casting a variable to a different type and using it in an expression
	// (https://www.boot.dev/lessons/73145333-7245-4643-ae6b-e65a5f719906)
	numMessagesFromDoris := 72
	costPerMessage := .02
	totalCost := costPerMessage * float64(numMessagesFromDoris)
	fmt.Printf("Doris spent %.2f on text messages today\n", totalCost)
	utils.PrintSectionEnd(false)
}

func c1_l8() {
	utils.PrintSectionStart("Chapter 1: Lesson 8 - Type Sizes", false)
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
	utils.PrintSectionEnd(false)
}

func c1_l11() {
	utils.PrintSectionStart("Chapter 1: Lesson 11 - Go Is Statically Typed", false)
	// Go is Statically typed
	// (https://www.boot.dev/lessons/b0807eaa-38e5-4d3f-8359-ffe5e1c9ae7e)
	// You can concatenate strings with the '+' operator, but not strings and ints/floats
	var username string = "potatohead"
	var password string = "12345"
	/*  You can't do this:
	var password int = 12345
	then concatenate it like:
	username+":"+password
	*/

	fmt.Println("Authorization: Basic", username+":"+string(password))
	utils.PrintSectionEnd(false)
}

func c1_l14() {
	utils.PrintSectionStart("Chapter 1: Lesson 14 - Same Line Declarations", false)
	// You can declare multiple variables on the same line,
	// 	the variable identifiers just need to appear in the same order as the init values
	// (https://www.boot.dev/lessons/6725ea2b-ce54-443f-a304-5ae8138b31eb)
	averageOpenRate, displayMessage := .23, "is the average open rate of your messages"
	fmt.Println(averageOpenRate, displayMessage)
	utils.PrintSectionEnd(false)
}

func c1_l17() {
	utils.PrintSectionStart("Chapter 1: Lesson 17 - Constants", false)
	// Constants cant use the walrus operator and must be primative type
	//(https://www.boot.dev/lessons/30beb009-2e1c-4cae-98b2-e9738101cd56)
	const premiumPlanName = "Premium Plan"
	const basicPlanName = "Basic Plan"
	fmt.Println("plan:", premiumPlanName)
	fmt.Println("plan:", basicPlanName)
	utils.PrintSectionEnd(false)
}

func c1_l18() {
	utils.PrintSectionStart("Chaper 1: Lesson 18 - Computed Constants", false)
	// constants must be known at compile time, meaning they should be static
	//	but the CAN be computed - as long as the computation can happen at compile time
	//	That means function calls elsewhere and other operations wouldn't work
	//	But operations using the base features of the language itself directly should
	const secondsInMinute = 60
	const minutesInHour = 60
	const secondsInHour = secondsInMinute * minutesInHour // computed!
	fmt.Println("number of seconds in an hour:", secondsInHour)
	utils.PrintSectionEnd(false)
}

func c1_l20() {
	utils.PrintSectionStart("Chaper 1: Lesson 20 - Formatting Strings in Go", false)
	// Go uses printf similarly to C.

	// fmt.Sprintf returns the formatted string
	str1 := fmt.Sprintf("I am %v years old", 37)
	fmt.Println(str1)

	// fmt.Printf prints a formatted string to standard out
	// %v - a nice default that uses a primitive value in the formatting
	fmt.Printf("I am %v years old", "many, many")
	fmt.Println()

	// %s - format a string
	fmt.Printf("I am %s years old", "way too many")
	fmt.Println()

	// %d - format an integer
	fmt.Printf("I am %d years old", 99)
	fmt.Println()

	// %f - format a float
	fmt.Printf("I am %f years old", 10.523)
	fmt.Println()

	// %.2f - format a float rounded to 2 decimal places
	fmt.Printf("I am %f years old", 10.523)
	fmt.Println()

	//Exercise:
	const name = "Saul Goodman"
	const openRate = 30.5
	var msg string = fmt.Sprintf("Hi %s, your open rate is %.1f percent%s", name, openRate, "\n")
	fmt.Print(msg)

	utils.PrintSectionEnd(false)
}

func c1_l21() {
	utils.PrintSectionStart("Chapter 1: Lesson 21 - Runes and String Encoding", false)
	// notes in the readme under 'Lesson 21 notes' section
	// (https://www.boot.dev/lessons/4ae3622b-70b8-45f9-84ee-9822e44c0fbc)
	const name = "🐻"
	fmt.Printf("constant 'name' byte length: %d\n", len(name))
	fmt.Printf("constant 'name' rune length: %d\n", utf8.RuneCountInString(name))
	fmt.Println("___")
	fmt.Printf("Hi %s, so good to have you back at the boneyard", name)
	utils.PrintSectionEnd(false)
}

func c1_l25() {
	utils.PrintSectionStart("Chapter 1: Lesson 25 - Format Practice", false)
	fname := "Dalinar"
	lname := "Kholin"
	age := 45
	messageRate := 0.5
	isSubscribed := false
	message := "Sometimes a hypocrite is nothing more than a man in the process of changing."

	userLog := fmt.Sprintf("Name: %s %s, Age: %d, Rate: %.1f, Is Subscribed: %t, Message: %s", fname, lname, age, messageRate, isSubscribed, message)

	fmt.Println(userLog)
	utils.PrintSectionEnd(false)
}
