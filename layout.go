package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

const sideWIdth = 220

type emailClientLayout struct {
	top, left, content fyne.CanvasObject
	seperators         [2]fyne.CanvasObject
}

func newEmailClientLayout(top, left, right, content fyne.CanvasObject, seperators [2]fyne.CanvasObject) fyne.Layout {
	return &emailClientLayout{top: top, left: left, content: content, seperators: seperators}
}

func (l *emailClientLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	topHeight := l.top.MinSize().Height
	l.top.Resize(fyne.NewSize(size.Width, topHeight))

	l.left.Move(fyne.NewPos(0, topHeight))
	l.left.Resize(fyne.NewSize(sideWIdth, size.Height-topHeight))

	// l.right.Move(fyne.NewPos(size.Width-sideWIdth, topHeight))
	// l.right.Resize(fyne.NewSize(sideWIdth, size.Height-topHeight))

	l.content.Move(fyne.NewPos(sideWIdth, topHeight))
	l.content.Resize(fyne.NewSize(size.Width-sideWIdth, size.Height-topHeight))

	seperatorThickness := theme.SeparatorThicknessSize()
	l.seperators[0].Move(fyne.NewPos(0, topHeight))
	l.seperators[0].Resize(fyne.NewSize(size.Width, seperatorThickness))

	l.seperators[1].Move(fyne.NewPos(sideWIdth, topHeight))
	l.seperators[1].Resize(fyne.NewSize(seperatorThickness, size.Height-topHeight))

	// l.seperators[2].Move(fyne.NewPos(size.Width-sideWIdth, topHeight))
	// l.seperators[2].Resize(fyne.NewSize(seperatorThickness, size.Height-topHeight))
}

func (l *emailClientLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	borders := fyne.NewSize(sideWIdth*2, l.top.MinSize().Height)
	return borders.AddWidthHeight(100, 100)
}
