package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func Toolbar(win fyne.Window, list *widget.Table) *widget.Toolbar {
	open := widget.NewToolbarAction(theme.FolderOpenIcon(), func() {

		//dialog.NewInformation("提示", "提示信息-------------------", win).Show()
		updateList(list)

	})
	new := widget.NewToolbarAction(theme.FolderNewIcon(), func() {
		//boy := Body{Name: "ASD", Age: 10}
		//bodys = append(bodys, &boy)

	})

	exit := widget.NewToolbarAction(theme.WindowCloseIcon(), func() {
		//fmt.Println()
		dialog.NewConfirm("提示", "你是否要退出！", func(b bool) {
			fmt.Println(b)
			if b {
				//win.Close()
			}
		}, win).Show()
	})

	//open.
	//open.OnActivated = f
	return widget.NewToolbar(open, new, widget.NewToolbarSpacer(), exit)
	//top.Theme().Color()

}
