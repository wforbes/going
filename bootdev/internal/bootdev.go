package bootdev

import (
	"bootdev/internal/utils"
	"fmt"
)

type Lesson struct {
	Number int
	Title  string
	Runner func()
}

func (l Lesson) RunLesson() {
	l.Runner()
}

type Chapter struct {
	Number  int
	Title   string
	Runner  func()
	Lessons []Lesson
}

func (c Chapter) RunLessons() {
	utils.PrintSectionStart(fmt.Sprintf("Chapter %d: %s", c.Number, c.Title), true)
	c.Runner()
	utils.PrintSectionEnd(true)
}

type Course struct {
	Title    string
	Chapters []Chapter
}

func (c Course) RunChapters() {
	utils.PrintSectionStart(fmt.Sprintf("Course: %s", c.Title), true)
	for _, chapter := range c.Chapters {
		chapter.RunLessons()
	}
	//utils.PrintSectionEnd(true)
}
