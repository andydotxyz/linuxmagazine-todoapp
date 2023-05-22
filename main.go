package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var (
	todos []string
	list  *widget.List
)

func main() {
	a := app.NewWithID("xyz.andy.todo")
	w := a.NewWindow("TODO")

	w.SetContent(loadUI())
	w.Resize(fyne.NewSize(200, 280))
	w.ShowAndRun()
}

func loadUI() fyne.CanvasObject {
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

				deleteTODO(check.Text)
			}
		})

	input := widget.NewEntry()
	input.SetPlaceHolder("New item")
	add := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		addTODO(input.Text)
		input.SetText("")
	})
	input.OnSubmitted = func(item string) {
		addTODO(item)
		input.SetText("")
	}
	head := container.NewBorder(nil, nil, nil, add, input)

	return container.NewBorder(head, nil, nil, nil, list)
}

func addTODO(todo string) {
	todos = append(todos, todo)
	list.Refresh()
}

func deleteTODO(todo string) {
	for i, text := range todos {
		if text != todo {
			continue
		}

		todos = append(todos[:i], todos[i+1:]...)
		break
	}
	list.Refresh()
}
