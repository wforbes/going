package bootdev

import (
	"fmt"
	"going/internal/utils"
)

type Chapter struct {
	Number int
	Title  string
	Runner func()
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
