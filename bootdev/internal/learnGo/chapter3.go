package learnGo

import (
	"bootdev/internal/utils"
	"errors"
	"fmt"
)

func runChapter3Lessons() {
	c3_l1()
	c3_l2()
	c3_l3()
	//lesson 4 and 5 talked about the difference in syntax from C family languages
	//	with quiz questions
	//4:(https://www.boot.dev/lessons/afb405db-9785-4444-94f5-e76866b5b6b7)
	//5:(https://www.boot.dev/lessons/34172c62-9ae4-46ec-88d8-d8bdeab39eb4)
	//lesson 6 showed the difference between C and Go for a function declaration
	// where the function takes a func with two params as the first, an int as the second
	// and returns an int. The C style requires you to read it in-out, where Go can be read right-left
	// 6:(https://www.boot.dev/lessons/221e3837-4eba-4171-a3fc-b32a1b3cd423)
	c3_l7()
	c3_l8()
	c3_l9()
	// lesson 10 and 11 were questions about named returns vs naked returns
	//	use named returns when there are many return values as it helps self-document the func
	//	use naked returns when the function is short and relatively simple
	c3_l12()
	// lesson 13 and 14 were questions about guard clauses aka early returns
	c3_l15() // functions as values
	c3_l16() // anon functions
	c3_l17() // defer
	c3_l18() // block scope
	c3_l19() // 'Processing Orders'
	c3_l20() // closures
	c3_l21() // currying
}

func c3_l1() {
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

func c3_l2() {
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
func c3_l3() {
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
func c3_l7() {
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
func c3_l8() {
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

func c3_l9() {
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

func c3_l12() {
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

func c3_l15() {
	utils.PrintSectionStart("Chapter 3: Lesson 15 - Functions As Values", false)
	// (https://www.boot.dev/lessons/3c0f1141-9d3e-4acd-bfe8-1ebf1b44121e)
	fmt.Println(`Functions are just another type, so they can be passed as values into other functions
- This allows for a function to take a function as a parameter and call it without having to know its implementation details
- It allows the function being passed in to be used interchangably with different functions`)
	fmt.Println()

	// The exercise asks us to create the format function, applying the 'formatter' func param 3 times and printing the result
	result := reformat("This is the message", addQuestionMark)
	fmt.Println(result)

	utils.PrintSectionEnd(false)
}
func reformat(message string, formatter func(string) string) string {
	for range 3 {
		message = formatter(message)
	}
	return "TEXTIO: " + message
}

// the unit test for chapter3Lesson15 has more functions to use as formatter
func addQuestionMark(s string) string {
	return s + "?"
}

func c3_l16() {
	utils.PrintSectionStart("Chapter 3: Lesson 16 - Anonymous Functions", false)
	//(https://www.boot.dev/lessons/b79955b7-eccf-4816-8490-1dd700f13c8e)
	fmt.Println(`Anon functions do not have a name and are used when defining a func that will only be used once
or to create a quick 'closure'. Useful when:
- When a function requires a func param, it can be defined in the function call
- When you want keep a function definition in a specific block scope
- Define it into a variable to be passed around`)
	fmt.Println()

	printReports(
		"Welcome to the Hotel California",
		"Such a lovely place",
		"Plenty of room at the Hotel California",
	)

	utils.PrintSectionEnd(false)
}

func printReports(intro, body, outro string) {
	printCostReport(func(s string) int {
		return len(s) * 2
	}, intro)
	// defining it into a variable
	lengthTimes3 := func(s string) int {
		return len(s) * 3
	}
	printCostReport(lengthTimes3, body)

	if true {
		// defining it in a specific scope
		lengthTimes4 := func(s string) int {
			return len(s) * 4
		}
		printCostReport(lengthTimes4, outro)
	}
}
func printCostReport(costCalculator func(string) int, message string) {
	cost := costCalculator(message)
	fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
	fmt.Println()
}

func c3_l17() {
	utils.PrintSectionStart("Chapter3: Lesson 17 - Defer", false)
	//(https://www.boot.dev/lessons/8ea7016d-a2fd-4900-b10d-178d6e8b2ecb)
	fmt.Println(`The 'defer' keyword allows a function to be executed automatically just before its enclosing function returns
The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

Deferred functions are typically used to clean up resources that are no longer being used. Often to close database connections, file handlers, etc.`)
	fmt.Println()

	// The exercise asked to place a defer statement in bootup() so that no matter where it returns,
	// it should print the 'bootup done' message before it returns
	testBootup(true, true)
	testBootup(false, true)
	testBootup(true, false)
	testBootup(false, false)

	utils.PrintSectionEnd(false)
}
func bootup() {
	ok := connectToDB()
	defer fmt.Println("TEXTIO BOOTUP DONE")
	if !ok {
		return
	}
	ok = connectToPaymentProvider()
	if !ok {
		return
	}
	fmt.Println("All systems ready!")
}

var shouldConnectToDB = true

func connectToDB() bool {
	fmt.Println("Connecting to database...")
	if shouldConnectToDB {
		fmt.Println("Connected!")
		return true
	}
	fmt.Println("Connection failed")
	return false
}

var shouldConnectToPaymentProvider = true

func connectToPaymentProvider() bool {
	fmt.Println("Connecting to payment provider...")
	if shouldConnectToPaymentProvider {
		fmt.Println("Connected!")
		return true
	}
	fmt.Println("Connection failed")
	return false
}
func testBootup(dbSuccess, paymentSuccess bool) {
	shouldConnectToDB = dbSuccess
	shouldConnectToPaymentProvider = paymentSuccess
	bootup()
	fmt.Println("====================================")
}

func c3_l18() {
	utils.PrintSectionStart("Chapter 3: Lesson 18 - Block Scope", false)
	//(https://www.boot.dev/lessons/43edcbd3-d84b-432e-898c-62b1463aca34)
	fmt.Println(`Go is 'block-scoped' which means vars declared in a block are only accessible there
and any nested blocks. There's also package scope but that will be in a future chapter.`)
	fmt.Println()

	email := "joebobksmith@joebob.com"
	fmt.Printf("Splitting %s\n", email)
	user, domain := splitEmail(email)
	fmt.Printf("username: [ %s ] , domain: [ %s ]\n", user, domain)

	utils.PrintSectionEnd(false)
}

// splitEmail demonstrates block-scope by accessing username and domain
//
//	inside of a nested block in this functions scope.
//	The function is initially presented with the username and domain initialized in their own block,
//	and this threw an error because they weren't in block-scope with the if statement's block body
func splitEmail(email string) (string, string) {
	username, domain := "", ""
	for i, r := range email {
		if r == '@' {
			username = email[:i]
			domain = email[i+1:]
			break
		}
	}
	return username, domain
}

func c3_l19() {
	utils.PrintSectionStart("Chapter 3: Lesson 19 - Processing Orders", false)
	//(https://www.boot.dev/lessons/09628165-d910-4344-b2bf-c9145d6e6317)
	fmt.Println(`This was an exercise that required making conditional determinations based
on the returns of a couple functions - checking if the user has enough money to buy an item
and that the item quantity they want to buy are in stock.`)
	fmt.Println()

	fmt.Println("Placing order for product#1 (x2) with balance=226.95")
	success, balance := placeOrder("1", 2, 226.95) // order 2 product #1 with $226.95 balance
	fmt.Printf("Transaction %s\n", utils.Ternary(success, "SUCCESSFUL", "FAILED"))
	fmt.Printf("Remaining balance: %.2f\n", balance)

	utils.PrintSectionEnd(false)
}
func placeOrder(productID string, quantity int, accountBalance float64) (bool, float64) {
	if quantity > amountInStock(productID) || accountBalance < calcPrice(productID, quantity) {
		return false, accountBalance
	}
	return true, accountBalance - calcPrice(productID, quantity)
}
func calcPrice(productID string, quantity int) float64 {
	return priceList(productID) * float64(quantity)
}
func priceList(productID string) float64 {
	if productID == "1" {
		return 1.50
	} else if productID == "2" {
		return 2.25
	} else if productID == "3" {
		return 3.00
	} else if productID == "4" {
		return 1.00
	} else if productID == "5" {
		return 2.50
	} else if productID == "6" {
		return 8.99
	} else if productID == "7" {
		return 22.50
	} else if productID == "8" {
		return 50.00
	} else if productID == "9" {
		return 999.99
	} else {
		return 0.00
	}
}
func amountInStock(productID string) int {
	if productID == "1" {
		return 11
	} else if productID == "2" {
		return 25
	} else if productID == "3" {
		return 4
	} else if productID == "4" {
		return 6
	} else if productID == "5" {
		return 50
	} else if productID == "6" {
		return 2
	} else if productID == "7" {
		return 0
	} else if productID == "8" {
		return 99
	} else if productID == "9" {
		return 1
	} else {
		return 0
	}
}

func c3_l20() {
	utils.PrintSectionStart("Chapter 3: Lesson 20 - Closures", false)
	//(https://www.boot.dev/lessons/f2c926e4-4e10-40d3-bffb-11683a8c3c1f)
	fmt.Println(`A closure is a function that references variables from outside its own body.
Here we used a closure to aggregate a sum across multiple operations. The running total
is kept up and modified within the 'adder' function`)
	fmt.Println()

	summation := adder()
	fmt.Println("Starting with 0...")
	fmt.Printf("+3 = %d\n", summation(3))
	fmt.Printf("+4 = %d\n", summation(4))
	fmt.Printf("+42 = %d\n", summation(42))
	fmt.Printf("+7 = %d\n", summation(7))

	utils.PrintSectionEnd(false)
}
func adder() func(int) int {
	sum := 0
	return func(toAdd int) int {
		sum += toAdd
		return sum
	}
}

func c3_l21() {
	utils.PrintSectionStart("Chapter 3: Lesson 21 - Currying", false)
	//(https://www.boot.dev/lessons/00e74458-f6e2-4f14-bc04-5ad5112551d6)
	fmt.Println(`Currying is a functional programming concept that allows a function
with multiple arguments to be transformed into a sequence of functions, each taking a single argument.

Here we use a getLogger func that returns a func which formats the two strings it gets in a specific way
and prints the result to the console...`)
	fmt.Println()

	dbErrors := []error{
		errors.New("out of memory"),
		errors.New("cpu is pegged"),
		errors.New("networking issue"),
		errors.New("invalid syntax"),
	}
	testLogger("Error on database server", dbErrors, colonDelimit)

	mailErrors := []error{
		errors.New("email too large"),
		errors.New("non alphanumeric symbols found"),
	}
	testLogger("Error on mail server", mailErrors, commaDelimit)

	utils.PrintSectionEnd(false)
}

// getLogger takes a function that formats two strings into
// a single string and returns a function that formats two strings but prints
// the result instead of returning it
func getLogger(formatter func(string, string) string) func(string, string) {
	return func(a, b string) {
		fmt.Println(formatter(a, b))
	}
}
func testLogger(first string, errors []error, formatter func(string, string) string) {
	defer fmt.Println("====================================")
	logger := getLogger(formatter)
	fmt.Println("Logs:")
	for _, err := range errors {
		logger(first, err.Error())
	}
}

// used as a formatter for c3_l21
func colonDelimit(first, second string) string {
	return first + ": " + second
}

// used as a formatter for c3_l21
func commaDelimit(first, second string) string {
	return first + ", " + second
}
