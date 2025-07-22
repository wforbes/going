// Fyne Getting Started: Updating Content
// https://www.youtube.com/watch?v=h2ZOdTA3ew4
package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("The time is 03:04:05")
	clock.SetText(formatted)
}

func main() {
	a := app.New()
	w := a.NewWindow("Clock")

	clock := widget.NewLabel("")
	updateTime(clock)
	w.SetContent(clock)
	go func() {
		ticker := time.NewTicker(time.Second)
		for range ticker.C {
			fyne.Do(func() {
				updateTime(clock)
			})

		}
	}()

	w.ShowAndRun()
}
