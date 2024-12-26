package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func updateTime(clock *widget.Label) {
	fmt := time.Now().Format("Time:03:04:05")
	clock.SetText(fmt)
}

func main() {

	//var win fyne.Window
	app := app.NewWithID(" 你好！")

	icon, err := fyne.LoadResourceFromPath("Icon-new.png")
	if err != nil {
		icon = theme.HomeIcon()
	}

	app.SetIcon(icon)

	win := MainWindow(app)

	SystemTray(app, win)

	win.ShowAndRun()

}
