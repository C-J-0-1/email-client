package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type emailClientTheme struct {
	fyne.Theme
}

func newClientTheme() fyne.Theme {
	return &emailClientTheme{
		Theme: theme.DefaultTheme(),
	}
}

func (t *emailClientTheme) Color(name fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	return t.Theme.Color(name, theme.VariantLight)
}

func (t *emailClientTheme) Size(name fyne.ThemeSizeName) float32 {
	if name == theme.SizeNameText {
		return 14
	}

	return t.Theme.Size(name)
}
