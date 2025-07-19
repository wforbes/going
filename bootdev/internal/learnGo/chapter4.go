package learnGo

import (
	"bootdev/internal/utils"
	"fmt"
)

func runChapter4Lessons() {
	c4_l1()
}

func c4_l1() {
	utils.PrintSectionStart("Chapter 4: Lesson 1 - Structs in Go", false)
	fmt.Println(`Structs in Go are used to represent structured data,
grouping different types of vars together. For example, representing a car:

type car struct {
	brand	string
	model	string
	doors	int
	mileage	int
}

The exercise here is to complete the definition of the 'messageToSend' struct
which needs 2 fields: phoneNumber (int) and message (string)`)
	fmt.Println()

	result := getMessageText(messageToSend{
		16612664837, "This weird learning repo isn't it?",
	})
	fmt.Println(result)

	utils.PrintSectionEnd(false)
}

type messageToSend struct {
	phoneNumber int
	message     string
}

func getMessageText(m messageToSend) string {
	return fmt.Sprintf("Sending message: '%s' to: %v", m.message, m.phoneNumber)
}
