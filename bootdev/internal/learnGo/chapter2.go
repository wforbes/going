package learnGo

import (
	"bootdev/internal/utils"
	"fmt"
)

func runChapter2Lessons() {
	c2_l1()
	c2_l2()
	// lesson 3 is just a question related to lesson 2
	c2_l4()
	c2_l5()
}

func c2_l1() {
	utils.PrintSectionStart("Chapter 2: Lesson 1 - Conditionals", false)
	//(https://www.boot.dev/lessons/e210dea3-0c70-41c1-871b-4aa5b3658917)

	messageLen := 10
	maxMessageLen := 20
	fmt.Println("Trying to send a message of length:", messageLen, "and a max length of:", maxMessageLen)

	if messageLen < maxMessageLen {
		fmt.Println("Message is within the limit")
	} else {
		fmt.Println("Message is too long")
	}

	utils.PrintSectionEnd(false)
}

func c2_l2() {
	utils.PrintSectionStart("Chapter 2: Lesson 2 - The Initial Statement of an If Block", false)
	//(https://www.boot.dev/lessons/84e5481a-ca9c-40a3-ad0f-64d275226859)

	var message string = "Hello, world!"
	maxMessageLen := 30

	if messageLen := len(message); messageLen > maxMessageLen {
		fmt.Println("Message is too long")
	} else {
		fmt.Println("Message is within the limit")
	}

	utils.PrintSectionEnd(false)
}

func c2_l4() {
	utils.PrintSectionStart("Chapter 2: Lesson 4 - Switch", false)
	//(https://www.boot.dev/lessons/f8b2fcea-078b-41be-b59b-8bece5ae923b)
	plan := "basic"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
	plan = "pro"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
	plan = "enterprise"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
	plan = "free"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
	plan = "unknown"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))

	utils.PrintSectionEnd(false)
}

func billingCost(plan string) float64 {
	switch plan {
	case "basic":
		return 10.00
	case "pro":
		return 20.0
	case "enterprise":
		return 50.0
	default:
		return 0.0
	}
}

func c2_l5() {
	utils.PrintSectionStart("Chapter 2: Lesson 5 - Calculate Balance", false)
	//(https://www.boot.dev/lessons/f6c53b4a-3bfe-49ca-8733-aa15c19fe7b1)

	var insufficientFundMessage string = "Purchase failed. Insufficient funds."
	var purchaseSuccessMessage string = "Purchase successful."
	var accountBalance float64 = 100.0
	var bulkMessageCost float64 = 75.0
	var isPremiumUser bool = true
	var discountRate float64 = 0.10
	var finalCost float64

	finalCost = bulkMessageCost
	if isPremiumUser {
		finalCost = bulkMessageCost * (1 - discountRate)
	}
	if finalCost <= accountBalance {
		accountBalance -= finalCost
		fmt.Println(purchaseSuccessMessage)
	} else {
		fmt.Println(insufficientFundMessage)
	}

	fmt.Println("Account balance:", accountBalance)

	utils.PrintSectionEnd(false)
}
