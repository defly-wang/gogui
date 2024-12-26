package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
)

func SystemTray(app fyne.App, win fyne.Window) {
	if desk, ok := app.(desktop.App); ok {

		show := fyne.NewMenuItem("显示窗口", func() {
			//app.Quit()
			win.Show()
		})
		show.Icon = theme.SearchIcon()

		menu := fyne.NewMenu("SysTray", show)

		desk.SetSystemTrayIcon(app.Icon())

		//desk.SetSystemTrayIcon()
		desk.SetSystemTrayMenu(menu)
		//win.SetMainMenu(fyne.NewMainMenu(menu))

	}

}
