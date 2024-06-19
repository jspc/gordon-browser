package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

const (
	LabelsColour fyne.ThemeColorName = "label0"
	TagsColour   fyne.ThemeColorName = "tag0"
)

type GordonTheme struct {
	fyne.Theme
}

func NewGordonTheme() *GordonTheme {
	return &GordonTheme{
		Theme: theme.DefaultTheme(),
	}
}

func (t GordonTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameInputBackground:
		return theme.DefaultTheme().Color(theme.ColorNameBackground, variant)

	case theme.ColorNameDisabled:
		return theme.DefaultTheme().Color(theme.ColorNameForeground, variant)

	case LabelsColour:
		return color.RGBA{R: 5, G: 102, B: 8, A: 255}

	case TagsColour:
		return color.RGBA{R: 90, G: 34, B: 139, A: 255}
	}

	return theme.DefaultTheme().Color(name, variant)
}
