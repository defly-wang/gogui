package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
)

func main() {

	myapp := app.NewWithID(" 测试 ")

	icon, err := fyne.LoadResourceFromPath("Icon.png")
	if err != nil {
		icon = theme.HomeIcon()
	}
	myapp.SetIcon(icon)

	mywin := myapp.NewWindow("测试")
	mywin.Resize(fyne.NewSize(800, 600))
	mywin.CenterOnScreen()

	mywin.SetCloseIntercept(func() {
		mywin.Hide()
	})
	SystemTray(myapp, mywin)

	mywin.ShowAndRun()
}

func SystemTray(app fyne.App, win fyne.Window) {
	if desk, ok := app.(desktop.App); ok {
		desk.SetSystemTrayIcon(app.Icon())

		showmenu := fyne.NewMenuItem("Open", func() {
			win.Show()
		})

		menu := fyne.NewMenu("TrayMenu", showmenu)
		desk.SetSystemTrayMenu(menu)

	}
}
