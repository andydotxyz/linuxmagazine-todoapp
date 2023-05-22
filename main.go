package main

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const joiner = "|"

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

func addTODO(todo string, p fyne.Preferences) {
	todos = append(todos, todo)
	list.Refresh()
	saveTODOs(todos, p)
}

func deleteTODO(todo string, p fyne.Preferences) {
	for i, text := range todos {
		if text != todo {
			continue
		}

		todos = append(todos[:i], todos[i+1:]...)
		break
	}
	list.Refresh()
	saveTODOs(todos, p)
}

func loadTODOs(p fyne.Preferences) []string {
	all := p.String("items")
	if all == "" {
		return []string{}
	}
	return strings.Split(all, joiner)
}

func saveTODOs(items []string, p fyne.Preferences) {
	allItems := strings.Join(items, joiner)
	p.SetString("items", allItems)
}
