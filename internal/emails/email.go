package emails

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func ReadEmail(u string) fyne.CanvasObject {
	return widget.NewLabel(u)
}
