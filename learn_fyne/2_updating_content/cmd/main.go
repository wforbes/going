// Fyne Getting Started: Updating Content
// https://www.youtube.com/watch?v=h2ZOdTA3ew4
// update: found that marking a moment in time with a button click
//
//	seemed to be causing a rounding up second because the UI updates
//	seconds approx .28s after the actual channel tick
package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("The time is 03:04:05.000")
	clock.SetText(formatted)
}

func main() {
	a := app.New()
	w := a.NewWindow("Clock")

	clock := widget.NewLabel("")
	updateTime(clock)

	markedTimeLabel := widget.NewLabel("Marked = --:--:--.---")
	button := widget.NewButton("Mark Time", func() {
		formattedTime := time.Now().Format("Marked = 03:04:05.000")
		fmt.Println(formattedTime)
		markedTimeLabel.SetText(formattedTime)
	})
	w.SetContent(container.NewVBox(
		clock,
		markedTimeLabel,
		button,
	))
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
