package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.NewWithID("xyz.andy.todo")
	w := a.NewWindow("TODO")

	w.SetContent(loadUI())
	w.Resize(fyne.NewSize(200, 280))
	w.ShowAndRun()
}

func loadUI() fyne.CanvasObject {
	list := widget.NewList(
		func() int {
			return 5
		},
		func() fyne.CanvasObject {
			return widget.NewCheck("TODO Item", func(bool) {})
		},
		func(id widget.ListItemID, o fyne.CanvasObject) {
		})

	input := widget.NewEntry()
	input.SetPlaceHolder("New item")
	add := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		log.Println("new item tapped")
	})
	head := container.NewBorder(nil, nil, nil, add, input)

	return container.NewBorder(head, nil, nil, nil, list)
}
