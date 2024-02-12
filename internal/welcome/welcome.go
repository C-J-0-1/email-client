package welcome

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type Welcome struct {
	title   string
	stack   []fyne.CanvasObject
	content *fyne.Container

	d dialog.Dialog
}

func NewWelcome(title string, content fyne.CanvasObject) *Welcome {
	w := &Welcome{title: title, stack: []fyne.CanvasObject{content}}
	w.content = container.NewStack(content)
	return w
}

func (w *Welcome) Hide() {
	w.d.Hide()
}

func (w *Welcome) Show(win fyne.Window) {
	w.d = dialog.NewCustomWithoutButtons(w.title, w.content, win)

	w.d.Show()
}

func (w *Welcome) Pop() {
	if len(w.stack) <= 1 {
		return
	}
	w.stack = w.stack[:len(w.stack)-1]

	w.content.Objects = []fyne.CanvasObject{w.stack[len(w.stack)-1]}
	w.content.Refresh()
}

func (w *Welcome) Push(title string, content fyne.CanvasObject) {
	w.stack = append(w.stack, w.wrap(title, content))

	w.content.Objects = []fyne.CanvasObject{w.stack[len(w.stack)-1]}
	w.content.Refresh()
}

func (w *Welcome) wrap(title string, content fyne.CanvasObject) fyne.CanvasObject {
	backButton := widget.NewButtonWithIcon("Back", theme.NavigateBackIcon(), w.Pop)

	return container.NewHBox(content, backButton)
}

func (w *Welcome) Resize(size fyne.Size) {
	if w.d == nil {
		return
	}

	w.d.Resize(size)
}

// func (w *Welcome) MakeWelcome() fyne.CanvasObject {
// 	content := widget.NewButtonWithIcon("Login", theme.LoginIcon(), func() {
// 		fmt.Println("logged in!!")
// 	})

// 	top := widget.NewLabel("")

// 	bottom := widget.NewLabel("")

// 	objs := []fyne.CanvasObject{top, content, bottom}
// 	return container.New(newWelcomeLayout(top, content, bottom), objs...)
// }
