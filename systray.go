package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

func SystemTray(app fyne.App, win fyne.Window) {
	if desk, ok := app.(desktop.App); ok {

		show := fyne.NewMenuItem("", nil)

		menu := fyne.NewMenu("SysTray", show)
		RefreshMenu(show, menu, win)

		show.Action = func() {
			if win.Content().Visible() {
				win.Hide()
				show.Label = "显示窗口"
			} else {
				win.Show()
				show.Label = "隐藏窗口"
			}
			//RefreshMenu(show, menu, win)
			menu.Refresh()
		}

		desk.SetSystemTrayIcon(app.Icon())
		desk.SetSystemTrayMenu(menu)

		win.SetCloseIntercept(func() {
			win.Hide()
			RefreshMenu(show, menu, win)
		})

	}
}

func RefreshMenu(show *fyne.MenuItem, menu *fyne.Menu, win fyne.Window) {

	if win.Content().Visible() {
		show.Label = "隐藏窗口"
	} else {
		show.Label = "显示窗口"
	}

	menu.Refresh()
}
