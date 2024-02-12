//go:generate fyne bundle -o bundled.go assets

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/binding"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(newClientTheme())

	w := a.NewWindow("Email Client")
	w.Resize(fyne.NewSize(1024, 768))

	ui := gui{win: w, text: binding.NewString()}
	w.SetContent(ui.makeGUI())
	w.SetMainMenu(ui.makeMenu())
	ui.text.AddListener(binding.NewDataListener(func() {
		text, _ := ui.text.Get()
		w.SetTitle(text)
	}))

	ui.readEmails()
	// ui.showWelcome(w)

	w.ShowAndRun()
}
