package learnGo

import (
	"fmt"
	"going/internal/utils"
)

func runChapter2Lessons() {
	chapter2Lesson1()
	chapter2Lesson2()
}

func chapter2Lesson1() {
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

func chapter2Lesson2() {
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
