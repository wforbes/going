package alexmux

import (
	"fmt"
	"yt_alexmux/internal/utils"
)

type Video struct {
	Title     string
	Tutorials []Tutorial
}

type Tutorial struct {
	Number int
	Title  string
	Runner func()
}

func (v Video) RunTutorials() {
	utils.PrintSectionStart(fmt.Sprintf("Video: %s", v.Title), true)
	for _, tutorial := range v.Tutorials {
		tutorial.Runner()
	}
	//utils.PrintSectionEnd(true)
}
