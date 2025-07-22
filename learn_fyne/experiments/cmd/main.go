package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var data = []string{"Experiment A", "Experiment B", "Experiment C"}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Experiments")

	displayText := canvas.NewText("uh", color.Black)

	mainLayout := layout.NewCenterLayout()

	mainContainer := container.New(mainLayout, displayText)

	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			itemContainer := container.NewHBox(widget.NewLabel(""))
			itemContainer.Resize(fyne.NewSize(200, 100))
			return itemContainer
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*fyne.Container).Objects[0].(*widget.Label).SetText(data[i])
			// with an button, instead of label
			/*o.(*fyne.Container).Objects[0].(*widget.Button).OnTapped = func() {
				// Handle the click for this specific item
				println("Button clicked for:", data[i])
			}*/
		})

	leftStack := layout.NewStackLayout()
	leftContainer := container.New(leftStack, list, layout.NewSpacer())

	//baseLayout := layout.NewHBoxLayout()
	baseContainer := container.NewHSplit(leftContainer, mainContainer)
	baseContainer.SetOffset(0.3)

	list.OnSelected = func(id widget.ListItemID) {
		fmt.Printf("Selected item %d \n", id)
		displayText.Text = data[id]
		baseContainer.Refresh()
	}

	myWindow.SetContent(baseContainer)
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.SetMaster()
	myWindow.ShowAndRun()
}
