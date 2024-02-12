package main

import (
	"fmt"
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/c-j-0-1/email-client/internal/emails"
	"github.com/c-j-0-1/email-client/internal/welcome"
)

type gui struct {
	win  fyne.Window
	text binding.String

	emailTree binding.StringTree
	content   *container.DocTabs
	openTabs  map[string]*container.TabItem
}

func makeBanner() fyne.CanvasObject {
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(), func() {}),
	)

	logo := canvas.NewImageFromResource(resourceEmailOrgPng)
	logo.FillMode = canvas.ImageFillContain

	return container.NewStack(toolbar, container.NewPadded(logo))
}

func (g *gui) makeGUI() fyne.CanvasObject {
	top := makeBanner()

	g.emailTree = binding.NewStringTree()
	emailList := widget.NewTreeWithData(g.emailTree, func(b bool) fyne.CanvasObject {
		return widget.NewLabel("new email")
	}, func(di binding.DataItem, b bool, co fyne.CanvasObject) {
		l := co.(*widget.Label)
		u, _ := di.(binding.String).Get()

		l.SetText(u)
	})
	emailList.OnSelected = func(uid widget.TreeNodeID) {
		u, err := g.emailTree.GetValue(uid)
		if err != nil {
			fmt.Println(err)
			return
		}

		// TODO read actual email
		l := emails.ReadEmail(u)
		// checking if tab is already opened
		if item, ok := g.openTabs[u]; ok {
			g.content.Select(item)
			return
		}

		email := container.NewTabItemWithIcon("email", theme.MailComposeIcon(), l)
		if g.openTabs == nil {
			g.openTabs = make(map[string]*container.TabItem)
		}
		// adding tab in map so we don't open same tab twice
		g.openTabs[u] = email

		g.content.Append(email)
		g.content.Select(email)

		// g.content.Objects = []fyne.CanvasObject{l}
		// g.content.Refresh()
	}

	left := widget.NewAccordion(
		widget.NewAccordionItem("Emails", emailList),
		widget.NewAccordionItem("Read", widget.NewLabel("To Read")),
	)
	// email accordian will open by default
	left.Open(0)
	left.MultiOpen = true

	window := container.NewPadded(widget.NewLabel("read email"))
	email := container.NewBorder(nil, nil, nil, nil, container.NewCenter(window))
	content := container.NewStack(canvas.NewRectangle(color.Gray{Y: 0xee}), email)

	g.content = container.NewDocTabs(container.NewTabItemWithIcon("email", theme.MailComposeIcon(), content))
	g.content.CloseIntercept = func(ti *container.TabItem) {
		key := ""
		for id, item := range g.openTabs {
			if item == ti {
				key = id
			}
		}

		if key != "" {
			delete(g.openTabs, key)
		}

		g.content.Remove(ti)
	}

	seperators := [2]fyne.CanvasObject{
		widget.NewSeparator(), widget.NewSeparator(),
	}
	objs := []fyne.CanvasObject{g.content, top, left, nil, seperators[0], seperators[1]}
	return container.New(newEmailClientLayout(top, left, nil, g.content, seperators), objs...)
}

func (g *gui) makeMenu() *fyne.MainMenu {
	return fyne.NewMainMenu(
		fyne.NewMenu("Options",
			fyne.NewMenuItem("Login", func() {
				g.text.Set("Logged in!!!")
			}),
		),
	)
}

func (g *gui) readEmails() {
	// TODO implemment main list email logic
	for i := 1; i < 6; i++ {
		g.emailTree.Append(binding.DataTreeRootID, strconv.Itoa(i), "email-"+strconv.Itoa(i))
	}
}

func (g *gui) showWelcome(w fyne.Window) {
	var welcomeWiz *welcome.Welcome

	login := widget.NewLabel("Login")
	loginButton := widget.NewButtonWithIcon("Login", theme.LoginIcon(), func() {
		loggedIn := widget.NewLabel("Logged in!!")

		welcomeWiz.Push("", loggedIn)
	})
	loginBox := container.NewHBox(login, loginButton)

	welcomeWiz = welcome.NewWelcome("Welcome", loginBox)
	welcomeWiz.Show(w)
	welcomeWiz.Resize(loginBox.MinSize().AddWidthHeight(80, 20))
}
