package main

import "fyne.io/fyne/v2"

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
