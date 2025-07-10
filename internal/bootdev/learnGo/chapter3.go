package learnGo

import (
	"fmt"
	"going/internal/utils"
)

func runChapter3Lessons() {
	chapter3Lesson1()
}

func chapter3Lesson1() {
	utils.PrintSectionStart("Chapter 3: Lesson 1 - Functions", false)
	//(https://www.boot.dev/lessons/9aedf839-7d94-43f7-82d0-1d27e5d0b79c)

	test1("Lane,", " happy birthday!")
	test1("Zuck,", " hope that Metaverse thing works out")
	test1("Go", " is fantastic")

	utils.PrintSectionEnd(false)
}

func concat(s1 string, s2 string) string {
	return s1 + s2
}

func test1(s1 string, s2 string) {
	fmt.Println(concat(s1, s2))
}
