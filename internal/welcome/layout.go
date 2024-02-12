package welcome

import "fyne.io/fyne/v2"

type welcomeLayout struct {
	top, content, bottom fyne.CanvasObject
}

func newWelcomeLayout(top, content, bottom fyne.CanvasObject) fyne.Layout {
	return &welcomeLayout{top: top, content: content, bottom: bottom}
}

func (l *welcomeLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	topHeight := l.top.MinSize().Height
	l.top.Resize(fyne.NewSize(size.Width, topHeight))

	l.content.Move(fyne.NewPos(0, topHeight*2))
	l.content.Resize(fyne.NewSize(size.Width, topHeight*2))

	l.bottom.Move(fyne.NewPos(0, topHeight*4))
	l.bottom.Resize(fyne.NewSize(size.Width, topHeight))
}

func (l *welcomeLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(l.top.MinSize().Width*10, l.top.MinSize().Height*6)
}
