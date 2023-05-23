package main

import (
	"strings"

	"fyne.io/fyne/v2"
)

const joiner = "|"

func loadTODOs(p fyne.Preferences) []string {
	all := p.StringWithFallback("items", "Do this item"+joiner+"Learn Fyne!"+joiner+"Build an app")
	if all == "" {
		return []string{}
	}
	return strings.Split(all, joiner)
}

func saveTODOs(items []string, p fyne.Preferences) {
	allItems := strings.Join(items, joiner)
	p.SetString("items", allItems)
}
