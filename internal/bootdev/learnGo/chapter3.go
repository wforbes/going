package learnGo

import (
	"fmt"
	"going/internal/utils"
)

func runChapter3Lessons() {
	chapter3Lesson1()
	chapter3Lesson2()
	chapter3Lesson3()
	//lesson 4 and 5 talked about the difference in syntax from C family languages
	//	with quiz questions
	//4:(https://www.boot.dev/lessons/afb405db-9785-4444-94f5-e76866b5b6b7)
	//5:(https://www.boot.dev/lessons/34172c62-9ae4-46ec-88d8-d8bdeab39eb4)
	//lesson 6 showed the difference between C and Go for a function declaration
	// where the function takes a func with two params as the first, an int as the second
	// and returns an int. The C style requires you to read it in-out, where Go can be read right-left
	// 6:(https://www.boot.dev/lessons/221e3837-4eba-4171-a3fc-b32a1b3cd423)
	chapter3Lesson7()
	chapter3Lesson8()
	chapter3Lesson9()
	// lesson 10 and 11 were questions about named returns vs naked returns
	//	use named returns when there are many return values as it helps self-document the func
	//	use naked returns when the function is short and relatively simple
	chapter3Lesson12()
}

func chapter3Lesson1() {
	utils.PrintSectionStart("Chapter 3: Lesson 1 - Functions", false)
	//(https://www.boot.dev/lessons/9aedf839-7d94-43f7-82d0-1d27e5d0b79c)

	potato("Lane,", " happy birthday!")
	potato("Zuck,", " hope that Metaverse thing works out")
	potato("Go", " is fantastic")

	utils.PrintSectionEnd(false)
}
func potato(s1 string, s2 string) {
	fmt.Println(concat(s1, s2))
}
func concat(s1 string, s2 string) string {
	return s1 + s2
}

func chapter3Lesson2() {
	utils.PrintSectionStart("Chapter 3: Lesson 2 - Multiple Parameters", false)
	//(https://www.boot.dev/lessons/7f503d3f-7425-496a-b50a-70815384a00c)
	// when multiple parameters are of the same type
	//	and are next to each other in the function signature
	//	the type only needs to be declared after the last parameter
	twoParameters("like", "this")
	manyParameters("like", "this", 42, "and", 100, 72, "this")
	utils.PrintSectionEnd(false)
}
func twoParameters(like, this string) {
	fmt.Println(like, this)
}
func manyParameters(a, b string, c int, d string, e, f int, g string) {
	// found a way to code 'g string' = success
	output := fmt.Sprintf("%s %s %d %s %d %d %s", a, b, c, d, e, f, g)
	fmt.Println(output)
}

// Chapter 3 Lesson 3...
//
//	this lesson introduces using unit tests within the boot.dev web ui
//	so I went ahead and tried to emulate this for any lesson here that does that.
//	The test for this will be found in chapter3_test.go on TestChapter3Lesson3 func
//	So from the root of the project the command to run the unit test for this lesson is:
//		go test -v ./internal/bootdev/learnGo -run=TestChapter3Lesson3
//	Or to run the whole learnGo unit test suite it would be:
//		go test -v ./internal/bootdev/learnGo
func chapter3Lesson3() {
	utils.PrintSectionStart("Chapter 3: Lesson 3 - Unit Test Lessons", false)
	// (https://www.boot.dev/lessons/a729ff01-7620-45a6-b330-7efb72bda67b)
	fmt.Println(getMonthlyPrice("basic"))
	fmt.Println(getMonthlyPrice("premium"))
	fmt.Println(getMonthlyPrice("enterprise"))
	fmt.Println(getMonthlyPrice("potato"))
	utils.PrintSectionEnd(false)
}

// Tested with 'TestChapter3Lesson3' unit test
func getMonthlyPrice(tier string) int {
	// NOTE: "To avoid pesky floating-point errors, we often store prices
	// in the currency's base unit. In this case, we are storing the prices
	// in pennies, and a dollar consists of 100 pennies."
	if tier == "basic" {
		return 10000
	} else if tier == "premium" {
		return 15000
	} else if tier == "enterprise" {
		return 50000
	}
	return 0
}

// go test -v ./internal/bootdev/learnGo -run=TestChapter3Lesson7
func chapter3Lesson7() {
	utils.PrintSectionStart("Chapter 3: Lesson 7 - Passing Variables by Value", false)
	//(https://www.boot.dev/lessons/351c4674-1c31-4148-b98f-1179dbcaac81)
	// Variables are passed into functions by value, not by reference
	//	by value = variables are copied into the function so the original value isn't altered by what the function does
	//	by reference = the references of variables are used in the function so the original value IS altered by the function

	//EXAMPLE:
	x := 5
	increment(x)
	fmt.Println(x) //its still going to be 5, even though increment did something to change it's copy of 'x' var

	//EXERCISE: {2, 89, 102, 26}  {2, 98, 104, 12}
	increase1 := monthlyBillIncrease(2, 89, 102)
	fmt.Printf("Monthly bill increased by %d (from %d to %d) the last month, costing %d per 'send'!\n", increase1, 89, 102, 2)
	increase2 := monthlyBillIncrease(2, 98, 104)
	fmt.Printf("Monthly bill increased by %d (from %d to %d) the last month, costing %d per 'send'!\n", increase2, 89, 102, 2)

	utils.PrintSectionEnd(false)
}
func increment(x int) {
	x++
}
func monthlyBillIncrease(costPerSend, numLastMonth, numThisMonth int) int {
	var lastMonthBill int = getBillForMonth(costPerSend, numLastMonth)
	var thisMonthBill int = getBillForMonth(costPerSend, numThisMonth)
	return thisMonthBill - lastMonthBill
}
func getBillForMonth(costPerSend, messagesSent int) int {
	return costPerSend * messagesSent
}

// go test -v ./internal/bootdev/learnGo -run=TestChapter3Lesson8
func chapter3Lesson8() {
	utils.PrintSectionStart("Chapter 3: Lesson 8 - Ignoring Return Values", false)
	//(https://www.boot.dev/lessons/185e65bf-8d3a-4419-abd0-258a457f0b88)

	fmt.Println(`Example:
	getPoint() returns 3 and 4 ...
	but if either of these values are set into '_' var
	in the caller that value will be ignored and inaccessible!`)

	x, _ := getPoint()

	fmt.Printf(`In the statement:
	x, _ := getPoint()
	the getPoint() func returned %v into the 'x' variable
	but ignored its 2nd return value by using '_' variable name
	like... you literally can't use _ in the caller`, x)
	fmt.Println(`Reasoning:
	You can avoid compiler errors like 'declared but not used'
	when your specific usecase of a function doesn't need one
	of the return values!`)

	utils.PrintSectionEnd(false)
}
func getPoint() (x int, y int) {
	return 3, 4
}
func getProductMessage(tier string) string {
	quantityMsg, priceMsg, _ := getProductInfo(tier)
	return "You get " + quantityMsg + " for " + priceMsg + "."
}

func getProductInfo(tier string) (string, string, string) {
	if tier == "basic" {
		return "1,000 texts per month", "$30 per month", "most popular"
	} else if tier == "premium" {
		return "50,000 texts per month", "$60 per month", "best value"
	} else if tier == "enterprise" {
		return "unlimited texts per month", "$100 per month", "customizable"
	} else {
		return "", "", ""
	}
}

func chapter3Lesson9() {
	utils.PrintSectionStart("Chapter 3: Lesson 9 - Named Return Values", false)
	//(https://www.boot.dev/lessons/bfd3eabd-58f5-4fe1-b59d-876b83bf52e8)
	// EXAMPLE: these two functions are the effectively the same:
	// 	getCoords1 defines the named return values of x and y
	//	which causes them to be initialized as 0
	//	and automatically returned if the return statement is empty (or naked)
	getCoords1()
	//	getCoords2 defines the return type - two unnamed ints
	//	which means they need to be initialized to be given a name
	//	and then returned by those names
	getCoords2()
	// use named returns when there are many values being returned
	//	use naked returns when the function is short and simple

	//EXERCISE: in yearsUntilEvents func, naming the return values
	//	allows for being able to immediately work with them as if
	//	they were initialized to 0, and then the empty return statement
	//	results in the named return values in the signature to be returned
	age := 13
	adult, drinking, carRental := yearsUntilEvents(age)
	fmt.Printf("You're %d years old? You'll be an adult in %d years!\n", age, adult)
	fmt.Printf("You're %d years old? You'll be able to drink in %d years!\n", age, drinking)
	fmt.Printf("You're %d years old? You'll be able to rent a car in %d years!\n", age, carRental)

	utils.PrintSectionEnd(false)
}
func getCoords1() (x, y int) {
	return
}
func getCoords2() (int, int) {
	var x int
	var y int
	return x, y
}
func yearsUntilEvents(age int) (yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental int) {
	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}
	return
}

func chapter3Lesson12() {
	utils.PrintSectionStart("Chapter 3: Lesson 12 - Explicit Returns", false)
	//(https://www.boot.dev/lessons/067df3cb-a240-4f10-8159-b04a737e5002)
	fmt.Println(`Even when named return values are used, we can still explicitly return values if we want
func getCoords() (x, y int) {
	return x, y // this is explicit
}

The return values can even be overwritten
func getCoords() (x, y int) {
    return 5, 6 // this is explicit, x and y are NOT returned
}

Or if we want the return values defined in the function signature we can use naked return
func getCoords() (x, y int) {
    return // implicitly returns x and y
}`)
	fmt.Println("")
	age := 17
	adult, drinking, carRental := yearsUntilEvents12(age)
	fmt.Printf("You're %d years old? You'll be an adult in %d years!\n", age, adult)
	fmt.Printf("You're %d years old? You'll be able to drink in %d years!\n", age, drinking)
	fmt.Printf("You're %d years old? You'll be able to rent a car in %d years!\n", age, carRental)

	utils.PrintSectionEnd(false)
}

func yearsUntilEvents12(age int) (yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental int) {
	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}
	return yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental //explicitly returning the values
}
