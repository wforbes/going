package learnGo

import (
	"fmt"
	"going/internal/utils"
)

func runChapter3Lessons() {
	chapter3Lesson1()
	chapter3Lesson2()
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
