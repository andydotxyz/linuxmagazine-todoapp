package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func loadUI(a fyne.App) fyne.CanvasObject {
	p := a.Preferences()
	list = widget.NewList(
		func() int {
			return len(todos)
		},
		func() fyne.CanvasObject {
			return widget.NewCheck("TODO Item", func(bool) {})
		},
		func(id widget.ListItemID, o fyne.CanvasObject) {
			check := o.(*widget.Check)
			check.SetChecked(false)
			check.Text = todos[id]
			check.Refresh()
			check.OnChanged = func(done bool) {
				if !done {
					return
				}

				deleteTODO(check.Text, p)
			}
		})

	input := widget.NewEntry()
	input.SetPlaceHolder("New item")
	add := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		addTODO(input.Text, p)
		input.SetText("")
	})
	input.OnSubmitted = func(item string) {
		addTODO(item, p)
		input.SetText("")
	}
	head := container.NewBorder(nil, nil, nil, add, input)

	return container.NewBorder(head, nil, nil, nil, list)
}
