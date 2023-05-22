package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.NewWithID("xyz.andy.todo")
	w := a.NewWindow("TODO")

	w.SetContent(widget.NewLabel("TODO App"))
	w.ShowAndRun()
}
