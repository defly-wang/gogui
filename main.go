package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

func main() {

	app := app.NewWithID(" 你好！")

	icon, err := fyne.LoadResourceFromPath("Icon.png")
	if err != nil {
		icon = theme.HomeIcon()
	}

	app.SetIcon(icon)

	win := MainWindow(app)
	win.CenterOnScreen()
	win.Show()

	SystemTray(app, win)

	app.Run()

}
