package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

var (
	todos []string
	list  *widget.List
)

func main() {
	a := app.NewWithID("xyz.andy.todo")
	w := a.NewWindow("TODO")

	todos = loadTODOs(a.Preferences())

	w.SetContent(loadUI(a))
	w.Resize(fyne.NewSize(200, 280))
	w.ShowAndRun()
}
